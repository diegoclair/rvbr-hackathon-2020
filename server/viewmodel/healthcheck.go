package viewmodel

import (
	"time"

	"github.com/RedVentures/rvbr-2020-hackathon-time-3-backend/domain/entity"
)

type HealthChecksReponse struct {
	HealthChecks []HealthCheckReponse `json:"health_checks"`
}

type HealthCheckReponse struct {
	UUID string    `json:"health_check_uuid"`
	Date time.Time `json:"health_check_date"`
	Note string    `json:"health_check_note"`

	DoctorName       string `json:"health_check_doctor_name"`
	DoctorSpeciality string `json:"health_check_doctor_speciality"`

	ExamName string `json:"health_check_exam_name"`
	ExamType string `json:"health_check_exam_type"`

	InstitutionName  string `json:"health_check_institution_name"`
	InstitutionCNPJ  string `json:"health_check_institution_cnpj"`
	InstitutionPhone string `json:"health_check_institution_phone"`
	InstitutionType  string `json:"health_check_institution_type"`
}

func ParseHealthChecksResponse(healths []entity.HealthCheck) HealthChecksReponse {

	body := make([]HealthCheckReponse, len(healths))
	for i, health := range healths {
		h := ParseHealthCheckResponse(health)
		body[i] = h
	}
	response := HealthChecksReponse{body}

	return response
}

func ParseHealthCheckResponse(health entity.HealthCheck) HealthCheckReponse {
	return HealthCheckReponse{
		UUID: health.UUID,
		Date: health.Date,
		Note: health.Note,

		DoctorName:       health.Doctor.Name,
		DoctorSpeciality: health.Doctor.Speciality,

		ExamName: health.Exam.Name,
		ExamType: health.Exam.Type,

		InstitutionName:  health.Institution.Name,
		InstitutionCNPJ:  health.Institution.CPNJ,
		InstitutionPhone: health.Institution.Phone,
		InstitutionType:  health.Institution.Type,
	}
}
