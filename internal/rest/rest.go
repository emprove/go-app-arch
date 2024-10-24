package rest

import (
	"go-app-arch/internal/service"
	"go-app-arch/internal/usecase"
)

type ProductHandlerAdm struct {
	sProduct *service.Product
}

type ProductHandler struct {
	sProduct *service.Product
}

type InfoHandler struct {
	uInfo *usecase.Info
}

func NewProductHandlerAdm(sp *service.Product) *ProductHandlerAdm {
	return &ProductHandlerAdm{
		sProduct: sp,
	}
}

func NewProductHandler(sp *service.Product) *ProductHandler {
	return &ProductHandler{
		sProduct: sp,
	}
}

func NewInfoHandler(u *usecase.Info) *InfoHandler {
	return &InfoHandler{
		uInfo: u,
	}
}
