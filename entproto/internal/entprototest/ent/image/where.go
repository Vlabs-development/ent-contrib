// Code generated by ent, DO NOT EDIT.

package image

import (
	"entgo.io/contrib/entproto/internal/entprototest/ent/predicate"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.Image {
	return predicate.Image(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.Image {
	return predicate.Image(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.Image {
	return predicate.Image(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.Image {
	return predicate.Image(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.Image {
	return predicate.Image(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.Image {
	return predicate.Image(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.Image {
	return predicate.Image(sql.FieldLTE(FieldID, id))
}

// URLPath applies equality check predicate on the "url_path" field. It's identical to URLPathEQ.
func URLPath(v string) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldURLPath, v))
}

// URLPathEQ applies the EQ predicate on the "url_path" field.
func URLPathEQ(v string) predicate.Image {
	return predicate.Image(sql.FieldEQ(FieldURLPath, v))
}

// URLPathNEQ applies the NEQ predicate on the "url_path" field.
func URLPathNEQ(v string) predicate.Image {
	return predicate.Image(sql.FieldNEQ(FieldURLPath, v))
}

// URLPathIn applies the In predicate on the "url_path" field.
func URLPathIn(vs ...string) predicate.Image {
	return predicate.Image(sql.FieldIn(FieldURLPath, vs...))
}

// URLPathNotIn applies the NotIn predicate on the "url_path" field.
func URLPathNotIn(vs ...string) predicate.Image {
	return predicate.Image(sql.FieldNotIn(FieldURLPath, vs...))
}

// URLPathGT applies the GT predicate on the "url_path" field.
func URLPathGT(v string) predicate.Image {
	return predicate.Image(sql.FieldGT(FieldURLPath, v))
}

// URLPathGTE applies the GTE predicate on the "url_path" field.
func URLPathGTE(v string) predicate.Image {
	return predicate.Image(sql.FieldGTE(FieldURLPath, v))
}

// URLPathLT applies the LT predicate on the "url_path" field.
func URLPathLT(v string) predicate.Image {
	return predicate.Image(sql.FieldLT(FieldURLPath, v))
}

// URLPathLTE applies the LTE predicate on the "url_path" field.
func URLPathLTE(v string) predicate.Image {
	return predicate.Image(sql.FieldLTE(FieldURLPath, v))
}

// URLPathContains applies the Contains predicate on the "url_path" field.
func URLPathContains(v string) predicate.Image {
	return predicate.Image(sql.FieldContains(FieldURLPath, v))
}

// URLPathHasPrefix applies the HasPrefix predicate on the "url_path" field.
func URLPathHasPrefix(v string) predicate.Image {
	return predicate.Image(sql.FieldHasPrefix(FieldURLPath, v))
}

// URLPathHasSuffix applies the HasSuffix predicate on the "url_path" field.
func URLPathHasSuffix(v string) predicate.Image {
	return predicate.Image(sql.FieldHasSuffix(FieldURLPath, v))
}

// URLPathEqualFold applies the EqualFold predicate on the "url_path" field.
func URLPathEqualFold(v string) predicate.Image {
	return predicate.Image(sql.FieldEqualFold(FieldURLPath, v))
}

// URLPathContainsFold applies the ContainsFold predicate on the "url_path" field.
func URLPathContainsFold(v string) predicate.Image {
	return predicate.Image(sql.FieldContainsFold(FieldURLPath, v))
}

// HasUserProfilePic applies the HasEdge predicate on the "user_profile_pic" edge.
func HasUserProfilePic() predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, UserProfilePicTable, UserProfilePicColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserProfilePicWith applies the HasEdge predicate on the "user_profile_pic" edge with a given conditions (other predicates).
func HasUserProfilePicWith(preds ...predicate.User) predicate.Image {
	return predicate.Image(func(s *sql.Selector) {
		step := newUserProfilePicStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Image) predicate.Image {
	return predicate.Image(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Image) predicate.Image {
	return predicate.Image(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Image) predicate.Image {
	return predicate.Image(sql.NotPredicates(p))
}
