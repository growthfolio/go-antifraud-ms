package utils

import (
	"errors"
	"strconv"
	"strings"
)

// HourToMinutes converte a hora para minutos desde a meia-noite.
func HourToMinutes(hour string) (int, error) {
	parts := strings.Split(hour, ":")
	if len(parts) != 2 {
		return 0, errors.New("formato de hora inválido")
	}

	h, err := strconv.Atoi(parts[0])
	if err != nil {
		return 0, err
	}

	m, err := strconv.Atoi(parts[1])
	if err != nil {
		return 0, err
	}

	return h*60 + m, nil
}
