package adapter

import (
	"MonobankTracker/pkg/currency/domain/monobank_action"
	"context"
	"database/sql"
)

//TODO: rename package domain.monobank_aktion
//TODO: rename NewPostgreSQLRepository

type Action struct {
	Id              string `json:"id"`
	Account         string `json:"account"`
	Time            int    `json:"time"`
	Description     string `json:"description"`
	Mcc             int    `json:"mcc"`
	OriginalMcc     int    `json:"originalMcc"`
	Hold            bool   `json:"hold"`
	Amount          int    `json:"amount"`
	OperationAmount int    `json:"operationAmount"`
	CurrencyCode    int    `json:"currencyCode"`
	CommissionRate  int    `json:"commissionRate"`
	CashbackAmount  int    `json:"cashbackAmount"`
	Balance         int    `json:"balance"`
	Comment         string `json:"comment"`
	ReceiptId       string `json:"receiptId"`
	InvoiceId       string `json:"invoiceId"`
	CounterEdrpou   string `json:"counterEdrpou"`
	CounterIban     string `json:"counterIban"`
	CounterName     string `json:"counterName"`
}

func (p *PostgreSQLRepository) NewAction(monobankAction *monobank_action.MonobankAction) *Action {
	return &Action{
		Id:              monobankAction.Data.StatementItem.Id,
		Account:         monobankAction.Data.Account,
		Time:            monobankAction.Data.StatementItem.Time,
		Description:     monobankAction.Data.StatementItem.Description,
		Mcc:             monobankAction.Data.StatementItem.Mcc,
		OriginalMcc:     monobankAction.Data.StatementItem.OriginalMcc,
		Hold:            monobankAction.Data.StatementItem.Hold,
		Amount:          monobankAction.Data.StatementItem.Amount,
		OperationAmount: monobankAction.Data.StatementItem.OperationAmount,
		CurrencyCode:    monobankAction.Data.StatementItem.CurrencyCode,
		CommissionRate:  monobankAction.Data.StatementItem.CommissionRate,
		CashbackAmount:  monobankAction.Data.StatementItem.CashbackAmount,
		Balance:         monobankAction.Data.StatementItem.Balance,
		Comment:         monobankAction.Data.StatementItem.Comment,
		ReceiptId:       monobankAction.Data.StatementItem.ReceiptId,
		InvoiceId:       monobankAction.Data.StatementItem.InvoiceId,
		CounterEdrpou:   monobankAction.Data.StatementItem.CounterEdrpou,
		CounterIban:     monobankAction.Data.StatementItem.CounterIban,
		CounterName:     monobankAction.Data.StatementItem.CounterName,
	}
}

type PostgreSQLRepository struct {
	db *sql.DB
}

func NewPostgreSQLRepository(db *sql.DB) monobank_action.Repository {
	return &PostgreSQLRepository{
		db: db,
	}
}

const insertQuery = `INSERT INTO "mono-action"(
                          id,
                          account,
                          "time",
                          description,
                          mcc,
                          originalMcc,
                          "hold",
                          amount, 
                          operationAmount,
                          currencyCode, 
                          commissionRate, 
                          cashbackAmount, 
                          balance, 
                          comment, 
                          receiptId, 
                          invoiceId,
    					  counterEdrpou, 
                          counterIban, 
                          counterName) 
					  values($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)`

func (p *PostgreSQLRepository) Create(monoAction Action) (*Action, error) {
	_, err := p.db.Exec(insertQuery,
		monoAction.Id, monoAction.Account, monoAction.Time, monoAction.Description, monoAction.Mcc, monoAction.OriginalMcc,
		monoAction.Hold, monoAction.Amount, monoAction.OperationAmount, monoAction.CurrencyCode, monoAction.CommissionRate,
		monoAction.CashbackAmount, monoAction.Balance, monoAction.Comment, monoAction.ReceiptId, monoAction.InvoiceId,
		monoAction.CounterEdrpou, monoAction.CounterIban, monoAction.CounterName)
	if err != nil {
		return nil, err
	}
	return &monoAction, nil
}

func (p *PostgreSQLRepository) All(ctx context.Context) ([]Action, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgreSQLRepository) GetByName(ctx context.Context, name string) (*Action, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgreSQLRepository) Update(ctx context.Context, id int, updated Action) (*Action, error) {
	//TODO implement me
	panic("implement me")
}

func (p *PostgreSQLRepository) Delete(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}
