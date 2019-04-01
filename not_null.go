package db

import "fmt"

// IsNotNullOperator TODO: NEEDS COMMENT INFO
func IsNotNullOperator(field string) Operator {
	return &isNotNullOperator{field}
}

type isNotNullOperator struct {
	field string
}

// Make TODO: NEEDS COMMENT INFO
func (o *isNotNullOperator) Make() (string, []interface{}) {
	return fmt.Sprintf("`%s` IS NULL"), []interface{}{}
}
