// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package entv2

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/entc/integration/migrate/entv2/pet"
	"entgo.io/ent/entc/integration/migrate/entv2/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// MixedString holds the value of the "mixed_string" field.
	MixedString string `json:"mixed_string,omitempty"`
	// MixedEnum holds the value of the "mixed_enum" field.
	MixedEnum user.MixedEnum `json:"mixed_enum,omitempty"`
	// Active holds the value of the "active" field.
	Active bool `json:"active,omitempty"`
	// Age holds the value of the "age" field.
	Age int `json:"age,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Nickname holds the value of the "nickname" field.
	Nickname string `json:"nickname,omitempty"`
	// Phone holds the value of the "phone" field.
	Phone string `json:"phone,omitempty"`
	// Buffer holds the value of the "buffer" field.
	Buffer []byte `json:"buffer,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// NewName holds the value of the "new_name" field.
	NewName string `json:"new_name,omitempty"`
	// NewToken holds the value of the "new_token" field.
	NewToken string `json:"new_token,omitempty"`
	// Blob holds the value of the "blob" field.
	Blob []byte `json:"blob,omitempty"`
	// State holds the value of the "state" field.
	State user.State `json:"state,omitempty"`
	// Status holds the value of the "status" field.
	Status user.Status `json:"status,omitempty"`
	// Workplace holds the value of the "workplace" field.
	Workplace string `json:"workplace,omitempty"`
	// Roles holds the value of the "roles" field.
	Roles []string `json:"roles,omitempty"`
	// DefaultExpr holds the value of the "default_expr" field.
	DefaultExpr string `json:"default_expr,omitempty"`
	// DefaultExprs holds the value of the "default_exprs" field.
	DefaultExprs string `json:"default_exprs,omitempty"`
	// CreatedAt holds the value of the "created_at" field.
	CreatedAt time.Time `json:"created_at,omitempty"`
	// DropOptional holds the value of the "drop_optional" field.
	DropOptional string `json:"drop_optional,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges       UserEdges `json:"edges"`
	blog_admins *int
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Car holds the value of the car edge.
	Car []*Car `json:"car,omitempty"`
	// Pets holds the value of the pets edge.
	Pets *Pet `json:"pets,omitempty"`
	// Friends holds the value of the friends edge.
	Friends []*User `json:"friends,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// CarOrErr returns the Car value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CarOrErr() ([]*Car, error) {
	if e.loadedTypes[0] {
		return e.Car, nil
	}
	return nil, &NotLoadedError{edge: "car"}
}

// PetsOrErr returns the Pets value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) PetsOrErr() (*Pet, error) {
	if e.loadedTypes[1] {
		if e.Pets == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: pet.Label}
		}
		return e.Pets, nil
	}
	return nil, &NotLoadedError{edge: "pets"}
}

