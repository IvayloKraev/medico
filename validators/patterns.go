package validators

const (
	LowerCasePattern   = `.*[[:lower:]]+.*[[:lower:]]+.*`
	UpperCasePattern   = `.*[[:upper:]]+.*[[:upper:]]+.*`
	DigitPattern       = `.*[[:digit:]]+.*[[:digit:]]+.*`
	SpecialCharPattern = `.*[[:punct:]]+.*[[:punct:]]+.*`
	NumOfCharPattern   = `^[[:graph:]]{12,}$`
	SpacePattern       = `[[:space:]]`
)

const (
	EmailPattern = `^[a-z0-9!#$%&'*+/=?^_` + "`" + `{|}~-]+(?:\\.[a-z0-9!#$%&'*+/=?^_` + "`" + `{|}~-]+)*@(?:[a-z0-9](?:[a-z0-9-]*[a-z0-9])?\.)+[a-z0-9](?:[a-z0-9-]*[a-z0-9])?`
)
