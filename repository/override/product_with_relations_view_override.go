package override

type ProductInstallments struct {
	Creditcard  []PaymentTypeInstallment `json:"creditcard"`
	PaymentSlip []PaymentTypeInstallment `json:"paymentSlip"`
	Pix         []PaymentTypeInstallment `json:"pix"`
}

type PaymentTypeInstallment struct {
	InstallmentTimeID   int32   `json:"installmentTimeId" example:"1"`
	InstallmentTimeName string  `json:"installmentTimeName" example:"1x"`
	Fee                 float64 `json:"fee" example:"3.22"`
	Tariff              float64 `json:"tariff" example:"7.00"`
}
