// Package htable is a data structure of rows and columns, with each row having the same number of
// items and each column holding the same type of data. The tables provide easy, fast construction
// and lookup of uniform lists. Following the go convention, all data within a column must have the
// same static type.
package htable

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/snhilde/dsa/data_structures/hlist"
)

const (
	nilTableString   = "<nil>"
	emptyTableString = "<empty>"
)

var (
	// This is the standard error message when trying to use an invalid table.
	errBadTable = fmt.Errorf("table must be created with New() first")
	// This is the standard error message when trying to use an invalid row.
	errBadRow = fmt.Errorf("row must be created with NewRow() first")
)

// Table is the main type in this package. It holds all the rows of data.
type Table struct {
	headers []string       // Column headers
	types   []reflect.Type // Types of each column (must be consistent for all rows)
	rows    *hlist.List    // Linked list of rows
}

// New creates a new table. headers denotes the name of each column. Each header must be unique and
// not empty.
func New(headers ...string) (*Table, error) {
	if len(headers) == 0 {
		return nil, fmt.Errorf("missing column headers")
	}

	// Validate the column headers.
	headerMap := make(map[string]int)
	for i, header := range headers {
		// Make sure every column has a header.
		if header == "" {
			return nil, fmt.Errorf("column %v has an empty header", i)
		}
		// Make sure none of the columns match each other.
		if c, found := headerMap[header]; found {
			return nil, fmt.Errorf("columns %v and %v have the same header", c+1, i+1)
		}
		headerMap[header] = i
	}

	t := new(Table)
	t.headers = make([]string, len(headers))
	copy(t.headers, headers)
	t.types = make([]reflect.Type, len(headers))
	t.rows = hlist.New()

	return t, nil
}

// Add creates a new row with the items and adds it to the end of the table.
func (t *Table) Add(items ...interface{}) error {
	// Build the row.
	row := NewRow(items...)

	return t.AddRow(row)
}

// Insert creates a new row with the items and inserts it into the table at the specified index.
func (t *Table) Insert(index int, items ...interface{}) error {
	// Build the row.
	row := NewRow(items...)

	return t.InsertRow(index, row)
}

// AddRow adds the row to the end of the table.
func (t *Table) AddRow(r *Row) error {
	// Validate that the overall length and the types of all the items are correct.
	if err := t.validateRow(r); err != nil {
		return err
	}

	// Add the row to the table.
	return t.rows.Append(r)
}

// InsertRow inserts the row into the table at the specified index.
func (t *Table) InsertRow(index int, r *Row) error {
	// Validate that the overall length and the types of all the items are correct.
	if err := t.validateRow(r); err != nil {
		return err
	}

	// Add the row to the table.
	return t.rows.Insert(index, r)
}

// RemoveRow deletes the row at the index from the table.
func (t *Table) RemoveRow(index int) {
	// We don't care if there are any problems with this operation. If the table is good and the row
	// exists, then it will be removed. That's all we really care about.
	if t != nil {
		t.rows.Remove(index)
	}
}

// Clear erases the rows in the table but leaves the column headers and column types.
func (t *Table) Clear() error {
	if t == nil {
		return errBadTable
	}

	return t.rows.Clear()
}

// ColumnToIndex returns the index of the column by header, or -1 if not found.
func (t *Table) ColumnToIndex(header string) int {
	if t == nil {
		return -1
	}

	for i, h := range t.headers {
		if header == h {
			return i
		}
	}

	// If we're here, then we didn't find anything.
	return -1
}

// String returns a formatted list of the items in the table, row by row.
func (t *Table) String() string {
	if t == nil {
		return nilTableString
	} else if t.Count() == 0 {
		return emptyTableString
	}

	b := new(strings.Builder)
	rowChan := t.rows.Yield(nil)
	if rowChan == nil {
		return emptyTableString
	}

	for r := range rowChan {
		row := r.(*Row)
		if row.Enabled() {
			tmp := new(strings.Builder)
			for i, item := range row.items {
				tmp.WriteString(fmt.Sprintf("%v: %v, ", t.headers[i], item))
			}
			s := tmp.String()
			s = strings.TrimSuffix(s, ", ")
			s = strings.Join([]string{"{", s, "}"}, "")
			b.WriteString(fmt.Sprintf("%v, ", s))
		}
	}

	s := b.String()
	if s == "" {
		s = emptyTableString
	}

	return strings.TrimSuffix(s, ", ")
}

