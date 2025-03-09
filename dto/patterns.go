package dto

const (
	lowerCasePattern   = `.*[[:lower:]]+.*[[:lower:]]+.*`
	upperCasePattern   = `.*[[:upper:]]+.*[[:upper:]]+.*`
	digitPattern       = `.*[[:digit:]]+.*[[:digit:]]+.*`
	specialCharPattern = `.*[[:punct:]]+.*[[:punct:]]+.*`
	spacePattern       = `[[:space:]]`
)

const (
	emailPattern = `^[a-z0-9!#$%&'*+/=?^_` + "`" + `{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_` + "`" + `{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?`
)

const (
	atcPattern = `\w\d\d\w\w\d\d`
)
