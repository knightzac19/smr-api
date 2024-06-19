// Code generated by ent, DO NOT EDIT.

package usermod

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/satisfactorymodding/smr-api/generated/ent/predicate"
)

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldEQ(FieldUserID, v))
}

// ModID applies equality check predicate on the "mod_id" field. It's identical to ModIDEQ.
func ModID(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldEQ(FieldModID, v))
}

// Role applies equality check predicate on the "role" field. It's identical to RoleEQ.
func Role(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldEQ(FieldRole, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...string) predicate.UserMod {
	return predicate.UserMod(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...string) predicate.UserMod {
	return predicate.UserMod(sql.FieldNotIn(FieldUserID, vs...))
}

// UserIDGT applies the GT predicate on the "user_id" field.
func UserIDGT(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldGT(FieldUserID, v))
}

// UserIDGTE applies the GTE predicate on the "user_id" field.
func UserIDGTE(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldGTE(FieldUserID, v))
}

// UserIDLT applies the LT predicate on the "user_id" field.
func UserIDLT(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldLT(FieldUserID, v))
}

// UserIDLTE applies the LTE predicate on the "user_id" field.
func UserIDLTE(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldLTE(FieldUserID, v))
}

// UserIDContains applies the Contains predicate on the "user_id" field.
func UserIDContains(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldContains(FieldUserID, v))
}

// UserIDHasPrefix applies the HasPrefix predicate on the "user_id" field.
func UserIDHasPrefix(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldHasPrefix(FieldUserID, v))
}

// UserIDHasSuffix applies the HasSuffix predicate on the "user_id" field.
func UserIDHasSuffix(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldHasSuffix(FieldUserID, v))
}

// UserIDEqualFold applies the EqualFold predicate on the "user_id" field.
func UserIDEqualFold(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldEqualFold(FieldUserID, v))
}

// UserIDContainsFold applies the ContainsFold predicate on the "user_id" field.
func UserIDContainsFold(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldContainsFold(FieldUserID, v))
}

// ModIDEQ applies the EQ predicate on the "mod_id" field.
func ModIDEQ(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldEQ(FieldModID, v))
}

// ModIDNEQ applies the NEQ predicate on the "mod_id" field.
func ModIDNEQ(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldNEQ(FieldModID, v))
}

// ModIDIn applies the In predicate on the "mod_id" field.
func ModIDIn(vs ...string) predicate.UserMod {
	return predicate.UserMod(sql.FieldIn(FieldModID, vs...))
}

// ModIDNotIn applies the NotIn predicate on the "mod_id" field.
func ModIDNotIn(vs ...string) predicate.UserMod {
	return predicate.UserMod(sql.FieldNotIn(FieldModID, vs...))
}

// ModIDGT applies the GT predicate on the "mod_id" field.
func ModIDGT(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldGT(FieldModID, v))
}

// ModIDGTE applies the GTE predicate on the "mod_id" field.
func ModIDGTE(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldGTE(FieldModID, v))
}

// ModIDLT applies the LT predicate on the "mod_id" field.
func ModIDLT(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldLT(FieldModID, v))
}

// ModIDLTE applies the LTE predicate on the "mod_id" field.
func ModIDLTE(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldLTE(FieldModID, v))
}

// ModIDContains applies the Contains predicate on the "mod_id" field.
func ModIDContains(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldContains(FieldModID, v))
}

// ModIDHasPrefix applies the HasPrefix predicate on the "mod_id" field.
func ModIDHasPrefix(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldHasPrefix(FieldModID, v))
}

// ModIDHasSuffix applies the HasSuffix predicate on the "mod_id" field.
func ModIDHasSuffix(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldHasSuffix(FieldModID, v))
}

// ModIDEqualFold applies the EqualFold predicate on the "mod_id" field.
func ModIDEqualFold(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldEqualFold(FieldModID, v))
}

// ModIDContainsFold applies the ContainsFold predicate on the "mod_id" field.
func ModIDContainsFold(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldContainsFold(FieldModID, v))
}

// RoleEQ applies the EQ predicate on the "role" field.
func RoleEQ(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldEQ(FieldRole, v))
}

// RoleNEQ applies the NEQ predicate on the "role" field.
func RoleNEQ(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldNEQ(FieldRole, v))
}

// RoleIn applies the In predicate on the "role" field.
func RoleIn(vs ...string) predicate.UserMod {
	return predicate.UserMod(sql.FieldIn(FieldRole, vs...))
}

// RoleNotIn applies the NotIn predicate on the "role" field.
func RoleNotIn(vs ...string) predicate.UserMod {
	return predicate.UserMod(sql.FieldNotIn(FieldRole, vs...))
}

// RoleGT applies the GT predicate on the "role" field.
func RoleGT(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldGT(FieldRole, v))
}

// RoleGTE applies the GTE predicate on the "role" field.
func RoleGTE(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldGTE(FieldRole, v))
}

// RoleLT applies the LT predicate on the "role" field.
func RoleLT(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldLT(FieldRole, v))
}

// RoleLTE applies the LTE predicate on the "role" field.
func RoleLTE(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldLTE(FieldRole, v))
}

// RoleContains applies the Contains predicate on the "role" field.
func RoleContains(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldContains(FieldRole, v))
}

// RoleHasPrefix applies the HasPrefix predicate on the "role" field.
func RoleHasPrefix(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldHasPrefix(FieldRole, v))
}

// RoleHasSuffix applies the HasSuffix predicate on the "role" field.
func RoleHasSuffix(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldHasSuffix(FieldRole, v))
}

// RoleEqualFold applies the EqualFold predicate on the "role" field.
func RoleEqualFold(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldEqualFold(FieldRole, v))
}

// RoleContainsFold applies the ContainsFold predicate on the "role" field.
func RoleContainsFold(v string) predicate.UserMod {
	return predicate.UserMod(sql.FieldContainsFold(FieldRole, v))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.UserMod {
	return predicate.UserMod(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, UserColumn),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.UserMod {
	return predicate.UserMod(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasMod applies the HasEdge predicate on the "mod" edge.
func HasMod() predicate.UserMod {
	return predicate.UserMod(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, ModColumn),
			sqlgraph.Edge(sqlgraph.M2O, false, ModTable, ModColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasModWith applies the HasEdge predicate on the "mod" edge with a given conditions (other predicates).
func HasModWith(preds ...predicate.Mod) predicate.UserMod {
	return predicate.UserMod(func(s *sql.Selector) {
		step := newModStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.UserMod) predicate.UserMod {
	return predicate.UserMod(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.UserMod) predicate.UserMod {
	return predicate.UserMod(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.UserMod) predicate.UserMod {
	return predicate.UserMod(sql.NotPredicates(p))
}