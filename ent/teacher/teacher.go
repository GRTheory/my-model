// Code generated by ent, DO NOT EDIT.

package teacher

const (
	// Label holds the string label denoting the teacher type in the database.
	Label = "teacher"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldTID holds the string denoting the t_id field in the database.
	FieldTID = "t_id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// EdgeStudents holds the string denoting the students edge name in mutations.
	EdgeStudents = "students"
	// Table holds the table name of the teacher in the database.
	Table = "teachers"
	// StudentsTable is the table that holds the students relation/edge.
	StudentsTable = "students"
	// StudentsInverseTable is the table name for the Student entity.
	// It exists in this package in order to avoid circular dependency with the "student" package.
	StudentsInverseTable = "students"
	// StudentsColumn is the table column denoting the students relation/edge.
	StudentsColumn = "teacher_students"
)

// Columns holds all SQL columns for teacher fields.
var Columns = []string{
	FieldID,
	FieldTID,
	FieldName,
	FieldAge,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultName holds the default value on creation for the "name" field.
	DefaultName string
)
