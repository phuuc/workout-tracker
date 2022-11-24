package repos

import (
	"context"

	"github.com/finnpn/workout-tracker/usecases/in"
)

//go:generate mockgen -destination=users_repo_mock.go -package=repos -source=users_repo.go

// UsersRepo ...
type UsersRepo interface {
	Register(ctx context.Context, record *in.Register) error
	Login(ctx context.Context, record *in.Login) error
}
