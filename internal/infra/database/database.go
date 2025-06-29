package database

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/rs/zerolog/log"
	"github.com/vnurhaqiqi/go-echo-starter/configs"
)

type TransactionBlock func(db *sqlx.Tx, c chan error)

type MySQLConn struct {
	MySQL *sqlx.DB
}

func ProvideMySQLConn(config *configs.Config) *MySQLConn {
	return &MySQLConn{
		MySQL: NewMySQLDBConnection(
			config.DB.MySQL.User,
			config.DB.MySQL.Password,
			config.DB.MySQL.Host,
			config.DB.MySQL.Port,
			config.DB.MySQL.Name,
			config.DB.MySQL.MaxConnLifetime,
			config.DB.MySQL.MaxIdleConn,
			config.DB.MySQL.MaxOpenConn,
			config.DB.MySQL.TimeZone,
		),
	}
}

func NewMySQLDBConnection(
	username,
	password,
	host,
	port,
	dbName string,
	maxConnLifetime time.Duration,
	maxIdleConn,
	maxOpenConn int,
	timeZone string) *sqlx.DB {
	conn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8&loc=%s&parseTime=true",
		username,
		password,
		host,
		port,
		dbName,
		timeZone,
	)

	db, err := sqlx.Connect("mysql", conn)
	if err != nil {
		log.
			Fatal().
			Err(err).
			Str("host", host).
			Str("port", port).
			Str("dbName", dbName).
			Msg("Failed connecting to MySQL database")
	} else {
		log.
			Info().
			Str("host", host).
			Str("port", port).
			Str("dbName", dbName).
			Msg("Connected to MySQL database")
	}

	db.SetConnMaxLifetime(maxConnLifetime)
	db.SetMaxIdleConns(maxIdleConn)
	db.SetMaxOpenConns(maxOpenConn)

	return db
}

func (m *MySQLConn) WithTransaction(block TransactionBlock) (err error) {
	e := make(chan error)
	tx, err := m.MySQL.Beginx()
	if err != nil {
		log.Err(err).
			Msg("[MySQLConn][WithTransaction] error begin transaction")
		return
	}
	go block(tx, e)
	err = <-e
	if err != nil {
		if errTx := tx.Rollback(); errTx != nil {
			log.Err(errTx).
				Msg("[MySQLConn][WithTransaction] error rollback transaction")
		}
		return
	}
	err = tx.Commit()
	return
}
