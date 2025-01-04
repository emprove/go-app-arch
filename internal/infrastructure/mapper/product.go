package mapper

import (
	"go-app-arch/internal/app/config"
)

type ProductMapper struct {
	Cfg *config.Cfg
}

func NewProductMapper(cfg *config.Cfg) *ProductMapper {
	return &ProductMapper{Cfg: cfg}
}
