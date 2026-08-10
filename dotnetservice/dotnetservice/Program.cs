using dotnetservice.DataAccess.Config;
using dotnetservice.Interfaces.Repositories;
using dotnetservice.Interfaces.Services;
using dotnetservice.Repositories;
using dotnetservice.Services.gRPC;
using dotnetservice.Services.REST;
using Microsoft.AspNetCore.Server.Kestrel.Core;

var builder = WebApplication.CreateBuilder(args);

// gRPC requires HTTP/2; allow both protocols so REST (HTTP/1.1) keeps working on the same endpoint
builder.WebHost.ConfigureKestrel(options =>
{
    options.ConfigureEndpointDefaults(o => o.Protocols = HttpProtocols.Http1AndHttp2);
});

// Add services to the container.
builder.Services.AddControllers();

// Learn more about configuring OpenAPI at https://aka.ms/aspnet/openapi
builder.Services.AddOpenApi();

string connectionString = builder.Configuration["DbSettings:ConnectionString"] ?? "";

builder.Services.AddDatabase(connectionString);

builder.Services.AddCors(options =>
            {
                options.AddDefaultPolicy(
                    builder =>
                    {
                        builder
                        .WithOrigins("*")
                        .AllowAnyMethod()
                        .WithHeaders("accept", "authorization", "content-type", "origin")
                        .AllowCredentials();
                    }
                );
            });

builder.Services.AddGrpc();

// Repositories
builder.Services.AddScoped<IUserRepository, UserRepository>();
builder.Services.AddScoped<IFileCounterRepository, FileCounterRepository>();

// Services
builder.Services.AddScoped<IUserService, UserService>();

var app = builder.Build();

// Configure the HTTP request pipeline.
if (app.Environment.IsDevelopment())
{
    app.MapOpenApi();
}

// Apply migrations
var dbSync = new SyncMigrations(app.Services);

dbSync.SyncPendingMigrations();

app.UseRouting();

app.MapControllers();
app.MapGrpcService<FileCounterService>();

await app.RunAsync();