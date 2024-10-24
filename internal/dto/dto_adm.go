package dto

import (
	"go-app-arch/internal/entity"
	"time"
)

type ProductFindListProductAdm struct {
	ID          int           `json:"id"`
	CurrencyIso string        `json:"currency_iso"`
	Category    string        `json:"category"`
	Name        string        `json:"name"`
	Price       float64       `json:"price"`
	Position    *int          `json:"position"`
	IsPublished bool          `json:"is_published"`
	Files       []entity.File `json:"files"`
}

type ProductFindListAdm struct {
	Products   []ProductFindListProductAdm `json:"products"`
	TotalCount int                         `json:"total_count"`
}

type ProductFindOneAdm struct {
	ID              int                         `json:"id"`
	CurrencyIso     string                      `json:"currency_iso"`
	Category        string                      `json:"category"`
	Name            string                      `json:"name"`
	Annotation      *string                     `json:"annotation"`
	Description     *string                     `json:"description"`
	Price           float64                     `json:"price"`
	Position        *int                        `json:"position"`
	Options         []string                    `json:"options"`
	CreatedAt       *time.Time                  `json:"created_at"`
	UpdatedAt       *time.Time                  `json:"updated_at"`
	VideoPath       *string                     `json:"video_path"`
	Slug            string                      `json:"slug"`
	MakingInDaysMin *int                        `json:"making_in_days_min"`
	MakingInDaysMax *int                        `json:"making_in_days_max"`
	NameEn          *string                     `json:"name_en"`
	AnnotationEn    *string                     `json:"annotation_en"`
	DescriptionEn   *string                     `json:"description_en"`
	SlugEn          string                      `json:"slug_en"`
	IsAvailable     bool                        `json:"is_available"`
	IsPublished     bool                        `json:"is_published"`
	Files           []entity.File               `json:"files"`
	Supplies        []entity.Supply             `json:"supplies"`
	RelatedProducts []ProductFindListProductAdm `json:"related_products"`
}

type ProductFindOneRowAdm struct {
	ID              int
	CurrencyIso     string
	Category        string
	Name            string
	Annotation      *string
	Description     *string
	Price           float64
	Position        *int
	Options         []string
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
	VideoPath       *string
	Slug            string
	MakingInDaysMin *int
	MakingInDaysMax *int
	NameEn          *string
	AnnotationEn    *string
	DescriptionEn   *string
	SlugEn          string
	IsAvailable     bool
	IsPublished     bool
	Files           []entity.File
	Supplies        []entity.Supply
	RelatedIDs      []int
}
