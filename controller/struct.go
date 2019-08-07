package controller

import (
	"time"
)

//Infos is ..
type Infos struct {
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
