package model

type BaseModel struct {
	ID        int64 `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	CreatedAt int64 `json:"created_at"`
	UpdatedAt int64 `json:"updated_at"`
	DeletedAt int64 `json:"-"`
}
