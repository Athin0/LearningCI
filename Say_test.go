package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSayHello(t *testing.T) {
	err := SayHello()
	require.NoError(t, err)
}
