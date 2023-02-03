package db_mysql

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/sharding"
	"gorm.io/gorm/logger"
	"gorm.io/plugin/opentelemetry/logging/logrus"
	"gorm.io/plugin/opentelemetry/tracing"
	"log"
	// constants "runedance_douyin/pkg/consts"
	"time"
)

var db *gorm.DB

// Init init DB
func Init() {
	var err error
	gormlogrus := logger.New(
		logrus.NewWriter(),
		logger.Config{
			SlowThreshold: time.Millisecond,
			Colorful:      false,
			LogLevel:      logger.Info,
		},
	)
	db, err = gorm.Open(mysql.Open("root:123456@tcp(localhost:3306)/douyin?charset=utf8&parseTime=True&loc=Local"),
		&gorm.Config{
			PrepareStmt: true,
			Logger:      gormlogrus,
		},
	)

	// sharding
	db.Use(sharding.Register(sharding.Config{
		ShardingKey: "user_to_user",
		NumberOfShards:      64,
		PrimaryKeyGenerator: sharding.PKSnowflake,
	}))

	if err != nil {
		log.Println("[gorm.Open error] err=", err)
		panic(err)
	}

	if err := db.Use(tracing.NewPlugin()); err != nil {
		panic(err)
	}
}