package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestSub(t *testing.T){
	t.Run("f", func (t *testing.T){
		result := Hello("fah")
		if result != "hello fah" {
			t.Fatal("must hello")
		}
		fmt.Println("failnow")
	})
}

func TestMain (t *testing.M){
	fmt.Println("before")
	t.Run()
	fmt.Println("after")
}
func TestHello(t *testing.T) {
	result := Hello("fah")
	if result != "hello fah" {
		t.Error("must hello")
	}

	fmt.Println("fail")
}
func TestHello2(t *testing.T) {
	result := Hello("fah")
	if result != "hello fah" {
		t.Fatal("must hello")
	}
	fmt.Println("failnow")
}
func TestHelloAssert(t *testing.T) {
	assert.Equal(t, "hello fah", Hello("fah"), "must hello")
	fmt.Println("fail")
}
func TestHelloRequire(t *testing.T) {
	require.Equal(t, "hello fah", Hello("fah"), "must hello")
	fmt.Println("failnow")
}

func TestSkip (t *testing.T){
	if runtime.GOOS == "windows"{
		t.Skip("windows")
	}
	fmt.Println("skip")

}