package storage

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/nurmuhammaddeveloper/API/models"
)

type DBManager struct {
	db *sql.DB
}

func NewDBManager(db *sql.DB) *DBManager {
	return &DBManager{
		db: db,
	}
}

func (strg *DBManager) Create(data *models.CreateStudentRequest) (*models.Student, error) {
	query := `INSERT INTO students(
		first_name,
		last_name,
        username,
		email,
		phone_number
	)VALUES ($1, $2, $3, $4, $5)
	RETURNING id, first_name, last_name, username, email, phone_number, created_at
	`
	insertedDataRow := strg.db.QueryRow(query,
		data.FirstName,
		data.LastName,
		data.UserName,
		data.Email,
		data.PhoneNumber,
	)
	var insertedDataResult models.Student
	err := insertedDataRow.Scan(
		&insertedDataResult.ID,
		&insertedDataResult.FirstName,
		&insertedDataResult.LastName,
		&insertedDataResult.UserName,
		&insertedDataResult.Email,
		&insertedDataResult.PhoneNumber,
		&insertedDataResult.CreatedAt,

	)
	if err != nil {
		return nil, err
	}
	return &insertedDataResult, nil
}
func (strg *DBManager) Update(data *models.UpdateStudentRequest) (*models.Student, error) {
	query := `UPDATE students SET
        first_name=$1,
        last_name=$2,
        username=$3,
        email=$4,
        phone_number=$5
		WHERE id=$6
		RETURNING id, first_name, last_name, username, email, phone_number, created_at
`
	updatedStudentRow := strg.db.QueryRow(query,
		data.FirstName,
		data.LastName,
		data.UserName,
		data.Email,
		data.PhoneNumber,
		data.ID,
	)
	var updatedStudentResult models.Student
	err := updatedStudentRow.Scan(
		&updatedStudentResult.ID,
		&updatedStudentResult.FirstName,
		&updatedStudentResult.LastName,
		&updatedStudentResult.UserName,
		&updatedStudentResult.Email,
		&updatedStudentResult.PhoneNumber,
		&updatedStudentResult.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &updatedStudentResult, err
}

func (strg *DBManager) Get(id int64) (*models.Student, error) {
	query := `SELECT 
				id,
				first_name,
				last_name,
				username,
				email,
                phone_number,
				created_at
			FROM students WHERE id=$1
	`
	Student := strg.db.QueryRow(query, id)
	var studentResult models.Student
	err := Student.Scan(
		&studentResult.ID,
		&studentResult.FirstName,
		&studentResult.LastName,
		&studentResult.UserName,
		&studentResult.Email,
		&studentResult.PhoneNumber,
		&studentResult.CreatedAt,
	)
	if err != nil {
		return nil, err
	}
	return &studentResult, err
}

func (strg *DBManager) Delete(id int64) error {
	query := `DELETE FROM students WHERE id=$1`
	result, err := strg.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsCount, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsCount < 1 {
		return errors.New("NO ROWS AFFECTED")
	}
	return nil
}

func (strg *DBManager) GetAll(params *models.GetStudentsQueryParam) (*models.GetStudentResult, error) {
	result := models.GetStudentResult{
		Students: make([]*models.Student, 0),
	}
	offset := (params.Page - 1) * params.Limit
	limit := fmt.Sprintf(" LIMIT %d OFFSET %d", params.Limit, offset)
	filter := "WHERE true"

	if params.FirstName != "" {
		filter += " AND first_name ilike '%" + params.FirstName + "%' "
	}
	if params.LastName != "" {
		filter += " AND last_name ilike '%" + params.LastName + "%' "
	}
	if params.LastName != "" {
		filter += " AND username ilike '%" + params.UserName + "%' "
	}
	query := `
		SELECT 
			id,
			first_name,
            last_name,
            username,
			email,
            phone_number,
            created_at
		FROM students
		` + filter + `
		ORDER BY created_at DESC
		` + limit
	rows, err := strg.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var student models.Student
		err := rows.Scan(
			&student.ID,
			&student.FirstName,
			&student.LastName,
			&student.UserName,
			&student.Email,
			&student.PhoneNumber,
			&student.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		result.Students = append(result.Students, &student)
	}
	queryCount := `SELECT count(1) FROM students ` + filter
	err = strg.db.QueryRow(queryCount).Scan(&result.Count)
	if err != nil {
		return nil, err
	}
	return &result, nil
}
