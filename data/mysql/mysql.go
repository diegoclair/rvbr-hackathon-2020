package mysql

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/GuiaBolso/darwin"
	"github.com/RedVentures/rvbr-2020-hackathon-time-3-backend/data/migrations"
	"github.com/RedVentures/rvbr-2020-hackathon-time-3-backend/domain/contract"
	"github.com/RedVentures/rvbr-2020-hackathon-time-3-backend/infra/config"
	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/go-sql-driver/mysql"
)

// DBManager is the MySQL connection manager
type DBManager struct {
	db *sql.DB
}

//Instance returns an instance of a RepoManager
func Instance() (contract.RepoManager, error) {
	cfg := config.GetDBConfig()

	dataSourceName := fmt.Sprintf("%s:root@tcp(%s:%s)/%s?charset=utf8&parseTime=true",
		cfg.Username, cfg.Host, cfg.Port, cfg.DBName,
	)

	log.Println("Connecting to database...")
	log.Println("Connection String: ", dataSourceName)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	if _, err = db.Exec("CREATE DATABASE IF NOT EXISTS health_db;"); err != nil {
		logger.Error("Create Database error: ", err)
		return nil, err
	}

	if _, err = db.Exec("USE health_db;"); err != nil {
		logger.Error("Default Database error: ", err)
		return nil, err
	}

	err = mysql.SetLogger(logger.GetLogger())
	if err != nil {
		return nil, err
	}
	logger.Info("Database successfully configured")

	logger.Info("Running the migrations")
	driver := darwin.NewGenericDriver(db, darwin.MySQLDialect{})

	d := darwin.New(driver, migrations.Migrations, nil)

	err = d.Migrate()
	if err != nil {
		logger.Error("Migrate Error: ", err)
		return nil, err
	}

	logger.Info("Migrations executed")

	instance := &DBManager{
		db: db,
	}

	return instance, nil
}

//Ping returns the ping set
func (c *DBManager) Ping() contract.PingRepo {
	return nil
}

//User returns the company set
func (c *DBManager) User() contract.UserRepo {
	return newUserRepo(c.db)
}

func (c *DBManager) HealthCheck() contract.HealthCheckRepo {
	return newHealthCheckRepo(c.db)
}
