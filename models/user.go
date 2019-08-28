package models

import (
	"errors"
	"github.com/rburawes/golang-gorm/config"
	"golang.org/x/crypto/bcrypt"
	"net/http"
)

// User object handles information about application's registered sessions.
type User struct {
	ID         uint
	Username   string
	Email      string
	Password   string
	Firstname  string
	Middlename string
	Lastname   string
}

// CheckUser looks for existing user using email or username
func CheckUser(un string) (User, error) {

	usr, ok := FindUser(un)

	if ok {
		return usr, errors.New("username or email is taken")
	}

	return usr, nil

}

// FindUser looks for registerd user by username.
func FindUser(un string) (User, bool) {

	u := User{}

	if config.Database.Where(&User{Username: un}).Find(&u).Error != nil {
		return u, false
	}

	return u, true

}

// SaveUser create new user entry.
func SaveUser(r *http.Request) (User, error) {

	// Get form values and validate
	u, err := validateUserForm(r)

	if err != nil {
		return u, err
	}

	bs, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.MinCost)

	if err != nil {
		return u, errors.New("the provided password is not valid")
	}

	u.Password = string(bs)

	if config.Database.Create(&u).Error != nil {
		return u, errors.New("unable to process registration")
	}

	return u, nil
}

// ValidateForm validates the submitted form for registration
func validateUserForm(r *http.Request) (User, error) {

	u := User{}
	p := r.FormValue("password")
	cp := r.FormValue("cpassword")
	f := r.FormValue("firstname")
	l := r.FormValue("lastname")
	e := r.FormValue("email")

	if p != cp {
		return u, errors.New("password does not match")
	}

	if e == "" || p == "" || cp == "" {
		return u, errors.New("fields cannot be empty")
	}

	_, err := CheckUser(e)

	if err != nil {
		return u, err
	}

	u.Username = e
	u.Email = e
	u.Firstname = f
	u.Lastname = l
	u.Password = p

	return u, nil

}

// ValidatePassword validates the input password against the one in the database.
func (u *User) ValidatePassword(p string) bool {

	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(p))
	if err != nil {
		return false
	}

	return true

}
