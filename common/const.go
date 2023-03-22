package common

import "log"

const (
	DbTypeRestaurant = 1
	DbTypeUser       = 2
)

func AppRecover() {
	if err := recover(); err != nil {
		log.Println("Recovered error:", err)
	}
}

type Requester interface {
	GetUserId() int
	GetEmail() string
	GetRole() string
}
