// Package htable is a data structure of rows and columns, with each row having the same number of items and each column
// holding the same type of data. The tables provide easy building and quick lookup for a fast implementation of storing
// and accessing uniform lists.
package htable

import (
	"github.com/snhilde/dsa/data_structures/hlist"
	"reflect"
	"errors"
	"fmt"
)


// Table is the main type in this package. It holds all the rows and columns of data.
type Table struct {
	types []reflect.Kind
	cols  []string
	rows   *hlist.List
}


// New creates a new table. The strings will denote the names of each column, used during lookup.
func New(cols ...string) (*Table, error) {
	if cols == nil || len(cols) == 0 {
		return nil, errors.New("Missing column headers")
	}

	// Make sure none of the columns match each other. Not the most efficient, but necessary.
	for i, v := range cols {
		for j, w := range cols[i+1:] {
			if v == w {
				return nil, errors.New(fmt.Sprintf("Columns %v and %v have the same header", i, j))
			}
		}
	}

	var t Table
	t.types = make([]reflect.Kind, len(cols))
	t.cols = cols
	t.rows = hlist.New()

	return &t, nil
}


// AddRow adds a new row of items to the end of the table.
func (t *Table) AddRow(items ...interface{}) error {
	r, err := t.newRow(items)
	if err != nil {
		return err
	}

	return t.rows.Append(r)
}

// InsertRow inserts a new row of items at the specified index.
func (t *Table) InsertRow(index int, items ...interface{}) error {
	r, err := t.newRow(items)
	if err != nil {
		return err
	}

	return t.rows.Insert(index, r)
}

// RemoveRow deletes a row from the table.
func (t *Table) RemoveRow(index int) {
	if t == nil {
		return
	}

	_ = t.rows.Remove(index)

	return nil
}

// Rows returns the number of rows in the table, or -1 on error.
func (t *Table) Rows() int {
	if t == nil {
		return -1
	}

	return t.rows.Length()
}


func tErr() error {
	return errors.New("Table must be created with New() first")
}

func (t *Table) newRow(items ...interface{}) *hlist.List, error {
	if t == nil {
		return nil, tErr()
	} else if len(items) != len(t.cols) {
		return nil, errors.New(fmt.Sprintf("Number of items (%v) does not match number of columns (%v)", len(items), len(t.cols)))
	}

	// Build out our row.
	r := hlist.New()
	for i, v := range items {
		rv := reflect.ValueOf(v)
		k := rv.Kind()
		switch k {
			// We want to use the largest type available, if there are options.
		case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
			k = reflect.Int64
		case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
			k = reflect.Uint64
		case reflect.Float32, reflect.Float64:
			k = reflect.Float64
		case reflect.Complext64, reflect.Complext128:
			k = reflect.Complext128
		}

		if t.Rows() == 0 {
			// This is the first row being added to the table. It will set the type of each column in the table.
			t.types[i] = k
		} else {
			// Make sure the type of this element matches the prototype.
			if k != t.types[i] {
				return nil, errors.New(fmt.Sprintf("Item %v (%v) does not match prototype (%v)", i, k, t.types[i]))
			}
		}
		r.Append(v)
	}

	return r, nil
}
