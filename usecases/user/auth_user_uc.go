package usecases

import (
	"context"

	"github.com/finnpn/workout-tracker/usecases/in"
)

func (v *AuthUserUc) Register(ctx context.Context, in *in.Register) error {
	if err := v.UserUsecase.validate(in); err != nil {
		return err
	}
	return v.UserUsecase.Register(ctx, in)
}
