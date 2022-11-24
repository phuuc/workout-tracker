package usecases

import (
	"context"
	"errors"

	"github.com/finnpn/workout-tracker/interfaces/repos"
	"github.com/finnpn/workout-tracker/interfaces/restapi/models"
	"github.com/finnpn/workout-tracker/pkg/helpers"
	"github.com/finnpn/workout-tracker/usecases/in"
)

type AuthUserUc struct {
	UserUsecase *UserUc
}

func NewAuthUserUc(userRepo repos.UsersRepo) *AuthUserUc {
	return &AuthUserUc{
		UserUsecase: NewUserUc(userRepo),
	}
}

type UserUc struct {
	userRepo repos.UsersRepo
}

func NewUserUc(userRepo repos.UsersRepo) *UserUc {
	return &UserUc{
		userRepo: userRepo,
	}
}

func (u *UserUc) Register(ctx context.Context, in *in.Register) error {

	hashedPw, err := helpers.HashPassword(in.Password)

	if err != nil {
		return err
	}
	in.Password = hashedPw
	err = u.userRepo.Register(ctx, in)
	if err != nil {
		return err
	}

	return nil
}

func (u *UserUc) validate(in *in.Register) error {
	if err := models.EmailRegexp(in.Email); err != nil {
		return err
	}
	if models.PasswordRegexp(in.Password) {
		return nil
	}
	return errors.New("not valid input")
}
