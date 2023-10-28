package database

import (
	"github.com/rs/zerolog/log"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"navicstein/private-gpt/internal/config"
	"navicstein/private-gpt/internal/database/model"
)

func Setup() (*gorm.DB, error) {
	// DB gorm connector
	var (
		DB  *gorm.DB
		err error
		cfg = config.GetConfig()
	)

	gormCfg := &gorm.Config{
		Logger: logger.Default.LogMode(logger.Error),
	}

	switch {
	case cfg.Switchers.DBDriver == "sqlite":
		DB, err = gorm.Open(sqlite.Open("db.sqlite"), gormCfg)
	default:
		panic("unknown driver")
	}

	if err != nil {
		return nil, err
	}

	log.Info().Str("driver", cfg.Switchers.DBDriver).Msg("Connection Opened to Database")
	_ = DB.AutoMigrate(&model.Document{}, &model.Conversation{})
	log.Debug().Msg("Database migrated successfully")

	return DB, nil
}
