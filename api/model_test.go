package api

import (
	"encoding/json"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMarshallBalanceType(t *testing.T) {
	balance := &Balance{
		Type:       BalanceAVAILABLE,
		Amount:     10000.12,
		CreditFlag: true,
	}

	blob, err := json.Marshal(balance)
	assert.Nil(t, err, "Marshall BalanceType error")

	str := string(blob)
	assert.True(t, strings.Contains(str, "\"type\":\"AVAILABLE\""))
	assert.True(t, strings.Contains(str, "\"amount\":10000.12"))
}

func TestUnmarshallBalanceType(t *testing.T) {
	blob := []byte(`
		{ 
			"BalanceType" : "CURRENT",
			"Amount" : 10000.12,
			"CreditFlag" : false
		}
	`)

	var balance Balance
	err := json.Unmarshal(blob, &balance)

	assert.Nil(t, err, "Unmarshall BalanceType error")
	assert.Equal(t, BalanceCURRENT, balance.Type)
	assert.Equal(t, 10000.12, balance.Amount)
}
