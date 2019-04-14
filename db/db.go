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
	"context"
	"database/sql"
	"runtime"
	"time"

	"github.com/jmoiron/sqlx"

	logging "github.com/op/go-logging"
)

// TODO GENERATOR
//     created_at timestamp without time zone NOT NULL,
// updated_at timestamp without time zone NOT NULL,
// CreatedAt            time.Time    `db:"created_at"`
// UpdatedAt            time.Time    `db:"updated_at"`

// ctx.tx.Selectx().Where().Execute(&connections)

var log = logging.MustGetLogger("go.dutchsec.com/db")

// Newerer TODO: NEEDS COMMENT INFO
type Newerer interface {
	isNew() bool
	setNew(bool)
}

func IsNoRowsErr(err error) bool {
	return err == sql.ErrNoRows
}

// New TODO: NEEDS COMMENT INFO
type New struct {
	new bool
}

func (d *New) isNew() bool {
	return d.new
}

func (d *New) setNew(v bool) {
	d.new = v
}

func Connect(driverName, dataSourceName string) (*DB, error) {
	db, err := sqlx.Connect(driverName, dataSourceName)
	if err != nil {
		return nil, err
	}

	return &DB{
		db,
	}, nil
}

// DB TODO: NEEDS COMMENT INFO
type DB struct {
	*sqlx.DB
}

type selectOption interface {
	Wrap(string, []interface{}) (string, []interface{})
}

// Begin TODO: NEEDS COMMENT INFO
func (db *DB) Begin(ctx context.Context, opts ...TxOptionFunc) *Tx {
	txOptions := &sql.TxOptions{}
	for _, fn := range opts {
		fn(txOptions)
	}

	tx, err := db.DB.BeginTxx(ctx, txOptions)
	if err != nil {
		log.Error("Error starting transaction: %s", err.Error())
		return nil
	}

	trace := make([]byte, 1024)
	count := runtime.Stack(trace, true)
	trace = trace[:count]

	return &Tx{
		Tx: tx,

		stacktrace: string(trace),
		time:       time.Now(),
	}
}

// Updater TODO: NEEDS COMMENT INFO
type Updater interface {
	Update(*sqlx.Tx) error
}

type InsertOrUpdater interface {
	InsertOrUpdate(*sqlx.Tx) error
}

// Inserter TODO: NEEDS COMMENT INFO
type Inserter interface {
	Insert(*sqlx.Tx) error
}

// Selecter TODO: NEEDS COMMENT INFO
type Selecter interface {
	Select(*sqlx.Tx, Query, ...interface{}) error
}

// Getter TODO: NEEDS COMMENT INFO
type Getter interface {
	Get(*sqlx.Tx, Queryx) error
}

// Deleter TODO: NEEDS COMMENT INFO
type Deleter interface {
	Delete(*sqlx.Tx) error
}
