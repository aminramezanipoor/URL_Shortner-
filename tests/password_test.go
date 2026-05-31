package tests

import (
	"testing"

	"golang.org/x/crypto/bcrypt"
)

func TestPasswordHash(t *testing.T) {

	password := "123456"

	hashedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(password),
		bcrypt.DefaultCost,
	)

	if err != nil {
		t.Fatal(err)
	}

	err = bcrypt.CompareHashAndPassword(
		hashedPassword,
		[]byte(password),
	)

	if err != nil {
		t.Fatal("password verification failed")
	}
}