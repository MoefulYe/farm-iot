// Code generated by ent, DO NOT EDIT.

package user

import (
	"entgo.io/ent/dialect/sql"
	"github.com/MoefulYe/farm-iot/database/postgres/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.User {
	return predicate.User(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.User {
	return predicate.User(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.User {
	return predicate.User(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.User {
	return predicate.User(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.User {
	return predicate.User(sql.FieldLTE(FieldID, id))
}

// Username applies equality check predicate on the "username" field. It's identical to UsernameEQ.
func Username(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUsername, v))
}

// Passwd applies equality check predicate on the "passwd" field. It's identical to PasswdEQ.
func Passwd(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPasswd, v))
}

// UsernameEQ applies the EQ predicate on the "username" field.
func UsernameEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldUsername, v))
}

// UsernameNEQ applies the NEQ predicate on the "username" field.
func UsernameNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldUsername, v))
}

// UsernameIn applies the In predicate on the "username" field.
func UsernameIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldUsername, vs...))
}

// UsernameNotIn applies the NotIn predicate on the "username" field.
func UsernameNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldUsername, vs...))
}

// UsernameGT applies the GT predicate on the "username" field.
func UsernameGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldUsername, v))
}

// UsernameGTE applies the GTE predicate on the "username" field.
func UsernameGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldUsername, v))
}

// UsernameLT applies the LT predicate on the "username" field.
func UsernameLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldUsername, v))
}

// UsernameLTE applies the LTE predicate on the "username" field.
func UsernameLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldUsername, v))
}

// UsernameContains applies the Contains predicate on the "username" field.
func UsernameContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldUsername, v))
}

// UsernameHasPrefix applies the HasPrefix predicate on the "username" field.
func UsernameHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldUsername, v))
}

// UsernameHasSuffix applies the HasSuffix predicate on the "username" field.
func UsernameHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldUsername, v))
}

// UsernameEqualFold applies the EqualFold predicate on the "username" field.
func UsernameEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldUsername, v))
}

// UsernameContainsFold applies the ContainsFold predicate on the "username" field.
func UsernameContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldUsername, v))
}

// PasswdEQ applies the EQ predicate on the "passwd" field.
func PasswdEQ(v string) predicate.User {
	return predicate.User(sql.FieldEQ(FieldPasswd, v))
}

// PasswdNEQ applies the NEQ predicate on the "passwd" field.
func PasswdNEQ(v string) predicate.User {
	return predicate.User(sql.FieldNEQ(FieldPasswd, v))
}

// PasswdIn applies the In predicate on the "passwd" field.
func PasswdIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldIn(FieldPasswd, vs...))
}

// PasswdNotIn applies the NotIn predicate on the "passwd" field.
func PasswdNotIn(vs ...string) predicate.User {
	return predicate.User(sql.FieldNotIn(FieldPasswd, vs...))
}

// PasswdGT applies the GT predicate on the "passwd" field.
func PasswdGT(v string) predicate.User {
	return predicate.User(sql.FieldGT(FieldPasswd, v))
}

// PasswdGTE applies the GTE predicate on the "passwd" field.
func PasswdGTE(v string) predicate.User {
	return predicate.User(sql.FieldGTE(FieldPasswd, v))
}

// PasswdLT applies the LT predicate on the "passwd" field.
func PasswdLT(v string) predicate.User {
	return predicate.User(sql.FieldLT(FieldPasswd, v))
}

// PasswdLTE applies the LTE predicate on the "passwd" field.
func PasswdLTE(v string) predicate.User {
	return predicate.User(sql.FieldLTE(FieldPasswd, v))
}

// PasswdContains applies the Contains predicate on the "passwd" field.
func PasswdContains(v string) predicate.User {
	return predicate.User(sql.FieldContains(FieldPasswd, v))
}

// PasswdHasPrefix applies the HasPrefix predicate on the "passwd" field.
func PasswdHasPrefix(v string) predicate.User {
	return predicate.User(sql.FieldHasPrefix(FieldPasswd, v))
}

// PasswdHasSuffix applies the HasSuffix predicate on the "passwd" field.
func PasswdHasSuffix(v string) predicate.User {
	return predicate.User(sql.FieldHasSuffix(FieldPasswd, v))
}

// PasswdEqualFold applies the EqualFold predicate on the "passwd" field.
func PasswdEqualFold(v string) predicate.User {
	return predicate.User(sql.FieldEqualFold(FieldPasswd, v))
}

// PasswdContainsFold applies the ContainsFold predicate on the "passwd" field.
func PasswdContainsFold(v string) predicate.User {
	return predicate.User(sql.FieldContainsFold(FieldPasswd, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.User) predicate.User {
	return predicate.User(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.User) predicate.User {
	return predicate.User(sql.NotPredicates(p))
}
