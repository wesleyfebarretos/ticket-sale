package override

type ProductCategory struct {
	Description *string `json:"description" example:"EVENT"`
	Name        string  `json:"name" example:"event"`
	ID          int32   `json:"id" example:"3"`
}
