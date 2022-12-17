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
)

func SubQuery(v interface{}) subQuery {
	return subQuery{
		v,
	}
}

type subQuery struct {
	v interface{}
}

func (f subQuery) Build() (Query, []interface{}) {
	if blder, ok := f.v.(Builder); ok {
		qry, params := blder.Build()
		return Query(fmt.Sprintf("(%s)", qry)), params
	} else {
		return Query(fmt.Sprintf("(%s)", f.v)), []interface{}{}
	}
}
