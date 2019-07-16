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

// Wrap TODO: NEEDS COMMENT INFO
func (o *orderByOption) Wrap(query string, params []interface{}) (string, []interface{}) {
	query = fmt.Sprintf("SELECT * FROM (%s) a ORDER BY %s", query, strings.Join(o.fields, ","))

	if o.desc {
		query = fmt.Sprintf("%s DESC", query)
	}

	return query, params
}
*/

type orderByOption struct {
	fields []string
	desc   bool
}
