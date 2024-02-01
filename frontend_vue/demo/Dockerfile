# Stage 1: Build the Vue.js application
FROM node:14 as build-stage

# Set the working directory in the Docker container
WORKDIR /app

# Copy package.json and package-lock.json (if available)
COPY package*.json ./

# Install project dependencies
RUN npm install

# Copy project files and folders to the current working directory (i.e. 'app' folder)
COPY . .

# Build the application for production
RUN npm run build

# Stage 2: Serve the application from Nginx
FROM nginx:stable

# Copy the built app to Nginx server
COPY --from=build-stage /app/dist /usr/share/nginx/html

# Expose port 80
EXPOSE 80

# Start Nginx server
CMD ["nginx", "-g", "daemon off;"]
