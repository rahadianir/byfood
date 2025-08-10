# ===== Stage 1: Build the Next.js app =====
FROM node:20-alpine AS builder

# Set working directory
WORKDIR /app

# Install dependencies first (better caching)
COPY package.json package-lock.json* ./
RUN npm ci

# Copy the rest of the source code
COPY . .

# Build Next.js for production
RUN npm run build

# ===== Stage 2: Production runtime =====
FROM node:20-alpine AS runner

WORKDIR /app

# Set NODE_ENV to production
ENV NODE_ENV=production

# Copy only necessary files from builder
COPY --from=builder /app/public ./public
COPY --from=builder /app/.next ./.next
COPY --from=builder /app/node_modules ./node_modules
COPY --from=builder /app/package.json ./package.json

# Next.js needs this for standalone mode
# COPY --from=builder /app/next.config.js ./next.config.js

EXPOSE 3000

# Run Next.js
CMD ["npm", "start"]
