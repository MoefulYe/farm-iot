// Code generated by ent, DO NOT EDIT.

package balance

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/MoefulYe/farm-iot/database/postgres/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Balance {
	return predicate.Balance(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Balance {
	return predicate.Balance(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Balance {
	return predicate.Balance(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Balance {
	return predicate.Balance(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Balance {
	return predicate.Balance(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Balance {
	return predicate.Balance(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Balance {
	return predicate.Balance(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Balance {
	return predicate.Balance(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Balance {
	return predicate.Balance(sql.FieldLTE(FieldID, id))
}

// When applies equality check predicate on the "when" field. It's identical to WhenEQ.
func When(v time.Time) predicate.Balance {
	return predicate.Balance(sql.FieldEQ(FieldWhen, v))
}

// Balance applies equality check predicate on the "balance" field. It's identical to BalanceEQ.
func Balance(v float64) predicate.Balance {
	return predicate.Balance(sql.FieldEQ(FieldBalance, v))
}

// WhenEQ applies the EQ predicate on the "when" field.
func WhenEQ(v time.Time) predicate.Balance {
	return predicate.Balance(sql.FieldEQ(FieldWhen, v))
}

// WhenNEQ applies the NEQ predicate on the "when" field.
func WhenNEQ(v time.Time) predicate.Balance {
	return predicate.Balance(sql.FieldNEQ(FieldWhen, v))
}

// WhenIn applies the In predicate on the "when" field.
func WhenIn(vs ...time.Time) predicate.Balance {
	return predicate.Balance(sql.FieldIn(FieldWhen, vs...))
}

// WhenNotIn applies the NotIn predicate on the "when" field.
func WhenNotIn(vs ...time.Time) predicate.Balance {
	return predicate.Balance(sql.FieldNotIn(FieldWhen, vs...))
}

// WhenGT applies the GT predicate on the "when" field.
func WhenGT(v time.Time) predicate.Balance {
	return predicate.Balance(sql.FieldGT(FieldWhen, v))
}

// WhenGTE applies the GTE predicate on the "when" field.
func WhenGTE(v time.Time) predicate.Balance {
	return predicate.Balance(sql.FieldGTE(FieldWhen, v))
}

// WhenLT applies the LT predicate on the "when" field.
func WhenLT(v time.Time) predicate.Balance {
	return predicate.Balance(sql.FieldLT(FieldWhen, v))
}

// WhenLTE applies the LTE predicate on the "when" field.
func WhenLTE(v time.Time) predicate.Balance {
	return predicate.Balance(sql.FieldLTE(FieldWhen, v))
}

// BalanceEQ applies the EQ predicate on the "balance" field.
func BalanceEQ(v float64) predicate.Balance {
	return predicate.Balance(sql.FieldEQ(FieldBalance, v))
}

// BalanceNEQ applies the NEQ predicate on the "balance" field.
func BalanceNEQ(v float64) predicate.Balance {
	return predicate.Balance(sql.FieldNEQ(FieldBalance, v))
}

// BalanceIn applies the In predicate on the "balance" field.
func BalanceIn(vs ...float64) predicate.Balance {
	return predicate.Balance(sql.FieldIn(FieldBalance, vs...))
}

// BalanceNotIn applies the NotIn predicate on the "balance" field.
func BalanceNotIn(vs ...float64) predicate.Balance {
	return predicate.Balance(sql.FieldNotIn(FieldBalance, vs...))
}

// BalanceGT applies the GT predicate on the "balance" field.
func BalanceGT(v float64) predicate.Balance {
	return predicate.Balance(sql.FieldGT(FieldBalance, v))
}

// BalanceGTE applies the GTE predicate on the "balance" field.
func BalanceGTE(v float64) predicate.Balance {
	return predicate.Balance(sql.FieldGTE(FieldBalance, v))
}

// BalanceLT applies the LT predicate on the "balance" field.
func BalanceLT(v float64) predicate.Balance {
	return predicate.Balance(sql.FieldLT(FieldBalance, v))
}

// BalanceLTE applies the LTE predicate on the "balance" field.
func BalanceLTE(v float64) predicate.Balance {
	return predicate.Balance(sql.FieldLTE(FieldBalance, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Balance) predicate.Balance {
	return predicate.Balance(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Balance) predicate.Balance {
	return predicate.Balance(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Balance) predicate.Balance {
	return predicate.Balance(sql.NotPredicates(p))
}
