// Code generated by ent, DO NOT EDIT.

package ent

import (
	"github.com/GRTheory/my-model/ent/schema"
	"github.com/GRTheory/my-model/ent/student"
	"github.com/GRTheory/my-model/ent/teacher"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	studentFields := schema.Student{}.Fields()
	_ = studentFields
	// studentDescNaem is the schema descriptor for naem field.
	studentDescNaem := studentFields[1].Descriptor()
	// student.DefaultNaem holds the default value on creation for the naem field.
	student.DefaultNaem = studentDescNaem.Default.(string)
	teacherFields := schema.Teacher{}.Fields()
	_ = teacherFields
	// teacherDescName is the schema descriptor for name field.
	teacherDescName := teacherFields[1].Descriptor()
	// teacher.DefaultName holds the default value on creation for the name field.
	teacher.DefaultName = teacherDescName.Default.(string)
}
