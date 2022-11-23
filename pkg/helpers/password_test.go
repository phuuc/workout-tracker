package helpers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestHashPassword(t *testing.T) {
	pw := RandStringRunes(8)
	hashed, err := HashPassword(pw)

	require.NoError(t, err)
	require.NotEmpty(t, hashed)

	err = CheckPassword(pw, hashed)

	require.NoError(t, err)

}
