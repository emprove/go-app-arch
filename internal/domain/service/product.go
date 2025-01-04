package service

import (
	"go-app-arch/internal/app"
	"go-app-arch/internal/app/config"
	"go-app-arch/internal/app/dto"
	"go-app-arch/internal/domain/repository"
	"go-app-arch/internal/validation"
)

type ProductServiceInterface interface {
	FindList(args *dto.ProductFindListArgs, currencyIso string) (*dto.ProductFindList, error)
	FindOne(args *dto.ProductFindOneArgs, currencyIso string) (*dto.ProductFindOne, error)
	FindListAdm(args *dto.ProductFindListAdmArgs) (*dto.ProductFindListAdm, error)
	FindOneAdm(args *dto.ProductFindOneAdmArgs) (*dto.ProductFindOneAdm, error)
}

type productService struct {
	DS          *config.DynamicState
	RepoProduct repository.Product
}

func NewProductService(ds *config.DynamicState, p repository.Product) ProductServiceInterface {
	return &productService{DS: ds, RepoProduct: p}
}

func (s *productService) FindList(args *dto.ProductFindListArgs, currencyIso string) (*dto.ProductFindList, error) {
	validator := validation.Validator{}
	invalid := args.Validate(&validator)
	if invalid {
		return nil, &app.ValidationError{Validator: validator}
	}

	res, err := s.RepoProduct.FindList(args, s.DS.Locale)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *productService) FindOne(args *dto.ProductFindOneArgs, currencyIso string) (*dto.ProductFindOne, error) {
	validator := validation.Validator{}
	invalid := args.Validate(&validator)
	if invalid {
		return nil, &app.ValidationError{Validator: validator}
	}

	row, err := s.RepoProduct.FindOne(args, s.DS.Locale)
	if err != nil {
		return nil, err
	}

	var res dto.ProductFindOne
	res.ID = row.ID
	res.CurrencyIso = currencyIso
	res.Category = row.Category
	res.Name = row.Name
	res.Annotation = row.Annotation
	res.Description = row.Description
	res.Options = row.Options
	res.VideoPath = row.VideoPath
	res.Slug = row.Slug
	res.MakingInDaysMin = row.MakingInDaysMin
	res.MakingInDaysMax = row.MakingInDaysMax
	res.IsAvailable = row.IsAvailable
	res.Files = row.Files
	res.Supplies = row.Supplies
	res.Price = row.Price

	// related products
	if len(row.RelatedIDs) > 0 {
		argsForList := &dto.ProductFindListArgs{
			IDs:     row.RelatedIDs,
			PerPage: 12,
			Page:    1,
		}
		list, err := s.FindList(argsForList, currencyIso)
		if err != nil {
			return nil, err
		}
		res.RelatedProducts = list.Products
	} else {
		res.RelatedProducts = make([]dto.ProductFindListProduct, 0)
	}

	return &res, nil
}

func (s *productService) FindListAdm(args *dto.ProductFindListAdmArgs) (*dto.ProductFindListAdm, error) {
	validator := validation.Validator{}
	invalid := args.Validate(&validator)
	if invalid {
		return nil, &app.ValidationError{Validator: validator}
	}

	res, err := s.RepoProduct.FindListAdm(args)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (s *productService) FindOneAdm(args *dto.ProductFindOneAdmArgs) (*dto.ProductFindOneAdm, error) {
	validator := validation.Validator{}
	invalid := args.Validate(&validator)
	if invalid {
		return nil, &app.ValidationError{Validator: validator}
	}

	row, err := s.RepoProduct.FindOneAdm(args)
	if err != nil {
		return nil, err
	}

	var res dto.ProductFindOneAdm
	res.ID = row.ID
	res.CurrencyIso = row.CurrencyIso
	res.Category = row.Category
	res.Name = row.Name
	res.Annotation = row.Annotation
	res.Description = row.Description
	res.Price = row.Price
	res.Position = row.Position
	res.Options = row.Options
	res.CreatedAt = row.CreatedAt
	res.UpdatedAt = row.UpdatedAt
	res.VideoPath = row.VideoPath
	res.Slug = row.Slug
	res.MakingInDaysMin = row.MakingInDaysMin
	res.MakingInDaysMax = row.MakingInDaysMax
	res.NameEn = row.NameEn
	res.AnnotationEn = row.AnnotationEn
	res.DescriptionEn = row.DescriptionEn
	res.SlugEn = row.SlugEn
	res.IsAvailable = row.IsAvailable
	res.IsPublished = row.IsPublished
	res.Files = row.Files
	res.Supplies = row.Supplies

	if len(row.RelatedIDs) > 0 {
		argsForList := &dto.ProductFindListAdmArgs{
			ID:      row.RelatedIDs,
			PerPage: 12,
			Page:    1,
		}
		list, err := s.FindListAdm(argsForList)
		if err != nil {
			return nil, err
		}
		res.RelatedProducts = list.Products
	} else {
		res.RelatedProducts = make([]dto.ProductFindListProductAdm, 0)
	}

	return &res, nil
}
