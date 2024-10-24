package mapper

import (
	"go-app-arch/internal/config"
)

type ProductMapper struct {
	Cfg *config.Cfg
}

func NewProductMapper(cfg *config.Cfg) *ProductMapper {
	return &ProductMapper{Cfg: cfg}
}
