package miraiex

import (
	"fmt"
	"testing"
)
import "github.com/stretchr/testify/require"

func TestMarketString(t *testing.T) {
	require.Equal(t, "BTCNOK", fmt.Sprintf("%s", BTCNOK))
}


/**

2021/02/08 19:33:32 Request body: {"market":"BTCNOK","type":"Bid","price":"370901.57","amount":"0.000100"}
2021/02/08 19:33:32 Response code400with body:{"name":"InsufficientFunds","message":"Insufficient funds"}
2021/02/08 19:33:32 Got status code 400
 */