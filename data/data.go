package data

import (
	"github.com/RedVentures/rvbr-2020-hackathon-time-3-backend/data/mysql"
	"github.com/RedVentures/rvbr-2020-hackathon-time-3-backend/domain/contract"
)

// Connect returns a instace of cassandra db
func Connect() (contract.RepoManager, error) {
	return mysql.Instance()
}
