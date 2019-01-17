package db

import "time"
import "github.com/jinzhu/gorm"

// Model 基本のテーブルモデル
type Model struct {
	gorm.Model
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

// User Userテーブル
type User struct {
	gorm.Model
	Name        string        `gorm:"not null"`
	Message     []Message     // Messageテーブルのforeign keyとする
	SendMessage []SendMessage // SendMessageテーブルのforeign keyとする
	Device      []Device      // Deviceテーブルのforeign keyとする
}

// Device Deviceテーブル
type Device struct {
	gorm.Model
	Name   string `gorm:"not null"`
	UserID uint   `gorm:"not null"`
}

//Message Messageテーブル
type Message struct {
	gorm.Model
	UserID      uint   `gorm:"not null"`
	Title       string `gorm:"not null"`
	Body        string `gorm:"not null"`
	BleID       uint   `gorm:"not null"`
	Due         time.Time
	SendMessage []SendMessage // SendMessageテーブルのforeign keyとする
}

//Ble BLEテーブル
type Ble struct {
	gorm.Model
	Name     string    `gorm:"not null"`
	AreaName string    `gorm:"not null"`
	Message  []Message // Messageテーブルのforeign keyとする
}

// SendMessage Send_messageテーブル
type SendMessage struct {
	gorm.Model
	UserID    uint `gorm:"not null"`
	MessageID uint `gorm:"not null"`
}
