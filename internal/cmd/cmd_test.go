package cmd

import "testing"

func TestCmd(t *testing.T) {
	scenarios := []struct {
		name     string
		input    string
		expected *Cmd
	}{
		{"standard", "echo Hello World", &Cmd{Name: "echo", Args: []string{"Hello", "World"}}},
	}

	for _, scenario := range scenarios {
		fn := func(t *testing.T) {
			expected, err := NewCmd(scenario.input)
			if err != nil {
				t.Error()
			}
			if expected.Name != scenario.expected.Name {
				t.Errorf("The expected name(%q) doesn't match with what it was got(%q)", expected.Name, scenario.expected.Name)
			}
		}
		t.Run(scenario.name, fn)
	}
}

func BenchmarkCmd(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NewCmd("echo Hello World")
	}
}
