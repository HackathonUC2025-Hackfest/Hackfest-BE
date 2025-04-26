package local

import (
	"time"

	"github.com/google/uuid"
)

type ResponseGetLocalBusinesses struct {
	ID          uuid.UUID         `json:"id"`
	Name        string            `json:"name"`
	Description string            `json:"description"`
	Address     string            `json:"address"`
	City        string            `json:"city"`
	Province    string            `json:"province"`
	Longitude   string            `json:"longitude"`
	Latitude    string            `json:"latitude"`
	Label       string            `json:"label"`
	OpenedTime  string            `json:"opened_time"`
	PhotoUrl    string            `json:"photo_url"`
	IsBusiness  bool              `json:"is_business"`
	CreatedAt   time.Time         `json:"created_at"`
	Reviews     []ResponseReviews `json:"reviews,omitempty"`
}

type ReseponseGetTourGuide struct {
	ID                          uuid.UUID         `json:"id"`
	Name                        string            `json:"name"`
	Description                 string            `json:"description"`
	Address                     string            `json:"address"`
	City                        string            `json:"city"`
	Province                    string            `json:"province"`
	Longitude                   float64           `json:"longitude"`
	Latitude                    float64           `json:"latitude"`
	PhotoUrl                    string            `json:"photo_url"`
	TourGuidePrice              int64             `json:"tour_guide_price"`
	TourGuideCount              int               `json:"tour_guide_count"`
	TourGuideDiscountPercentage float32           `json:"tour_guide_discount_percentage"`
	Price                       int64             `json:"price"`
	DiscountPercentage          float32           `json:"discount_percentage"`
	CreatedAt                   time.Time         `json:"created_at"`
	Reviews                     []ResponseReviews `json:"reviews,omitempty"`
}

type ResponseReviews struct {
	ID        uuid.UUID `json:"id"`
	Star      int       `json:"star"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
	PhotoURL  string    `json:"photo_url"`
}

type QueryParamRequestGetLocals struct {
	City string
	Type string
}

type RequestGenerateSnapLink struct {
	UserID   string ``
	TAID     string ``
	BookedAt string `json:"booked_at" validate:"required"`
}

type ResponseGenerateSnapLink struct {
	TAID       string `json:"ta_id"`
	PaymentUrl string `json:"payment_url"`
}
