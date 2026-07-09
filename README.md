# Portfolio Backend (Go)

REST API backend untuk website portfolio, dibangun dengan Go.

## Tech Stack

- **Framework**: [Gin](https://github.com/gin-gonic/gin)
- **ORM**: [GORM](https://gorm.io/)
- **Database**: SQLite
- **Auth**: JWT (HS256)

## Arsitektur

```
cmd/api/          → Entry point (main.go)
internal/
├── config/       → Environment configuration
├── database/     → Database connection & seeder
├── dto/          → Request/Response DTOs dengan validasi
├── models/       → GORM models & API response wrapper
├── repositories/ → Database access layer
├── services/     → Business logic layer
├── handlers/     → HTTP handler layer
├── middleware/    → Auth & CORS middleware
├── routes/       → Route registration
└── utils/        → JWT & password utilities
```

## Setup

### Prerequisites

- Go 1.21+
- GCC (untuk SQLite driver)

### Installation

```bash
# Clone repository
git clone <repo-url>
cd porto-backend-golang

# Copy environment config
cp .env.example .env
# Edit .env dan isi JWT_SECRET dengan string acak yang panjang

# Download dependencies
go mod download

# Jalankan server (akan auto-seed admin user)
make run
```

### Default Admin Credentials

Saat pertama kali dijalankan, seeder akan membuat admin user:
- **Username**: `admin`
- **Password**: `admin123`

> ⚠️ Ganti password default ini segera setelah pertama kali login!

## API Endpoints

### Public

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/api/login` | Login, dapatkan JWT token |
| `POST` | `/api/register` | Register user baru |
| `GET` | `/api/projects` | Lihat semua project |

### Admin (butuh JWT)

| Method | Endpoint | Description |
|--------|----------|-------------|
| `POST` | `/api/admin/projects` | Tambah project baru |
| `PUT` | `/api/admin/projects/:id` | Update project |
| `DELETE` | `/api/admin/projects/:id` | Hapus project |

### Authentication

Kirim JWT token di header:
```
Authorization: Bearer <token>
```

## Development

```bash
# Jalankan server
make run

# Build binary
make build

# Jalankan tests
make test

# Tidy dependencies
make tidy
```

## License

MIT
