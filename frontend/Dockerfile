# Stage 1: Build the React application
FROM node:20.19.0-alpine AS builder
WORKDIR /app

# Copy package.json and yarn.lock
COPY package.json yarn.lock ./

# Install dependencies with Yarn
RUN yarn install --frozen-lockfile

# Copy the rest of the source code
COPY . .

# Build the application for production
RUN yarn build

# Stage 2: Serve the built application with Nginx
FROM nginx:1.25-alpine

# Copy our custom Nginx configuration
COPY nginx.conf /etc/nginx/conf.d/default.conf

# Copy the built static files from the builder stage
COPY --from=builder /app/dist /usr/share/nginx/html

# Expose port 80
EXPOSE 80

# Start Nginx
CMD ["nginx", "-g", "daemon off;"]
