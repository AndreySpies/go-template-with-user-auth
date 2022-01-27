package request

type CreateUser struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
	Email     string `json:"email" validate:"required,email"`
	Password  string `json:"password" validate:"required,min=8,gte=8,lte=24"`
	Birthday  string `json:"birthday" validate:"required"`
	Location  string `json:"location"`
	CreatedAt string `json:"created_at"`
}

type Login struct {
	Email    string `json:"email" validate:"required"`
	Password string `json:"password" validate:"required"`
}
