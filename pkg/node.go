package pkg

import (
	"strconv"
)

type Node struct {
	Value     string
	Left_val  *Node
	Right_val *Node
}

func (n *Node) Calculate() float64 {
	lval := 0.0
	rval := 0.0

	if n.Left_val != nil {
		lval = n.Left_val.Calculate()
	}

	if n.Right_val != nil {
		rval = n.Right_val.Calculate()
	}

	if n.Left_val == nil && n.Right_val == nil {
		value, _ := strconv.ParseFloat(n.Value, 64)
		return value
	}

	return n.DoOperation(lval, rval)
}

func (n *Node) DoOperation(lval float64, rval float64) float64 {
	return RealizeOperation(n.Value, lval, rval)
}
