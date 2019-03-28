package db

type queryxOption interface {
	Wrap(string, []interface{}) (string, []interface{})
}

// Query TODO: NEEDS COMMENT INFO
type Query string

// Queryx TODO: NEEDS COMMENT INFO
type Queryx struct {
	Query  Query
	Params []interface{}
}
