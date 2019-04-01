package db

import "fmt"

// IsNotNull TODO: NEEDS COMMENT INFO
func IsNotNull(field string) Operator {
	return &isNotNullOperator{field}
}

type isNotNullOperator struct {
	field string
}

// Make TODO: NEEDS COMMENT INFO
func (o *isNotNullOperator) Make() (string, []interface{}) {
	return fmt.Sprintf("`%s` IS NULL"), []interface{}{}
}
