package repository

import (
	"context"
	"database/sql"
	"errors"
	"strings"

	"github.com/HackathonUC2025-Hackfest/Hackfest-BE/internal/domain/local"
)

func (r *localRepository) GetAllTourGuides(ctx context.Context, city string, out *[]local.TouristAttractions) error {
	query := `
	SELECT
		id, name, description, address, city, province, longitude, latitude, photo_url, tour_guide_price, tour_guide_count, tour_guide_discount_percentage, price, discount_percentage, created_at, updated_at
	FROM tourist_attractions
	WHERE 1=1`

	if city != "" {
		query += " AND LOWER(city) LIKE $1"
		city = "%" + strings.ToLower(city) + "%"
	}

	rows, err := r.q.QueryxContext(ctx, query, city)
	if err != nil {
		return err
	}
	defer rows.Close()

	var result []local.TouristAttractions
	for rows.Next() {
		var item local.TouristAttractions
		if err := rows.StructScan(&item); err != nil {
			return err
		}

		result = append(result, item)
	}

	*out = result
	return nil
}

func (r *localRepository) GetTAByID(ctx context.Context, data *local.TouristAttractions) error {
	query := `SELECT
	id, name, description, address, city, province, longitude, latitude, photo_url, tour_guide_price, tour_guide_count, tour_guide_discount_percentage, price, discount_percentage, created_at, updated_at
	FROM tourist_attractions
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

func (r *localRepository) GetBookingsByTAID(ctx context.Context, TAID string, out *[]local.TourGuideBookings) error {
	query := `
	SELECT
		tb.id, tb.payment_url, tb.star, tb.content, tb.created_at, tb.updated_at, tb.status, tb.user_id, tb.tourist_attraction_id
	FROM tourguide_bookings tb INNER JOIN users u ON u.id = tb.user_id
	WHERE tb.tourist_attraction_id = $1`

	rows, err := r.q.QueryxContext(ctx, query, TAID)
	if err != nil {
		return err
	}
	defer rows.Close()

	var result []local.TourGuideBookings
	for rows.Next() {
		var item local.TourGuideBookings
		if err := rows.StructScan(&item); err != nil {
			return err
		}
		result = append(result, item)
	}

	*out = result
	return nil
}

func (r *localRepository) GetFullBook(ctx context.Context, attractionID string, dates *[]string) error {
	query := `
	SELECT
	  DATE(tb.booked_at) AS date
	FROM
	  tourguide_bookings tb
	JOIN
	  tourist_attractions ta ON ta.id = tb.tourist_attraction_id
	WHERE
	  EXTRACT(MONTH FROM tb.booked_at) = 4
	  AND EXTRACT(YEAR FROM tb.booked_at) = 2025
	  AND tb.tourist_attraction_id = $1
	GROUP BY
	  DATE(tb.booked_at),
	  ta.tour_guide_count
	HAVING
	  COUNT(*) >= ta.tour_guide_count
	ORDER BY
	  DATE(tb.booked_at);
	`

	rows, err := r.q.QueryxContext(ctx, query, attractionID)
	if err != nil {
		return err
	}
	defer rows.Close()

	var results []string
	for rows.Next() {
		var date string
		if err := rows.Scan(&date); err != nil {
			return err
		}
		results = append(results, date)
	}

	*dates = results
	return nil
}

func (r *localRepository) CreateBooking(ctx context.Context, data local.TourGuideBookings) error {
	query := `INSERT INTO tourguide_bookings (
		id, payment_url, star, content, booked_at, status, user_id, tourist_attraction_id
	) VALUES (
		:id, :payment_url, :star, :content, :booked_at, :status, :user_id, :tourist_attraction_id
	)`

	_, err := r.q.NamedExecContext(ctx, query, data)
	if err != nil {
		return err
	}

	return nil
}
