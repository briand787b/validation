package validation

import (
	"errors"
)

type ValidationError error

var (
	ErrNoUsername           = ValidationError(errors.New("You must supply a username"))
	ErrNoEmail              = ValidationError(errors.New("You must supply an email"))
	ErrNoPassword           = ValidationError(errors.New("You must supply a password"))
	ErrPasswordTooShort     = ValidationError(errors.New("Your password is too short"))
	ErrUsernameExists       = ValidationError(errors.New("That username is taken"))
	ErrEmailExists          = ValidationError(errors.New("That email address has an account"))
	ErrCredentialsIncorrect = ValidationError(errors.New("We couldn't find a user with the supplied username and password combination"))
	ErrPasswordIncorrect    = ValidationError(errors.New("Password did not match"))
	ErrInvalidImageType     = ValidationError(errors.New("Please upload only jpeg, gif, or png images"))
	ErrNoImage              = ValidationError(errors.New("Please select and image to upload"))
	ErrImageURLInvalid      = ValidationError(errors.New("Couldn't download image from the URL you provided"))
)

func IsValidationError(err error) bool {
	_, ok := err.(ValidationError)
	return ok
}

