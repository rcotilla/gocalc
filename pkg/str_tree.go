package pkg

import (
	"slices"
)

func Calculate(str string) float64 {

	tokens := strToTockens(str)
	tree := arrayToTree(tokens)

	return tree.Calculate()
}

func strToTockens(str string) []string {
	var res []string

	number := ""
	operatorsToStringArray := OperatorsToStringArray()
	for _, c := range str {
		if slices.Contains(operatorsToStringArray, string(c)) {
			if number != "" {
				res = append(res, number)
				number = ""
			}
			res = append(res, string(c))
		} else if c != ' ' {
			number += string(c)
		}
	}

	if number != "" {
		res = append(res, number)
	}

	return res
}

func arrayToTree(arr []string) *Node {
	if !isValidExpression(arr) {
		panic("Invalid expression")
	}
	node := &Node{}
	value := ""
	value_pos := 0

	arr = removeGroupings(arr)
	inGroup := 0

	if len(arr) == 1 {
		node.Value = arr[0]
	} else {
		for pos := len(arr) - 1; pos >= 0; pos-- {
			if IsRightGrouping(arr[pos]) {
				inGroup++
				continue
			}

			if IsLeftGrouping(arr[pos]) {
				inGroup--
				continue
			}

			if IsOperator(arr[pos]) && inGroup == 0 {
				if value == "" ||
					(GetPrecedence(arr[pos]) > GetPrecedence(value)) {
					value = arr[pos]
					value_pos = pos
				}

				continue
			}
		}

		node.Value = arr[value_pos]
		node.Left_val = arrayToTree(arr[:value_pos])
		node.Right_val = arrayToTree(arr[value_pos+1:])
	}

	return node
}

func isValidExpression(arr []string) bool {
	operatorStack := []string{}
	for i := 0; i <= len(arr)-1; i++ {
		if IsOperator(arr[i]) {
			if IsLeftGrouping(arr[i]) {
				if i != 0 && !IsOperator(arr[i-1]) {
					return false
				}
				operatorStack = append(operatorStack, arr[i])
				continue
			}

			if IsRightGrouping(arr[i]) {
				if i != len(arr)-1 && !IsOperator(arr[i+1]) {
					return false
				}
				if len(operatorStack) == 0 {
					return false
				}
				if !IsLeftGrouping(operatorStack[len(operatorStack)-1]) && !IsSameGrouping(arr[i], operatorStack[len(operatorStack)-1]) {
					return false
				}
				operatorStack = operatorStack[:len(operatorStack)-1]

				continue
			}

			// TODO: Fix this
			// if i > 0 || i < len(arr)-1 && !IsGrouping(arr[i]) {
			// 	if !IsGrouping(arr[i-1]) || (!IsGrouping(arr[i+1]) && IsOperator(arr[i-1])) {
			// 		return false
			// 	}
			// }
		}

	}

	return true
}

func removeGroupings(arr []string) []string {
	if IsLeftGrouping(arr[0]) && IsRightGrouping(arr[len(arr)-1]) {
		return removeGroupings(arr[1 : len(arr)-1])
	}

	return arr
}
