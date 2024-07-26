package override

type CreditcardFlag struct {
	Id          int32
	Name        string
	Description *string
	Regex       string
}

type CreditcardType struct {
	Id   int32
	Name string
}