// SetItem changes the value of the item at the specified coordinates.
func (t *Table) SetItem(header string, index int, value interface{}) error {
	if t == nil {
		return errBadTable
	} else if index < 0 || t.Rows() <= index {
		return fmt.Errorf("invalid index")
	}

	// Grab our row.
	r := t.rows.Item(index)
	if r == nil {
		return fmt.Errorf("missing row")
	}

	// Figure out the index of the column.
	i := t.ColumnToIndex(header)
	if i < 0 || i >= len(t.headers) {
		return fmt.Errorf("invalid column")
	}

	// Change the value.
	nr := r.(*Row)
	nr.SetItem(i, value)

	// Add the row back in to the table, which will also validate the new value.
	t.RemoveRow(index)
	return t.InsertRow(index, nr)
}

// SetHeader changes the specified column's header.
func (t *Table) SetHeader(oldHeader, newHeader string) error {
	if t == nil {
		return errBadTable
	} else if newHeader == "" {
		return fmt.Errorf("missing new header")
	}

	for i, header := range t.headers {
		if oldHeader == header {
			t.headers[i] = newHeader
			return nil
		}
	}

	// If we're here, then we didn't find the column.
	return fmt.Errorf("invalid header")
}

// Headers returns a copy of the table's column headers.
func (t *Table) Headers() []string {
	if t == nil {
		return nil
	}

	headers := make([]string, len(t.headers))
	copy(headers, t.headers)

	return headers
}

// Copy makes an exact and separate copy of the table. If the table to copy is nil, it returns nil.
func (t *Table) Copy() *Table {
	if t == nil {
		return nil
	}

	// Make a new table and copy over the headers.
	nt, _ := New(t.Headers()...)

	// Copy over the column types. The slice was already made in New.
	copy(nt.types, t.types)

	// Copy over the rows. We'll discard the list already created.
	nt.rows, _ = t.rows.Copy()

	return nt
}

// Rows returns the number of rows in the table, or -1 on error. This includes all rows, regardless
// of enabled status.
func (t *Table) Rows() int {
	if t == nil {
		return -1
	}

	return t.rows.Length()
}

// Enabled returns the number of enabled rows in the table, or -1 on error.
func (t *Table) Enabled() int {
	if t == nil {
		return -1
	}

	rowChan := t.rows.Yield(nil)
	if rowChan == nil {
		return 0
	}

	numEnabled := 0
	for r := range rowChan {
		row := r.(*Row)
		if row.Enabled() {
			numEnabled++
		}
	}

	return numEnabled
}

// Disabled returns the number of disabled rows in the table, or -1 on error.
func (t *Table) Disabled() int {
	if t == nil {
		return -1
	}

	rowChan := t.rows.Yield(nil)
	if rowChan == nil {
		return 0
	}

	numDisabled := 0
	for r := range rowChan {
		row := r.(*Row)
		if !row.Enabled() {
			numDisabled++
		}
	}

	return numDisabled
}

// Columns returns the number of columns in the table, or -1 on error.
func (t *Table) Columns() int {
	if t == nil {
		return -1
	}

	return len(t.headers)
}

// Count returns the number of items in the table, or -1 on error. This includes all items,
// regardless of enabled status.
func (t *Table) Count() int {
	c := t.Columns()
	r := t.Rows()

	if c == -1 || r == -1 {
		return -1
	}

	return c * r
}

// Same checks whether or not the tables point to the same memory.
func (t *Table) Same(nt *Table) bool {
	if t == nil || nt == nil {
		return false
	}

	if t == nt {
		return true
	}

	return false
}

