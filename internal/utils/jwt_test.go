package utils

import (
	"testing"
)

func TestGenerateAndValidateToken(t *testing.T) {
	secret := "test-secret-key-for-unit-tests"
	userID := "550e8400-e29b-41d4-a716-446655440000"

	// Generate token
	tokenString, err := GenerateToken(userID, secret)
	if err != nil {
		t.Fatalf("GenerateToken gagal: %v", err)
	}

	if tokenString == "" {
		t.Fatal("Token tidak boleh kosong")
	}

	// Validate token
	token, err := ValidateToken(tokenString, secret)
	if err != nil {
		t.Fatalf("ValidateToken gagal: %v", err)
	}

	if !token.Valid {
		t.Fatal("Token seharusnya valid")
	}
}

func TestValidateToken_InvalidSecret(t *testing.T) {
	secret := "correct-secret"
	wrongSecret := "wrong-secret"
	userID := "user-123"

	tokenString, err := GenerateToken(userID, secret)
	if err != nil {
		t.Fatalf("GenerateToken gagal: %v", err)
	}

	// Validasi dengan secret yang salah harus gagal
	_, err = ValidateToken(tokenString, wrongSecret)
	if err == nil {
		t.Fatal("ValidateToken seharusnya gagal dengan secret yang salah")
	}
}

func TestValidateToken_InvalidTokenString(t *testing.T) {
	secret := "test-secret"

	// Token string yang tidak valid
	_, err := ValidateToken("ini.bukan.token.valid", secret)
	if err == nil {
		t.Fatal("ValidateToken seharusnya gagal dengan token string yang tidak valid")
	}
}

func TestValidateToken_EmptyToken(t *testing.T) {
	secret := "test-secret"

	_, err := ValidateToken("", secret)
	if err == nil {
		t.Fatal("ValidateToken seharusnya gagal dengan token kosong")
	}
}
