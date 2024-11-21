package medicoErrors

const (
	InvalidEmail    = "INVALID EMAIL"
	InvalidPassword = "INVALID PASSWORD"
)

const (
	IncorrectEmail         = "the provided email is not correct"
	NotEnoughLowerCase     = "password must contain at least 2 lower case letters"
	NotEnoughUpperCase     = "password must contain at least 2 upper case letters"
	NotEnoughDigits        = "password must contain at least 2 digits"
	NotEnoughSpecialChars  = "password must contain at least 2 special characters"
	NotEnoughNumberOfChars = "password must be at least 12 characters"
	IncludedWhiteSpace     = "password must not include spaces"
)
