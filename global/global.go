package global

import (
	"github.com/go-redis/redis"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"wm-take-out/config"
)

var (
	Config *config.AllConfig // 全局Config
	Log    logger.ILog
	DB     *gorm.DB
	Redis  *redis.Client
)
