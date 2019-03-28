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
