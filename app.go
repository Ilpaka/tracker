package main

import (
	"context"
	"errors"
	"time"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// App struct
type App struct {
	ctx context.Context
	db  *gorm.DB
}

// NewApp creates a new App application struct
func NewApp(db *gorm.DB) *App {
	return &App{db: db}
}

// startup is called when the app starts
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// User methods
type LoginRequest struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
	UserID  uint   `json:"userId"`
}

// Register creates a new user
func (a *App) Register(req LoginRequest) LoginResponse {
	// Check if user exists
	var existingUser User
	if err := a.db.Where("username = ?", req.Username).First(&existingUser).Error; err == nil {
		return LoginResponse{Success: false, Message: "User already exists"}
	}

	// Hash password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return LoginResponse{Success: false, Message: "Error creating user"}
	}

	user := User{
		Username: req.Username,
		Password: string(hashedPassword),
	}

	if err := a.db.Create(&user).Error; err != nil {
		return LoginResponse{Success: false, Message: "Error creating user"}
	}

	return LoginResponse{Success: true, Message: "User created successfully", UserID: user.ID}
}

// Login authenticates a user
func (a *App) Login(req LoginRequest) LoginResponse {
	var user User
	if err := a.db.Where("username = ?", req.Username).First(&user).Error; err != nil {
		return LoginResponse{Success: false, Message: "Invalid username or password"}
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
		return LoginResponse{Success: false, Message: "Invalid username or password"}
	}

	return LoginResponse{Success: true, Message: "Login successful", UserID: user.ID}
}

// Habit methods
type HabitRequest struct {
	UserID      uint   `json:"userId"`
	Name        string `json:"name"`
	Description string `json:"description"`
	Icon        string `json:"icon"`
}

type HabitResponse struct {
	ID          uint      `json:"id"`
	Name        string    `json:"name"`
	Description string    `json:"description"`
	Icon        string    `json:"icon"`
	CreatedAt   time.Time `json:"createdAt"`
}

// CreateHabit creates a new habit
func (a *App) CreateHabit(req HabitRequest) (HabitResponse, error) {
	habit := Habit{
		UserID:      req.UserID,
		Name:        req.Name,
		Description: req.Description,
		Icon:        req.Icon,
	}

	if err := a.db.Create(&habit).Error; err != nil {
		return HabitResponse{}, err
	}

	return HabitResponse{
		ID:          habit.ID,
		Name:        habit.Name,
		Description: habit.Description,
		Icon:        habit.Icon,
		CreatedAt:   habit.CreatedAt,
	}, nil
}

// GetHabits returns all habits for a user
func (a *App) GetHabits(userID uint) ([]HabitResponse, error) {
	var habits []Habit
	if err := a.db.Where("user_id = ?", userID).Find(&habits).Error; err != nil {
		return nil, err
	}

	response := make([]HabitResponse, len(habits))
	for i, habit := range habits {
		response[i] = HabitResponse{
			ID:          habit.ID,
			Name:        habit.Name,
			Description: habit.Description,
			Icon:        habit.Icon,
			CreatedAt:   habit.CreatedAt,
		}
	}

	return response, nil
}

// DeleteHabit deletes a habit
func (a *App) DeleteHabit(habitID uint, userID uint) error {
	result := a.db.Where("id = ? AND user_id = ?", habitID, userID).Delete(&Habit{})
	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return errors.New("habit not found")
	}
	return nil
}

// Track methods
type TrackRequest struct {
	HabitID uint   `json:"habitId"`
	UserID  uint   `json:"userId"`
	Date    string `json:"date"` // YYYY-MM-DD format
}

type TrackResponse struct {
	ID        uint      `json:"id"`
	HabitID   uint      `json:"habitId"`
	Date      time.Time `json:"date"`
	Completed bool      `json:"completed"`
}

// ToggleTrack toggles habit completion for a specific date
func (a *App) ToggleTrack(req TrackRequest) (TrackResponse, error) {
	// Parse date
	date, err := time.Parse("2006-01-02", req.Date)
	if err != nil {
		return TrackResponse{}, err
	}

	// Check if habit belongs to user
	var habit Habit
	if err := a.db.Where("id = ? AND user_id = ?", req.HabitID, req.UserID).First(&habit).Error; err != nil {
		return TrackResponse{}, errors.New("habit not found")
	}

	// Check if track exists
	var track Track
	err = a.db.Where("habit_id = ? AND date = ?", req.HabitID, date).First(&track).Error

	if err == gorm.ErrRecordNotFound {
		// Create new track
		track = Track{
			HabitID:   req.HabitID,
			Date:      date,
			Completed: true,
		}
		if err := a.db.Create(&track).Error; err != nil {
			return TrackResponse{}, err
		}
	} else if err != nil {
		return TrackResponse{}, err
	} else {
		// Toggle existing track
		track.Completed = !track.Completed
		if err := a.db.Save(&track).Error; err != nil {
			return TrackResponse{}, err
		}
	}

	return TrackResponse{
		ID:        track.ID,
		HabitID:   track.HabitID,
		Date:      track.Date,
		Completed: track.Completed,
	}, nil
}

// GetTracks returns all tracks for a habit
func (a *App) GetTracks(habitID uint, userID uint) ([]TrackResponse, error) {
	// Check if habit belongs to user
	var habit Habit
	if err := a.db.Where("id = ? AND user_id = ?", habitID, userID).First(&habit).Error; err != nil {
		return nil, errors.New("habit not found")
	}

	var tracks []Track
	if err := a.db.Where("habit_id = ?", habitID).Find(&tracks).Error; err != nil {
		return nil, err
	}

	response := make([]TrackResponse, len(tracks))
	for i, track := range tracks {
		response[i] = TrackResponse{
			ID:        track.ID,
			HabitID:   track.HabitID,
			Date:      track.Date,
			Completed: track.Completed,
		}
	}

	return response, nil
}
