// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// StudentsColumns holds the columns for the "students" table.
	StudentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "s_id", Type: field.TypeInt, Unique: true},
		{Name: "naem", Type: field.TypeString, Default: "unknown"},
		{Name: "age", Type: field.TypeInt},
		{Name: "teacher_students", Type: field.TypeInt},
	}
	// StudentsTable holds the schema information for the "students" table.
	StudentsTable = &schema.Table{
		Name:       "students",
		Columns:    StudentsColumns,
		PrimaryKey: []*schema.Column{StudentsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "students_teachers_students",
				Columns:    []*schema.Column{StudentsColumns[4]},
				RefColumns: []*schema.Column{TeachersColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TeachersColumns holds the columns for the "teachers" table.
	TeachersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "t_id", Type: field.TypeInt, Unique: true},
		{Name: "name", Type: field.TypeString, Default: "unknown"},
		{Name: "age", Type: field.TypeInt},
	}
	// TeachersTable holds the schema information for the "teachers" table.
	TeachersTable = &schema.Table{
		Name:       "teachers",
		Columns:    TeachersColumns,
		PrimaryKey: []*schema.Column{TeachersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		StudentsTable,
		TeachersTable,
	}
)

func init() {
	StudentsTable.ForeignKeys[0].RefTable = TeachersTable
}
