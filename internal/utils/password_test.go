package utils

import (
	"testing"
)

func TestHashAndCheckPassword(t *testing.T) {
	password := "mySecurePassword123"

	hashed, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword gagal: %v", err)
	}

	if hashed == "" {
		t.Fatal("Hashed password tidak boleh kosong")
	}

	if hashed == password {
		t.Fatal("Hashed password tidak boleh sama dengan plain text")
	}

	// Password yang benar harus match
	if !CheckPassword(hashed, password) {
		t.Error("CheckPassword seharusnya return true untuk password yang benar")
	}

	// Password yang salah harus tidak match
	if CheckPassword(hashed, "wrongPassword") {
		t.Error("CheckPassword seharusnya return false untuk password yang salah")
	}
}

func TestHashPassword_DifferentHashesForSamePassword(t *testing.T) {
	password := "samePassword"

	hash1, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword pertama gagal: %v", err)
	}

	hash2, err := HashPassword(password)
	if err != nil {
		t.Fatalf("HashPassword kedua gagal: %v", err)
	}

	// bcrypt harus menghasilkan hash yang berbeda setiap kali (karena random salt)
	if hash1 == hash2 {
		t.Error("Dua hash dari password yang sama seharusnya berbeda (random salt)")
	}

	// Tapi keduanya harus valid
	if !CheckPassword(hash1, password) {
		t.Error("Hash pertama seharusnya valid")
	}
	if !CheckPassword(hash2, password) {
		t.Error("Hash kedua seharusnya valid")
	}
}
