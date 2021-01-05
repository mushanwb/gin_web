package model

type BaseModel struct {
	ID        int64 `gorm:"primaryKey;autoIncrement;not null" json:"id"`
	CreatedAt int32 `json:"created_at"`
	UpdatedAt int32 `json:"updated_at"`
	DeletedAt int32 `json:"-"`
}
