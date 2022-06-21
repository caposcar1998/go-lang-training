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

}

func TestEmpoyee_Update(t *testing.T) {}

func TestEmployee_RetrieveAll(t *testing.T) {

}
