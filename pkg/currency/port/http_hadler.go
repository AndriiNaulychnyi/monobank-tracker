package port

import (
	"MonobankTracker/pkg/currency/application"
	"MonobankTracker/pkg/currency/domain/monobank_action"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func NewHttpHandler(service application.Service) *httpHandler {
	return &httpHandler{
		service: service,
	}
}

type httpHandler struct {
	service application.Service
}

func (h *httpHandler) RegisterApiHandlers(router gin.IRouter) {
	router.GET("/currency", h.getCurrency)
	router.GET("/currency/:name", h.getCurrencyByName)
	router.GET("/webhook", h.confirmWebHook)
	router.POST("/webhook", h.monobankAction)
	//router.GET("/get/all-action", h.allMonoAction)
}
func (h *httpHandler) getCurrencyByName(c *gin.Context) {
	currencyName := c.Param("name")
	currency, err := h.service.GetCurrencyByName(currencyName)
	if err != nil {
		c.Error(err)
		return
	}
	c.JSON(http.StatusOK, currency)
}
func (h *httpHandler) getCurrency(c *gin.Context) {
	currencies, err := h.service.GetCurrencies()
	if err != nil {
		c.AbortWithError(http.StatusInternalServerError, err)
		return
	}
	c.JSON(http.StatusOK, currencies)
}

func (h *httpHandler) confirmWebHook(c *gin.Context) {
	log.Print("Confirm WebHook")
	c.JSON(http.StatusOK, gin.H{})
}

func (h *httpHandler) monobankAction(c *gin.Context) {
	log.Println("POST: action in monobank")
	monobankAction := &monobank_action.MonobankAction{}
	err := c.ShouldBindJSON(monobankAction)
	if err != nil {
		c.Error(err)
		return
	}
	err = h.service.MonobankAction(monobankAction)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}
	log.Print(monobankAction)
	c.JSON(http.StatusOK, gin.H{})
}
