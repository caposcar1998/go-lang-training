package employee

import (
	"database/sql"
	"log"
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

func NewMock() (*sql.DB, sqlmock.Sqlmock, error) {
	db, mock, err := sqlmock.New()
	if err != nil {
		log.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}

	return db, mock, err
}

func TestEmployee_RetrieveId(t *testing.T) {
	db, mock, err := NewMock()
	defer db.Close()

	query := "SELECT * FROM employee WHERE id = //?"
	mock.ExpectBegin()
	mock.ExpectQuery(query).WithArgs(u.ID).WillReturnRows()
	mock.ExpectCommit()

	user, err := employee.EmployeeRetrieve(u.ID)

	assert.NotNil(t, user)
	assert.NoError(t, err)
}

func TestEmployee_Save(t *testing.T) {

}

func TestEmpoyee_Update(t *testing.T) {}

func TestEmployee_RetrieveAll(t *testing.T) {

}
