package controller

import (
	"time"
)

//Thongtins is ..
type Thongtins struct {
	ID         int
	Username   string
	Password   string
	Email      string
	Birth      time.Time
	Sex        string
	Phone      string
	NationalID int
	Height     float64
}
