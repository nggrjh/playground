package internal

import "fmt"

func text(count int, char string) string {
	if count < 1 {
		return ""
	}

	var s string
	for i := 0; i < count; i++ {
		s += char
	}
	return s
}

func main() {
	input := 80
	star := "*"

	for i := 1; i < input+1; i++ {
		t := text(input-i, " ") + text(i-1, star) + text(i, star)
		fmt.Println(t)
	}
}
