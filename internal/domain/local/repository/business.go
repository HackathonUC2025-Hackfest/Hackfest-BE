package repository

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/HackathonUC2025-Hackfest/Hackfest-BE/internal/domain/local"
)

func (r *localRepository) GetAllLocalBusinesses(ctx context.Context, data local.QueryParamRequestGetLocals, out *[]local.Locals) error {
	query := `
	SELECT
		id, name, description, address, city, province, longitude, latitude, label, opened_time, photo_url, is_business, created_at, updated_at
	FROM locals
	WHERE 1=1`

	if data.City != "" {
		query += " AND LOWER(city) LIKE :city"
		data.City = "%" + strings.ToLower(data.City) + "%"
	}

	if data.Type == "business" {
		query += " AND is_business = true"
	} else {
		query += " AND is_business = false"
	}

	rows, err := r.q.NamedQueryContext(ctx, query, &data)
	if err != nil {
		return err
	}
	defer rows.Close()

	var result []local.Locals
	for rows.Next() {
		var item local.Locals
		if err := rows.StructScan(&item); err != nil {
			return err
		}
		result = append(result, item)
	}

	*out = result
	return nil
}

func (r *localRepository) GetLocalBusinessByID(ctx context.Context, data *local.Locals) error {
	query := `SELECT
	id, name, description, address, city, province, longitude, latitude, label, opened_time, photo_url, is_business, created_at, updated_at
	FROM locals
	WHERE id = $1
	`

	row := r.q.QueryRowxContext(ctx, query, data.ID)
	if err := row.StructScan(data); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return local.ErrLBNotFound
		} else {
			return err
		}
	}

	return nil
}
