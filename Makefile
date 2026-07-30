.PHONY: run build test tidy clean db-up db-down db-reset prod-up prod-down

# Jalankan server dalam mode development
run:
	go run ./cmd/api/

# Build binary
build:
	go build -o bin/server ./cmd/api/

# Jalankan semua tests
test:
	go test -v ./...

# Jalankan tests dengan coverage
test-cover:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html

# Tidy dependencies
tidy:
	go mod tidy

# Bersihkan build artifacts
clean:
	rm -rf bin/
	rm -f coverage.out coverage.html

# Go vet
vet:
	go vet ./...

# ============================================
# Docker — Development (DB only)
# ============================================

# Start PostgreSQL container
db-up:
	docker compose up -d

# Stop PostgreSQL container
db-down:
	docker compose down

# Reset database (hapus volume, mulai fresh)
db-reset:
	docker compose down -v
	docker compose up -d

# ============================================
# Docker — Production (App + DB)
# ============================================

# Build & start semua services
prod-up:
	docker compose -f docker-compose.prod.yml up -d --build

# Stop semua services
prod-down:
	docker compose -f docker-compose.prod.yml down
