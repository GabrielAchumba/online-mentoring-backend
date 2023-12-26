// Code generated by ent, DO NOT EDIT.

package privacy

import (
	"api/ent/entgenerated"
	"context"

	"entgo.io/ent/entql"
	"entgo.io/ent/privacy"
)

var (
	// Allow may be returned by rules to indicate that the policy
	// evaluation should terminate with allow decision.
	Allow = privacy.Allow

	// Deny may be returned by rules to indicate that the policy
	// evaluation should terminate with deny decision.
	Deny = privacy.Deny

	// Skip may be returned by rules to indicate that the policy
	// evaluation should continue to the next rule.
	Skip = privacy.Skip
)

// Allowf returns a formatted wrapped Allow decision.
func Allowf(format string, a ...any) error {
	return privacy.Allowf(format, a...)
}

// Denyf returns a formatted wrapped Deny decision.
func Denyf(format string, a ...any) error {
	return privacy.Denyf(format, a...)
}

// Skipf returns a formatted wrapped Skip decision.
func Skipf(format string, a ...any) error {
	return privacy.Skipf(format, a...)
}

// DecisionContext creates a new context from the given parent context with
// a policy decision attach to it.
func DecisionContext(parent context.Context, decision error) context.Context {
	return privacy.DecisionContext(parent, decision)
}

// DecisionFromContext retrieves the policy decision from the context.
func DecisionFromContext(ctx context.Context) (error, bool) {
	return privacy.DecisionFromContext(ctx)
}

type (
	// Policy groups query and mutation policies.
	Policy = privacy.Policy

	// QueryRule defines the interface deciding whether a
	// query is allowed and optionally modify it.
	QueryRule = privacy.QueryRule
	// QueryPolicy combines multiple query rules into a single policy.
	QueryPolicy = privacy.QueryPolicy

	// MutationRule defines the interface which decides whether a
	// mutation is allowed and optionally modifies it.
	MutationRule = privacy.MutationRule
	// MutationPolicy combines multiple mutation rules into a single policy.
	MutationPolicy = privacy.MutationPolicy
	// MutationRuleFunc type is an adapter which allows the use of
	// ordinary functions as mutation rules.
	MutationRuleFunc = privacy.MutationRuleFunc

	// QueryMutationRule is an interface which groups query and mutation rules.
	QueryMutationRule = privacy.QueryMutationRule
)

// QueryRuleFunc type is an adapter to allow the use of
// ordinary functions as query rules.
type QueryRuleFunc func(context.Context, entgenerated.Query) error

// Eval returns f(ctx, q).
func (f QueryRuleFunc) EvalQuery(ctx context.Context, q entgenerated.Query) error {
	return f(ctx, q)
}

// AlwaysAllowRule returns a rule that returns an allow decision.
func AlwaysAllowRule() QueryMutationRule {
	return privacy.AlwaysAllowRule()
}

// AlwaysDenyRule returns a rule that returns a deny decision.
func AlwaysDenyRule() QueryMutationRule {
	return privacy.AlwaysDenyRule()
}

// ContextQueryMutationRule creates a query/mutation rule from a context eval func.
func ContextQueryMutationRule(eval func(context.Context) error) QueryMutationRule {
	return privacy.ContextQueryMutationRule(eval)
}

// OnMutationOperation evaluates the given rule only on a given mutation operation.
func OnMutationOperation(rule MutationRule, op entgenerated.Op) MutationRule {
	return privacy.OnMutationOperation(rule, op)
}

// DenyMutationOperationRule returns a rule denying specified mutation operation.
func DenyMutationOperationRule(op entgenerated.Op) MutationRule {
	rule := MutationRuleFunc(func(_ context.Context, m entgenerated.Mutation) error {
		return Denyf("entgenerated/privacy: operation %s is not allowed", m.Op())
	})
	return OnMutationOperation(rule, op)
}

// The EmailCredentialQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type EmailCredentialQueryRuleFunc func(context.Context, *entgenerated.EmailCredentialQuery) error

// EvalQuery return f(ctx, q).
func (f EmailCredentialQueryRuleFunc) EvalQuery(ctx context.Context, q entgenerated.Query) error {
	if q, ok := q.(*entgenerated.EmailCredentialQuery); ok {
		return f(ctx, q)
	}
	return Denyf("entgenerated/privacy: unexpected query type %T, expect *entgenerated.EmailCredentialQuery", q)
}

// The EmailCredentialMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type EmailCredentialMutationRuleFunc func(context.Context, *entgenerated.EmailCredentialMutation) error

// EvalMutation calls f(ctx, m).
func (f EmailCredentialMutationRuleFunc) EvalMutation(ctx context.Context, m entgenerated.Mutation) error {
	if m, ok := m.(*entgenerated.EmailCredentialMutation); ok {
		return f(ctx, m)
	}
	return Denyf("entgenerated/privacy: unexpected mutation type %T, expect *entgenerated.EmailCredentialMutation", m)
}

// The LoginSessionQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type LoginSessionQueryRuleFunc func(context.Context, *entgenerated.LoginSessionQuery) error

// EvalQuery return f(ctx, q).
func (f LoginSessionQueryRuleFunc) EvalQuery(ctx context.Context, q entgenerated.Query) error {
	if q, ok := q.(*entgenerated.LoginSessionQuery); ok {
		return f(ctx, q)
	}
	return Denyf("entgenerated/privacy: unexpected query type %T, expect *entgenerated.LoginSessionQuery", q)
}

// The LoginSessionMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type LoginSessionMutationRuleFunc func(context.Context, *entgenerated.LoginSessionMutation) error

