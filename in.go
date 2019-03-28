package db

import (
	"fmt"
)

type inOperator struct {
	qry Queryx

	field string
}

// OrOperator TODO: NEEDS COMMENT INFO
func InOperator(field string, qry Queryx) Operator {
	return &inOperator{
		field: field,
		qry:   qry,
	}
}

// Make TODO: NEEDS COMMENT INFO
func (o *inOperator) Make() (string, []interface{}) {
	return fmt.Sprintf("%s IN (%s)", o.field, o.qry.Query), o.qry.Params
}
