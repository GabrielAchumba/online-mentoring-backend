// Code generated by ent, DO NOT EDIT.

package runtime

import (
	"api/ent/entgenerated/emailcredential"
	"api/ent/entgenerated/loginsession"
	"api/ent/entgenerated/superuserprofile"
	"api/ent/entgenerated/user"
	"api/ent/entgenerated/userpublicprofile"
	"api/ent/schema"
	"context"

	"entgo.io/ent"
	"entgo.io/ent/privacy"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	emailcredential.Policy = privacy.NewPolicies(schema.EmailCredential{})
	emailcredential.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := emailcredential.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	emailcredentialFields := schema.EmailCredential{}.Fields()
	_ = emailcredentialFields
	// emailcredentialDescEmail is the schema descriptor for email field.
	emailcredentialDescEmail := emailcredentialFields[0].Descriptor()
	// emailcredential.EmailValidator is a validator for the "email" field. It is called by the builders before save.
	emailcredential.EmailValidator = emailcredentialDescEmail.Validators[0].(func(string) error)
	loginsessionMixin := schema.LoginSession{}.Mixin()
	loginsession.Policy = privacy.NewPolicies(loginsessionMixin[1], schema.LoginSession{})
	loginsession.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := loginsession.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	superuserprofileMixin := schema.SuperuserProfile{}.Mixin()
	superuserprofile.Policy = privacy.NewPolicies(superuserprofileMixin[1], superuserprofileMixin[2], schema.SuperuserProfile{})
	superuserprofile.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := superuserprofile.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	user.Policy = privacy.NewPolicies(schema.User{})
	user.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := user.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
	userpublicprofileMixin := schema.UserPublicProfile{}.Mixin()
	userpublicprofile.Policy = privacy.NewPolicies(userpublicprofileMixin[1], schema.UserPublicProfile{})
	userpublicprofile.Hooks[0] = func(next ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			if err := userpublicprofile.Policy.EvalMutation(ctx, m); err != nil {
				return nil, err
			}
			return next.Mutate(ctx, m)
		})
	}
}

const (
	Version = "v0.12.4"                                         // Version of ent codegen.
	Sum     = "h1:LddPnAyxls/O7DTXZvUGDj0NZIdGSu317+aoNLJWbD8=" // Sum of ent codegen.
)
