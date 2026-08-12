# syntax=docker/dockerfile:1.7

############################
# 1. Base runtime
############################
FROM mcr.microsoft.com/dotnet/aspnet:10.0 AS base

WORKDIR /app

# REST
EXPOSE 5144

# gRPC
EXPOSE 5145

# ASP.NET Core container images provide a non-root "app" user.
USER app

############################
# 2. Build
############################
FROM mcr.microsoft.com/dotnet/sdk:10.0 AS build

ARG BUILD_CONFIGURATION=Release

WORKDIR /src

# Copy project files first to maximize Docker layer caching.
COPY ["dotnetservice/dotnetsevice.csproj", "dotnetservice/"]

# If you have other projects:
COPY ["dotnetservice.DataAccess/dotnetservice.DataAccess.csproj", "dotnetservice.DataAccess/"]
COPY ["dotnetservice.Interfaces/dotnetservice.Interfaces.csproj", "dotnetservice.Interfaces/"]
COPY ["dotnetservice.Models/dotnetservice.Models.csproj", "dotnetservice.Models/"]
COPY ["dotnetservice.Repositories/dotnetservice.Repositories.csproj", "dotnetservice.Repositories/"]
COPY ["dotnetservice.Services/dotnetservice.Services.csproj", "dotnetservice.Services/"]

RUN --mount=type=cache,id=nuget,target=/root/.nuget/packages \
    dotnet restore "dotnetservice/dotnetsevice.csproj"

# Copy source
COPY . .

WORKDIR "/src/dotnetservice"

RUN --mount=type=cache,id=nuget,target=/root/.nuget/packages \
    dotnet build "dotnetsevice.csproj" \
    --configuration $BUILD_CONFIGURATION \
    --no-restore \
    --output /app/build


############################
# 3. Publish
############################
FROM build AS publish

ARG BUILD_CONFIGURATION=Release

RUN --mount=type=cache,id=nuget,target=/root/.nuget/packages \
    dotnet publish "dotnetsevice.csproj" \
    --configuration $BUILD_CONFIGURATION \
    --no-restore \
    --output /app/publish \
    --no-self-contained


############################
# 4. Final image
############################
FROM base AS final

WORKDIR /app

COPY --from=publish /app/publish .

ENTRYPOINT ["dotnet", "dotnetsevice.dll"]