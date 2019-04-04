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

type orOperator struct {
	operators []Operator
}

// Or TODO: NEEDS COMMENT INFO
func Or(operators ...Operator) Operator {
	return &orOperator{operators}
}

// Make TODO: NEEDS COMMENT INFO
func (o *orOperator) Make() (string, []interface{}) {
	queries := []string{}
	params := []interface{}{}

	for _, operator := range o.operators {
		operatorQuery, operatorParams := operator.Make()

		queries = append(queries, fmt.Sprintf("(%s)", operatorQuery))
		params = append(params, operatorParams...)
	}

	return strings.Join(queries, " OR "), params
}
