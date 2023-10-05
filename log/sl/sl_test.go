package sl

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_Err(t *testing.T) {
	err := errors.New("test error")

	attr := Err(err)
	require.Equal(t, "error", attr.Key)
	require.Equal(t, attr.String(), "error=test error")
}
