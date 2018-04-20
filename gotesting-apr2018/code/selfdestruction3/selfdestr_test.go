package selfdestruction

import (
	"fmt"
	"testing"

	"github.com/DATA-DOG/godog"
)

func ExampleKlingon_InitSelfDestruction() {
	kruge := Klingon{}
	kruge.InitSelfDestruction(10, "wo’ batlhvaD")
	// Output:
	// wo’ batlhvaD
	// booooooom
}

func ExampleHuman_InitSelfDestruction() {
	kirk := Human{}
	kirk.InitSelfDestruction("Destruct sequence 1, code 1-1 A")
	// Output:
	// self destruction in = 10 s
	// self destruction in = 9 s
	// self destruction in = 8 s
	// self destruction in = 7 s
	// self destruction in = 6 s
	// self destruction in = 5 s
	// self destruction in = 4 s
	// self destruction in = 3 s
	// self destruction in = 2 s
	// self destruction in = 1 s
	// booooooom
}

func TestKlingon_InitSelfDestruction(t *testing.T) {
	wantErr := false
	kruge := Klingon{}
	err := kruge.InitSelfDestruction(10, "wo’ batlhvaD")
	if (err != nil) != wantErr {
		t.Errorf("Klingon InitSelfDestructionIn() got: %s, expected: %v\n", err, wantErr)
	}
}

func TestHuman_InitSelfDestruction(t *testing.T) {
	wantErr := false
	kirk := Human{}
	err := kirk.InitSelfDestruction("Destruct sequence 1, code 1-1 A")
	if (err != nil) != wantErr {
		t.Errorf("Human InitSelfDestructionIn() got: %s, expected: %v\n", err, wantErr)
	}
}

func (s *Starship) thereAreHumanOfficers(number int) error {
	h := []Human{}
	for i := 0; i < number; i++ {
		h = append(h, Human{})
	}
	s.Humans = h
	return nil
}

func (s *Starship) officersEnterTheCorrectDestructSequence() error {
	err := s.SelfDestruction()
	if err == ErrWrongInitCode {
		return err
	}
	return nil
}

func (s Starship) itIsForbiddenToRunSelfdestruction() error {
	err := s.SelfDestruction()
	if len(s.Humans) < 2 && err != ErrLessOfficers {
		return fmt.Errorf("expected %s, but got %s", ErrLessOfficers, err)
	}
	return nil
}

func (s Starship) itIsAllowedToRunSelfdestruction() error {
	err := s.SelfDestruction()
	if len(s.Humans) > 1 && err != nil {
		return fmt.Errorf("expected no error, but got %s", err)
	}
	return nil
}

func FeatureContext(s *godog.Suite) {
	enterprise := &Starship{}
	s.Step(`^there are (\d+) human officers$`, enterprise.thereAreHumanOfficers)
	s.Step(`^(\d+) officers enter the destruct sequence$`, enterprise.officersEnterTheCorrectDestructSequence)
	s.Step(`^it is forbidden to run self-destruction$`, enterprise.itIsForbiddenToRunSelfdestruction)
	s.Step(`^it is allowed to run self-destruction$`, enterprise.itIsAllowedToRunSelfdestruction)

}
