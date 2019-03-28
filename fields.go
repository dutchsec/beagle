package db

import (
	"fmt"
	"strings"
)

// Where TODO: NEEDS COMMENT INFO
func (qx Queryx) Fields(fields ...string) Queryx {
	q := string(qx.Query)
	params := qx.Params

	q, params = Fields(fields).Wrap(q, params)

	return Queryx{
		Query(q),
		params,
	}
}

func Fields(fields []string) selectOption {
	return &fieldsOption{fields}
}

type fieldsOption struct {
	fields []string
}

// Wrap TODO: NEEDS COMMENT INFO
func (o *fieldsOption) Wrap(query string, params []interface{}) (string, []interface{}) {
	query = fmt.Sprintf("SELECT %s FROM (%s) a", strings.Join(o.fields, ","), query)
	return query, params
}
