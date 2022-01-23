package config

type User struct{
	Name string `json:"name" form:"name" query:"name"`
	LastName string `json:"lastname" form:"lastname" query:"lastname"`
	Email string `json:"email" form:"email" query:"email"`
	Password string `json:"password" form:"password" query:"password"`
	Age int `json:"age" form:"age" query:"age"`
	IsActive bool `json:"isactive" form:"isactive" query:"isactive"`
}

type UserList struct{
	UserList []User
}
