package models

type User struct {
	BaseModel
	Email    string   `gorm:"index;not null"`
	Username string   `gorm:"index;not null"`
	Password string   `gorm:"not null"`
	ApiKeys  []ApiKey `gorm:"foreignKey:UserId"`
}
