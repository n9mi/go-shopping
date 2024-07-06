package test

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func CobaTest(t *testing.T) {
	require.Equal(t, 2*2, 4)
}
