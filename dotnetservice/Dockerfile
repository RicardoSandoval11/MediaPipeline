# syntax=docker/dockerfile:1.7

############################
# 1. Build
############################
FROM mcr.microsoft.com/dotnet/sdk:10.0 AS build
ARG BUILD_CONFIGURATION=Release

WORKDIR /src

COPY ["dotnetservice/dotnetsevice.csproj", "dotnetservice/"]
COPY ["dotnetservice.DataAccess/dotnetservice.DataAccess.csproj", "dotnetservice.DataAccess/"]
COPY ["dotnetservice.Interfaces/dotnetservice.Interfaces.csproj", "dotnetservice.Interfaces/"]
COPY ["dotnetservice.Models/dotnetservice.Models.csproj", "dotnetservice.Models/"]
COPY ["dotnetservice.Repositories/dotnetservice.Repositories.csproj", "dotnetservice.Repositories/"]
COPY ["dotnetservice.Services/dotnetservice.Services.csproj", "dotnetservice.Services/"]

# Agregamos -r linux-x64 aquí para que el restore incluya los assets del RID objetivo
RUN --mount=type=cache,id=nuget,target=/root/.nuget/packages \
    dotnet restore "dotnetservice/dotnetsevice.csproj" -r linux-x64

COPY . .
WORKDIR "/src/dotnetservice"

############################
# 2. Publish
############################
FROM build AS publish
ARG BUILD_CONFIGURATION=Release

RUN --mount=type=cache,id=nuget,target=/root/.nuget/packages \
    dotnet publish "dotnetsevice.csproj" \
    --configuration $BUILD_CONFIGURATION \
    --no-restore \
    --output /app/publish \
    --self-contained true \
    -r linux-x64 \
    -p:PublishSingleFile=true \
    -p:PublishTrimmed=false

############################
# 3. Final Image (Chiseled)
############################
FROM mcr.microsoft.com/dotnet/nightly/runtime-deps:10.0-noble-chiseled AS final

WORKDIR /app
EXPOSE 5144
EXPOSE 5145

USER app

COPY --from=publish /app/publish .

ENTRYPOINT ["./dotnetsevice"]