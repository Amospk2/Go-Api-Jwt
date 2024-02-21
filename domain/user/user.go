package user

type Users struct {
	Id    string `json:"id,omitempty"`
	Name  string `json:"name,omitempty"`
	Email string `json:"email,omitempty"`
	Age   string `json:"age,omitempty"`
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
