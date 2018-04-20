package selfdestruction

import (
	"errors"
	"fmt"
)

var (
	ErrWrongInitCode       = errors.New("incorrect self-destruction code")
	counter          uint8 = 10
)

type (
	Klingon struct{}
	Human   struct{}
)

func (k Klingon) InitSelfDestruction(lastWords string) error {
	fmt.Println(lastWords)
	fmt.Println("booooooom")
	return nil
}

func (h Human) InitSelfDestruction(initCode string) error {
	if initCode == "Destruct sequence 1, code 1-1 A" {
		for i := counter; i > 0; i-- {
			fmt.Printf("self destruction in = %d s\n", i)
		}
		fmt.Println("booooooom")
		return nil
	}
	return ErrWrongInitCode
}
