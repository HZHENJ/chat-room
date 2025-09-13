package db

import (
	"chatting-room/pkg/conf"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	gormopentracing "gorm.io/plugin/opentracing"
)

var (
	DB *gorm.DB
)

// Init Mysql DB
func Init() {
	var err error
	DB, err = gorm.Open(postgres.Open(conf.GetConf().GetString("db.postgres.dsn")),
		&gorm.Config{
			PrepareStmt:            true,
			SkipDefaultTransaction: true,
		},
	)
	if err != nil {
		panic(err)
	}

	if err = DB.Use(gormopentracing.New()); err != nil {
		panic(err)
	}

	// MVP version：User、Conversation、ConversationMember
	if err := DB.AutoMigrate(
		&User{},
		&Conversation{},
		&Member{},
	); err != nil {
		panic("failed to migrate tables: " + err.Error())
	}

}
