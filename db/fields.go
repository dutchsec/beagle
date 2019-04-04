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
