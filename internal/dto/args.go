package dto

import (
	"go-app-arch/internal/validation"
)

type ProductFindListArgs struct {
	IDs      []int
	Category string
	PerPage  int
	Page     int
}

func (args *ProductFindListArgs) Validate(validator *validation.Validator) bool {
	validator.CheckField(args.PerPage > 0, "perPage", "perPage must be > 0")
	validator.CheckField(args.Page > 0, "page", "page must be > 0")
	return validator.HasErrors()
}

type ProductFindOneArgs struct {
	Slug string
}

func (args *ProductFindOneArgs) Validate(validator *validation.Validator) bool {
	validator.CheckField(len(args.Slug) > 0, "slug", "slug length must be > 0")
	return validator.HasErrors()
}

type GetConfigArgs struct {
	ProductCategoryFindArgs *ProductCategoryFindArgs
}

type ProductCategoryFindArgs struct {
	IsAvailable *bool
}
