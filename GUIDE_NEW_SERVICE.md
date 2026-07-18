# 📘 Panduan: Menambahkan Service Baru

Panduan ini menjelaskan langkah-langkah lengkap untuk menambahkan fitur/service baru
ke dalam project ini. Kita akan menggunakan contoh fitur **"Blog"** sebagai ilustrasi.

---

## Urutan File yang Harus Dibuat

```
Langkah 1 → internal/models/blog.go           (Struktur tabel database)
Langkah 2 → internal/dto/blog_dto.go           (Validasi input dari client)
Langkah 3 → internal/repositories/blog_repository.go  (Query ke database)
Langkah 4 → internal/services/blog_service.go  (Logika bisnis)
Langkah 5 → internal/handlers/blog_handler.go  (Terima & kirim HTTP response)
Langkah 6 → Sambungkan semuanya (database.go, main.go, routes.go)
```

### Diagram Alur Data

```
Client Request
     │
     ▼
┌─────────────┐     ┌──────────────┐     ┌─────────────────┐     ┌──────────┐
│   Handler    │ ──▶ │   Service    │ ──▶ │   Repository    │ ──▶ │ Database │
│ (HTTP layer) │     │ (bisnis)     │     │ (query DB)      │     │ (SQLite) │
└─────────────┘     └──────────────┘     └─────────────────┘     └──────────┘
     │
     ▼
Client Response (JSON)
```

---

## Langkah 1: Buat Model (`internal/models/blog.go`)

Model mendefinisikan bentuk tabel di database.

```go
package models

import (
	"time"
	"gorm.io/gorm"
)

type Blog struct {
	ID        string         `json:"id" gorm:"primaryKey;type:varchar(36)"`
	Title     string         `json:"title" gorm:"not null;size:200"`
	Content   string         `json:"content" gorm:"type:text"`
	ImageURL  string         `json:"image_url" gorm:"size:500"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"-" gorm:"index"`
}
```

### Aturan Penting:
- `ID` selalu `string` dengan tag `gorm:"primaryKey;type:varchar(36)"` (untuk UUID).
- `json:"-"` artinya field tersebut **tidak akan dikirim** ke client (contoh: `DeletedAt`).
- Gunakan `gorm:"type:text"` untuk field yang isinya bisa panjang.
- Gunakan `gorm:"size:200"` untuk membatasi panjang maksimal karakter.

---

## Langkah 2: Buat DTO (`internal/dto/blog_dto.go`)

DTO (Data Transfer Object) mendefinisikan **data apa saja yang boleh dikirim oleh client**
dan otomatis memvalidasinya. Ini BUKAN model database, melainkan "penjaga pintu masuk".

```go
package dto

// Untuk membuat blog baru (POST)
type CreateBlogRequest struct {
	Title    string `json:"title" binding:"required,min=3,max=200"`
	Content  string `json:"content" binding:"required,min=10"`
	ImageURL string `json:"image_url" binding:"omitempty,url"`
}

// Untuk mengubah blog (PUT) — semua field opsional
type UpdateBlogRequest struct {
	Title    string `json:"title" binding:"omitempty,min=3,max=200"`
	Content  string `json:"content" binding:"omitempty,min=10"`
	ImageURL string `json:"image_url" binding:"omitempty,url"`
}
```

### Tag Validasi yang Sering Dipakai:
| Tag | Arti |
|-----|------|
| `binding:"required"` | Wajib diisi, kalau kosong langsung error |
| `binding:"omitempty"` | Boleh kosong, tapi kalau diisi harus sesuai aturan |
| `min=3` | Minimal 3 karakter |
| `max=200` | Maksimal 200 karakter |
| `email` | Harus format email yang valid |
| `url` | Harus format URL yang valid |

---

## Langkah 3: Buat Repository (`internal/repositories/blog_repository.go`)

Repository **hanya** bertugas menjalankan query ke database. 
Tidak boleh ada logika bisnis di sini.

```go
package repositories

import (
	"portfolio-backend/internal/models"
	"gorm.io/gorm"
)

type BlogRepository struct {
	db *gorm.DB
}

func NewBlogRepository(db *gorm.DB) *BlogRepository {
	return &BlogRepository{db: db}
}

func (r *BlogRepository) Create(blog *models.Blog) error {
	return r.db.Create(blog).Error
}

func (r *BlogRepository) FindAll() ([]models.Blog, error) {
	var blogs []models.Blog
	err := r.db.Order("created_at DESC").Find(&blogs).Error
	return blogs, err
}

func (r *BlogRepository) FindByID(id string) (*models.Blog, error) {
	var blog models.Blog
	result := r.db.Where("id = ?", id).First(&blog)
	if result.Error != nil {
		return nil, result.Error
	}
	return &blog, nil
}

func (r *BlogRepository) Update(blog *models.Blog) error {
	return r.db.Save(blog).Error
}

func (r *BlogRepository) Delete(id string) (int64, error) {
	result := r.db.Where("id = ?", id).Delete(&models.Blog{})
	return result.RowsAffected, result.Error
}
```

