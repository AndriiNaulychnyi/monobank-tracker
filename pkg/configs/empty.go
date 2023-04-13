package configs

type T struct {
	Type string `json:"type"`
	Data struct {
		Account       string `json:"account"`
		StatementItem struct {
			Id              string `json:"id"`
			Time            int    `json:"time"`
			Description     string `json:"description"`
			Mcc             int    `json:"mcc"`
			OriginalMcc     int    `json:"originalMcc"`
			Amount          int    `json:"amount"`
			OperationAmount int    `json:"operationAmount"`
			CurrencyCode    int    `json:"currencyCode"`
			CommissionRate  int    `json:"commissionRate"`
			CashbackAmount  int    `json:"cashbackAmount"`
			Balance         int    `json:"balance"`
			Hold            bool   `json:"hold"`
		} `json:"statementItem"`
	} `json:"data"`
}
type T2 struct {
	Type string `json:"type"`
	Data struct {
		Account       string `json:"account"`
		StatementItem struct {
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
		} `json:"statementItem"`
	} `json:"data"`
}
type T3 struct {
	Type string `json:"type"`
	Data struct {
		Account       string `json:"account"`
		StatementItem struct {
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
		} `json:"statementItem"`
	} `json:"data"`
}
