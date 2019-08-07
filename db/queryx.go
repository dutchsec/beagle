// Copyright 2019 The DutchSec Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
package db

import (
	"fmt"
	"strings"
)

type Queryx struct {
	tableName string

	type_ string

	countRows bool

	fields []Field

	builder []interface{}
}

func (tq Queryx) Dump() string {
	q, _ := tq.Build()
	return string(q)
}

func (tq Queryx) CountRows() Queryx {
	tq.countRows = true
	return tq
}

/*
	if _, err = manager.DbMap.Select(&total, "SELECT FOUND_ROWS()"); err != nil {
		return &countryList, 0, err
	}
*/

func (tq Queryx) Build() (Query, []interface{}) {
	fields := make([]string, len(tq.fields))
	for i, field := range tq.fields {
		fields[i] = string(field)
	}

	b := strings.Builder{}

	if tq.type_ == "SELECT" {
		b.WriteString("SELECT ")

		if tq.countRows {
			b.WriteString("SQL_CALC_FOUND_ROWS ")
		}

		b.WriteString(fmt.Sprintf("%s ", strings.Join(fields, ",")))
	} else if tq.type_ == "DELETE" {
		b.WriteString("DELETE ")

		b.WriteString(fmt.Sprintf("%s.* ", tq.tableName))
	} else {
		// failed

	}

	b.WriteString("FROM ")

	b.WriteString(fmt.Sprintf("%s ", tq.tableName))

	params := []interface{}{}

	for _, expr := range tq.builder {
		// b.WriteString(expr.String())
		if w, ok := expr.(where); ok {

			whereStmt, whereParams := w.Make()
			params = append(params, whereParams...)

			if whereStmt == "" {
			} else {
				b.WriteString("WHERE ")
				b.WriteString(whereStmt)
				b.WriteString(" ")
			}
		} else if tjq, ok := expr.(tableJoinQuery); ok {
			if tjq.joinType == "LEFT" {
				b.WriteString("LEFT JOIN ")
			} else if tjq.joinType == "RIGHT" {
				b.WriteString("RIGHT JOIN ")
			} else {
				b.WriteString("JOIN ")
			}

			b.WriteString(fmt.Sprintf("%s ", tjq.tableName))
			b.WriteString("ON ")

			// whereStmt, whereParams := tjq.on.Make()
			// params = append(params, whereParams...)

			b.WriteString(fmt.Sprintf("%s=%s ", tjq.left, tjq.right))
		} else if lo, ok := expr.(limitOption); ok {
			b.WriteString("LIMIT ")

			b.WriteString(fmt.Sprintf("%d, %d", lo.offset, lo.count))
		} else if gb, ok := expr.(groupBy); ok {
			b.WriteString("GROUP BY ")

			// TODO: join fields
			b.WriteString(fmt.Sprintf("%s ", gb[0]))
		} else if ob, ok := expr.(orderByOption); ok {
			b.WriteString("ORDER BY ")

			// TODO: join fields
			b.WriteString(fmt.Sprintf("%s ", ob.fields[0]))

			// TODO: join fields
			if ob.desc {
				b.WriteString("DESC ")
			}
		}

		fmt.Printf("%#v\n", expr)
	}

	fmt.Println(b.String())

	qry := b.String()

	return Query(qry), params
}

func SelectQuery(tableName string) Queryx {
	return Queryx{
		tableName: tableName,
		type_:     "SELECT",
	}
}

func DeleteQuery(tableName string) Queryx {
	return Queryx{
		tableName: tableName,
		type_:     "DELETE",
	}
}

func (tq Queryx) Fields(fields ...Field) Queryx {
	tq.fields = fields
	return tq
}

type where Operator

func (tq Queryx) Where(operator Operator) Queryx {
	w := where(operator)

	tq.builder = append(tq.builder, w)
	return tq
}

// TODO: keep track of fields on construct at once, make it an arry
func (tq Queryx) OrderBy(fields ...Field) Queryx {
	sfields := make([]string, len(fields))
	for i, field := range fields {
		sfields[i] = string(field)
	}

	ob := orderByOption{sfields, false}
	tq.builder = append(tq.builder, ob)
	return tq
}

// TODO: keep track of fields on construct at once, make it an arry
func (tq Queryx) OrderByDesc(fields ...Field) Queryx {
	sfields := make([]string, len(fields))
	for i, field := range fields {
		sfields[i] = string(field)
	}

	ob := orderByOption{sfields, true}
	tq.builder = append(tq.builder, ob)
	return tq
}

func (tq Queryx) Limit(offset, count int) Queryx {
	lo := limitOption{offset, count}
	tq.builder = append(tq.builder, lo)
	return tq
}

type groupBy []Field

func (tq Queryx) GroupBy(fields ...Field) Queryx {
	gb := groupBy(fields)
	tq.builder = append(tq.builder, gb)
	return tq
}
