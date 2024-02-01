# Start from a base image with the .NET SDK
FROM mcr.microsoft.com/dotnet/sdk:latest

# Set the working directory
WORKDIR /app

# Copy the project files
COPY *.csproj ./
RUN dotnet restore

# Copy everything else and build
COPY . ./
RUN dotnet publish -c Release -o out

# Expose the necessary port
EXPOSE 5002

# Run the application
ENTRYPOINT ["dotnet", "out/service1_csharp.dll"]
