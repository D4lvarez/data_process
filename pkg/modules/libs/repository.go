package libs

import "database/sql"

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) Repository {
	return Repository{
		db: db,
	}
}

func (r Repository) Execute(query string) (sql.Result, error) {
	return r.db.Exec(query)
}

func (r Repository) Insert(query string, values ...any) (sql.Result, error) {
	return r.db.Exec(query, values...)
}

func (r Repository) FetchAll(query string, values ...any) (*sql.Rows, error) {
	return r.db.Query(query, values...)
}

func (r Repository) FetchOne(query string, values ...any) *sql.Row {
	return r.db.QueryRow(query, values...)
}

func (r Repository) Update(query string, values ...any) (sql.Result, error) {
	return r.db.Exec(query, values...)
}

func (r Repository) Delete(query string, values ...any) (sql.Result, error) {
	return r.db.Exec(query, values...)
}
