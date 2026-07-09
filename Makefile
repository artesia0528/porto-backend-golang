.PHONY: run build test tidy clean

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
