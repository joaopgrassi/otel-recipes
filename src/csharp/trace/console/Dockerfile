FROM mcr.microsoft.com/dotnet/sdk:7.0 AS build
WORKDIR /source

# Copy csproj and restore as distinct layers
COPY Console.csproj ./

# Copy everything else and build
COPY . ./
RUN dotnet publish -c Release -o /app --no-cache

# final stage/image
FROM mcr.microsoft.com/dotnet/runtime:7.0
WORKDIR /app
COPY --from=build /app ./
ENTRYPOINT ["dotnet", "Console.dll"]
