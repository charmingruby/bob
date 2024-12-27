package postgres

import (
	"database/sql"

	"{{ .SourcePath }}/shared/custom_err/database_err"
	"{{ .SourcePath }}/shared/custom_err/database_err/sql_err"
	"{{ .SourcePath }}/{{ .Module }}/core/model"
	"github.com/jmoiron/sqlx"
)

const (
	create{{ .UpperCaseModel }}   = "create {{ .LowerCaseModel }}"
	find{{ .UpperCaseModel }}ByID = "find {{ .LowerCaseModel }} by id"
	delete{{ .UpperCaseModel }}   = "delete {{ .LowerCaseModel }}"
)

func {{ .LowerCaseModel }}Queries() map[string]string {
	return map[string]string{
		create{{ .UpperCaseModel }}: `INSERT INTO {{ .LowerCaseModel }}s
		(id, name)
		VALUES ($1, $2)
		RETURNING *`,
		find{{ .UpperCaseModel }}ByID: `SELECT * FROM {{ .LowerCaseModel }}s 
		WHERE id = $1`,
		delete{{ .UpperCaseModel }}: `UPDATE {{ .LowerCaseModel }}s
		SET deleted_at = $1
		WHERE id = $2 AND deleted_at IS NULL`,
	}
}

func New{{ .UpperCaseModel }}Repository(db *sqlx.DB) (*{{ .UpperCaseModel }}Repository, error) {
	stmts := make(map[string]*sqlx.Stmt)

	for queryName, statement := range {{ .LowerCaseModel }}Queries() {
		stmt, err := db.Preparex(statement)
		if err != nil {
			return nil,
				sql_err.NewPreparationErr(queryName, "{{ .LowerCaseModel }}", err)
		}

		stmts[queryName] = stmt
	}

	return &{{ .UpperCaseModel }}Repository{
		db:    db,
		stmts: stmts,
	}, nil
}

type {{ .UpperCaseModel }}Repository struct {
	db    *sqlx.DB
	stmts map[string]*sqlx.Stmt
}

func (r *{{ .UpperCaseModel }}Repository) statement(queryName string) (*sqlx.Stmt, error) {
	stmt, ok := r.stmts[queryName]

	if !ok {
		return nil,
			sql_err.NewStatementNotPreparedErr(queryName, "{{ .LowerCaseModel }}")
	}

	return stmt, nil
}

func (r *{{ .UpperCaseModel }}Repository) Store(model *model.{{ .UpperCaseModel }}) error {
	stmt, err := r.statement(create{{ .UpperCaseModel }})
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(
		model.ID,
		model.Name,
	); err != nil {
		return database_err.NewPersistenceErr(err, "{{ .LowerCaseModel }} store", "postgres")
	}

	return nil
}

func (r *{{ .UpperCaseModel }}Repository) FindByID(id string) (*model.{{ .UpperCaseModel }}, error) {
	stmt, err := r.statement(find{{ .UpperCaseModel }}ByID)
	if err != nil {
		return nil, err
	}

	var {{ .LowerCaseModel }} model.{{ .UpperCaseModel }}
	if err := stmt.Get(&{{ .LowerCaseModel }}, id); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}

		return nil, database_err.NewPersistenceErr(err, "{{ .LowerCaseModel }} find_by_id", "postgres")
	}

	return &{{ .LowerCaseModel }}, nil
}

func (r *{{ .UpperCaseModel }}Repository) Delete(model *model.{{ .UpperCaseModel }}) error {
	stmt, err := r.statement(delete{{ .UpperCaseModel }})
	if err != nil {
		return err
	}

	if _, err := stmt.Exec(model.DeletedAt, model.ID); err != nil {
		return database_err.NewPersistenceErr(err, "{{ .LowerCaseModel }} delete", "postgres")
	}

	return nil
}