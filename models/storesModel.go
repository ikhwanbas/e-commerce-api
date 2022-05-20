package models

import (
	"ecommerce/utils/token"
	"time"

	"github.com/google/uuid"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type (
	Stores struct {
		ID            uuid.UUID     `gorm:"primary_key" json:"id"`
		StoreName     string        `json:"store_name"`
		Email         string        `gorm:"unique" json:"email"`
		Password      string        `json:"password"`
		StoresProfile StoresProfile `gorm:"foreignKey:ID"`
		ProductsID    []Products
	}

	StoresProfile struct {
		ID          uuid.UUID `gorm:"primary_key" json:"id"`
		Address     string    `json:"address"`
		PhoneNumber int       `json:"phone_number"`
		City        string    `json:"city"`
		PostalCode  int       `json:"postal_code"`
		CreatedAt   time.Time `json:"created_at"`
		UpdatedAt   time.Time `json:"updated_at"`
	}

	StoreRegisterInput struct {
		StoreName string `json:"store_name"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	StoreLoginInput struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	StoreProfileInput struct {
		Fullname    string `json:"fullname"`
		Address     string `json:"address"`
		PhoneNumber int    `json:"phone_number"`
		City        string `json:"city"`
		PostalCode  int    `json:"postal_code"`
	}
)

func (s *Stores) SaveStore(db *gorm.DB) (*Stores, error) {
	//turn password into hash
	hashedPassword, errPassword := bcrypt.GenerateFromPassword([]byte(s.Password), bcrypt.DefaultCost)
	if errPassword != nil {
		return &Stores{}, errPassword
	}
	// generate uuid
	s.ID = uuid.New()
	// hashed password
	s.Password = string(hashedPassword)

	// create store data
	var err error = db.Create(&s).Error
	if err != nil {
		return &Stores{}, err
	}

	// generate store profile data
	storeProfile := StoresProfile{ID: s.ID}
	err = db.Create(&storeProfile).Error
	if err != nil {
		return &Stores{}, err
	}

	return s, nil
}

func StoreVerifyPassword(password, hashedPassword string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func StoreLoginCheck(email string, password string, db *gorm.DB) (string, error) {

	var err error

	s := Stores{}

	err = db.Model(Stores{}).Where("email = ?", email).Take(&s).Error

	if err != nil {
		return "", err
	}

	err = StoreVerifyPassword(password, s.Password)

	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}

	token, err := token.GenerateToken(s.ID)

	if err != nil {
		return "", err
	}

	return token, nil

}
