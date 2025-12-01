package main

import (
	"time"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// User model
type User struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	Username  string    `gorm:"unique;not null" json:"username"`
	Password  string    `gorm:"not null" json:"-"`
	Habits    []Habit   `gorm:"foreignKey:UserID" json:"habits"`
}

// Habit model
type Habit struct {
	ID          uint      `gorm:"primarykey" json:"id"`
	CreatedAt   time.Time `json:"createdAt"`
	UserID      uint      `gorm:"not null" json:"userId"`
	Name        string    `gorm:"not null" json:"name"`
	Description string    `json:"description"`
	Icon        string    `json:"icon"`
	Tracks      []Track   `gorm:"foreignKey:HabitID" json:"tracks"`
}

// Track model - represents daily completion
type Track struct {
	ID        uint      `gorm:"primarykey" json:"id"`
	CreatedAt time.Time `json:"createdAt"`
	HabitID   uint      `gorm:"not null" json:"habitId"`
	Date      time.Time `gorm:"not null;index" json:"date"`
	Completed bool      `gorm:"default:true" json:"completed"`
}

// InitDB initializes the database
func InitDB() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("tracker.db"), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	// Auto migrate schemas
	err = db.AutoMigrate(&User{}, &Habit{}, &Track{})
	if err != nil {
		return nil, err
	}

	return db, nil
}