### Pola Umum Repository:
| Fungsi | Query GORM | Keterangan |
|--------|-----------|------------|
| `Create` | `db.Create(&model)` | Simpan data baru |
| `FindAll` | `db.Find(&models)` | Ambil semua data |
| `FindByID` | `db.Where("id = ?", id).First(&model)` | Ambil 1 data |
| `Update` | `db.Save(&model)` | Update data (seluruh field) |
| `Delete` | `db.Where("id = ?", id).Delete(&model)` | Hapus (soft delete) |

---

## Langkah 4: Buat Service (`internal/services/blog_service.go`)

Service berisi **logika bisnis**. Di sinilah tempat kamu menaruh aturan-aturan
seperti: "generate UUID", "cek apakah data valid", "kirim notifikasi", dll.

```go
package services

import (
	"errors"
	"portfolio-backend/internal/dto"
	"portfolio-backend/internal/models"
	"portfolio-backend/internal/repositories"

	"github.com/google/uuid"
)

type BlogService struct {
	blogRepo *repositories.BlogRepository
}

func NewBlogService(blogRepo *repositories.BlogRepository) *BlogService {
	return &BlogService{blogRepo: blogRepo}
}

func (s *BlogService) GetAll() ([]models.Blog, error) {
	return s.blogRepo.FindAll()
}

func (s *BlogService) GetByID(id string) (*models.Blog, error) {
	blog, err := s.blogRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("blog tidak ditemukan")
	}
	return blog, nil
}

func (s *BlogService) Create(req dto.CreateBlogRequest) (*models.Blog, error) {
	blog := &models.Blog{
		ID:       uuid.New().String(), // ← JANGAN LUPA: Generate UUID
		Title:    req.Title,
		Content:  req.Content,
		ImageURL: req.ImageURL,
	}

	if err := s.blogRepo.Create(blog); err != nil {
		return nil, errors.New("gagal membuat blog")
	}
	return blog, nil
}

func (s *BlogService) Update(id string, req dto.UpdateBlogRequest) (*models.Blog, error) {
	blog, err := s.blogRepo.FindByID(id)
	if err != nil {
		return nil, errors.New("blog tidak ditemukan")
	}

	// Update hanya field yang diisi (tidak kosong)
	if req.Title != "" {
		blog.Title = req.Title
	}
	if req.Content != "" {
		blog.Content = req.Content
	}
	if req.ImageURL != "" {
		blog.ImageURL = req.ImageURL
	}

	if err := s.blogRepo.Update(blog); err != nil {
		return nil, errors.New("gagal memperbarui blog")
	}
	return blog, nil
}

func (s *BlogService) Delete(id string) error {
	rowsAffected, err := s.blogRepo.Delete(id)
	if err != nil {
		return errors.New("gagal menghapus blog")
	}
	if rowsAffected == 0 {
		return errors.New("blog tidak ditemukan")
	}
	return nil
}
```

### Hal yang Harus Diingat di Service:
- ✅ Selalu generate UUID: `ID: uuid.New().String()`
- ✅ Kembalikan pesan error yang manusiawi (bukan error mentah dari DB)
- ✅ Cek `rowsAffected == 0` pada Delete untuk mendeteksi data yang tidak ada
- ❌ JANGAN import `gin` atau `http` di sini (itu tugas Handler)

---

## Langkah 5: Buat Handler (`internal/handlers/blog_handler.go`)

Handler **hanya** bertugas:
1. Membaca request HTTP (parsing JSON, ambil parameter URL)
2. Memanggil Service
3. Mengirim response JSON

```go
package handlers

import (
	"net/http"
	"portfolio-backend/internal/dto"
	"portfolio-backend/internal/models"
	"portfolio-backend/internal/services"

	"github.com/gin-gonic/gin"
)

type BlogHandler struct {
	blogService *services.BlogService
}

func NewBlogHandler(blogService *services.BlogService) *BlogHandler {
	return &BlogHandler{blogService: blogService}
}

// GET /api/blogs (publik)
func (h *BlogHandler) GetAll(c *gin.Context) {
	blogs, err := h.blogService.GetAll()
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, "Gagal mengambil data blog")
		return
	}
	models.SuccessResponse(c, http.StatusOK, "Berhasil", blogs)
}

// POST /api/admin/blogs (butuh login)
func (h *BlogHandler) Create(c *gin.Context) {
	var req dto.CreateBlogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Input tidak valid: "+err.Error())
		return
	}

	blog, err := h.blogService.Create(req)
	if err != nil {
		models.ErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	models.SuccessResponse(c, http.StatusCreated, "Blog berhasil dibuat", blog)
}

// PUT /api/admin/blogs/:id (butuh login)
func (h *BlogHandler) Update(c *gin.Context) {
	id := c.Param("id")

	var req dto.UpdateBlogRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		models.ErrorResponse(c, http.StatusBadRequest, "Input tidak valid: "+err.Error())
		return
	}

	blog, err := h.blogService.Update(id, req)
	if err != nil {
		models.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	models.SuccessResponse(c, http.StatusOK, "Blog berhasil diperbarui", blog)
}

// DELETE /api/admin/blogs/:id (butuh login)
func (h *BlogHandler) Delete(c *gin.Context) {
	id := c.Param("id")

	if err := h.blogService.Delete(id); err != nil {
		models.ErrorResponse(c, http.StatusNotFound, err.Error())
		return
	}
	models.SuccessResponse(c, http.StatusOK, "Blog berhasil dihapus", nil)
}
```

