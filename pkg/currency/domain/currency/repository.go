package currency

type Repository interface {
	// All methods that save and retrieve Currency domain
	GetCurrencyByName(string) (*Currency, error)
	//SaveMonobankAction(action *MonobankAction)
	//GetCurrencyTable() (Table, error)
}
