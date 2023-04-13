package application

import (
	"MonobankTracker/pkg/currency/domain/currency"
	"MonobankTracker/pkg/currency/domain/monobank_action"
	"log"
)

type Service interface {
	GetCurrencies() ([]CurrencyDto, error)
	GetCurrencyByName(currencyName string) (*CurrencyDto, error)
	MonobankAction(monoActoin *monobank_action.MonobankAction) error
}

type CurrencyDto struct {
	Name          string                        `json:"name"`
	Code          int                           `json:"code"`
	ExchangeRates map[int]currency.ExchangeRate `json:"exchangeRates"`
}

// Service implementation

func NewService(currencyRepository currency.Repository, monoActionRepository monobank_action.Repository) Service {
	return &service{
		currencyRepository:   currencyRepository,
		monoActionRepository: monoActionRepository,
	}
}

type service struct {
	currencyRepository   currency.Repository
	monoActionRepository monobank_action.Repository
}

func (s *service) GetCurrencies() ([]CurrencyDto, error) {
	//table, err := s.currencyRepository.GetCurrencyTable()
	//if err != nil {
	//	return nil, err
	//}
	//table.GetCurrency("USD")

	//TODO implement me
	panic("implement me")
}

func (s *service) GetCurrencyByName(currencyName string) (*CurrencyDto, error) {
	currencyByName, err := s.currencyRepository.GetCurrencyByName(currencyName)
	if err != nil {
		return nil, err
	}

	return &CurrencyDto{
		Name:          currencyByName.Name,
		Code:          currencyByName.Id,
		ExchangeRates: currencyByName.ExchangeRates,
	}, nil
}

func (s *service) MonobankAction(monobankAction *monobank_action.MonobankAction) error {
	action := s.monoActionRepository.NewAction(monobankAction)
	_, err := s.monoActionRepository.Create(action)
	if err != nil {
		log.Print(err)
		return err
	}
	return nil
}

//
//func (s *service) GetAllMonobankAction() *map[int]currency.MonobankAction {
//	mapActions := currency.GetAllMonobankActions()
//	return &mapActions
//}
