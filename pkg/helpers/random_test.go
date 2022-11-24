package helpers

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRandNumber(t *testing.T) {
	n := 6

	str := RandNumber(n)
	require.Equal(t, len(str), n)
}
