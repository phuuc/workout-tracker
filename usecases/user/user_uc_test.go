package usecases

import (
	"context"
	"errors"
	"testing"

	"github.com/finnpn/workout-tracker/interfaces/repos"
	"github.com/finnpn/workout-tracker/usecases/in"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
)

func TestUserUc_Register_Success(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUsersRepo := repos.NewMockUsersRepo(mockCtrl)

	mockUsersRepo.EXPECT().Register(gomock.Any(), gomock.Any()).Return(nil)

	userUc := &UserUc{
		userRepo: mockUsersRepo,
	}

	err := userUc.Register(context.Background(), &in.Register{
		Email:    "12345@gmail.com",
		Password: "12345678",
	})

	require.NoError(t, err)
}

func TestUserUc_Register_Fail(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUsersRepo := repos.NewMockUsersRepo(mockCtrl)

	mockUsersRepo.EXPECT().Register(gomock.Any(), gomock.Any()).Return(errors.New(""))

	userUc := &UserUc{
		userRepo: mockUsersRepo,
	}

	err := userUc.Register(context.Background(), &in.Register{
		Email:    "12345@gmail.com",
		Password: "12345678",
	})

	require.Error(t, err)
}

func TestUserUc_validate_Success(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUsersRepo := repos.NewMockUsersRepo(mockCtrl)

	userUc := &UserUc{
		userRepo: mockUsersRepo,
	}

	err := userUc.validate(&in.Register{
		Email:    "12345@gmail.com",
		Password: "12345678",
	})

	require.NoError(t, err)
}

func TestUserUc_validate_Fail(t *testing.T) {
	t.Parallel()
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mockUsersRepo := repos.NewMockUsersRepo(mockCtrl)

	userUc := &UserUc{
		userRepo: mockUsersRepo,
	}

	err := userUc.validate(&in.Register{
		Email:    "12345@gmail.com",
		Password: "1234568",
	})

	require.Error(t, err)
}