// FriendsOrErr returns the Friends value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) FriendsOrErr() ([]*User, error) {
	if e.loadedTypes[2] {
		return e.Friends, nil
	}
	return nil, &NotLoadedError{edge: "friends"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldBuffer, user.FieldBlob, user.FieldRoles:
			values[i] = new([]byte)
		case user.FieldActive:
			values[i] = new(sql.NullBool)
		case user.FieldID, user.FieldAge:
			values[i] = new(sql.NullInt64)
		case user.FieldMixedString, user.FieldMixedEnum, user.FieldName, user.FieldDescription, user.FieldNickname, user.FieldPhone, user.FieldTitle, user.FieldNewName, user.FieldNewToken, user.FieldState, user.FieldStatus, user.FieldWorkplace, user.FieldDefaultExpr, user.FieldDefaultExprs, user.FieldDropOptional:
			values[i] = new(sql.NullString)
		case user.FieldCreatedAt:
			values[i] = new(sql.NullTime)
		case user.ForeignKeys[0]: // blog_admins
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldMixedString:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mixed_string", values[i])
			} else if value.Valid {
				u.MixedString = value.String
			}
		case user.FieldMixedEnum:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field mixed_enum", values[i])
			} else if value.Valid {
				u.MixedEnum = user.MixedEnum(value.String)
			}
		case user.FieldActive:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field active", values[i])
			} else if value.Valid {
				u.Active = value.Bool
			}
		case user.FieldAge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age", values[i])
			} else if value.Valid {
				u.Age = int(value.Int64)
			}
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				u.Description = value.String
			}
		case user.FieldNickname:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field nickname", values[i])
			} else if value.Valid {
				u.Nickname = value.String
			}
		case user.FieldPhone:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field phone", values[i])
			} else if value.Valid {
				u.Phone = value.String
			}
		case user.FieldBuffer:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field buffer", values[i])
			} else if value != nil {
				u.Buffer = *value
			}
		case user.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				u.Title = value.String
			}
		case user.FieldNewName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field new_name", values[i])
			} else if value.Valid {
				u.NewName = value.String
			}
		case user.FieldNewToken:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field new_token", values[i])
			} else if value.Valid {
				u.NewToken = value.String
			}
		case user.FieldBlob:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field blob", values[i])
			} else if value != nil {
				u.Blob = *value
			}
		case user.FieldState:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field state", values[i])
			} else if value.Valid {
				u.State = user.State(value.String)
			}
		case user.FieldStatus:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field status", values[i])
			} else if value.Valid {
				u.Status = user.Status(value.String)
			}
		case user.FieldWorkplace:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field workplace", values[i])
			} else if value.Valid {
				u.Workplace = value.String
			}
		case user.FieldRoles:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field roles", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &u.Roles); err != nil {
					return fmt.Errorf("unmarshal field roles: %w", err)
				}
			}
		case user.FieldDefaultExpr:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field default_expr", values[i])
			} else if value.Valid {
				u.DefaultExpr = value.String
			}
		case user.FieldDefaultExprs:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field default_exprs", values[i])
			} else if value.Valid {
				u.DefaultExprs = value.String
			}
		case user.FieldCreatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field created_at", values[i])
			} else if value.Valid {
				u.CreatedAt = value.Time
			}
		case user.FieldDropOptional:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field drop_optional", values[i])
			} else if value.Valid {
				u.DropOptional = value.String
			}
		case user.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field blog_admins", value)
			} else if value.Valid {
				u.blog_admins = new(int)
				*u.blog_admins = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryCar queries the "car" edge of the User entity.
func (u *User) QueryCar() *CarQuery {
	return (&UserClient{config: u.config}).QueryCar(u)
}

// QueryPets queries the "pets" edge of the User entity.
func (u *User) QueryPets() *PetQuery {
	return (&UserClient{config: u.config}).QueryPets(u)
}

// QueryFriends queries the "friends" edge of the User entity.
func (u *User) QueryFriends() *UserQuery {
	return (&UserClient{config: u.config}).QueryFriends(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	_tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("entv2: User is not a transactional entity")
	}
	u.config.driver = _tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v, ", u.ID))
	builder.WriteString("mixed_string=")
	builder.WriteString(u.MixedString)
	builder.WriteString(", ")
	builder.WriteString("mixed_enum=")
	builder.WriteString(fmt.Sprintf("%v", u.MixedEnum))
	builder.WriteString(", ")
	builder.WriteString("active=")
	builder.WriteString(fmt.Sprintf("%v", u.Active))
	builder.WriteString(", ")
	builder.WriteString("age=")
	builder.WriteString(fmt.Sprintf("%v", u.Age))
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(u.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(u.Description)
	builder.WriteString(", ")
	builder.WriteString("nickname=")
	builder.WriteString(u.Nickname)
	builder.WriteString(", ")
	builder.WriteString("phone=")
	builder.WriteString(u.Phone)
	builder.WriteString(", ")
	builder.WriteString("buffer=")
	builder.WriteString(fmt.Sprintf("%v", u.Buffer))
	builder.WriteString(", ")
	builder.WriteString("title=")
	builder.WriteString(u.Title)
	builder.WriteString(", ")
	builder.WriteString("new_name=")
	builder.WriteString(u.NewName)
	builder.WriteString(", ")
	builder.WriteString("new_token=")
	builder.WriteString(u.NewToken)
	builder.WriteString(", ")
	builder.WriteString("blob=")
	builder.WriteString(fmt.Sprintf("%v", u.Blob))
	builder.WriteString(", ")
	builder.WriteString("state=")
	builder.WriteString(fmt.Sprintf("%v", u.State))
	builder.WriteString(", ")
	builder.WriteString("status=")
	builder.WriteString(fmt.Sprintf("%v", u.Status))
	builder.WriteString(", ")
	builder.WriteString("workplace=")
	builder.WriteString(u.Workplace)
	builder.WriteString(", ")
	builder.WriteString("roles=")
	builder.WriteString(fmt.Sprintf("%v", u.Roles))
	builder.WriteString(", ")
	builder.WriteString("default_expr=")
	builder.WriteString(u.DefaultExpr)
	builder.WriteString(", ")
	builder.WriteString("default_exprs=")
	builder.WriteString(u.DefaultExprs)
	builder.WriteString(", ")
	builder.WriteString("created_at=")
	builder.WriteString(u.CreatedAt.Format(time.ANSIC))
	builder.WriteString(", ")
	builder.WriteString("drop_optional=")
	builder.WriteString(u.DropOptional)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
