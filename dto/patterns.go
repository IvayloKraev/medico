package dto

const (
	lowerCasePattern   = `.*[[:lower:]]+.*[[:lower:]]+.*`
	upperCasePattern   = `.*[[:upper:]]+.*[[:upper:]]+.*`
	digitPattern       = `.*[[:digit:]]+.*[[:digit:]]+.*`
	specialCharPattern = `.*[[:punct:]]+.*[[:punct:]]+.*`
	numOfCharPattern   = `^[[:graph:]]{12,}$`
	spacePattern       = `[[:space:]]`
)

const (
	emailPattern = `^[a-z0-9!#$%&'*+/=?^_` + "`" + `{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_` + "`" + `{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?`
)
