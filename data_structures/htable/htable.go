// Package htable is a data structure of rows and columns, with each row having the same number of items and each column
// holding the same type of data. The tables provide easy building and quick lookup for a fast implementation of storing
// and accessing uniform lists.
package htable

import (
	"github.com/snhilde/dsa/data_structures/hlist"
	"errors"
)


// Table is the main type in this package. It holds all the rows and columns of data.
type Table struct {
	cols   []string
	rows    *hlist.List
}


// New creates a new table. The strings will denote the names of each column, used during lookup.
func New(cols ...string) *Table {
	var t Table

	t.cols = cols
	t.rows = hlist.New()

	return &t
}
