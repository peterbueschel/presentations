package translator

import (
	"errors"
)

var ErrNaKN = errors.New("input must be greater than 0 and less than 10")

func ToOldKlingonNumber(input uint8) (string, error) {
	if input > 9 || input < 1 {
		return "NaKN", ErrNaKN
	}
	return "", nil
}

//