### Pola Response:
- Selalu gunakan `models.SuccessResponse(...)` dan `models.ErrorResponse(...)`.
- Gunakan HTTP status code yang tepat:
  - `200 OK` → berhasil ambil/update/hapus
  - `201 Created` → berhasil buat data baru
  - `400 Bad Request` → input dari client salah
  - `404 Not Found` → data tidak ditemukan
  - `500 Internal Server Error` → error di server

---

## Langkah 6: Sambungkan Semuanya

Ini langkah terakhir dan **paling sering terlupa**! Ada 3 file yang harus diubah:

### 6a. `internal/database/database.go` — Daftarkan Model

```diff
 func Connect(cfg *config.Config) (*gorm.DB, error) {
     // ...
-    if err := db.AutoMigrate(&models.User{}, &models.Project{}, &models.Message{}); err != nil {
+    if err := db.AutoMigrate(&models.User{}, &models.Project{}, &models.Message{}, &models.Blog{}); err != nil {
         return nil, fmt.Errorf("gagal migrasi database: %w", err)
     }
     // ...
 }
```

### 6b. `cmd/api/main.go` — Dependency Injection

Tambahkan 3 baris baru di bagian "wire semua layer":

```diff
 // 5. Dependency Injection: wire semua layer
 userRepo := repositories.NewUserRepository(db)
 projectRepo := repositories.NewProjectRepository(db)
 messageRepo := repositories.NewMessageRepository(db)
+blogRepo := repositories.NewBlogRepository(db)

 authService := services.NewAuthService(userRepo, cfg.JWTSecret)
 projectService := services.NewProjectService(projectRepo)
 messageService := services.NewMessageService(messageRepo)
+blogService := services.NewBlogService(blogRepo)

 authHandler := handlers.NewAuthHandler(authService)
 projectHandler := handlers.NewProjectHandler(projectService)
 messageHandler := handlers.NewMessageHandler(messageService)
+blogHandler := handlers.NewBlogHandler(blogService)

 // 6. Register routes
-routes.SetupRoutes(r, cfg.JWTSecret, authHandler, projectHandler, messageHandler)
+routes.SetupRoutes(r, cfg.JWTSecret, authHandler, projectHandler, messageHandler, blogHandler)
```

### 6c. `internal/routes/routes.go` — Daftarkan Route

```diff
-func SetupRoutes(r *gin.Engine, jwtSecret string, authHandler *handlers.AuthHandler, projectHandler *handlers.ProjectHandler, messageHandler *handlers.MessageHandler) {
+func SetupRoutes(r *gin.Engine, jwtSecret string, authHandler *handlers.AuthHandler, projectHandler *handlers.ProjectHandler, messageHandler *handlers.MessageHandler, blogHandler *handlers.BlogHandler) {
     api := r.Group("/api")
     {
         // Publik
         api.POST("/login", authHandler.Login)
         api.GET("/projects", projectHandler.GetProjects)
         api.POST("/contact", messageHandler.Create)
+        api.GET("/blogs", blogHandler.GetAll)   // ← publik: siapa saja bisa baca

         admin := api.Group("/admin")
         admin.Use(middleware.AuthRequired(jwtSecret))
         {
             // Projects
             admin.POST("/projects", projectHandler.CreateProject)
             // ...

+            // Blogs
+            admin.POST("/blogs", blogHandler.Create)
+            admin.PUT("/blogs/:id", blogHandler.Update)
+            admin.DELETE("/blogs/:id", blogHandler.Delete)
         }
     }
 }
```

---

## ✅ Checklist Sebelum Test

Setelah selesai, pastikan:

- [ ] Model sudah ditambahkan di `AutoMigrate` (`database.go`)
- [ ] Repository, Service, Handler sudah di-*wire* di `main.go`
- [ ] Handler baru sudah dikirim ke parameter `SetupRoutes`
- [ ] Route sudah didaftarkan (publik di luar `admin`, admin di dalam `admin`)
- [ ] Jalankan `go build ./cmd/api/` — harus tanpa error
- [ ] Hapus `portfolio.db` jika menambahkan model baru (agar tabel baru terbuat)
- [ ] Jalankan `make run` dan test dengan curl

---

## Ringkasan Singkat

```
1. models/xxx.go        → Bentuk tabel (ID string UUID)
2. dto/xxx_dto.go       → Validasi input (binding tags)
3. repositories/xxx.go  → Query database (GORM)
4. services/xxx.go      → Logika bisnis (uuid.New, error handling)
5. handlers/xxx.go      → HTTP layer (parse request, send response)
6. Sambungkan:
   → database.go        → AutoMigrate
   → main.go            → Repo → Service → Handler (DI)
   → routes.go          → Daftarkan endpoint
```