// Row returns the index and Row of the first enabled row that contains the item in the specified
// column, or -1 and nil if not found or error.
func (t *Table) Row(header string, item interface{}) (int, *Row) {
	if t == nil {
		return -1, nil
	}

	// Find out which column we need to match on.
	c := -1
	for i, h := range t.headers {
		if header == h {
			c = i
			break
		}
	}

	// Make sure we found a column.
	if c < 0 {
		return -1, nil
	}

	// Get our iterator to go through the rows.
	i := 0
	quit := make(chan struct{})
	rowChan := t.rows.Yield(quit)
	if rowChan == nil {
		return -1, nil
	}

	for v := range rowChan {
		row := v.(*Row)
		if row.Enabled() {
			if reflect.DeepEqual(item, row.items[c]) {
				// Break out of the list iteration. If Yield's goroutine has already exited (because
				// the list was fully traversed), then it won't receive the message to quit. We'll
				// try to send the quit message, and then we'll exit.
				select {
				case quit <- struct{}{}:
				default:
				}
				return i, row
			}
		}
		i++
	}

	// If we're here, then we didn't find anything.
	return -1, nil
}

// RowByIndex returns the Row at the specified index, or nil if not found or error.
func (t *Table) RowByIndex(index int) *Row {
	if t == nil {
		return nil
	}

	// Grab our row.
	row := t.rows.Item(index)
	if row == nil {
		return nil
	}
	return row.(*Row)
}

// Item returns the item at the specified coordinates, or nil if there is no item at the coordinates.
func (t *Table) Item(header string, index int) interface{} {
	if t == nil {
		return nil
	}

	// Figure out the index of the column.
	i := t.ColumnToIndex(header)
	if i < 0 || i >= len(t.headers) {
		return nil
	}

	// Grab our row.
	r := t.rows.Item(index)
	if r == nil {
		return nil
	}
	row := r.(*Row)

	return row.Item(i)
}

// Matches returns true if the value matches the item at the specified coordinates or false if there
// is no match. Matching can occur on disabled rows.
func (t *Table) Matches(header string, index int, value interface{}) bool {
	item := t.Item(header, index)
	if item == nil {
		return false
	}

	return reflect.DeepEqual(value, item)
}

// Toggle sets the row at the specified index to either be checked or skipped during table lookups
// (like Row and Count).
func (t *Table) Toggle(index int, enabled bool) error {
	if t == nil {
		return errBadTable
	}

	r := t.rows.Item(index)
	if r == nil {
		return fmt.Errorf("invalid index")
	}

	row := r.(*Row)
	row.ToggleRow(enabled)

	return nil
}

// SortByColumn sorts the table on the specified column. The comparison function less is given the
// values of two different items in a column and should return true only if the left item should be
// sorted before the right item.
func (t *Table) SortByColumn(header string, less func(interface{}, interface{}) bool) error {
	if t == nil {
		return errBadTable
	} else if less == nil {
		return fmt.Errorf("missing comparison callback")
	}

	// Figure out the index of the column.
	i := t.ColumnToIndex(header)
	if i < 0 || i > len(t.headers) {
		return fmt.Errorf("invalid column")
	}

	// Sort the rows.
	err := t.rows.Sort(func(left, right interface{}) bool {
		leftRow := left.(*Row)
		rightRow := right.(*Row)
		return less(leftRow.Item(i), rightRow.Item(i))
	})

	return err
}

// SortByRow sorts the table by rows. The comparison function less is given pointers to two Row
// objects and should return true only if the left row should be sorted before the right row.
func (t *Table) SortByRow(less func(*Row, *Row) bool) error {
	if t == nil {
		return errBadTable
	} else if less == nil {
		return fmt.Errorf("missing comparison callback")
	}

	// Sort the rows.
	err := t.rows.Sort(func(left, right interface{}) bool {
		leftRow := left.(*Row)
		rightRow := right.(*Row)

		// Return copies of the rows to prevent changes from affecting the table.
		leftCopy := leftRow.Copy()
		rightCopy := rightRow.Copy()

		return less(leftCopy, rightCopy)
	})

	return err
}

