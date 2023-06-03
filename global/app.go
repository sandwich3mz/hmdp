package global

import (
	"github.com/go-eden/routine"
	"github.com/go-redis/redis/v8"
	"hmdp/dto"
)

type Application struct {
	Redis *redis.Client
}

var App = new(Application)
var UserDTO dto.UserDTO
var UserInfo = routine.NewLocalStorage()
