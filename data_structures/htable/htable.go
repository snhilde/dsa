// Package htable is a data structure of rows and columns, with each row having the same number of items and each column
// holding the same type of data. The tables provide easy building and quick lookup for a fast implementation of storing
// and accessing uniform lists.
package htable

import (
	"github.com/snhilde/dsa/data_structures/hlist"
	"reflect"
	"errors"
	"fmt"
	"strings"
)


// Table is the main type in this package. It holds all the rows of data.
type Table struct {
	cols  []string       // Column headers
	types []reflect.Kind // Types of each column (must be consistent for all rows)
	rows   *hlist.List   // Linked list of rows
}


// New creates a new table. The strings will denote the names of each column, used during lookup.
func New(cols ...string) (*Table, error) {
	if cols == nil || len(cols) == 0 {
		return nil, errors.New("Missing column headers")
	}

	// Validate the column headers.
	for i, v := range cols {
		// Make sure every column has a header.
		if v == "" {
			return nil, errors.New(fmt.Sprintf("Column %v has an empty header", i))
		}
		// Make sure none of the columns match each other. Not the most efficient, but necessary.
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
	row, err := t.newRow(items...)
	if err != nil {
		return err
	}

	return t.rows.Append(row)
}

// InsertRow inserts a new row of items at the specified index.
func (t *Table) InsertRow(index int, items ...interface{}) error {
	row, err := t.newRow(items...)
	if err != nil {
		return err
	}

	return t.rows.Insert(index, row)
}

// RemoveRow deletes a row from the table.
func (t *Table) RemoveRow(index int) error {
	if t == nil {
		return tErr()
	}

	v := t.rows.Remove(index)
	if v == nil {
		// hlist.Remove will return the value at the index. Because our rows can never be nil, if we receive a nil
		// value, then it means an error occurred.
		return errors.New(fmt.Sprintf("Failed to remove row %v", index))
	}

	// All good
	return nil
}

// Rows returns the number of rows in the table, or -1 on error.
func (t *Table) Rows() int {
	if t == nil {
		return -1
	}

	return t.rows.Length()
}

// Columns returns the number of columns in the table, or -1 on error.
func (t *Table) Columns() int {
	if t == nil {
		return -1
	}

	return len(t.cols)
}

// Count returns the number of items in the table, or -1 on error.
func (t *Table) Count() int {
	r := t.Rows()
	c := t.Columns()

	if r == -1 || c == -1 {
		return -1
	}

	return r * c
}

// String returns a printable list of values, grouped by row.
func (t *Table) String() string {
	if t == nil {
		return "<nil>"
	} else if t.Count() == 0 {
		return "<empty>"
	}

	var b strings.Builder
	i := 0
	n := t.Rows()
	rows := t.rows.YieldAll()
	for row := range rows {
		vs := row.(*Row)
		b.WriteString(fmt.Sprintf("%v", vs.v))
		if i != n - 1 {
			b.WriteString(", ")
		}
		i++
	}

	return b.String()
}

// Row returns the index of the first row that contains the item in the specified column, or -1 on error or not found.
func (t *Table) Row(col string, item interface{}) int {
	// Find out which column we need to match on. (This will also catch a nil table.)
	c := t.findCol(col)
	if c == -1 {
		return -1
	}

	// Get our iterator to go through the rows.
	i := 0
	quit := make(chan interface{})
	rows := t.rows.Yield(quit)
	for row := range rows {
		vs := row.(*Row)
		if reflect.DeepEqual(item, vs.v[c]) {
			// Break out of the list iteration. If Yield's goroutine has already exited (because the list was fully
			// traversed), then it won't receive the message to quit. We'll try to send the quit message, and then
			// we'll exit.
			select {
			case quit <- 0:
			default:
				break
			}
			return i
		}
		i++
	}

	// If we're here, then we didn't find anything.
	return -1
}

// Item returns the item at the specified coordinates, or nil if there is no item at the coordinates.
func (t *Table) Item(row int, col string) interface{} {
	// Find out which column we need to match on. (This will also catch a nil table.)
	c := t.findCol(col)
	if c == -1 {
		return nil
	}

	// Grab our row.
	r := t.rows.Value(row)
	if r == nil {
		return nil
	}

	// Get the value of the column.
	vs := r.(*Row)
	if len(vs.v) <= c {
		// Shouldn't ever happen, but our index is greater than the number of columns in the row.
		return nil
	}
	return vs.v[c]
}

// Matches returns true if the value matches the item at the specified coordinates or false if there is no match.
func (t *Table) Matches(row int, col string, v interface{}) bool {
	item := t.Item(row, col)
	return reflect.DeepEqual(v, item)
}


// Row holds all the data for each row in the table.
type Row struct {
	v []interface{}
}


func tErr() error {
	return errors.New("Table must be created with New() first")
}

func (t *Table) findCol(col string) int {
	if t == nil {
		return -1
	}

	for i, v := range t.cols {
		if col == v {
			return i
		}
	}

	// We didn't find the column.
	return -1
}

func (t *Table) newRow(items ...interface{}) (*Row, error) {
	if t == nil {
		return nil, tErr()
	} else if len(items) != t.Columns() {
		return nil, errors.New(fmt.Sprintf("Number of items (%v) does not match number of columns (%v)", len(items), len(t.cols)))
	}

	first := false
	if t.Rows() == 0 {
		// This is the first row being added to the table. It will set the type of each column in the table.
		first = true
	}

	// Build out our row.
	row := new(Row)
	row.v = make([]interface{}, t.Columns())
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
		case reflect.Complex64, reflect.Complex128:
			k = reflect.Complex128
		}

		if first {
			t.types[i] = k
		} else {
			// Make sure the type of this element matches the prototype.
			if k != t.types[i] {
				return nil, errors.New(fmt.Sprintf("Item %v's type (%v) does not match column's prototype (%v)", i, k, t.types[i]))
			}
		}
		row.v[i] = v
	}

	return row, nil
}
