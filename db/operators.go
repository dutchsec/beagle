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

// Operator TODO: NEEDS COMMENT INFO
type Operator interface {
	Make() (string, []interface{})
}

// Active TODO: NEEDS COMMENT INFO
// this is not ok.
func Active() Operator {
	return &equalOperator{"active", "1"}
}

func True(fld Field) Operator {
	return &equalOperator{fld, "1"}
}
