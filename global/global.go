package global

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"wm-take-out/config"
	"wm-take-out/pkg/log"
)

var (
	Config *config.AllConfig // 全局Config
	Log    log.ILog
	DB     *gorm.DB
	Redis  *redis.Client
)
