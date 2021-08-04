package romantoint

import "testing"

func TestRomanToInt_Example1(t *testing.T) {
	s := "III"

	num := romanToInt(s)

	if num != 3 {
		t.Errorf("Expected: 3. Actual: %v", num)
	}
}

func TestRomanToInt_Example2(t *testing.T) {
	s := "IV"

	num := romanToInt(s)

	if num != 4 {
		t.Errorf("Expected: 4. Actual: %v", num)
	}
}

func TestRomanToInt_Example3(t *testing.T) {
	s := "IX"

	num := romanToInt(s)

	if num != 9 {
		t.Errorf("Expected: 9. Actual: %v", num)
	}
}

func TestRomanToInt_Example4(t *testing.T) {
	s := "LVIII"

	num := romanToInt(s)

	if num != 58 {
		t.Errorf("Expected: 58. Actual: %v", num)
	}
}

func TestRomanToInt_Example5(t *testing.T) {
	s := "MCMXCIV"

	num := romanToInt(s)

	if num != 1994 {
		t.Errorf("Expected: 58. Actual: %v", num)
	}
}
