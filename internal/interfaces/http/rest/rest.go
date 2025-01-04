package rest

import (
	"go-app-arch/internal/app/usecase"
	"go-app-arch/internal/domain/service"
)

type ProductHandlerAdm struct {
	sProduct service.ProductServiceInterface
}

type ProductHandler struct {
	sProduct service.ProductServiceInterface
}

type InfoHandler struct {
	uInfo *usecase.Info
}

func NewProductHandlerAdm(sp service.ProductServiceInterface) *ProductHandlerAdm {
	return &ProductHandlerAdm{
		sProduct: sp,
	}
}

func NewProductHandler(sp service.ProductServiceInterface) *ProductHandler {
	return &ProductHandler{
		sProduct: sp,
	}
}

func NewInfoHandler(u *usecase.Info) *InfoHandler {
	return &InfoHandler{
		uInfo: u,
	}
}
