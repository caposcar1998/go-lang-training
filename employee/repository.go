package employee

import "context"

type Repository interface {
	Employees(ctx context.Context, pos Position) ([]Employee, error)
	EmployeeRetrieve(ctx context.Context, id int) (*Employee, error)
	Save(ctx context.Context, e *Employee) error
	RetrieveAllValues(ctx context.Context, table string) ([]Employee, error)
	UpdateEmployee(ctx context.Context) error
}
