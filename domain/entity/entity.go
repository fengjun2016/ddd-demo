package entity

import (
	"ddd-demo/infrastructure/auth"
	"github.com/badoux/checkmail"
	"html"
	"strings"
	"time"
)

type User struct {
	ID        int64     `gorm:"primary_key:auto_increment" json:"id"`
	FirstName string    `gorm:"size:100" json:"first_name"`
	LastName  string    `gorm:"size:100" json:"last_name"`
	Email     string    `gorm:"size:100" json:"email"`
	Password  string    `gorm:"size:100" json:"password"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"deleted_at"`
}

type Food struct {
	ID        int64     `gorm:"primary_key:auto_increment" json:"id"`
	FoodName string    `gorm:"size:100" json:"food_name"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
	DeletedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"deleted_at"`
}

type PublicUser struct {
	ID        int64  `gorm:"primary_key:auto_increment" json:"id"`
	FirstName string `gorm:"size:100;not null;" json:"first_name"`
	LastName  string `gorm:"size:100;not null;" json:"last_name"`
}

type PublicFood struct {
	ID        int64  `gorm:"primary_key:auto_increment" json:"id"`
	FoodName string `gorm:"size:100;not null;" json:"food_name"`
}

//BeforeSave is a  gorm hook
func (u *User) BeforeSave() error {
	hashPassword, err := auth.Encrypt(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashPassword)
	return nil
}

type Users []User
type Foods []Food

//So that we dont expose the user's email address and password to the world
func (users Users) PublicUsers() []interface{} {
	result := make([]interface{}, len(users))
	for index, user := range users {
		result[index] = user.PublicUser()
	}
	return result
}

//So that we dont expose the user's email address and password to the world
func (foods Foods) PublicFoods() []interface{} {
	result := make([]interface{}, len(foods))
	for index, food := range foods {
		result[index] = food.PublicFood()
	}
	return result
}
func (f *Food) PublicFood() interface{}  {
	return &PublicFood{
		ID: f.ID,
		FoodName: f.FoodName,
	}
}
//So that we dont expose the user's email address and password to the world
func (u *User) PublicUser() interface{} {
	return &PublicUser{
		ID:        u.ID,
		FirstName: u.FirstName,
		LastName:  u.LastName,
	}
}

func (u *User) Prepare() {
	u.FirstName = html.EscapeString(strings.TrimSpace(u.FirstName))
	u.LastName = html.EscapeString(strings.TrimSpace(u.LastName))
	u.Email = html.EscapeString(strings.TrimSpace(u.Email))
	u.CreatedAt = time.Now()
	u.UpdatedAt = time.Now()
}

func (u *User) Validate(action string) map[string]string {
	var errorMessages = make(map[string]string)
	var err error

	switch strings.ToLower(action) {
	case "update":
		if u.Email == "" {
			errorMessages["email_required"] = "email required"
		}

		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				errorMessages["invalidate_email"] = "email email"
			}
		}
	case "login":
		if u.Password == "" {
			errorMessages["password_required"] = ""
		}

		if u.Email == "" {
			errorMessages["email_required"] = "email is required"
		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				errorMessages["invalid_email"] = "please provide a valid email"
			}
		}
	case "forgotpassword":
		if u.Email == "" {
			errorMessages["email_required"] = "email required"
		}
		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				errorMessages["invalid_email"] = "please provide a valid email"
			}
		}
	default:
		if u.FirstName == "" {
			errorMessages["firstname_required"] = "first name required"
		}

		if u.LastName == "" {
			errorMessages["lastname_required"] = "last name is required"
		}

		if u.Password == "" {
			errorMessages["password_required"] = "password is required"
		}

		if u.Password != "" && len(u.Password) < 6 {
			errorMessages["invalid_password"] = "password should be at least 6 characters"
		}

		if u.Email == "" {
			errorMessages["email_required"] = "email is required"
		}

		if u.Email != "" {
			if err = checkmail.ValidateFormat(u.Email); err != nil {
				errorMessages["invalid_email"] = "please provide a valid email"
			}
		}
	}

	return errorMessages
}
