package db

import (
	"fmt"
	"strings"
)

type orOperator struct {
	operators []Operator
}

// OrOperator TODO: NEEDS COMMENT INFO
func OrOperator(operators ...Operator) Operator {
	return &orOperator{operators}
}

// Make TODO: NEEDS COMMENT INFO
func (o *orOperator) Make() (string, []interface{}) {
	queries := []string{}
	params := []interface{}{}

	for _, operator := range o.operators {
		operatorQuery, operatorParams := operator.Make()

		queries = append(queries, fmt.Sprintf("(%s)", operatorQuery))
		params = append(params, operatorParams...)
	}

	return strings.Join(queries, " OR "), params
}
