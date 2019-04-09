// sufix project main.go
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var sufix string
	fmt.Scanln(&sufix)
	postfix := sufixtopostfix(sufix)
	answer := calculate(postfix)
	fmt.Println(answer)
}

func sufixtopostfix(sufix string) string {
	var postfix string
	var stack [50]string
	var prechar string
	top := -1

	suflen := len(sufix)

	for i := 0; i < suflen; i++ {
		char := string(sufix[i])

		switch char {

		//左括号进栈
		case "(":
			top++
			stack[top] = "("

		//右括号出栈
		case ")":
			for top != -1 {
				prechar = stack[top]
				top--
				if prechar == "(" {
					break
				}
				postfix += prechar
			}

		//number
		case "0", "1", "2", "3", "4", "5", "6", "7", "8", "9":
			postfix += char
			postfix += "#"

		//add or sub
		case "+", "-":
			for top != -1 {
				if stack[top] != "(" {
					prechar = stack[top]
					postfix += prechar
					top--
				} else {
					break
				}
			}
			top++
			stack[top] = char
		//mui or div
		case "*", "/":
			for top != -1 {
				prechar = stack[top]
				if prechar == "*" || prechar == "/" {
					postfix += prechar
					top--
				} else {
					break
				}
			}
			top++
			stack[top] = char
		}
	}
	//if stack is not empty
	for top != -1 {
		prechar = stack[top]
		postfix += prechar
		top--
	}
	return postfix
}

func calculate(postfix string) int {
	var stack [50]int
	top := -1
	var a, b, c, d, e int
	var prechar string
	strlen := len(postfix)

	for i := 0; i < strlen; i++ {
		char := string(postfix[i])
		switch char {
		case "+":
			a = stack[top]
			top--
			b = stack[top]
			top--
			c = a + b
			top++
			stack[top] = c
		case "-":
			a = stack[top]
			top--
			b = stack[top]
			top--
			c = b - a
			top++
			stack[top] = c
		case "*":
			a = stack[top]
			top--
			b = stack[top]
			top--
			c = a * b
			top++
			stack[top] = c
		case "/":
			a = stack[top]
			top--
			b = stack[top]
			top--
			c = b / a
			top++
			stack[top] = c
		case "#":
			d, _ = strconv.Atoi(prechar)
			top++
			stack[top] = d
			prechar = ""
		default:
			prechar += char
		}
	}
	e = stack[top]
	top--
	return e
}
