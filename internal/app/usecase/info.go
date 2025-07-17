package usecase

import (
	"go-app-arch/internal/app/config"
	"go-app-arch/internal/app/dto"
	"go-app-arch/internal/domain/repository"
)

type Info struct {
	cfg          *config.Cfg
	DS           *config.DynamicState
	repoSettings repository.Settings
}

type InfoInterface interface {
	GetLocales() []config.Locale
	GetConfig(args *dto.GetConfigArgs, locale string) (*dto.GetConfigRes, error)
}

func NewInfo(cfg *config.Cfg, ds *config.DynamicState, rs repository.Settings) *Info {
	return &Info{cfg: cfg, DS: ds, repoSettings: rs}
}

func (u *Info) GetLocales() []config.Locale {
	return u.cfg.GetLocales()
}

func (u *Info) GetConfig(args *dto.GetConfigArgs, locale string) (*dto.GetConfigRes, error) {
	settings, err := u.repoSettings.GetAll()
	if err != nil {
		return nil, err
	}

	productCats, err := u.repoSettings.FindProductCategories(args.ProductCategoryFindArgs, u.DS.GetLocale())
	if err != nil {
		return nil, err
	}
	productOpts, err := u.repoSettings.FindProductOptions(u.DS.GetLocale())
	if err != nil {
		return nil, err
	}

	res := &dto.GetConfigRes{
		Settings:          settings,
		ProductCategories: productCats,
		ProductOptions:    productOpts,
	}

	return res, nil
}
