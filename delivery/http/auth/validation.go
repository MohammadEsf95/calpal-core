package auth

import (
	"calpal-core/entity"
	"errors"
	"regexp"
)

func ValidateUser(user entity.User) error {
	if err := validateFirstName(user.FirstName); err != nil {
		return err
	}
	if err := validateLastName(user.LastName); err != nil {
		return err
	}
	if err := validateEmail(user.Email); err != nil {
		return err
	}
	if err := validatePassword(user.Password); err != nil {
		return err
	}
	return nil
}

func validateFirstName(firstName string) error {
	if len(firstName) < 2 || len(firstName) > 50 {
		return errors.New("first name must be between 2 and 50 characters")
	}
	return nil
}

func validateLastName(lastName string) error {
	if len(lastName) < 2 || len(lastName) > 50 {
		return errors.New("last name must be between 2 and 50 characters")
	}
	return nil
}

func validateEmail(email string) error {
	// Simple email regex for validation
	const emailRegex = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	re := regexp.MustCompile(emailRegex)
	if !re.MatchString(email) {
		return errors.New("invalid email format")
	}
	return nil
}

func validatePassword(password string) error {
	if len(password) < 8 {
		return errors.New("password must be at least 8 characters long")
	}
	return nil
}
