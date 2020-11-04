package mysql

import (
	"database/sql"

	"github.com/RedVentures/rvbr-2020-hackathon-time-3-backend/domain/entity"
	"github.com/diegoclair/go_utils-lib/logger"
	"github.com/diegoclair/go_utils-lib/resterrors"
)

type healthCheckRepo struct {
	db *sql.DB
}

func newHealthCheckRepo(db *sql.DB) *healthCheckRepo {
	return &healthCheckRepo{
		db: db,
	}
}

func (r healthCheckRepo) FindAllByUserUUID(userUUID string) ([]entity.HealthCheck, resterrors.RestErr) {
	query := `
		SELECT 
			h.health_check_id,
			h.health_check_created_at,
			h.health_check_updated_at,
			h.uuid,
			h.health_check_date,
			h.note,
			
			d.uuid,
			d.name,
			d.speciality,
					
			e.uuid,
			e.name,
			e.type,
			
			i.uuid,
			i.name,
			i.cnpj,
			i.phone,
			i.type

		FROM health_db.tab_health_check AS h
			
			JOIN health_db.tab_doctor AS d
				ON d.doctor_deleted_at IS NULL
				AND d.doctor_id = h.doctor_id

			JOIN health_db.tab_exam AS e
				ON e.exam_deleted_at IS NULL
				AND e.exam_id = h.exam_id
						
			JOIN health_db.tab_institution AS i
				ON i.institution_deleted_at IS NULL
				AND i.institution_id = h.institution_id

			JOIN health_db.tab_user AS u
				ON u.user_deleted_at IS NULL
				AND u.user_id = h.user_id
						
			WHERE u.uuid = ?
		;
	`

	stmt, errorPrepare := r.db.Prepare(query)
	if errorPrepare != nil {
		logger.Error("FindAllByUserUUID: Prepare error", errorPrepare)
		return nil, resterrors.NewInternalServerError("Database error")
	}
	defer stmt.Close()

	rows, err := stmt.Query(
		userUUID,
	)
	if err != nil {
		logger.Error("FindAllByUserUUID: Prepare error", errorPrepare)
		return nil, resterrors.NewInternalServerError("Database error")
	}

	healths, err := r.parseEntities(rows)
	if err != nil {
		logger.Error("FindAllByUserUUID: Parse Entities error", errorPrepare)
		return nil, resterrors.NewInternalServerError("Database error")
	}

	return healths, nil
}

func (r healthCheckRepo) parseEntities(rows *sql.Rows) ([]entity.HealthCheck, resterrors.RestErr) {
	healths := make([]entity.HealthCheck, 0)
	for rows.Next() {

		health, err := r.parseEntity(rows)
		if err != nil {
			return nil, err
		}

		healths = append(healths, health)
	}

	return healths, nil
}

func (r healthCheckRepo) parseEntity(row *sql.Rows) (entity.HealthCheck, resterrors.RestErr) {

	var health entity.HealthCheck
	err := row.Scan(
		&health.ID,
		&health.CreatedAt,
		&health.UpdatedAt,
		&health.UUID,
		&health.Date,
		&health.Note,

		&health.Doctor.UUID,
		&health.Doctor.Name,
		&health.Doctor.Speciality,

		&health.Exam.UUID,
		&health.Exam.Name,
		&health.Exam.Type,

		&health.Institution.UUID,
		&health.Institution.Name,
		&health.Institution.CPNJ,
		&health.Institution.Phone,
		&health.Institution.Type,
	)
	if err != nil {
		logger.Error("FindAllByUserUUID: Parse Entity error", err)
		return entity.HealthCheck{}, resterrors.NewInternalServerError("Database error")
	}

	return health, nil
}
