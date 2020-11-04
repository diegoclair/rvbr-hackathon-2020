package mysql

import (
	"database/sql"

	"github.com/RedVentures/rvbr-2020-hackathon-time-3-backend/domain/entity"
	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/go_utils-lib/mysqlutils"
	"github.com/diegoclair/go_utils-lib/resterrors"
)

type userRepo struct {
	db *sql.DB
}

// newUserRepo returns a instance of dbrepo
func newUserRepo(db *sql.DB) *userRepo {
	return &userRepo{
		db: db,
	}
}

var queryBase string = `
		SELECT 
			  u.user_id
			, u.user_created_at
			, u.user_updated_at
			, u.uuid 
			, u.name 
    		, u.cpf	
    		, u.email
		FROM health_db.tab_user AS u
	`

func (r userRepo) FindByEmailAndPassword(email string, password string) (entity.User, resterrors.RestErr) {
	query := queryBase + `
			WHERE u.user_deleted_at IS NULL
				AND ( u.email = ? AND u.password = ? );
	`

	stmt, errorPrepare := r.db.Prepare(query)
	if errorPrepare != nil {
		logger.Error("FindByEmailAndPassword: Prepare error", errorPrepare)
		return entity.User{}, resterrors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	user, err := r.parseEntity(stmt.QueryRow(
		email,
		password,
	))
	if err != nil {
		return entity.User{}, err
	}

	return user, nil
}

func (r userRepo) parseEntity(row *sql.Row) (entity.User, resterrors.RestErr) {

	var user entity.User
	err := row.Scan(
		&user.ID,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.UUID,
		&user.Name,
		&user.CPF,
		&user.Email,
	)
	if err != nil {
		logger.Error("FindByEmailAndPassword: Parse Entity error", err)
		return entity.User{}, mysqlutils.HandleMySQLError(err)
	}

	return user, nil
}

func (r userRepo) GetUserByUUID(userUUID string) (user entity.User, err resterrors.RestErr) {
	query := queryBase + `
			WHERE u.user_deleted_at IS NULL
			  AND u.uuid 			= ?;
	`

	stmt, errorPrepare := r.db.Prepare(query)
	if errorPrepare != nil {
		logger.Error("GetUserByUUID: Prepare error", errorPrepare)
		return user, resterrors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	user, err = r.parseEntity(stmt.QueryRow(
		userUUID,
	))
	if err != nil {
		return user, err
	}

	return user, nil
}
