package selfdestruction

import (
	"testing"
)

func ExampleKlingon_InitSelfDestruction() {
	kruge := Klingon{}
	kruge.InitSelfDestruction("wo’ batlhvaD")
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
	err := kruge.InitSelfDestruction("wo’ batlhvaD")
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
