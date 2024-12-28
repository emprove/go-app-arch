package repository

import (
	"go-app-arch/internal/dto"
	"go-app-arch/internal/entity"
)

type Product interface {
	FindOneAdm(args *dto.ProductFindOneAdmArgs) (*dto.ProductFindOneRowAdm, error)
	FindListAdm(args *dto.ProductFindListAdmArgs) (*dto.ProductFindListAdm, error)
	FindList(args *dto.ProductFindListArgs, locale string) (*dto.ProductFindList, error)
	FindOne(args *dto.ProductFindOneArgs, locale string) (*dto.ProductFindOneRow, error)
}

type Settings interface {
	GetAll() ([]entity.Settings, error)
	FindProductCategories(args *dto.ProductCategoryFindArgs, locale string) ([]entity.ProductCategory, error)
	FindProductOptions(locale string) ([]entity.ProductOption, error)
}

type User interface {
	FindOneByToken(token string) (*entity.User, error)
}
