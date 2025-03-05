package calculation

import (
	"fmt"
	"strconv"
	"strings"
	//"structs"

	//"time"
	//"github.com/nastts/calc_go2/internal/orchestrator"
	"github.com/nastts/calc_go2/internal/structs"
)



type Task struct{
	ID string `json:"id"`
	Arg1 float64 `json:"arg1"`
	Arg2 float64 `json:"arg2"`
	Operation string `json:"operation"`
	OperationTime int `json:"operationTime"`
} 

func Tokenize(expression string) []string {
	var tokens []string
	var currentToken strings.Builder

	for _, char := range expression {
		if char == ' ' {
			continue
	}
		if Operator(string(char)) || char == '(' || char == ')' {
			if currentToken.Len() > 0 {
				tokens = append(tokens, currentToken.String())
				currentToken.Reset()
			}
			tokens = append(tokens, string(char))
		} else {
			currentToken.WriteRune(char)
		}
	}

	if currentToken.Len() > 0 {
		tokens = append(tokens, currentToken.String())
	}

	return tokens
	}

func ParseFloat(token string) bool {
 	_, err := strconv.ParseFloat(token, 64)
 	return err == nil
}

func Operator(token string) bool {
 	return token == "+" || token == "-" || token == "*" || token == "/"
}

func precedence(op string) int {
	switch op {
		case "*", "/":
			return 2
		case "+", "-":
			return 1
		
		default:
			return 0
	}
}

func evaluateRPN(tokens []string) (float64, error) {
 	stack := []float64{}

	for _, token := range tokens {
		if ParseFloat(token) {
			num, _ := strconv.ParseFloat(token, 64)
			stack = append(stack, num)
		} else if Operator(token) {
			if len(stack) < 2 {
				return 0, ErrInternalServerError
		}
		b := stack[len(stack)-1]
		a := stack[len(stack)-2]
		stack = stack[:len(stack)-2]

		switch token {
			case "+":
				stack = append(stack, a+b)
			case "-":
				stack = append(stack, a-b)
			case "*":
				stack = append(stack, a*b)
			case "/":
				if b == 0 {
					return 0, ErrExpressionIsNotValid
				}
				stack = append(stack, a/b)
			default:
				return 0, ErrExpressionIsNotValid
			}
		} else {
			return 0, ErrInternalServerError
		}
	}

	if len(stack) != 1 {
		return 0, ErrInternalServerError
	}
	return stack[0], nil
}




func Calc(expression string) (float64, error) {
	tokens := Tokenize(expression)
	if len(tokens) == 0 {
		return 0, ErrInternalServerError
 	}

 	output := []string{}
 	operatorStack := []string{}

	for _, token := range tokens {
		if ParseFloat(token) {
		output = append(output, token)
		} else if Operator(token) {
		for len(operatorStack) > 0 && precedence(operatorStack[len(operatorStack)-1]) >= precedence(token) {
			output = append(output, operatorStack[len(operatorStack)-1])
			operatorStack = operatorStack[:len(operatorStack)-1]
		}
		operatorStack = append(operatorStack, token)
		} else if token == "(" {
			operatorStack = append(operatorStack, token)
		} else if token == ")" {
			for len(operatorStack) > 0 && operatorStack[len(operatorStack)-1] != "(" {
				output = append(output, operatorStack[len(operatorStack)-1])
				operatorStack = operatorStack[:len(operatorStack)-1]
		}
			if len(operatorStack) == 0 {
				return 0, ErrExpressionIsNotValid
			}
			operatorStack = operatorStack[:len(operatorStack)-1] 
			} else {
				return 0, ErrExpressionIsNotValid
			}
	}

	for len(operatorStack) > 0 {
		if operatorStack[len(operatorStack)-1] == "(" {
			return 0, ErrExpressionIsNotValid
		}
		output = append(output, operatorStack[len(operatorStack)-1])
		operatorStack = operatorStack[:len(operatorStack)-1]
	}

	return evaluateRPN(output)
	}


func ParseExpression(expression string)(*structs.Node, error){
	tokens := Tokenize(expression)
	if len(tokens) == 0{
		return nil, fmt.Errorf("пустое выражение")
	}
	var outputQueue []*structs.Node
	var operations []string
	for _, token := range tokens{
		if Operator(token){
			for len(operations) > 0 {
				op := operations[len(operations)-1]
				if Operator(op) && precedence(op) >= precedence(token){
					operations = operations[:len(operations)-1]
					if len(outputQueue)<2{
						return nil, fmt.Errorf("incorrect")
					}
					left := outputQueue[len(outputQueue)-2]
					right := outputQueue[len(outputQueue)-1]
					outputQueue = outputQueue[:len(outputQueue)-2]
					node := &structs.Node{
						Operator: op,
						Left: left,
						Right: right,
						Computed: false,
					}
					outputQueue = append(outputQueue, node)
				}else {
					break
				}
			}
			operations = append(operations, token)
		}else if token == "("{
			operations = append(operations, token)
		}else if token == ")"{
			for len(operations) >0{
				op := operations[len(operations)-1]
				operations = operations[:len(operations)-1]
				if op == "("{
					break
				}
				if len(outputQueue) < 2{
					return nil, fmt.Errorf("incorrect")
				}
				left := outputQueue[len(outputQueue)-2]
				right := outputQueue[len(outputQueue)-1]
					outputQueue = outputQueue[:len(outputQueue)-2]
					node := &structs.Node{
						Operator: op,
						Left: left,
						Right: right,
						Computed: false,
					}
					outputQueue = append(outputQueue, node)
			}
		}else{
			val, err := strconv.ParseFloat(token, 64)
			if err != nil{
				return nil, fmt.Errorf("incorrect")
			}
			node := &structs.Node{
				Value: val,
				Computed: true,
			}
			outputQueue = append(outputQueue, node)
		}
	}
	for len(operations) > 0 {
		op := operations[len(operations)-1]
		operations = operations[:len(operations)-1]
		if op == "(" || op == ")" {
			return nil, fmt.Errorf("скобки")
		}
		if len(outputQueue) < 2 {
			return nil, fmt.Errorf("некорректное выражение")
		}
		right := outputQueue[len(outputQueue)-2]
		left := outputQueue[len(outputQueue)-1]
		outputQueue = outputQueue[:len(outputQueue)-2]
		node := &structs.Node{
			Operator: op,
			Left:     left,
			Right:    right,
			Computed: false,
		}
		outputQueue = append(outputQueue, node)
	}
	for _, op := range operations{
		if op == "("{
			return nil, fmt.Errorf("некорректное выражение — лишняя (")
		}
	}
	if len(outputQueue) != 1 {
		return nil, fmt.Errorf("некорректное выражение")
	}
	return outputQueue[0], nil
}


