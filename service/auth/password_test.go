package auth

import (
	"testing"
)

var passwordToHash = "password"

func TestHashPassword(t *testing.T) {
	hash, err := HashPassword(passwordToHash)
	if err != nil {
		t.Errorf("error hashing password %v", err)
	}

	if hash == "" {
		t.Errorf("expected hash to be not empty")
	}

	if hash == passwordToHash {
		t.Errorf("expected hash to be different from password")
	}
}

func TestComparePassword(t *testing.T) {
	hash, err := HashPassword(passwordToHash)
	if err != nil {
		t.Errorf("error hashing password %v", err)
	}

	if !ComparePasswords(hash, []byte(passwordToHash)) {
		t.Errorf("expected password to match hash")
	}

	if ComparePasswords(hash, []byte("fake_password")) {
		t.Errorf("expected password to not match hash")
	}
}
