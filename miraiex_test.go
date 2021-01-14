package main

import (
	"fmt"
	"testing"
)
import "github.com/stretchr/testify/require"

func TestMarketString(t *testing.T) {
	require.Equal(t, "BTCNOK", fmt.Sprintf("%s", BTCNOK))
}
