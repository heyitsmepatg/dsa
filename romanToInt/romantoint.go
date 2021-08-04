package romantoint

func romanToInt(s string) int {
	sum := 0
	for i := len(s) - 1; i >= 0; i-- {
		switch numeral := string(s[i]); numeral {
		case "I":
			sum++
		case "V":
			if i > 0 && (s[i-1]) == byte('I') {
				sum += 4
				i--
			} else {
				sum += 5
			}
		case "X":
			if i > 0 && (s[i-1]) == byte('I') {
				sum += 9
				i--
			} else {
				sum += 10
			}
		case "L":
			if i > 0 && (s[i-1]) == byte('X') {
				sum += 40
				i--
			} else {
				sum += 50
			}
		case "C":
			if i > 0 && (s[i-1]) == byte('X') {
				sum += 90
				i--
			} else {
				sum += 100
			}
		case "D":
			if i > 0 && (s[i-1]) == byte('C') {
				sum += 400
				i--
			} else {
				sum += 500
			}
		case "M":
			if i > 0 && (s[i-1]) == byte('C') {
				sum += 900
				i--
			} else {
				sum += 1000
			}
		}
	}
	return sum
}
