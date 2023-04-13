package currency

func NewTable(currencies []Currency) *Table {
	return &Table{}
}

type Table struct {
	currencies []Currency
}

func (t *Table) GetCurrency(name string) (Currency, error) {

	panic("implement me")
}
