package db

import (
	"fmt"
	"strings"
)

// Where TODO: NEEDS COMMENT INFO
func (qx Queryx) OrderBy(fields ...string) Queryx {
	q := string(qx.Query)
	params := qx.Params

	q, params = OrderBy(fields...).Wrap(q, params)

	return Queryx{
		Query(q),
		params,
	}
}

// Where TODO: NEEDS COMMENT INFO
func (qx Queryx) OrderByDesc(fields ...string) Queryx {
	q := string(qx.Query)
	params := qx.Params

	q, params = OrderByDesc(fields...).Wrap(q, params)

	return Queryx{
		Query(q),
		params,
	}
}

// OrderByDesc TODO: NEEDS COMMENT INFO
func OrderByDesc(fields ...string) selectOption {
	return &orderByOption{fields, true}
}

// OrderBy TODO: NEEDS COMMENT INFO
func OrderBy(fields ...string) selectOption {
	return &orderByOption{fields, false}
}

type orderByOption struct {
	fields []string
	desc   bool
}

// Wrap TODO: NEEDS COMMENT INFO
func (o *orderByOption) Wrap(query string, params []interface{}) (string, []interface{}) {
	query = fmt.Sprintf("SELECT * FROM (%s) a ORDER BY %s", query, strings.Join(o.fields, ","))

	if o.desc {
		query = fmt.Sprintf("%s DESC", query)
	}

	return query, params
}