// EvalMutation calls f(ctx, m).
func (f LoginSessionMutationRuleFunc) EvalMutation(ctx context.Context, m entgenerated.Mutation) error {
	if m, ok := m.(*entgenerated.LoginSessionMutation); ok {
		return f(ctx, m)
	}
	return Denyf("entgenerated/privacy: unexpected mutation type %T, expect *entgenerated.LoginSessionMutation", m)
}

// The SuperuserProfileQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type SuperuserProfileQueryRuleFunc func(context.Context, *entgenerated.SuperuserProfileQuery) error

// EvalQuery return f(ctx, q).
func (f SuperuserProfileQueryRuleFunc) EvalQuery(ctx context.Context, q entgenerated.Query) error {
	if q, ok := q.(*entgenerated.SuperuserProfileQuery); ok {
		return f(ctx, q)
	}
	return Denyf("entgenerated/privacy: unexpected query type %T, expect *entgenerated.SuperuserProfileQuery", q)
}

// The SuperuserProfileMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type SuperuserProfileMutationRuleFunc func(context.Context, *entgenerated.SuperuserProfileMutation) error

// EvalMutation calls f(ctx, m).
func (f SuperuserProfileMutationRuleFunc) EvalMutation(ctx context.Context, m entgenerated.Mutation) error {
	if m, ok := m.(*entgenerated.SuperuserProfileMutation); ok {
		return f(ctx, m)
	}
	return Denyf("entgenerated/privacy: unexpected mutation type %T, expect *entgenerated.SuperuserProfileMutation", m)
}

// The UserQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type UserQueryRuleFunc func(context.Context, *entgenerated.UserQuery) error

// EvalQuery return f(ctx, q).
func (f UserQueryRuleFunc) EvalQuery(ctx context.Context, q entgenerated.Query) error {
	if q, ok := q.(*entgenerated.UserQuery); ok {
		return f(ctx, q)
	}
	return Denyf("entgenerated/privacy: unexpected query type %T, expect *entgenerated.UserQuery", q)
}

// The UserMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type UserMutationRuleFunc func(context.Context, *entgenerated.UserMutation) error

// EvalMutation calls f(ctx, m).
func (f UserMutationRuleFunc) EvalMutation(ctx context.Context, m entgenerated.Mutation) error {
	if m, ok := m.(*entgenerated.UserMutation); ok {
		return f(ctx, m)
	}
	return Denyf("entgenerated/privacy: unexpected mutation type %T, expect *entgenerated.UserMutation", m)
}

// The UserPublicProfileQueryRuleFunc type is an adapter to allow the use of ordinary
// functions as a query rule.
type UserPublicProfileQueryRuleFunc func(context.Context, *entgenerated.UserPublicProfileQuery) error

// EvalQuery return f(ctx, q).
func (f UserPublicProfileQueryRuleFunc) EvalQuery(ctx context.Context, q entgenerated.Query) error {
	if q, ok := q.(*entgenerated.UserPublicProfileQuery); ok {
		return f(ctx, q)
	}
	return Denyf("entgenerated/privacy: unexpected query type %T, expect *entgenerated.UserPublicProfileQuery", q)
}

// The UserPublicProfileMutationRuleFunc type is an adapter to allow the use of ordinary
// functions as a mutation rule.
type UserPublicProfileMutationRuleFunc func(context.Context, *entgenerated.UserPublicProfileMutation) error

// EvalMutation calls f(ctx, m).
func (f UserPublicProfileMutationRuleFunc) EvalMutation(ctx context.Context, m entgenerated.Mutation) error {
	if m, ok := m.(*entgenerated.UserPublicProfileMutation); ok {
		return f(ctx, m)
	}
	return Denyf("entgenerated/privacy: unexpected mutation type %T, expect *entgenerated.UserPublicProfileMutation", m)
}

type (
	// Filter is the interface that wraps the Where function
	// for filtering nodes in queries and mutations.
	Filter interface {
		// Where applies a filter on the executed query/mutation.
		Where(entql.P)
	}

	// The FilterFunc type is an adapter that allows the use of ordinary
	// functions as filters for query and mutation types.
	FilterFunc func(context.Context, Filter) error
)

// EvalQuery calls f(ctx, q) if the query implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalQuery(ctx context.Context, q entgenerated.Query) error {
	fr, err := queryFilter(q)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

// EvalMutation calls f(ctx, q) if the mutation implements the Filter interface, otherwise it is denied.
func (f FilterFunc) EvalMutation(ctx context.Context, m entgenerated.Mutation) error {
	fr, err := mutationFilter(m)
	if err != nil {
		return err
	}
	return f(ctx, fr)
}

var _ QueryMutationRule = FilterFunc(nil)

func queryFilter(q entgenerated.Query) (Filter, error) {
	switch q := q.(type) {
	case *entgenerated.EmailCredentialQuery:
		return q.Filter(), nil
	case *entgenerated.LoginSessionQuery:
		return q.Filter(), nil
	case *entgenerated.SuperuserProfileQuery:
		return q.Filter(), nil
	case *entgenerated.UserQuery:
		return q.Filter(), nil
	case *entgenerated.UserPublicProfileQuery:
		return q.Filter(), nil
	default:
		return nil, Denyf("entgenerated/privacy: unexpected query type %T for query filter", q)
	}
}

func mutationFilter(m entgenerated.Mutation) (Filter, error) {
	switch m := m.(type) {
	case *entgenerated.EmailCredentialMutation:
		return m.Filter(), nil
	case *entgenerated.LoginSessionMutation:
		return m.Filter(), nil
	case *entgenerated.SuperuserProfileMutation:
		return m.Filter(), nil
	case *entgenerated.UserMutation:
		return m.Filter(), nil
	case *entgenerated.UserPublicProfileMutation:
		return m.Filter(), nil
	default:
		return nil, Denyf("entgenerated/privacy: unexpected mutation type %T for mutation filter", m)
	}
}