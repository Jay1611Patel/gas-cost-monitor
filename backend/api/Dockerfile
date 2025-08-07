# Use an official Node.js runtime as a parent image
FROM node:18-alpine

# Set working directory
WORKDIR /usr/app

# Copy package.json and yarn.lock from backend/api
COPY package.json yarn.lock ./

# Install dependencies
RUN yarn install --frozen-lockfile

# Copy the rest of the application code
COPY . .

# Default command (API entrypoint, not reporter.js)
CMD ["node", "src/app.js"]
