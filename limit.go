package db

import (
	"fmt"
)

func (qx Queryx) Limit(offset, count int) Queryx {
	q := string(qx.Query)
	params := qx.Params

	q, params = Limit(offset, count).Wrap(q, params)

	return Queryx{
		Query(q),
		params,
	}
}

// Limit TODO: NEEDS COMMENT INFO
func Limit(offset, count int) selectOption {
	return &limitOption{offset, count}
}

type limitOption struct {
	offset int
	count  int
}

// Wrap TODO: NEEDS COMMENT INFO
func (o *limitOption) Wrap(query string, params []interface{}) (string, []interface{}) {
	query = fmt.Sprintf("SELECT a.* FROM (%s) a LIMIT ?, ?", query)
	params = append(params, o.offset)
	params = append(params, o.count)
	return query, params
}
