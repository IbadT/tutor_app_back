package shared

// PasswordHasher defines password hashing operations
type PasswordHasher interface {
	HashPassword(password string) (string, error)
	ComparePassword(password, hash string) bool
}
