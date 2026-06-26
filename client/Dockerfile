# Stage 1: Build SvelteKit application
FROM node:20-alpine AS builder

WORKDIR /app

# Copy package files first to cache dependencies layer
COPY package.json package-lock.json ./

# Install all dependencies (including devDependencies)
RUN npm ci

# Copy the rest of the application source code
COPY . .

# Build the production bundle
RUN npm run build

# Stage 2: Create minimal runner image
FROM node:20-alpine AS runner

WORKDIR /app

# Install sirv-cli globally to serve the static assets
RUN npm install -g sirv-cli

# Copy built app files from the builder stage
COPY --from=builder /app/build ./build

# Expose server port
EXPOSE 3000

# Set default runtime environment variables
ENV PORT=3000

# Start static server in SPA mode (redirecting non-file routes to index.html)
CMD ["sirv", "build", "--port", "3000", "--single", "--host", "0.0.0.0"]
