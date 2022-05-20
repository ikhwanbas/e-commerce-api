package models

import (
	"ecommerce/utils/token"
	"html"
	"strings"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	Users struct {
		ID           uuid.UUID    `gorm:"primary_key" json:"id"`
		Username     string       `json:"username"`
		Email        string       `gorm:"unique" json:"email"`
		Password     string       `json:"password"`
		UsersProfile UsersProfile `gorm:"foreignKey:ID"`
		// db migration
		CartsID []Carts
	}

	UsersProfile struct {
		ID          uuid.UUID `gorm:"primary_key" json:"id"`
		Fullname    string    `json:"fullname"`
		Address     string    `json:"address"`
		PhoneNumber int       `json:"phone_number"`
		City        string    `json:"city"`
		PostalCode  int       `json:"postal_code"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	RegisterInput struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	LoginInput struct {
		Username string `json:"username"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	UsersProfileInput struct {
		Fullname    string `json:"fullname"`
		Address     string `json:"address"`
		PhoneNumber int    `json:"phone_number"`
		City        string `json:"city"`
		PostalCode  int    `json:"postal_code"`
	}
)

// function for register and login
func VerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func LoginCheck(email string, username string, password string, db *gorm.DB) (string, error) {

	var err error

	u := Users{}

	err = db.Model(Users{}).Where("username = ?", username).Or("email = ?", email).Take(&u).Error

	if err != nil {
		return "", err
	}

	err = VerifyPassword(password, u.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(u.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}

func (u *Users) SaveUser(db *gorm.DB) (*Users, error) {
	//turn password into hash
	hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if errPassword != nil {
		return &Users{}, errPassword
	}
	// generate uuid
	u.ID = uuid.New()
	// hashed password
	u.Password = string(hashedPassword)
	//remove spaces in username
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))

	// create user data
	var err error = db.Create(&u).Error
	if err != nil {
		return &Users{}, err
	}

	// generate user profile data
	userProfile := UsersProfile{ID: u.ID}
	err = db.Create(&userProfile).Error
	if err != nil {
		return &Users{}, err
	}

	return u, nil
}
