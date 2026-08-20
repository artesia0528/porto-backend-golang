package utils

import (
	"errors"
	"fmt"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

const (
	MaxFileSize    = 5 << 20 // 5MB
	uploadsBaseDir = "uploads"
)

var allowedImageTypes = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
	".gif":  true,
	".webp": true,
}

// SaveUploadedFile menyimpan file dari form upload ke disk.
// Mengembalikan path relatif file (misal: "/uploads/projects/abc.png").
func SaveUploadedFile(file *multipart.FileHeader, subDir string) (string, error) {
	if file.Size > MaxFileSize {
		return "", errors.New("ukuran file melebihi 5MB")
	}

	ext := strings.ToLower(filepath.Ext(file.Filename))
	if !allowedImageTypes[ext] {
		return "", fmt.Errorf("tipe file %s tidak diizinkan, gunakan: jpg, jpeg, png, gif, webp", ext)
	}

	dir := filepath.Join(uploadsBaseDir, subDir)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return "", errors.New("gagal membuat direktori upload")
	}

	filename := uuid.New().String() + ext
	savePath := filepath.Join(dir, filename)

	src, err := file.Open()
	if err != nil {
		return "", errors.New("gagal membuka file upload")
	}
	defer src.Close()

	dst, err := os.Create(savePath)
	if err != nil {
		return "", errors.New("gagal menyimpan file")
	}
	defer dst.Close()

	if _, err := io.Copy(dst, src); err != nil {
		return "", errors.New("gagal menulis file ke disk")
	}

	// Path relatif dengan forward slash untuk URL
	relPath := "/" + filepath.ToSlash(filepath.Join(uploadsBaseDir, subDir, filename))
	return relPath, nil
}

// DeleteFile menghapus file dari disk berdasarkan path relatif.
// Aman dipanggil meskipun file tidak ada.
func DeleteFile(relPath string) {
	if relPath == "" {
		return
	}
	// Hapus leading slash agar filepath.Join bisa bekerja
	cleanPath := strings.TrimPrefix(relPath, "/")
	os.Remove(cleanPath)
}

// InitUploadDirs membuat semua subdirektori upload jika belum ada.
func InitUploadDirs(subDirs []string) error {
	for _, sub := range subDirs {
		dir := filepath.Join(uploadsBaseDir, sub)
		if err := os.MkdirAll(dir, 0755); err != nil {
			return fmt.Errorf("gagal membuat direktori %s: %w", dir, err)
		}
	}
	return nil
}
