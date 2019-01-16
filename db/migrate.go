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
	Name        string
	DeviceID    uint
	Password    string
	Message     Message     // Messageテーブルのforeign keyとする
	SendMessage SendMessage // SendMessageテーブルのforeign keyとする
}

// Device Deviceテーブル
type Device struct {
	gorm.Model
	Name string
	User User // Userテーブルのforeign keyとする
}

//Message Messageテーブル
type Message struct {
	gorm.Model
	UserID      uint
	Title       string
	Body        string
	BleID       uint
	Due         time.Time
	SendMessage SendMessage // SendMessageテーブルのforeign keyとする
}

//Ble BLEテーブル
type Ble struct {
	gorm.Model
	Name     string
	AreaName string
	Message  Message // Messageテーブルのforeign keyとする
}

// SendMessage Send_messageテーブル
type SendMessage struct {
	gorm.Model
	UserID    uint
	MessageID uint
}
