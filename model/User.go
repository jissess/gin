package model

type User struct {
	Id        uint64 `json:"id" gorm:"primary_key autoIncrement"`
	UserName  string `json:"user_name" gorm:"type:varchar(20)"`
	Telephone string `json:"telephone" gorm:"type:varchar(11);unique"`
	Password  string `json:"password" gorm:"type:varchar(100)"`
	CreatedAt Time   `json:"created_at" gorm:"type:datetime"`
	UpdatedAt Time   `json:"updated_at" gorm:"type:datetime"`
	DeletedAt Time   `json:"deleted_at" gorm:"type:datetime"`
}
