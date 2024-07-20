package electronic

type androidPhone struct {
	brand     string
	model     string
	phoneType string
	os        string
}

func (a androidPhone) Brand() string {
	return a.brand
}

func (a androidPhone) Model() string {
	return a.model
}

func (a androidPhone) Type() string {
	return a.phoneType
}

func (a androidPhone) OS() string {
	return a.os
}

func NewAndroidPhone(brand, model, os string) androidPhone {
	return androidPhone{
		phoneType: "smartphone",
		brand:     brand,
		model:     model,
		os:        os,
	}
}
