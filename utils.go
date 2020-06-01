package main

import (
	"github.com/google/uuid"
)

//generateUUID генерирует UUID (128 bits)
func generateUUID() string {
	return uuid.New().String()
}
