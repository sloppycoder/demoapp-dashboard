package app

import (
	"dashboard/api"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"

	"golang.org/x/sync/errgroup"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"

	log "github.com/sirupsen/logrus"
)

var (
	DummyCustomer = &api.Customer{LoginName: "skip"}
	DummyCasa     = []*api.CasaAccount{{AccountId: "skip"}}
)

const (
	StatsReportingPeriod = 60
	Timeout              = 5 * time.Second
)

func DashboardHandler(ctx echo.Context) error {
	login := ctx.Param("login")
	dashboard, err := GetDashboard(login)
	if err != nil {
		message := fmt.Sprintf("unable to retrieve dashboard due to %v", err)
		return ctx.String(http.StatusInternalServerError, message)
	}
	return ctx.JSON(http.StatusOK, dashboard)
}

func GetDashboard(customerId string) (*api.Dashboard, error) {
	dashboard := &api.Dashboard{}

	errs := errgroup.Group{}
	errs.Go(func() error {
		cust, err := GetCustomer(customerId)
		if err == nil {
			dashboard.Customer = cust
		}
		return err
	})
	errs.Go(func() error {
		casa, err := GetCasaAccounts(customerId)
		if err == nil {
			dashboard.Casa = casa
		}
		return err
	})
	if err := errs.Wait(); err != nil {
		return nil, err
	}

	return dashboard, nil
}

func GetCustomer(customerId string) (*api.Customer, error) {
	addr := os.Getenv("CUST_SVC_ADDR")
	if addr == "" || addr == "false" {
		return DummyCustomer, nil
	}

	log.Infof("retrieving customer for %s", customerId)

	req, err := http.NewRequest("GET", addr+"/customers/"+customerId, nil)
	if err != nil {
		log.Warnf("Got error %v when creating request", err)
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		log.Infof("Got error %v when calling customer svc", err)
		return nil, err
	}
	defer resp.Body.Close()

	var customer api.Customer
	if err := json.NewDecoder(resp.Body).Decode(&customer); err != nil {
		log.Infof("Got error %v when decode response", err)
		return nil, err
	}

	return &customer, nil
}

func GetCasaAccounts(customerId string) ([]*api.CasaAccount, error) {
	return DummyCasa, nil
}

func StartServer(addr string) {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/dashboard/:login", DashboardHandler)

	e.Logger.Fatal(e.Start(addr))
}
