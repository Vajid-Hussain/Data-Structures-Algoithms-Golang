package main

import "fmt"

func main() {
	s := "{[]}"
	var stack = []rune{}

	remove := func(val rune) bool {
		if len(stack) == 0 {
			return false
		}

		switch {
		case val == '}':
			if stack[len(stack)-1] != '{' {
				return false
			}
		case val == ']':
			if stack[len(stack)-1] != '[' {
				return false
			}
		case val == ')':
			if stack[len(stack)-1] != '(' {
				return false
			}
		}
		return true
	}
	
	var i = -1
	for _, val := range s {
		if val == '[' || val == '(' || val == '{' {
			stack = append(stack, val)
			i++
			continue
		} else {
			if !remove(val) {
				fmt.Println("##", "not satisfied")
				break
			}
		}
		stack = stack[:i]
		i--
	}

}
