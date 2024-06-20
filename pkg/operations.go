package pkg

import (
	"math"
	"slices"
)

type Operator struct {
	Operator   string
	Precedence int
}

var Operators = []Operator{
	{"+", 4},
	{"-", 4},
	{"*", 3},
	{"/", 3},
	{"^", 2},
	{"v", 2},
	{"(", 1},
	{")", 1},
	{"[", 1},
	{"]", 1},
	{"{", 1},
	{"}", 1},
}

func GetPrecedence(op string) int {
	for _, operator := range Operators {
		if operator.Operator == op {
			return operator.Precedence
		}
	}

	return 0
}

func IsOperator(op string) bool {
	return slices.Contains(OperatorsToStringArray(), op)
}

func OperatorsToStringArray() []string {
	var res []string
	for _, operator := range Operators {
		res = append(res, operator.Operator)
	}
	return res
}

func IsGrouping(op string) bool {
	return IsLeftGrouping(op) || IsRightGrouping(op)
}

func IsLeftGrouping(op string) bool {
	return op == "(" || op == "[" || op == "{"
}

func IsRightGrouping(op string) bool {
	return op == ")" || op == "]" || op == "}"
}

func IsSameGrouping(op1 string, op2 string) bool {
	res := false

	if IsParenthesis(op1) && IsParenthesis(op2) {
		res = true
	}
	if IsBracket(op1) && IsBracket(op2) {
		res = true
	}
	if IsCurlyBrace(op1) && IsCurlyBrace(op2) {
		res = true
	}

	return res
}

func IsParenthesis(op string) bool {
	return op == "(" || op == ")"
}

func IsBracket(op string) bool {
	return op == "[" || op == "]"
}

func IsCurlyBrace(op string) bool {
	return op == "{" || op == "}"
}

func RealizeOperation(op string, lval float64, rval float64) float64 {
	res := 0.0

	switch op {
	case "+":
		res = lval + rval
	case "-":
		res = lval - rval
	case "*":
		res = lval * rval
	case "/":
		res = lval / rval
	case "^":
		res = math.Pow(lval, rval)
	case "v":
		res = math.Pow(rval, 1.0/lval)
	}

	return res
}
