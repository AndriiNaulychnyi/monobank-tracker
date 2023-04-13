package currency

import (
	"MonobankTracker/pkg/currency/adapter"
	"MonobankTracker/pkg/currency/application"
	"MonobankTracker/pkg/currency/port"
	"database/sql"
	"github.com/gin-gonic/gin"
)

func RegisterEndpoints(router gin.IRouter, dbConnection *sql.DB) {
	currencyRepository := adapter.NewMonobankCurrencyRepository()
	monobankActionRepository := adapter.NewPostgreSQLRepository(dbConnection)
	service := application.NewService(currencyRepository, monobankActionRepository)
	handler := port.NewHttpHandler(service)
	handler.RegisterApiHandlers(router)
}
