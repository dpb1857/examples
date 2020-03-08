// Code generated by goa v3.1.1, DO NOT EDIT.
//
// divider HTTP client CLI support package
//
// Command:
// $ goa gen goa.design/examples/error/design -o
// $(GOPATH)/src/goa.design/examples/error

package client

import (
	"fmt"
	"strconv"

	divider "goa.design/examples/error/gen/divider"
)

// BuildIntegerDividePayload builds the payload for the divider integer_divide
// endpoint from CLI flags.
func BuildIntegerDividePayload(dividerIntegerDivideA string, dividerIntegerDivideB string) (*divider.IntOperands, error) {
	var err error
	var a int
	{
		var v int64
		v, err = strconv.ParseInt(dividerIntegerDivideA, 10, 64)
		a = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for a, must be INT")
		}
	}
	var b int
	{
		var v int64
		v, err = strconv.ParseInt(dividerIntegerDivideB, 10, 64)
		b = int(v)
		if err != nil {
			return nil, fmt.Errorf("invalid value for b, must be INT")
		}
	}
	v := &divider.IntOperands{}
	v.A = a
	v.B = b

	return v, nil
}

// BuildDividePayload builds the payload for the divider divide endpoint from
// CLI flags.
func BuildDividePayload(dividerDivideA string, dividerDivideB string) (*divider.FloatOperands, error) {
	var err error
	var a float64
	{
		a, err = strconv.ParseFloat(dividerDivideA, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for a, must be FLOAT64")
		}
	}
	var b float64
	{
		b, err = strconv.ParseFloat(dividerDivideB, 64)
		if err != nil {
			return nil, fmt.Errorf("invalid value for b, must be FLOAT64")
		}
	}
	v := &divider.FloatOperands{}
	v.A = a
	v.B = b

	return v, nil
}
