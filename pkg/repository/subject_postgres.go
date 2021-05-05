package repository

import (
	"backend_consumer/pkg/domain"
	"fmt"
	"github.com/jmoiron/sqlx"
	"strings"
)

type SubjectItemPostgres struct {
	db *sqlx.DB
}

func NewSubjectPostgres(db *sqlx.DB) *SubjectItemPostgres {
	return &SubjectItemPostgres{db: db}
}

func (r *SubjectItemPostgres) GetAllSubjects() ([]domain.Subject, error) {
	var subjects []domain.Subject

	//SELECT st.subject_name, st.full_name_subject, st.inn, st.KPP, ft.form_name FROM %s st, %s ft WHERE st.form_of_ownership_id=ft.id

	query := fmt.Sprintf(`SELECT * FROM %s`, subjectTable)

	if err := r.db.Select(&subjects, query); err != nil {
		return nil, err
	}
	return subjects, nil
}

func (r *SubjectItemPostgres) CreateSubject(subject domain.Subject) (int, error) {
	var id int
	query := fmt.Sprintf("INSERT INTO %s (subject_name, full_name_subject, inn, KPP, form_of_ownership_id) VALUES ($1, $2, $3, $4, $5) RETURNING id", subjectTable)
	row := r.db.QueryRow(query, subject.SubjectName, subject.FullNameSubject, subject.INN, subject.KPP, subject.FormOfOwnerShipId)

	if err := row.Scan(&id); err != nil {
		return 0, nil
	}
	return id, nil
}

func (r *SubjectItemPostgres) UpdateSubject(subjectId int, subject domain.SubjectInput) error {
	setValues := make([]string, 0)
	args := make([]interface{}, 0)
	argId := 1

	if subject.SubjectName != nil {
		setValues = append(setValues, fmt.Sprintf("subject_name=$%d", argId))
		args = append(args, *subject.SubjectName)
		argId++
	}
	if subject.FullNameSubject != nil {
		setValues = append(setValues, fmt.Sprintf("full_name_subject=$%d", argId))
		args = append(args, *subject.FullNameSubject)
		argId++
	}
	if subject.INN != nil {
		setValues = append(setValues, fmt.Sprintf("inn=$%d", argId))
		args = append(args, *subject.INN)
		argId++
	}
	if subject.KPP != nil {
		setValues = append(setValues, fmt.Sprintf("KPP=$%d", argId))
		args = append(args, *subject.KPP)
		argId++
	}
	if subject.FormOfOwnerShipId != nil {
		setValues = append(setValues, fmt.Sprintf("form_of_ownership_id=$%d", argId))
		args = append(args, *subject.FormOfOwnerShipId)
		argId++
	}

	setQuery := strings.Join(setValues, ", ")

	query := fmt.Sprintf(`UPDATE %s SET %s WHERE id=$%d`, subjectTable, setQuery, argId)

	args = append(args, subjectId)

	_, err := r.db.Exec(query, args...)
	return err

}

func (r *SubjectItemPostgres) DeleteSubject(subjectId int) error {
	query := fmt.Sprintf(`DELETE FROM %s WHERE id = $1`, subjectTable)
	_, err := r.db.Exec(query, subjectId)
	return err
}
