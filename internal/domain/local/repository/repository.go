package repository

import (
	"context"
	"database/sql"
	"errors"

	"github.com/HackathonUC2025-Hackfest/Hackfest-BE/internal/domain/local"
	"github.com/jmoiron/sqlx"
)

var (
	errFailedToCommit   = errors.New("FAILED_TO_COMMIT_TRANSACTION")
	errFailedToRollback = errors.New("FAILED_TO_ROLLBACK_TRANSACTION")
)

type repository struct {
	DB *sqlx.DB
}

type RepositoryItf interface {
	NewClient(tx bool) (localRepositoryItf, error)
}

type localRepository struct {
	q namedExt
}

type localRepositoryItf interface {
	Commit() error
	Rollback() error
	GetAllLocalBusinesses(ctx context.Context, data local.QueryParamRequestGetLocals, out *[]local.Locals) error
	GetLocalBusinessByID(ctx context.Context, data *local.Locals) error
	GetReviewsByLocalsID(ctx context.Context, localsID string, out *[]local.Review) error
	GetAllTourGuides(ctx context.Context, city string, out *[]local.TouristAttractions) error
	GetTAByID(ctx context.Context, data *local.TouristAttractions) error
	GetBookingsByTAID(ctx context.Context, TAID string, out *[]local.TourGuideBookings) error
	CreateBooking(ctx context.Context, data local.TourGuideBookings) error
	GetFullBook(ctx context.Context, attractionID string, dates *[]string) error
}

type namedExt interface {
	sqlx.ExtContext
	NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error)
	NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error)
}

type txWrapper struct {
	*sqlx.Tx
}

func (t *txWrapper) NamedExecContext(ctx context.Context, query string, arg interface{}) (sql.Result, error) {
	return sqlx.NamedExecContext(ctx, t.Tx, query, arg)
}

func (t *txWrapper) NamedQueryContext(ctx context.Context, query string, arg interface{}) (*sqlx.Rows, error) {
	return sqlx.NamedQueryContext(ctx, t.Tx, query, arg)
}

func New(db *sqlx.DB) RepositoryItf {
	return &repository{db}
}

func (r *repository) NewClient(tx bool) (localRepositoryItf, error) {
	var db namedExt

	db = r.DB
	if tx {
		rawTx, err := r.DB.Beginx()
		if err != nil {
			return nil, err
		}
		db = &txWrapper{rawTx}
	}

	return &localRepository{db}, nil
}

func (r *localRepository) Commit() error {
	switch q := r.q.(type) {
	case *txWrapper:
		return q.Tx.Commit()
	case *sqlx.DB:
		return nil
	default:
		return errFailedToCommit
	}
}

func (r *localRepository) Rollback() error {
	switch q := r.q.(type) {
	case *txWrapper:
		return q.Tx.Rollback()
	case *sqlx.DB:
		return nil
	default:
		return errFailedToRollback
	}
}
