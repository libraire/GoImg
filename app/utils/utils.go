package utils

import (
	"fmt"
	"os"
	"strings"
)

func cleanEmailAddress(email string) string {
	email = strings.ReplaceAll(email, "@", "")
	email = strings.ReplaceAll(email, ".", "")
	email = strings.ToLower(email)
	return email
}

func GenerateImageSourcePath(email string, position uint) string {
	return fmt.Sprintf("/images/%s/%d", cleanEmailAddress(email), position)
}

func MakeUserImageFolder(email string) error {
	folder := fmt.Sprintf("/images/%s", cleanEmailAddress(email))
	return os.MkdirAll(folder, os.ModePerm)
}
