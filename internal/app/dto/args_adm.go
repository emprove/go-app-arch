package dto

import "go-app-arch/internal/validation"

type ProductFindOneAdmArgs struct {
	ID int
}

func (args *ProductFindOneAdmArgs) Validate(validator *validation.Validator) bool {
	validator.CheckField(args.ID > 0, "id", "id must be > 0")
	return validator.HasErrors()
}

type ProductFindListAdmArgs struct {
	ID          []int
	Categories  []string
	Name        string
	Price       *int
	IsPublished *bool
	PerPage     int
	Page        int
}

func (args *ProductFindListAdmArgs) Validate(validator *validation.Validator) bool {
	if len(args.ID) > 0 {
		for _, v := range args.ID {
			validator.CheckField(v > 0, "id", "id must be > 0")
		}
	}
	if args.Name != "" {
		validator.CheckField(len(args.Name) < 255, "name", "name is too long")
	}
	if args.Price != nil {
		validator.CheckField(*args.Price > 0, "price", "price must be > 0")
	}
	validator.CheckField(args.PerPage > 0, "perPage", "perPage must be > 0")
	validator.CheckField(args.Page > 0, "page", "page must be > 0")
	return validator.HasErrors()
}
