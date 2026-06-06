package service

import (
	"fmt"
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

func Parse(s string) (string, error) {
	if len(s) == 0 {
		return "", fmt.Errorf("no data")
	}

	if strings.HasPrefix(s, ".") || strings.HasPrefix(s, "-") {
		return morse.ToText(s), nil
	}
	return morse.ToMorse(s), nil
}
