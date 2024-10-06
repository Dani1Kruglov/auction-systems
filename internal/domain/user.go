package domain

import (
	"errors"
)

type User struct {
	Id              int
	Name            string
	Role            string
	Email           string
	Password        string
	NotificationUrl string
}

func (u *User) Validate() error {
	if u.Name == "" {
		return errors.New("name is required")
	}
	if u.Email == "" {
		return errors.New("email is required")
	}
	if u.Password == "" {
		return errors.New("password is required")
	}
	if u.Role != SellerRole && u.Role != BuyerRole {
		return errors.New("role is invalid")
	}
	if u.NotificationUrl == "" {
		return errors.New("notification_url is required")
	}
	return nil
}
