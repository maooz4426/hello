package hello

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Hello(t *testing.T) {
	msg, err := Hello("test")
	require.NoError(t, err)
	require.Equal(t, msg, "Hello test")
}
