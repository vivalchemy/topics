# Use an official Bun base image
FROM oven/bun:alpine

# Set working directory
WORKDIR /app

# Copy package.json, bun.lockb, and tsconfig.json to the container
COPY package.json bun.lockb tsconfig.json ./

# Install dependencies
RUN bun install;\
  bun run migrate;\
  bun run db

# Copy the rest of the application files
COPY . .

# Expose the port your development server will run on (e.g., 3000)
EXPOSE 3000

# Command to start the development server
CMD ["bun", "run", "index.ts"]
