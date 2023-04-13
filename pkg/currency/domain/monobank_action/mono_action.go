package monobank_action

type MonobankAction struct {
	Type string `json:"type"`
	Data Data   `json:"data"`
}
type Data struct {
	Account       string        `json:"account"`
	StatementItem StatementItem `json:"statementItem"`
}
type StatementItem struct {
	Id              string `json:"id"`
	Time            int    `json:"time"`
	Description     string `json:"description"`
	Comment         string `json:"comment"`
	Mcc             int    `json:"mcc"`
	OriginalMcc     int    `json:"originalMcc"`
	Amount          int    `json:"amount"`
	OperationAmount int    `json:"operationAmount"`
	CurrencyCode    int    `json:"currencyCode"`
	CommissionRate  int    `json:"commissionRate"`
	CashbackAmount  int    `json:"cashbackAmount"`
	Balance         int    `json:"balance"`
	Hold            bool   `json:"hold"`
	ReceiptId       string `json:"receiptId"`
	InvoiceId       string `json:"invoiceId"`
	CounterEdrpou   string `json:"counterEdrpou"`
	CounterIban     string `json:"counterIban"`
	CounterName     string `json:"counterName"`
}
