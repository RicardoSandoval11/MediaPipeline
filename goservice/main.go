package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/RicardoSandoval11/MediaPipeline/goservice/conf"
	"github.com/RicardoSandoval11/MediaPipeline/goservice/db"
	"github.com/RicardoSandoval11/MediaPipeline/goservice/db/migrations"
	"github.com/RicardoSandoval11/MediaPipeline/goservice/pkg/infra"
	"github.com/RicardoSandoval11/MediaPipeline/goservice/pkg/upload"
	httptransport "github.com/go-kit/kit/transport/http"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	ctx, cancel := signal.NotifyContext(
		context.Background(),
		os.Interrupt,
		syscall.SIGTERM,
	)

	defer cancel()

	cfg, err := conf.GetEnv()
	if err != nil {
		log.Fatalf("could not get env variables: %v", err.Error())
		os.Exit(1)
	}

	conn, err := grpc.Dial(cfg.FileCounterGrpcAddr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("could not establish to gRPC server: %v", err)
		os.Exit(1)
	}
	defer conn.Close()

	dbInstance := db.GetDatabase(cfg.DbConnectionString)

	if err := migrations.Run(dbInstance); err != nil {
		log.Fatalf("could not apply migrations to database: %v", err.Error())
		os.Exit(1)
	}

	fileRepository := upload.NewPostgresqlRepository(dbInstance)
	fileCounterEndpoint := upload.NewFileCounterGRPCClient(conn)
	fileService := upload.NewUploadService(fileCounterEndpoint, fileRepository)

	fileEndpoint := upload.MakeUploadFileEndpoint(fileService)
	fileUploadHandler := httptransport.NewServer(
		fileEndpoint,
		upload.DecodeUploadFileRequest,
		upload.EncodeUploadFileResponse,
	)

	getFileEndpoint := upload.MakeGetFileEndpoint(fileService)
	getFileHandler := httptransport.NewServer(
		getFileEndpoint,
		upload.DecodeGetFileRequest,
		upload.EncodeGetFileResponse,
	)

	infraService := infra.NewInfraService()
	livenessEndpoint := infra.MakeLivenessEndpoint(infraService)
	readinessEndpoint := infra.MakeReadinessEndpoint(infraService)

	livenessHandler := httptransport.NewServer(
		livenessEndpoint,
		infra.DecodeLivenessRequest,
		infra.EncodeLivenessResponse,
	)

	readinessHandler := httptransport.NewServer(
		readinessEndpoint,
		infra.DecodeReadinessRequest,
		infra.EncodeReadinessResponse,
	)

	mux := http.NewServeMux()
	mux.Handle("/api/v1/file/upload-file", fileUploadHandler)
	mux.Handle("/api/v1/file/get-file", getFileHandler)
	mux.Handle("/api/v1/infra/readyz", readinessHandler)
	mux.Handle("/api/v1/infra/livez", livenessHandler)

	errs := make(chan error)

	srv := &http.Server{
		Addr:    cfg.ApplicationPort,
		Handler: mux,
	}

	go func() {
		errs <- srv.ListenAndServe()
	}()

	select {
	case e := <-errs:
		{
			log.Fatal(
				"Application was not able to start",
				"error", e,
			)
		}
	case <-ctx.Done():
		{
			fmt.Println("Shutting done...")
		}
	}
}
