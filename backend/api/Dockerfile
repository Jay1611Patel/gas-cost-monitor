
FROM node:20-alpine

# Set the working directory
WORKDIR /usr/src/app

# Copy package.json and yarn.lock to install dependencies first
COPY package.json yarn.lock ./

# Install dependencies
RUN yarn install --production --frozen-lockfile

# Copy the rest of the application code
COPY . .

# Expose the API port
EXPOSE 3000

# Command to run the application
CMD ["node", "src/app.js"]
