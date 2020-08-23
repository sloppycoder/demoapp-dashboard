package app

import (
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestGetDashboard(t *testing.T) {
	os.Setenv("CUST_SVC_ADDR", "false")
	os.Setenv("CASA_SVC_ADDR", "false")

	dashboard, err := GetDashboard("10001000")
	assert.Nil(t, err, "DashboardHandler returned error")
	assert.Equal(t, dashboard.Customer.LoginName, "skip")
}
