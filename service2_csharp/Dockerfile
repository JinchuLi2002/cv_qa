# Use the official .NET Core 5.0 SDK image as the base image
FROM base-csharp as build

# Copy the project files to the container
COPY . .

# Build the C# application
RUN dotnet publish -c Release -o out

# Expose the port your C# application listens on (if needed)
EXPOSE 5003

# Run the C# application
ENTRYPOINT ["dotnet", "out/service1_csharp.dll"]
