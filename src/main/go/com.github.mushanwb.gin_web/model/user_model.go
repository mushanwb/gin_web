package model

type User struct {
	BaseModel

	Name     string `gorm:"type:varchar(64); not null; unique" json:"name"`
	Phone    string `gorm:"type:varchar(11); not null; unique" json:"phone"`
	Password string `gorm:"type:varchar(255); not null;" json:"-"`
}
