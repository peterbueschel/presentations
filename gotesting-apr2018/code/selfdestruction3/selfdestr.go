package selfdestruction

import (
	"errors"
	"fmt"
)

var (
	ErrWrongInitCode = errors.New("incorrect self-destruction code")
	ErrLessOfficers  = errors.New("two officers are required")
	counter          = 10
)

type (
	Klingon  struct{}
	Human    struct{}
	Starship struct {
		Humans []Human
	}
)

func (k Klingon) InitSelfDestruction(counter uint8, lastWords string) error {
	fmt.Println(lastWords)
	fmt.Println("booooooom")
	return nil
}

func (h Human) InitSelfDestruction(initCode string) error {
	if len(initCode) > 0 {
		return nil
	}
	return ErrWrongInitCode
}

func (s Starship) SelfDestruction() error {
	if len(s.Humans) < 2 {
		return ErrLessOfficers
	}
	for n, h := range s.Humans {
		initCode := fmt.Sprintf("Destruct sequence %d", n)
		if err := h.InitSelfDestruction(initCode); err != nil {
			return err
		}
	}
	for i := counter; i > 0; i-- {
		fmt.Printf("self destruction in = %d s\n", i)
	}
	fmt.Println("booooooom")
	return nil
}
