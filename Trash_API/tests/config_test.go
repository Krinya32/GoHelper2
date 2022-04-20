package tests

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

//func TestPort(t *testing.T) {
//	got := Abs(-1)
//	if got != 1 {
//		t.Errorf("Abs(-1) = %d; want 1", got)
//	}
//}

func TestDebug(t *testing.T) {
	assert.True(t, true)
	assert.False(t, false)
	assert.Equal(t, '1', '1')
}
func ExampleHello() {
	fmt.Println("hello")
	// Output: hello
}

func ExampleSalutations() {
	fmt.Println("hello, and")
	fmt.Println("goodbye")
	// Output:
	// hello, and
	// goodbye
}
