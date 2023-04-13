package adapter

import (
	"MonobankTracker/pkg/currency/domain/currency"
	"encoding/json"
	"log"
	"net/http"
)

func NewMonobankCurrencyRepository() currency.Repository {
	return &currencyRepository{}
}

type currencyRepository struct {
}

type monobankCurrency struct {
	CurrencyCodeA int     `json:"currencyCodeA"`
	CurrencyCodeB int     `json:"currencyCodeB"`
	Date          int     `json:"date"`
	RateSell      float64 `json:"rateSell"`
	RateBuy       float64 `json:"rateBuy"`
	RateCross     float64 `json:"rateCross"`
}

//var monobankActions = map[int64]currency.MonobankAction{}

func (c *currencyRepository) getCurrenciesFromMonobank() ([]monobankCurrency, error) {
	resp, err := http.Get("https://api.monobank.ua/bank/currency")
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var monoCurrency []monobankCurrency
	err = json.NewDecoder(resp.Body).Decode(&monoCurrency)
	if err != nil {
		log.Print(err)
		return nil, err
	}
	return monoCurrency, nil
}

func (c *currencyRepository) GetCurrencyByName(name string) (*currency.Currency, error) {
	monoCurrency, err := c.getCurrenciesFromMonobank()
	if err != nil {
		return nil, err
	}
	log.Print(monoCurrency)
	currencyByName := c.parsMonobankCurrency(monoCurrency, name)
	return currencyByName, nil
}

func (c *currencyRepository) parsMonobankCurrency(monoCurrencies []monobankCurrency, name string) *currency.Currency {
	currencyCode := currency.GetCurrencyCodeByName(name)
	currencyByName := &currency.Currency{
		Id:            currencyCode,
		Name:          name,
		ExchangeRates: make(map[int]currency.ExchangeRate),
	}

	for _, monoCurrency := range monoCurrencies {
		if monoCurrency.CurrencyCodeB != currencyCode {
			continue
		}
		currencyByName.ExchangeRates[monoCurrency.CurrencyCodeA] = currency.ExchangeRate{
			Id:        monoCurrency.CurrencyCodeA,
			Name:      currency.GetCurrencyNameByCode(monoCurrency.CurrencyCodeA),
			RateBuy:   monoCurrency.RateBuy,
			RateSell:  monoCurrency.RateSell,
			RateCross: monoCurrency.RateCross,
		}
		currencyByName.Date = monoCurrency.Date
	}
	return currencyByName
}

//func (c *currencyRepository) SaveMonobankAction(action *currency.MonobankAction) {
//	currency.SaveMonobankAction(action)
//	//monobankActions[time.Now().Unix()] = *action
//	print(action)
//}
