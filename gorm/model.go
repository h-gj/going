package gorm

type User struct {
	ID       uint `gorm:"primaryKey"`
	Username string
	Age      uint8
}
