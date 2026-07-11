package downloader

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSanitizeFilename(t *testing.T) {
	fileName := "a<b>c:d\\e\"f/g\\h|i?j*k"
	sanitized := SanitizeFilename(fileName)
	if sanitized != "abcdefghijk" {
		t.Error("Invalid characters must get stripped")
	}

	fileName = "aB Cd"
	sanitized = SanitizeFilename(fileName)
	if sanitized != "aB Cd" {
		t.Error("Casing and whitespaces must be preserved")
	}

	fileName = "~!@#$%^&()[].,"
	sanitized = SanitizeFilename(fileName)
	if sanitized != "~!@#$%^&()[].," {
		t.Error("The common harmless symbols should remain valid")
	}
}

func TestSanitizeDirname(t *testing.T) {
	inputDirNames := []string{"Kurzgesagt – In a Nutshell", "foo/bar\\baz", "  hello world  ", ""}
	expectedDirNames := []string{"Kurzgesagt – In a Nutshell", "foobarbaz", "hello world", "Unknown"}

	for i := range inputDirNames {
		sanitizedDirName := SanitizeDirname(inputDirNames[i])
		assert.Equal(t, expectedDirNames[i], sanitizedDirName)
	}
}
