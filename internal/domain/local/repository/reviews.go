package repository

import (
	"context"

	"github.com/HackathonUC2025-Hackfest/Hackfest-BE/internal/domain/local"
)

func (r *localRepository) GetReviewsByLocalsID(ctx context.Context, localsID string, out *[]local.Review) error {
	query := `
	SELECT
		r.id, r.star, r.content, r.created_at, r.updated_at, u.photo_url
	FROM reviews r INNER JOIN users u ON u.id = r.user_id
	WHERE r.local_id = $1`

	rows, err := r.q.QueryxContext(ctx, query, localsID)
	if err != nil {
		return err
	}
	defer rows.Close()

	var result []local.Review
	for rows.Next() {
		var item local.Review
		if err := rows.StructScan(&item); err != nil {
			return err
		}
		result = append(result, item)
	}

	*out = result
	return nil
}
