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

/*
type wrapper struct {
	alias string
}

func (w *wrapper) NextAlias() string {
	w.alias = w.alias + 1
}

*/

/*
// TODO: Join is not complete and working yet
// Join TODO: NEEDS COMMENT INFO

func (qx Queryx) Join(qry Queryx, left, right string) Queryx {
	q := string(qx.Query)
	params := qx.Params

	q, params = Join(qry, left, right).Wrap(q, params)

	return Queryx{
		Query(q),
		params,
	}
}

func Join(qry Queryx, left, right string) selectOption {
	return &joinOption{qry, left, right}
}

type joinOption struct {
	qry   Queryx
	left  string
	right string
}

// Wrap TODO: NEEDS COMMENT INFO
// what fields do we want to return?
// do we need an genereated AllField?
func (o *joinOption) Wrap(query string, params []interface{}) (string, []interface{}) {
	// automatic column naming
	query = fmt.Sprintf("SELECT a.* FROM (%s) a JOIN (%s) b ON a.`%s` = b.`%s`", query, o.qry.Query, o.left, o.right)
	params = append(params, o.qry.Params...)
	return query, params
}

/*
type joinOperator struct {
	qry Queryx

	left  string
	right string
}

// OrOperator TODO: NEEDS COMMENT INFO
func JoinOperator(qry Queryx, left, right string) Operator {
	return &joinOperator{
		qry:   qry,
		left:  left,
		right: right,
	}
}

// TODO: Use named fields instead of counted
// Make TODO: NEEDS COMMENT INFO
func (o *joinOperator) Make() (string, []interface{}) {
	return fmt.Sprint("JOIN (%s) b ON a.`%s` = b.`%s`", o.qry.Query), o.qry.Params
}
*/
