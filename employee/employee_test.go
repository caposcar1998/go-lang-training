package employee

import (
	"database/sql"
	"log"
	"regexp"
	"testing"
	"time"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/stretchr/testify/assert"
)

var u = Employee{
	ID:          1,
	FullName:    "Oscar Contreras",
	Position:    1,
	Salary:      500.00,
	Joined:      time.Now(),
	OnProbation: false,
	CreatedAt:   time.Now(),
}

var u2 = Employee{
	ID:          2,
	FullName:    "Hector Contreras",
	Position:    2,
	Salary:      600.00,
	Joined:      time.Now(),
	OnProbation: true,
	CreatedAt:   time.Now(),
}

var u3 = Employee{
	ID:          2,
	FullName:    "Hector Contreras",
	Position:    2,
	Salary:      600.00,
	Joined:      time.Now(),
	OnProbation: true,
}

func NewMock() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock, err
}

func TestEmployee_RetrieveId(t *testing.T) {
	db, mock, _ := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := regexp.QuoteMeta("SELECT * FROM employee WHERE id = ?")

	rows := sqlmock.NewRows([]string{"id", "full_name", "position", "salary", "joined", "on_probation", "created_at"}).
		AddRow(u.ID, u.FullName, u.Position, u.Salary, u.Joined, u.OnProbation, u.CreatedAt)

	mock.ExpectQuery(query).WithArgs(u.ID).WillReturnRows(rows)

	user, err := repo.EmployeeRetrieve(u.ID)
	assert.NotNil(t, user)
	assert.NoError(t, err)
}

func TestEmployee_Save(t *testing.T) {

	db, mock, _ := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()
	query := regexp.QuoteMeta("INSERT INTO employee (full_name, position,salary,joined,on_probation) VALUES (? ,? , ? , ? , ?)")
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(u.FullName, u.Position, u.Salary, u.Joined, u.OnProbation).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.Save(&u3)
	assert.NoError(t, err)
}

func TestEmpoyee_Update(t *testing.T) {
	db, mock, _ := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()
	query := regexp.QuoteMeta("UPDATE employee SET full_name = ?, position = ?, salary = ?, joined = ?, on_probation = ? WHERE id = ? ")
	prep := mock.ExpectPrepare(query)
	prep.ExpectExec().WithArgs(u.FullName, u.Position, u.Salary, u.Joined, u.OnProbation).WillReturnResult(sqlmock.NewResult(0, 1))

	err := repo.UpdateEmployee(u3, "1")
	assert.NoError(t, err)
}

func TestEmployee_RetrieveAll(t *testing.T) {
	db, mock, _ := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := regexp.QuoteMeta("SELECT * FROM employee")

	rows := sqlmock.NewRows([]string{"id", "full_name", "position", "salary", "joined", "on_probation", "created_at"}).
		AddRow(u.ID, u.FullName, u.Position, u.Salary, u.Joined, u.OnProbation, u.CreatedAt).
		AddRow(u2.ID, u2.FullName, u2.Position, u2.Salary, u2.Joined, u2.OnProbation, u2.CreatedAt)

	mock.ExpectQuery(query).WithArgs().WillReturnRows(rows)

	user, err := repo.RetrieveAllValues("employee")
	assert.NotNil(t, user)
	assert.NoError(t, err)
}

func TestEmployee_RetrievePosition(t *testing.T) {
	db, mock, _ := NewMock()
	repo := &repository{db}
	defer func() {
		repo.Close()
	}()

	query := regexp.QuoteMeta("SELECT * FROM employee WHERE position = ?")

	rows := sqlmock.NewRows([]string{"id", "full_name", "position", "salary", "joined", "on_probation", "created_at"}).
		AddRow(u.ID, u.FullName, u.Position, u.Salary, u.Joined, u.OnProbation, u.CreatedAt).
		AddRow(u2.ID, u2.FullName, u2.Position, u2.Salary, u2.Joined, u2.OnProbation, u2.CreatedAt)

	mock.ExpectQuery(query).WithArgs("1").WillReturnRows(rows)

	user, err := repo.Employees("1")
	assert.NotNil(t, user)
	assert.NoError(t, err)
}
