# ============================================
# Stage 1: Build
# ============================================
FROM golang:1.26-alpine AS builder

WORKDIR /app

# Copy dependency files dulu (untuk cache layer)
COPY go.mod go.sum ./
RUN go mod download

# Copy seluruh source code
COPY . .

# Build binary
RUN CGO_ENABLED=0 GOOS=linux go build -o /app/server ./cmd/api

# ============================================
# Stage 2: Run (image kecil, tanpa Go toolchain)
# ============================================
FROM alpine:3.21

RUN apk --no-cache add ca-certificates tzdata

WORKDIR /app

# Copy binary dari builder stage
COPY --from=builder /app/server .

# Buat direktori uploads (untuk menyimpan file yang di-upload)
RUN mkdir -p /app/uploads/projects /app/uploads/blogs /app/uploads/experiences

EXPOSE 8080

CMD ["./server"]
