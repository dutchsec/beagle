package db

import (
	"fmt"
	"strings"
)

type andOperator struct {
	operators []Operator
}

// TODO: the AndOperator should check and combine multiple and operators and should merge them!

// AndOperator TODO: NEEDS COMMENT INFO
func And(operators ...Operator) Operator {
	return &andOperator{operators}
}

// Make TODO: NEEDS COMMENT INFO
func (o *andOperator) Make() (string, []interface{}) {
	queries := []string{}
	params := []interface{}{}

	for _, operator := range o.operators {
		operatorQuery, operatorParams := operator.Make()

		queries = append(queries, fmt.Sprintf("(%s)", operatorQuery))
		params = append(params, operatorParams...)
	}

	return strings.Join(queries, " AND "), params
}
