package models

type Users struct {
	Id        int64   `gorm:"primary_key, AUTO_INCREMENT"`
	Username  string  `gorm:"not null" form:"username" json:"username"`
	Password  string  `gorm:"not null" form:"password" json:"password"`
	Firstname string  `gorm:"not null" form:"firstname" json:"firstname"`
	Lastname  string  `gorm:"not null" form:"lastname" json:"lastname"`
	Posts     []Posts `gorm:"ForeignKey:UserID"`
}

func (b *Users) TableName() string {
	return "user"
}

type Posts struct {
	Id     int64  `gorm:"primary_key, AUTO_INCREMENT"`
	Userid int64  `gorm:"not null" form:"userid" json:"userid"`
	Title  string `gorm:"not null" form:"title" json:"title"`
	Body   string `gorm:"not null" form:"body" json:"body"`
}

func (b *Posts) TableName() string {
	return "post"
}

type Login struct {
	Username string `gorm:"not null" form:"username" json:"username"`
	Password string `gorm:"not null" form:"password" json:"password"`
}
