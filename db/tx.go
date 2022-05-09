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
	"bytes"
	"database/sql"
	"fmt"
	"reflect"
	"runtime"
	"strings"
	"sync"
	"time"

	"github.com/jmoiron/sqlx"
)

// Tx TODO: NEEDS COMMENT INFO
type Tx struct {
	Tx *sqlx.Tx

	counter uint64

	m          sync.Mutex
	stacktrace string
	time       time.Time

	statementsCache sync.Map

	queries []string
}

func (tx *Tx) Preparex(query Query) (*sqlx.Stmt, error) {
	tx.queries = append(tx.queries, string(query))

	if stmt, ok := tx.statementsCache.Load(string(query)); ok {
		return stmt.(*sqlx.Stmt), nil
	}

	stmt, err := tx.Tx.Preparex(string(query))
	if err != nil {
		return nil, err
	}

	tx.statementsCache.Store(string(query), stmt)
	return stmt, nil
}

func findMethod() string {
	trace := make([]byte, 1024)

	count := runtime.Stack(trace, true)
	trace = trace[:count]

	parts := bytes.Split(trace, []byte("\n"))

	return strings.TrimSpace(string(parts[6]))
}

func (tx *Tx) Commit() error {
	tx.m.Lock()
	defer tx.m.Unlock()

	log.Infof("[%d] tx", tx.counter)
	defer log.Infof("[%d] tx finished", tx.counter)

	err := tx.Tx.Commit()
	if err != nil {
		log.Errorf("[%d] Could not commit transaction (%s): %s: %p", tx.counter, findMethod(), err, tx.Tx)
		return err
	}

	now := time.Now()

	if now.Sub(tx.time) > 1*time.Second {
		log.Warningf("[%d] Transaction commit (%s) took long, took: %s, queries=\n * %v.", tx.counter, findMethod(), now.Sub(tx.time), strings.Join(tx.queries, "\n * "))
	}

	log.Debugf("[%d] Transaction commit (%s), took: %v. %p", tx.counter, findMethod(), now.Sub(tx.time), tx.Tx)

	tx.Tx = nil
	return err
}

func (tx *Tx) Rollback() error {
	tx.m.Lock()
	defer tx.m.Unlock()

	err := tx.Tx.Rollback()
	log.Errorf("[%d] Transaction rollback, took: %v", tx.counter, time.Since(tx.time))
	return err
}

/*
TableQuery("table").Fields(AllFields).Where().OrderBy().GroupBy(FieldX)
TableQuery("table").Fields(Count("xyz")).Where(db.And()).OrderBy(Fieldxyz).GroupBy(FieldOther)
TableQuery("events").Fields("test1").Join("otherTable").On(db.Equal("test", "test2"))
TableQuery("events").Fields("test1").LeftJoin("otherTable").On(db.Equal("test", "test2"))
TableQuery("events").Fields("test1").RightJoin("otherTable").On(db.Equal("test", "test2"))
QueryActions() = Query()
*/

/*
func Count(field Field) Field {

}
*/

/*
type fields []Field

func (f *fields) String() {
}
*/

// Selectx TODO: NEEDS COMMENT INFO
func (tx *Tx) Selectx(o interface{}, qy Queryx, options ...selectOption) error {
	tx.m.Lock()
	defer tx.m.Unlock()

	q, params := qy.Build()

	start := time.Now()

	defer func() {
		now := time.Now()
		if now.Sub(start) > 1*time.Second {
			log.Warningf("Query took too long %v: %s", now.Sub(start), q)
		}
	}()

	/*
		for _, option := range options {
			q, params = option.Wrap(q, params)
		}
	*/

	if u, ok := o.(Selecter); ok {
		err := u.Select(tx.Tx, q, params...)
		if err != nil {
			log.Errorf("Error executing query: %s: %s", q, err.Error())
		}

		return err
	}

	stmt, err := tx.Preparex(q)
	if err != nil {
		log.Errorf("Error executing query: %s: %s", q, err.Error())
		return err
	}

	return stmt.Select(o, params...)
}

// Selectx TODO: NEEDS COMMENT INFO
/*
func (tx *Tx) Selectx(o interface{}, qx Queryx, options ...selectOption) error {
	q := string(qx.Query)
	params := qx.Params

	start := time.Now()

	defer func() {
		now := time.Now()
		if now.Sub(start) > 1*time.Second {
			log.Warningf("Query took too long %v: %s", now.Sub(start), q)
		}
	}()

	for _, option := range options {
		q, params = option.Wrap(q, params)
	}

	if u, ok := o.(Selecter); ok {
		err := u.Select(tx.Tx, Query(q), params...)
		if err != nil {
			log.Errorf("Error executing query: %s: %s", q, err.Error())
		}

		return err
	}

	stmt, err := tx.Preparex(q)
	if err != nil {
		log.Errorf("Error executing query: %s: %s", q, err.Error())
		return err
	}

	return stmt.Select(o, params...)
}
*/

// Exists TODO: NEEDS COMMENT INFO
func (tx *Tx) Exists(qy Queryx) (bool, error) {
	tx.m.Lock()
	defer tx.m.Unlock()

	q, params := qy.Build()

	stmt, err := tx.Preparex(Query(fmt.Sprintf("SELECT EXISTS(%s)", string(q))))
	if err != nil {
		log.Errorf("Error preparing query: %s: %s", q, err.Error())
		return false, err
	}

	exists := false

	err = stmt.Get(&exists, params...)
	if err != nil {
		log.Errorf("Error executing query: %s: %s", q, err.Error())
		return false, err
	}

	return exists, err
}

