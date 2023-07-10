package guitar

import (
	"context"
	"database/sql"

	"github.com/Bobs-code/Guitar-API/models"
	"github.com/Bobs-code/Guitar-API/repository"
)

func NewSQLGuitarRepo(Conn *sql.DB) repository.GuitarRepo {
	return &postgresGuitRepo{
		Conn: Conn,
	}
}

type postgresGuitRepo struct {
	Conn *sql.DB
}

func (m *postgresGuitRepo) fetch(ctx context.Context, query string, args ...interface{}) ([]*models.Post, error) {
	rows, err := m.Conn.QueryContext(ctx, query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	payload := make([]*models.Guitar, 0)
	for rows.Next() {
		data := new(models.Guitar)

		err := rows.Scan(
			&data.ID,
			&data.Brand_id,
			&data.Description,
		)
	}
}
