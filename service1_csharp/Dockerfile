# Use the official .NET Core 5.0 SDK image as the base image
FROM mcr.microsoft.com/dotnet/sdk:5.0

# Set the working directory inside the container
WORKDIR /app

# Copy the project files to the container
COPY . .

# Build the C# application
RUN dotnet publish -c Release -o out

# Expose the port your C# application listens on (if needed)
EXPOSE 5002

# Run the C# application
ENTRYPOINT ["dotnet", "out/service1_csharp.dll"]
