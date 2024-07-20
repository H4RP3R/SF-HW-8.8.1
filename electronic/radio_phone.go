package electronic

type radioPhone struct {
	brand     string
	model     string
	phoneType string
	btnCnt    int
}

func (r radioPhone) Brand() string {
	return r.brand
}

func (r radioPhone) Model() string {
	return r.model
}

func (r radioPhone) Type() string {
	return r.phoneType
}

func (r radioPhone) ButtonsCount() int {
	return r.btnCnt
}

func NewRadioPhone(brand, model string, btnCnt int) radioPhone {
	return radioPhone{
		phoneType: "station",
		brand:     brand,
		model:     model,
		btnCnt:    btnCnt,
	}
}
