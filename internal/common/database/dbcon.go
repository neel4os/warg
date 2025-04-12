package database

import (
	"sync"

	"github.com/neel4os/warg/internal/common/config"
	"github.com/rs/zerolog/log"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	instance *DataConn
	once     sync.Once
)

type DataConn struct {
	Db *gorm.DB
}

func newDataConn(cfg config.Config) *DataConn {
	log.Debug().Caller().Msg("Creating new db connection")
	dbconfig := cfg.DbConfig
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN: dbconfig.GetDbDsn(),
	}), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	return &DataConn{
		Db: db,
	}
}

func GetDataConn(cfg config.Config) *DataConn {
	once.Do(func() {
		instance = newDataConn(cfg)
	})
	return instance
}

func (d *DataConn) GetDb() *gorm.DB {
	return d.Db
}

func (d *DataConn) Close() {
	// close the db connection
	sqlDB, err := d.Db.DB()
	if err != nil {
		panic(err)
	}
	err = sqlDB.Close()
	if err != nil {
		panic(err)
	}
}

func (d *DataConn) Ping() {
	// ping the db connection
	sqlDB, err := d.Db.DB()
	if err != nil {
		panic(err)
	}
	err = sqlDB.Ping()
	if err != nil {
		panic(err)
	}
}
