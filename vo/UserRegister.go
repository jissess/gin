package vo

type UserRegister struct {
	UserName  string `json:"user_name"`
	Telephone string `json:"telephone" binding:"required"`
	Password  string `json:"password" binding:"required"`
}
