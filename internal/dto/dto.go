package dto

import "go-app-arch/internal/entity"

type GetConfigRes struct {
	Settings          []entity.Settings        `json:"settings"`
	ProductCategories []entity.ProductCategory `json:"product_categories"`
	ProductOptions    []entity.ProductOption   `json:"product_options"`
}

type ProductFindListProduct struct {
	ID          int           `json:"id"`
	CurrencyIso string        `json:"currency_iso"`
	Category    string        `json:"category"`
	Name        string        `json:"name"`
	Annotation  *string       `json:"annotation"`
	Description *string       `json:"description"`
	Price       float64       `json:"price"`
	Position    *int          `json:"position"`
	Slug        string        `json:"slug"`
	Files       []entity.File `json:"files"`
}

type ProductFindList struct {
	Products   []ProductFindListProduct `json:"products"`
	TotalCount int                      `json:"total_count"`
}

type ProductFindOne struct {
	ID              int                      `json:"id"`
	CurrencyIso     string                   `json:"currency_iso"`
	Category        string                   `json:"category"`
	Name            string                   `json:"name"`
	Annotation      *string                  `json:"annotation"`
	Description     *string                  `json:"description"`
	Price           float64                  `json:"price"`
	Options         []string                 `json:"options"`
	VideoPath       string                   `json:"video_path"`
	Slug            string                   `json:"slug"`
	MakingInDaysMin *int                     `json:"making_in_days_min"`
	MakingInDaysMax *int                     `json:"making_in_days_max"`
	IsAvailable     bool                     `json:"is_available"`
	Files           []entity.File            `json:"files"`
	Supplies        []entity.Supply          `json:"supplies"`
	RelatedProducts []ProductFindListProduct `json:"related_products"`
}

type ProductFindOneRow struct {
	ID              int
	CurrencyIso     string
	Category        string
	Name            string
	Annotation      *string
	Description     *string
	Price           float64
	Options         []string
	VideoPath       string
	Slug            string
	MakingInDaysMin *int
	MakingInDaysMax *int
	IsAvailable     bool
	Files           []entity.File
	Supplies        []entity.Supply
	RelatedIDs      []int
}
