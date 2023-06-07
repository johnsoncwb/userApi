package api

import (
	"net/http"

	"github.com/johnsoncwb/userApi/internal/adapters/external"
	"github.com/johnsoncwb/userApi/internal/config"
	"github.com/johnsoncwb/userApi/internal/core/service"
	"github.com/newrelic/go-agent/v3/integrations/nrhttprouter"
)

func LoadRoutes() {
	externalService := external.NewRequests()
	services := service.NewUserService(*externalService)
	handler := NewUsersHandlers(*services)

	router := nrhttprouter.New(config.GlobalConfig.NewRelicApp)

	var basePath = "/v1"

	router.GET(basePath+"/users", handler.GetAll)
	router.GET(basePath+"/user/:id", handler.GetById)

	_ = http.ListenAndServe("localhost:9000", router)
}
