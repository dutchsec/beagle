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
	"crypto/rand"
	"fmt"
	"unicode"
)

type Field string

func (s Field) Alias(alias string) {
	// NOT IMPLEMENTED YET
}

func sanitize(s string) (string, error) {
	field := ""

	state := 0

	// we need to tokenize

	i := 0

	valid := func(r rune) bool {
		return unicode.IsLetter(r) || unicode.IsNumber(r) || r == '_'
	}

	runes := []rune(s)

	for i < len(runes) {
		switch {
		case valid(runes[i]):
			if state == 0 {
				field += string(runes[i])
				i++

				state = 1
			} else if state == 1 {
				field += string(runes[i])
				i++
			} else if state == 2 {
				field += string(runes[i])
				i++
			} else {
				return "", fmt.Errorf("Warning, unexpected character in field name: %c (%s)", runes[i], s)
			}
		case runes[i] == '`':
			if state == 0 {
				field += string(runes[i])
				i++

				state = 2
			} else if state == 2 {
				field += string(runes[i])
				i++

				state = 3
			} else {
				return "", fmt.Errorf("Warning, unexpected character in field name: %c (%s)", runes[i], s)
			}
		case runes[i] == '.':
			if state == 3 {
				field += string(runes[i])
				i++

				state = 0
			} else {
				return "", fmt.Errorf("Warning, unexpected character in field name: %c (%s)", runes[i], s)
			}
		default:
			return "", fmt.Errorf("Warning, unexpected character in field name: %c (%s)", runes[i], s)

		}
	}

	if state != 3 && state != 1 {
		return "", fmt.Errorf("Unexpected state %d, invalid field name: %s", state, s)
	}

	return field, nil
}

var chars = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890-"

func RandomString(length int) string {
	ll := len(chars)
	b := make([]byte, length)
	rand.Read(b) // generates len(b) random bytes
	for i := 0; i < length; i++ {
		b[i] = chars[int(b[i])%ll]
	}
	return string(b)
}

func (f Field) Build() (Query, []interface{}) {
	// we're sanitizing the column names, to prevent
	// accidental mistakes in eg where statements.
	// Ideally we'd refactor using fields.
	// in case of unallowed values, we'll just return a
	// random string
	s, err := sanitize(string(f))
	if err != nil {
		fmt.Println(err.Error())
		return Query(fmt.Sprintf(`"%s"`, RandomString(32))), []interface{}{}
	}

	return Query(s), []interface{}{}
}
