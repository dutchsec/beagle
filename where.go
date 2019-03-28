package db

import (
	"fmt"
)

// Where TODO: NEEDS COMMENT INFO
func Where(operator Operator) queryxOption {
	return &whereOption{operator}
}

// Where TODO: NEEDS COMMENT INFO
func (qx Queryx) Where(operator Operator) Queryx {
	q := string(qx.Query)
	params := qx.Params

	q, params = Where(operator).Wrap(q, params)

	return Queryx{
		Query(q),
		params,
	}
}

type whereOption struct {
	operator Operator
}

// Wrap TODO: NEEDS COMMENT INFO
func (o *whereOption) Wrap(query string, params []interface{}) (string, []interface{}) {
	whereQuery, whereParams := o.operator.Make()
	query = fmt.Sprintf("SELECT * FROM (%s) a WHERE %s", query, whereQuery)

	params = append(params, whereParams...)
	return query, params
}
