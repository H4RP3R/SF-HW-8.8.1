package electronic

type applePhone struct {
	brand     string
	model     string
	phoneType string
	os        string
}

func (a applePhone) Brand() string {
	return a.brand
}

func (a applePhone) Model() string {
	return a.model
}

func (a applePhone) Type() string {
	return a.phoneType
}

func (a applePhone) OS() string {
	return a.os
}

func NewApplePhone(model, os string) applePhone {
	return applePhone{
		brand:     "apple",
		phoneType: "smartphone",
		model:     model,
		os:        os,
	}
}
