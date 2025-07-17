package rest

import (
	"go-app-arch/internal/app/usecase"
	"go-app-arch/internal/domain/service"
)

type ProductHandlerAdm struct {
	sProduct service.ProductAdmService
}

type ProductHandler struct {
	sProduct service.ProductService
}

type InfoHandler struct {
	uInfo *usecase.Info
}

func NewProductHandlerAdm(sp service.ProductAdmService) *ProductHandlerAdm {
	return &ProductHandlerAdm{
		sProduct: sp,
	}
}

func NewProductHandler(sp service.ProductService) *ProductHandler {
	return &ProductHandler{
		sProduct: sp,
	}
}

func NewInfoHandler(u *usecase.Info) *InfoHandler {
	return &InfoHandler{
		uInfo: u,
	}
}
