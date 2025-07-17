package rest

import (
	"context"
	"errors"
	"net/http"
	"time"

	"go-app-arch/internal/app"
	"go-app-arch/internal/app/dto"
	"go-app-arch/internal/typefmt"
)

func (h *ProductHandler) FindList(w http.ResponseWriter, r *http.Request) {
	currencyIso := r.Header.Get("Currency")
	if currencyIso == "" {
		BadRequest(w, r, errors.New("currency header is required"))
		return
	}

	r.ParseForm()
	args := &dto.ProductFindListArgs{
		IDs:      typefmt.StrToIntSlice(r.Form["id[]"]),
		Category: r.Form.Get("category"),
		PerPage:  12,
		Page:     1,
	}
	if perPage, ok := typefmt.StrToInt(r.Form.Get("per_page")); ok {
		args.PerPage = perPage
	}
	if page, ok := typefmt.StrToInt(r.Form.Get("page")); ok {
		args.Page = page
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
	defer cancel()

	res, err := h.sProduct.FindList(ctx, args, currencyIso)
	if err != nil {
		var vErr *app.ValidationError
		if errors.As(err, &vErr) {
			FailedValidation(w, r, vErr.Validator)
			return
		}
		ServerError(w, r, err)
		return
	}

	JSON(w, http.StatusOK, res)
}

func (h *ProductHandler) FindOne(w http.ResponseWriter, r *http.Request) {
	slug := r.URL.Query().Get("slug")
	if slug == "" {
		BadRequest(w, r, errors.New("slug is required"))
		return
	}

	currencyIso := r.Header.Get("Currency")
	if currencyIso == "" {
		BadRequest(w, r, errors.New("currency header is required"))
		return
	}

	args := &dto.ProductFindOneArgs{
		Slug: slug,
	}

	ctx, cancel := context.WithTimeout(r.Context(), time.Second*5)
	defer cancel()

	res, err := h.sProduct.FindOne(ctx, args, currencyIso)
	if err != nil {
		var vErr *app.ValidationError
		if errors.As(err, &vErr) {
			FailedValidation(w, r, vErr.Validator)
			return
		}
		ServerError(w, r, err)
		return
	}

	JSON(w, http.StatusOK, res)
}

func (h *InfoHandler) GetLocales(w http.ResponseWriter, r *http.Request) {
	res := h.uInfo.GetLocales()
	JSON(w, http.StatusOK, res)
}

func (h *InfoHandler) GetConfig(w http.ResponseWriter, r *http.Request) {
	isAvailable := true
	pcFindArgs := &dto.ProductCategoryFindArgs{
		IsAvailable: &isAvailable,
	}
	args := &dto.GetConfigArgs{
		ProductCategoryFindArgs: pcFindArgs,
	}

	res, err := h.uInfo.GetConfig(args, h.uInfo.DS.GetLocale())
	if err != nil {
		ServerError(w, r, err)
		return
	}

	JSON(w, http.StatusOK, res)
}
