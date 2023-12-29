package utils

import "golang.org/x/crypto/bcrypt"

// HashPassword generates a hashed version of the provided password using bcrypt algorithm.
// It takes a string password as input and returns the hashed password as a string along with any error encountered during the process.
func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	return string(bytes), err
}

// CheckPasswordHash compares a password with its corresponding hash and returns true if they match.
// It uses bcrypt.CompareHashAndPassword to perform the comparison.
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}
