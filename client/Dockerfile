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

# Set default Node environment to production
ENV NODE_ENV=production

# Copy package configuration
COPY package.json package-lock.json ./

# Install only production dependencies to keep the image slim
RUN npm ci --omit=dev

# Copy built app files from the builder stage
COPY --from=builder /app/build ./build

# Expose SvelteKit default runtime port
EXPOSE 3000

# Set default runtime environment variables
ENV PORT=3000

# Start SvelteKit production server
CMD ["node", "build/index.js"]
