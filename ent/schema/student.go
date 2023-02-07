package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// Student holds the schema definition for the Student entity.
type Student struct {
	ent.Schema
}

// Fields of the Student.
func (Student) Fields() []ent.Field {
	return []ent.Field{
		field.Int("s_id").Unique(),
		field.String("naem").Default("unknown"),
		field.Int("age").Nillable(),
	}
}

// Edges of the Student. This function defines a relations which 
// is a back-reference between one teacher and many students.
func (Student) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("teachers", Teacher.Type).
			Ref("students").
			Unique().
			Required(),
	}
}
