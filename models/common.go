package models

import (
	"github.com/google/uuid"
	"time"
)

// ModelID represents the type of the primary key / foreign key
type ModelID uuid.UUID

// Text represents the type of text fields
type Text string

// DateTime represents the type of time as unix timestamp
type DateTime time.Time

// WholeQuantity represents the type of quantity as unsigned integer
type WholeQuantity uint16

type WeekDay time.Weekday
