package models

import "errors"

var (
	ErrNoRecord = errors.New("model: no matching record found")

	ErrInvalidCredentials = errors.New("models: invalid credentials")
	// Add a new ErrDuplicateEmail error. We'll use this later if a user
	// tries to signup with an email address that's already in use.
	ErrDuplicateEmail = errors.New("models: duplicate email")
)
