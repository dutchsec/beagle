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

type orderByOption struct {
	fields []Field
	desc   bool
}

// TODO: keep track of fields on construct at once, make it an arry
// Query().OrderBy(fld).OrderByDesc(fld2)
func (tq Queryx) OrderBy(fields ...Field) Queryx {
	ob := orderByOption{fields, false}
	tq.builder = append(tq.builder, ob)
	return tq
}

// TODO: keep track of fields on construct at once, make it an arry
func (tq Queryx) OrderByDesc(fields ...Field) Queryx {
	ob := orderByOption{fields, true}
	tq.builder = append(tq.builder, ob)
	return tq
}
