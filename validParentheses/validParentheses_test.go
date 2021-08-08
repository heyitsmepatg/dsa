package validparentheses

import "testing"

func TestValidParentheses_Example1(t *testing.T) {
	input := "()"
	output := isValid(input)
	if !output {
		t.Errorf("Expected: true. Actual: %v", output)
	}
}

func TestValidParentheses_Example2(t *testing.T) {
	input := "()[]{}"
	output := isValid(input)
	if !output {
		t.Errorf("Expected: true. Actual: %v", output)
	}
}

func TestValidParentheses_Example3(t *testing.T) {
	input := "(]"
	output := isValid(input)
	if output {
		t.Errorf("Expected: false. Actual: %v", output)
	}
}

func TestValidParentheses_Example4(t *testing.T) {
	input := "([)]"
	output := isValid(input)
	if output {
		t.Errorf("Expected: false. Actual: %v", output)
	}
}

func TestValidParentheses_Example5(t *testing.T) {
	input := "{[]}"
	output := isValid(input)
	if !output {
		t.Errorf("Expected: true. Actual: %v", output)
	}
}

func TestValidParentheses_Example6(t *testing.T) {
	input := "{{}[][[[]]]}"
	output := isValid(input)
	if !output {
		t.Errorf("Expected: true. Actual: %v", output)
	}
}

func TestValidParentheses_Example7(t *testing.T) {
	input := "{"
	output := isValid(input)
	if output {
		t.Errorf("Expected: false. Actual: %v", output)
	}
}

func TestValidParentheses_Example8(t *testing.T) {
	input := "}{"
	output := isValid(input)
	if output {
		t.Errorf("Expected: false. Actual: %v", output)
	}
}

func TestValidParentheses_Example9(t *testing.T) {
	input := "(("
	output := isValid(input)
	if output {
		t.Errorf("Expected: false. Actual: %v", output)
	}
}
func TestValidParentheses_Example10(t *testing.T) {
	input := "()(("
	output := isValid(input)
	if output {
		t.Errorf("Expected: false. Actual: %v", output)
	}
}

func TestValidParentheses_Example11(t *testing.T) {
	input := "(){}}{"
	output := isValid(input)
	if output {
		t.Errorf("Expected: false. Actual: %v", output)
	}
}
