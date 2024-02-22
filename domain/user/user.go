package user

type Users struct {
	Id    string `json:"id,omitempty"`
	Name  string `json:"name" validate:"required,max=32"`
	Email string `json:"email" validate:"required,max=32"`
	Age   string `json:"age" validate:"required"`
}

func NewUser(id string,
	name string,
	email string,
	age string,
) *Users {
	return &Users{
		Id:    id,
		Name:  name,
		Email: email,
		Age:   age,
	}

}
