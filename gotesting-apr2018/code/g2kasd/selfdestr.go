package selfdestruction

import (
	"errors"
	"fmt"

	"bitbucket.ops.expertcity.com/oa/go-test-presentation/code/oldcounting"
)

var (
	ErrWrongInitCode = errors.New("incorrect self-destruction code")

	klingonWarning = ""
)

type (
	Klingon struct{}
	Human   struct {
		Captian bool
	}
	Starship struct { //Rearange
		Humans   []Human
		Klingons []Klingon
	}
)

func (h Human) InitSelfDestruction(counter uint8, initCode string, ship Starship) ([]string, error) {
	output := []string{}
	if initCode == "Destruct sequence 0, code 1-1 A" {
		if len(ship.Klingons) > 0 {
			output = append(output, klingonWarning)
		}
		for i := counter; i > 0; i-- {
			if kn, err := translator.ToOldKlingonNumber(i); err == nil {
				output = append(output, fmt.Sprintf("%s", kn))
				continue
			}
			output = append(output, fmt.Sprintf("self destruction in = %d s", i))

		}
		return output, nil
	}
	return output, ErrWrongInitCode
}
