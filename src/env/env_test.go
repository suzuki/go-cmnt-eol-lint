package env

import (
	"context"
	"os"
	"strconv"
	"testing"
)

func TestConfig(t *testing.T) {
	var (
		allowedEOLs = "ABC"
		minWords    = "10"
	)

	os.Setenv("ALLOWED_EOLS", allowedEOLs)
	os.Setenv("MIN_WORDS", minWords)

	ctx := context.Background()
	ctx = WithConfig(ctx)

	config := GetConfig(ctx)

	if config.AllowedEOLs != allowedEOLs {
		t.Errorf("AllowedEOLs is not match. want=%s got=%s", allowedEOLs, config.AllowedEOLs)
	}
	if strconv.Itoa(config.MinWords) != minWords {
		t.Errorf("MinWords is not match. want=%s got=%d", minWords, config.MinWords)
	}
}
