// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/GRTheory/my-model/ent/student"
	"github.com/GRTheory/my-model/ent/teacher"
)

// Student is the model entity for the Student schema.
type Student struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// SID holds the value of the "s_id" field.
	SID int `json:"s_id,omitempty"`
	// Naem holds the value of the "naem" field.
	Naem string `json:"naem,omitempty"`
	// Age holds the value of the "age" field.
	Age *int `json:"age,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the StudentQuery when eager-loading is set.
	Edges            StudentEdges `json:"edges"`
	teacher_students *int
}

// StudentEdges holds the relations/edges for other nodes in the graph.
type StudentEdges struct {
	// Teachers holds the value of the teachers edge.
	Teachers *Teacher `json:"teachers,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// TeachersOrErr returns the Teachers value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e StudentEdges) TeachersOrErr() (*Teacher, error) {
	if e.loadedTypes[0] {
		if e.Teachers == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: teacher.Label}
		}
		return e.Teachers, nil
	}
	return nil, &NotLoadedError{edge: "teachers"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Student) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case student.FieldID, student.FieldSID, student.FieldAge:
			values[i] = new(sql.NullInt64)
		case student.FieldNaem:
			values[i] = new(sql.NullString)
		case student.ForeignKeys[0]: // teacher_students
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Student", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Student fields.
func (s *Student) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case student.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			s.ID = int(value.Int64)
		case student.FieldSID:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field s_id", values[i])
			} else if value.Valid {
				s.SID = int(value.Int64)
			}
		case student.FieldNaem:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field naem", values[i])
			} else if value.Valid {
				s.Naem = value.String
			}
		case student.FieldAge:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field age", values[i])
			} else if value.Valid {
				s.Age = new(int)
				*s.Age = int(value.Int64)
			}
		case student.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field teacher_students", value)
			} else if value.Valid {
				s.teacher_students = new(int)
				*s.teacher_students = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryTeachers queries the "teachers" edge of the Student entity.
func (s *Student) QueryTeachers() *TeacherQuery {
	return NewStudentClient(s.config).QueryTeachers(s)
}

// Update returns a builder for updating this Student.
// Note that you need to call Student.Unwrap() before calling this method if this Student
// was returned from a transaction, and the transaction was committed or rolled back.
func (s *Student) Update() *StudentUpdateOne {
	return NewStudentClient(s.config).UpdateOne(s)
}

// Unwrap unwraps the Student entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (s *Student) Unwrap() *Student {
	_tx, ok := s.config.driver.(*txDriver)
	if !ok {
		panic("ent: Student is not a transactional entity")
	}
	s.config.driver = _tx.drv
	return s
}

// String implements the fmt.Stringer.
func (s *Student) String() string {
	var builder strings.Builder
	builder.WriteString("Student(")
	builder.WriteString(fmt.Sprintf("id=%v, ", s.ID))
	builder.WriteString("s_id=")
	builder.WriteString(fmt.Sprintf("%v", s.SID))
	builder.WriteString(", ")
	builder.WriteString("naem=")
	builder.WriteString(s.Naem)
	builder.WriteString(", ")
	if v := s.Age; v != nil {
		builder.WriteString("age=")
		builder.WriteString(fmt.Sprintf("%v", *v))
	}
	builder.WriteByte(')')
	return builder.String()
}

// Students is a parsable slice of Student.
type Students []*Student

func (s Students) config(cfg config) {
	for _i := range s {
		s[_i].config = cfg
	}
}