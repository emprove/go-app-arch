package rest

import (
	"go-app-arch/internal/service"
	"go-app-arch/internal/usecase"
)

type ProductHandlerAdm struct {
	sProduct *service.ProductService
}

type ProductHandler struct {
	sProduct *service.ProductService
}

type InfoHandler struct {
	uInfo *usecase.Info
}

func NewProductHandlerAdm(sp *service.ProductService) *ProductHandlerAdm {
	return &ProductHandlerAdm{
		sProduct: sp,
	}
}

func NewProductHandler(sp *service.ProductService) *ProductHandler {
	return &ProductHandler{
		sProduct: sp,
	}
}

func NewInfoHandler(u *usecase.Info) *InfoHandler {
	return &InfoHandler{
		uInfo: u,
	}
}
