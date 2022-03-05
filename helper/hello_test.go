package helper

import (
	"fmt"
	"runtime"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func BenchmarkTable(b *testing.B) {
	names := []struct{
		Name string
	}{
		{
			"fahmi",
		},
		{
			"aziz",
		},
		{
			"ega",
		},
		{
			Name: "soni",
		},
	}

	for _, data := range names {
		b.Run(data.Name, func(b *testing.B) {
			for i:=0; i < b.N; i++ {
				Hello(data.Name)
			}
		})
	}
}

func BenchmarkSub(b *testing.B) {
	b.Run("f", func(b *testing.B) {
		for i:= 0; i < b.N; i++ {
			Hello("fahmi")
		}
	})

	b.Run("a", func(b *testing.B) {
		for i:=0;i < b.N; i++{
			Hello("aziz")
		}
	})
}

func BenchmarkHelloWorld(b *testing.B) {
	for i:= 0; i < b.N; i++{
		Hello("fahmi")
	}
}

func TestTable (t *testing.T){
	expected:= "hell "
	names:= []string{
		"fahmi",
		"soni",
		"ega",
	}

	type StructTest struct{
		name string
		request string
		expected string
	}

	var tests []StructTest
 
	for _, name:= range names{
		tests = append(tests, StructTest{
			name: name,
			request: name,
			expected: expected + name,
		})
	}

	for _, data:= range tests {
		t.Run(data.name, func(t *testing.T) {
			result:= Hello(data.request)
			ex:= &data.expected
			assert.Equal(t, *ex, result)
		})
	}
}
func TestSub(t *testing.T){
	t.Run("f", func (t *testing.T){
		result := Hello("fah")
		if result != "hello fah" {
			t.Fatal("must hello")
		}
		fmt.Println("failnow")
	})
	t.Run("j", func (t *testing.T){
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