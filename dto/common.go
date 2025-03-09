package dto

type Modeler interface {
	FromModel(interface{}) error
	ToModel(interface{}) error
}

type Validator interface {
	Validate() error
}

type Defaulter interface {
	ToDefault()
}
