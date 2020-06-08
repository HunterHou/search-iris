package datamodels

type Supplier struct {
	Name string
	Cnt  int
}

func NewSupplier(name string) Supplier {
	return Supplier{
		Name: name,
		Cnt:  1,
	}
}
func (su *Supplier) Plus() {
	su.Cnt = su.Cnt + 1
}
