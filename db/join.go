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

type tableJoinQuery struct {
	tableName string

	tq Queryx

	joinType string

	op Operator
}

func (tq Queryx) LeftJoin(tableName string) tableJoinQuery {
	tjq := tableJoinQuery{
		tableName: tableName,
		tq:        tq,
		joinType:  "LEFT",
	}

	return tjq
}

func (tq Queryx) RightJoin(tableName string) tableJoinQuery {
	tjq := tableJoinQuery{
		tableName: tableName,
		tq:        tq,
		joinType:  "RIGHT",
	}

	return tjq
}

func (tq Queryx) Join(tableName string) tableJoinQuery {
	tjq := tableJoinQuery{
		tableName: tableName,
		tq:        tq,
		joinType:  "",
	}

	return tjq
}

func (tjq tableJoinQuery) On(operator Operator) Queryx {
	tjq.op = operator

	tq := tjq.tq
	tq.builder = append(tq.builder, tjq)
	return tq
}
