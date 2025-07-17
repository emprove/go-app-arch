package repository

import (
	"context"
	"go-app-arch/internal/app/dto"
	"go-app-arch/internal/domain/entity"
)

type Product interface {
	FindOneAdm(ctx context.Context, args *dto.ProductFindOneAdmArgs) (*dto.ProductFindOneRowAdm, error)
	FindListAdm(ctx context.Context, args *dto.ProductFindListAdmArgs) (*dto.ProductFindListAdm, error)
	FindList(ctx context.Context, args *dto.ProductFindListArgs, locale string) (*dto.ProductFindList, error)
	FindOne(ctx context.Context, args *dto.ProductFindOneArgs, locale string) (*dto.ProductFindOneRow, error)
}

type Settings interface {
	GetAll() ([]entity.Settings, error)
	FindProductCategories(args *dto.ProductCategoryFindArgs, locale string) ([]entity.ProductCategory, error)
	FindProductOptions(locale string) ([]entity.ProductOption, error)
}

type User interface {
	FindOneByToken(token string) (*entity.User, error)
}
