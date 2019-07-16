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
*/

type whereOption struct {
	operator Operator
}

/*
// Wrap TODO: NEEDS COMMENT INFO
func (o *whereOption) Wrap(query string, params []interface{}) (string, []interface{}) {
	whereQuery, whereParams := o.operator.Make()
	query = fmt.Sprintf("%s WHERE %s", query, whereQuery)

	params = append(params, whereParams...)
	return query, params
}
*/
