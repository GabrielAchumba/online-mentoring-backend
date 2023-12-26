// Code generated by ent, DO NOT EDIT.

package entgenerated

import (
	"api/ent/entgenerated/loginsession"
	"api/ent/entgenerated/user"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
)

// LoginSession is the model entity for the LoginSession schema.
type LoginSession struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// OwnerID holds the value of the "owner_id" field.
	OwnerID int `json:"owner_id,omitempty"`
	// LastLoginTime holds the value of the "last_login_time" field.
	LastLoginTime time.Time `json:"last_login_time,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the LoginSessionQuery when eager-loading is set.
	Edges        LoginSessionEdges `json:"edges"`
	selectValues sql.SelectValues
}

// LoginSessionEdges holds the relations/edges for other nodes in the graph.
type LoginSessionEdges struct {
	// Owner holds the value of the owner edge.
	Owner *User `json:"owner,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
	// totalCount holds the count of the edges above.
	totalCount [1]map[string]int
}

// OwnerOrErr returns the Owner value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e LoginSessionEdges) OwnerOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.Owner == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Owner, nil
	}
	return nil, &NotLoadedError{edge: "owner"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*LoginSession) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case loginsession.FieldID, loginsession.FieldOwnerID:
			values[i] = new(sql.NullInt64)
		case loginsession.FieldLastLoginTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the LoginSession fields.
func (ls *LoginSession) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case loginsession.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			ls.ID = int(value.Int64)
		case loginsession.FieldOwnerID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field owner_id", values[i])
			} else if value.Valid {
				ls.OwnerID = int(value.Int64)
			}
		case loginsession.FieldLastLoginTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field last_login_time", values[i])
			} else if value.Valid {
				ls.LastLoginTime = value.Time
			}
		default:
			ls.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the LoginSession.
// This includes values selected through modifiers, order, etc.
func (ls *LoginSession) Value(name string) (ent.Value, error) {
	return ls.selectValues.Get(name)
}

// QueryOwner queries the "owner" edge of the LoginSession entity.
func (ls *LoginSession) QueryOwner() *UserQuery {
	return NewLoginSessionClient(ls.config).QueryOwner(ls)
}

// Update returns a builder for updating this LoginSession.
// Note that you need to call LoginSession.Unwrap() before calling this method if this LoginSession
// was returned from a transaction, and the transaction was committed or rolled back.
func (ls *LoginSession) Update() *LoginSessionUpdateOne {
	return NewLoginSessionClient(ls.config).UpdateOne(ls)
}

// Unwrap unwraps the LoginSession entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ls *LoginSession) Unwrap() *LoginSession {
	_tx, ok := ls.config.driver.(*txDriver)
	if !ok {
		panic("entgenerated: LoginSession is not a transactional entity")
	}
	ls.config.driver = _tx.drv
	return ls
}

// String implements the fmt.Stringer.
func (ls *LoginSession) String() string {
	var builder strings.Builder
	builder.WriteString("LoginSession(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ls.ID))
	builder.WriteString("owner_id=")
	builder.WriteString(fmt.Sprintf("%v", ls.OwnerID))
	builder.WriteString(", ")
	builder.WriteString("last_login_time=")
	builder.WriteString(ls.LastLoginTime.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// LoginSessions is a parsable slice of LoginSession.
type LoginSessions []*LoginSession