// CSV returns a representation of the table as rows of comma-separated values, with each row
// delineated by \r\n newlines.
func (t *Table) CSV() string {
	if t == nil {
		return ""
	}

	b := new(strings.Builder)

	// Add the headers.
	b.WriteString(strings.Join(t.Headers(), ","))

	// Add the rows.
	rows := t.rows.Items()
	if rows == nil {
		return b.String()
	}
	for _, r := range rows {
		row := r.(*Row)
		if row.Enabled() {
			items := row.Items()
			itemValues := make([]string, len(items))
			for i, item := range items {
				itemValues[i] = fmt.Sprintf("%v", item)
			}
			b.WriteString("\r\n")
			b.WriteString(strings.Join(itemValues, ","))
		}
	}

	return b.String()
}

// Row holds all the data for each row in the table.
type Row struct {
	items   []interface{}
	enabled bool
}

// NewRow creates a new row with the given items.
func NewRow(items ...interface{}) *Row {
	if len(items) == 0 {
		return nil
	}

	r := new(Row)

	// Add the items.
	r.items = make([]interface{}, len(items))
	copy(r.items, items)

	// Default to enabling this row.
	r.enabled = true

	return r
}

// String returns a formatted list of the items in the row.
func (r *Row) String() string {
	if r == nil {
		return nilTableString
	} else if len(r.items) == 0 {
		return emptyTableString
	}

	b := new(strings.Builder)
	for _, item := range r.items {
		b.WriteString(fmt.Sprintf("%v, ", item))
	}

	s := strings.TrimSuffix(b.String(), ", ")

	return strings.Join([]string{"{", s, "}"}, "")
}

// SetItem changes the value of the item in the specified column index.
func (r *Row) SetItem(index int, value interface{}) error {
	if r == nil {
		return errBadRow
	} else if index < 0 || index >= r.Count() {
		return fmt.Errorf("invalid column")
	}

	r.items[index] = value

	return nil
}

// Count returns the number of items in the row, or -1 on error.
func (r *Row) Count() int {
	if r == nil {
		return -1
	}

	return len(r.items)
}

// Item returns the item's value at the specified index in this row, or nil if not found or error.
func (r *Row) Item(index int) interface{} {
	if r == nil || index < 0 || index >= len(r.items) {
		return nil
	}

	return r.items[index]
}

// Items returns a copy of the items in the row.
func (r *Row) Items() []interface{} {
	if r == nil {
		return nil
	}

	items := make([]interface{}, len(r.items))
	copy(items, r.items)

	return items
}

// Enabled returns true if the row is enabled, otherwise false.
func (r *Row) Enabled() bool {
	if r == nil {
		return false
	}

	return r.enabled
}

// ToggleRow sets the row as either enabled or disabled.
func (r *Row) ToggleRow(enabled bool) {
	if r != nil {
		r.enabled = enabled
	}
}

// Matches returns true if the value matches the item in the specified column or false if there is
// no match. Matching can occur on disabled rows.
func (r *Row) Matches(index int, value interface{}) bool {
	item := r.Item(index)
	if item == nil {
		return false
	}

	return reflect.DeepEqual(value, item)
}

// Copy makes a complete and separate copy of the row. If the row to copy is nil, it returns nil.
func (r *Row) Copy() *Row {
	if r == nil {
		return nil
	}

	nr := NewRow(r.Items()...)
	nr.ToggleRow(r.Enabled())

	return nr
}

func (t *Table) validateRow(r *Row) error {
	if t == nil {
		return errBadTable
	} else if n := t.Columns(); n != len(r.items) {
		return fmt.Errorf("number of items (%v) does not match number of columns (%v)", len(r.items), n)
	}

	// Validate the items.
	for i, item := range r.items {
		// Make sure that none of the items is this table itself.
		if nt, ok := item.(*Table); ok {
			if t.Same(nt) {
				return fmt.Errorf("can't add table to itself")
			}
		}

		// Validate the types.
		typeof := reflect.TypeOf(item)
		if t.Rows() == 0 {
			// This is the first row being added to the table. It will set the type of each column
			// in the table.
			t.types[i] = typeof
		} else if typeof != t.types[i] {
			return fmt.Errorf("item %v's type (%v) does not match column's prototype (%v)", i, typeof, t.types[i])
		}
	}

	return nil
}