// Countx TODO: NEEDS COMMENT INFO
func (tx *Tx) Countx(qy Queryx) (int, error) {
	tx.m.Lock()
	defer tx.m.Unlock()

	q, params := qy.Build()

	stmt, err := tx.Preparex(q)
	if err != nil {
		log.Errorf("Error preparing query: %s: %s", q, err.Error())
		return 0, err
	}

	count := 0

	err = stmt.Get(&count, params...)
	if err != nil {
		log.Errorf("Error executing query: %s: %s", q, err.Error())
	}

	return count, err
}

/*
// Countx TODO: NEEDS COMMENT INFO
func (tx *Tx) Countx(qx Queryx) (int, error) {
	stmt, err := tx.Preparex(fmt.Sprintf("SELECT COUNT(*) FROM (%s) q", string(qx.Query)))
	if err != nil {
		log.Errorf("Error preparing query: %s: %s", qx.Query, err.Error())
		return 0, err
	}

	count := 0

	err = stmt.Get(&count, qx.Params...)
	if err != nil {
		log.Errorf("Error executing query: %s: %s", qx.Query, err.Error())
	}

	return count, err
}
*/

// Exec TODO: NEEDS COMMENT INFO
/*
func (tx *Tx) Exec(qx Queryx) error {
	stmt, err := tx.Preparex(qx.Query)
	if err != nil {
		return err
	}

	_, err = stmt.Exec(qx.Params...)
	return err
}
*/
func (tx *Tx) Execute(qy Queryx) error {
	tx.m.Lock()
	defer tx.m.Unlock()

	q, params := qy.Build()
	log.Debugf("[%d] Executing query: %s", tx.counter, q)

	stmt, err := tx.Preparex(q)
	if err != nil {
		log.Errorf("[%d] Error preparing query: %s: %s", tx.counter, q, err.Error())
		return err
	}

	_, err = stmt.Exec(params...)
	if err != nil {
		log.Errorf("[%d] Error executing query: %s: %s", tx.counter, q, err.Error())
		return err
	}

	return err
}

// Getx TODO: NEEDS COMMENT INFO
func (tx *Tx) Getx(o interface{}, qy Queryx) error {

	tx.m.Lock()
	defer tx.m.Unlock()

	q, params := qy.Build()
	log.Debugf("[%d] Executing query: %s", tx.counter, q)

	if u, ok := o.(Getter); ok {
		err := u.Get(tx.Tx, q, params)
		if IsNoRowsErr(err) {
		} else if err != nil {
			log.Errorf("[%d] Error executing query: %s: %s", tx.counter, q, err.Error())
		}

		return err
	}

	log.Error("No getter found for object: %s", reflect.TypeOf(o))
	return ErrNoGetterFound
}

// Getx TODO: NEEDS COMMENT INFO
/*
func (tx *Tx) Getx(o interface{}, qx Queryx) error {
	if u, ok := o.(Getter); ok {
		err := u.Get(tx.Tx, qx)
		if err != nil {
			log.Errorf("Error executing query: %s: %s", qx.Query, err.Error())
		}

		return err
	}

	log.Error("No getter found for object: %s", reflect.TypeOf(o))
	return ErrNoGetterFound
}

// Get TODO: NEEDS COMMENT INFO
func (tx *Tx) Get(o interface{}, qx Queryx) error {
	if u, ok := o.(Getter); ok {
		err := u.Get(tx.Tx, qx)
		if err != nil {
			log.Errorf("Error executing query: %s: %s", qx.Query, err.Error())
		}

		return err
	}

	log.Error("No getter found for object: %s", reflect.TypeOf(o))
	return ErrNoGetterFound
}
*/

// Update TODO: NEEDS COMMENT INFO
func (tx *Tx) NamedExec(query string, arg interface{}) (sql.Result, error) {
	tx.m.Lock()
	defer tx.m.Unlock()

	log.Debugf("[%d] Executing query: %s", tx.counter, query)

	return tx.Tx.NamedExec(query, arg)
}

// Update TODO: NEEDS COMMENT INFO
func (tx *Tx) InsertOrUpdate(o interface{}) error {
	log.Debugf("[%d] Executing insert or update", tx.counter)
	if u, ok := o.(InsertOrUpdater); ok {
		return u.InsertOrUpdate(tx.Tx)
	}

	log.Error("No InsertOrUpdate found for object: %s", reflect.TypeOf(o))
	return ErrNoInsertOrUpdaterFound
}

// Update TODO: NEEDS COMMENT INFO
func (tx *Tx) Update(o interface{}) error {
	log.Debugf("[%d] Executing update", tx.counter)
	if u, ok := o.(Updater); ok {
		return u.Update(tx.Tx)
	}

	log.Error("No updater found for object: %s", reflect.TypeOf(o))
	return ErrNoUpdaterFound
}

// Delete TODO: NEEDS COMMENT INFO
func (tx *Tx) Delete(o interface{}) error {
	log.Debugf("[%d] Executing delete", tx.counter)

	if u, ok := o.(Deleter); ok {
		return u.Delete(tx.Tx)
	}

	log.Error("No deleter found for object: %s", reflect.TypeOf(o))
	return ErrNoDeleterFound
}

// Insert TODO: NEEDS COMMENT INFO
func (tx *Tx) Insert(o interface{}) error {
	log.Debugf("[%d] Executing insert", tx.counter)

	if u, ok := o.(Inserter); ok {
		err := u.Insert(tx.Tx)
		if err != nil {
			log.Error(err.Error())
		}
		return err
	}

	log.Error("No inserter found for object: %s", reflect.TypeOf(o))
	return ErrNoInserterFound
}

type TxOptionFunc func(opt *sql.TxOptions)

func ReadOnly() TxOptionFunc {
	return func(opt *sql.TxOptions) {
		opt.ReadOnly = true
	}
}
