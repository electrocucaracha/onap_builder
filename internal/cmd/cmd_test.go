package cmd

import (
	"reflect"
	"testing"
)

func TestNewCmd(t *testing.T) {
	fn := func(t *testing.T) {
		input := "echo Hello World"
		expected := Cmd{
			Name: "echo",
			Args: []string{"Hello", "World"},
		}

		result, err := NewCmd(input)
		if err != nil {
			t.Fatalf("NewCmd returned an error (%s)", err)
		}
		if expected.Name != result.Name {
			t.Fatalf("NewCmd returned name %s, expected %s", result.Name, expected.Name)
		}
		if !reflect.DeepEqual(expected.Args, result.Args) {
			t.Fatalf("NewCmd returned args %v, expected %v", result.Args, expected.Args)
		}
	}
	t.Run("Standard", fn)

	fn = func(t *testing.T) {
		_, err := NewCmd("")
		if err == nil {
			t.Fatalf("NewCmd didn't returned an error")
		}
	}
	t.Run("Empty", fn)
}

func TestWithArg(t *testing.T) {
	input := "echo Hello"
	result, err := NewCmd(input)
	if err != nil {
		t.Fatalf("NewCmd returned an error (%s)", err)
	}

	result.WithArg("")
	expected := []string{"Hello"}
	if !reflect.DeepEqual(expected, result.Args) {
		t.Fatalf("WithArg returned args %v, expected %v", result.Args, expected)
	}

	expected = []string{"Hello", "World"}
	result.WithArg("World")
	if !reflect.DeepEqual(expected, result.Args) {
		t.Fatalf("WithArg returned args %v, expected %v", result.Args, expected)
	}
}

func TestWithArgs(t *testing.T) {
	input := "echo"
	result, err := NewCmd(input)
	if err != nil {
		t.Fatalf("NewCmd returned an error (%s)", err)
	}

	expected := []string{"Hello", "World"}
	result.WithArgs("Hello", "World")
	if !reflect.DeepEqual(expected, result.Args) {
		t.Fatalf("WithArg returned args %v, expected %v", result.Args, expected)
	}
}

func BenchmarkNewCmd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewCmd("echo Hello World")
	}
}
