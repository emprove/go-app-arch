package postgres

import "time"

type Currency struct {
	ID          int64 `sql:"primary_key"`
	Title       string
	IsoAlfa     string
	Symbol      *string
	Position    *int16
	IsAvailable bool
	CreatedAt   *time.Time
	UpdatedAt   *time.Time
}

type File struct {
	ID        int64 `sql:"primary_key"`
	Name      *string
	Path      string
	PathThumb *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type ProductHasFile struct {
	ProductID int64
	FileID    int64
	Position  *int16
}

type ProductHasRelated struct {
	ProductID int64
	RelatedID int64
}

type ProductSupply struct {
	ID        int64 `sql:"primary_key"`
	ProductID int64
	Options   *string
	Quantity  int16
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type Product struct {
	ID              int64 `sql:"primary_key"`
	Category        *string
	Article         *string
	Status          int16
	Name            string
	Annotation      *string
	Description     *string
	Price           float64
	Position        *int16
	Options         *string
	Meta            *string
	CreatedAt       *time.Time
	UpdatedAt       *time.Time
	VideoPath       *string
	Slug            string
	MakingInDaysMin *int16
	MakingInDaysMax *int16
	NameEn          *string
	AnnotationEn    *string
	DescriptionEn   *string
	SlugEn          string
	IsAvailable     bool
	IsPublished     bool
}

type Settings struct {
	ID        int64 `sql:"primary_key"`
	Key       string
	Meta      *string
	CreatedAt *time.Time
	UpdatedAt *time.Time
}

type User struct {
	ID            int64 `sql:"primary_key"`
	Name          string
	Email         string
	Password      string
	RememberToken *string
	AccessToken   *string
	Permissions   *string
	CreatedAt     *time.Time
	UpdatedAt     *time.Time
}
