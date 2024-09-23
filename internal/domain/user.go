package domain

type User struct {
	ID       uint   `gorm:"primaryKey"`
	Login    string `gorm:"size:32;not null;unique;check:char_length(login) >= 5" json:"name"`
	Password string `gorm:"size:255;not null"`
}
