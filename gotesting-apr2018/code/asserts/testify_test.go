package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func verifyTestify(t *testing.T) {
	fmt.Printf("%+v\n", assert.Equal(t, 123, "123") == assert.Equal(t, 123, 123))
}

func main() {
	verifyTestify(&testing.T{})
}
