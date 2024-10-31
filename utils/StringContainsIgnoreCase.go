package utils

import (
	"strings"
)

func StringContainsIgnoreCase(str, substr string) bool {

	return strings.Contains(strings.ToLower(str), strings.ToLower(substr))
}
