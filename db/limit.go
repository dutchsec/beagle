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

// Wrap TODO: NEEDS COMMENT INFO
func (o *limitOption) Wrap(query string, params []interface{}) (string, []interface{}) {
	query = fmt.Sprintf("%s LIMIT ?, ?", query)
	params = append(params, o.offset)
	params = append(params, o.count)
	return query, params
}
*/

type limitOption struct {
	offset int
	count  int
}
