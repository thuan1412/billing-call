// Code generated by ent, DO NOT EDIT.

package ent

import (
	"calling-bill/ent/call"
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Call is the model entity for the Call schema.
type Call struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Duration holds the value of the "duration" field.
	Duration int `json:"duration,omitempty"`
	// BlockCount holds the value of the "block_count" field.
	BlockCount int `json:"block_count,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the CallQuery when eager-loading is set.
	Edges CallEdges `json:"edges"`
}

// CallEdges holds the relations/edges for other nodes in the graph.
type CallEdges struct {
	// User holds the value of the user edge.
	User []*User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading.
func (e CallEdges) UserOrErr() ([]*User, error) {
	if e.loadedTypes[0] {
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Call) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case call.FieldID, call.FieldDuration, call.FieldBlockCount:
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Call", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Call fields.
func (c *Call) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case call.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			c.ID = int(value.Int64)
		case call.FieldDuration:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field duration", values[i])
			} else if value.Valid {
				c.Duration = int(value.Int64)
			}
		case call.FieldBlockCount:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field block_count", values[i])
			} else if value.Valid {
				c.BlockCount = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Call entity.
func (c *Call) QueryUser() *UserQuery {
	return (&CallClient{config: c.config}).QueryUser(c)
}

// Update returns a builder for updating this Call.
// Note that you need to call Call.Unwrap() before calling this method if this Call
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Call) Update() *CallUpdateOne {
	return (&CallClient{config: c.config}).UpdateOne(c)
}

// Unwrap unwraps the Call entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Call) Unwrap() *Call {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Call is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Call) String() string {
	var builder strings.Builder
	builder.WriteString("Call(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("duration=")
	builder.WriteString(fmt.Sprintf("%v", c.Duration))
	builder.WriteString(", ")
	builder.WriteString("block_count=")
	builder.WriteString(fmt.Sprintf("%v", c.BlockCount))
	builder.WriteByte(')')
	return builder.String()
}

// Calls is a parsable slice of Call.
type Calls []*Call

func (c Calls) config(cfg config) {
	for _i := range c {
		c[_i].config = cfg
	}
}
