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

import "fmt"

// Equal TODO: NEEDS COMMENT INFO
func Like(left interface{}, right interface{}) Operator {
	return &likeOperator{left, right}
}

type likeOperator struct {
	left  interface{}
	right interface{}
}

// Make TODO: NEEDS COMMENT INFO
func (o *likeOperator) Make() (string, []interface{}) {
	allParams := []interface{}{}

	leftValue := "?"

	if blder, ok := o.left.(Builder); ok {
		value, params := blder.Build()
		allParams = append(allParams, params...)
		leftValue = string(value)
	} else {
		allParams = append(allParams, o.left)
	}

	rightValue := "?"

	if blder, ok := o.right.(Builder); ok {
		value, params := blder.Build()
		allParams = append(allParams, params...)
		rightValue = string(value)
	} else {
		allParams = append(allParams, o.right)
	}

	return fmt.Sprintf("%s LIKE %s", leftValue, rightValue), allParams
}
