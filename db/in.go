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

type inOperator struct {
	qry Queryx

	field Field
}

// OrOperator TODO: NEEDS COMMENT INFO
func InOperator(field Field, qry Queryx) Operator {
	return &inOperator{
		field: field,
		qry:   qry,
	}
}

// Make TODO: NEEDS COMMENT INFO
func (o *inOperator) Make() (string, []interface{}) {
	qry, params := o.qry.Build()

	return fmt.Sprintf("%s IN (%s)", string(o.field), qry), params
}

type inStringOperator struct {
	params []interface{}

	field Field
}

// OrOperator TODO: NEEDS COMMENT INFO
func In(field Field, params []interface{}) Operator {
	return &inStringOperator{
		field:  field,
		params: params,
	}
}

// Make TODO: NEEDS COMMENT INFO
func (o *inStringOperator) Make() (string, []interface{}) {
	placeholders := make([]string, len(o.params))
	for i := 0; i < len(o.params); i++ {
		placeholders[i] = "?"
	}

	// fmt.Println(fmt.Sprintf("%s IN (%s)", o.field, strings.Join(placeholders, ",")), o.params)
	return fmt.Sprintf("%s IN (%s)", string(o.field), strings.Join(placeholders, ",")), o.params
}
