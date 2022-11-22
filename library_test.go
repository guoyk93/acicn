package acicn

import (
	"github.com/guoyk93/gg"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestLoad(t *testing.T) {
	lib, err := Load(gg.M{})
	require.NoError(t, err)
	for _, item := range lib.Repos {
		gg.Log(item)
	}
}
