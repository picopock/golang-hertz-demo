package dal

import (
	"hertz_demo/internal/dal/mysql"
	"hertz_demo/internal/dal/redis"
)

func Init() {
	mysql.Init()
	redis.Init()
}
