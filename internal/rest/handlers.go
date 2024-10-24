package rest

import (
	"errors"
	"net/http"

	"go-app-arch/internal/app"
	"go-app-arch/internal/dto"
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

	res, err := h.sProduct.FindList(args, currencyIso)
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

	res, err := h.sProduct.FindOne(args, currencyIso)
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

	res, err := h.uInfo.GetConfig(args, h.uInfo.DS.Locale)
	if err != nil {
		ServerError(w, r, err)
		return
	}

	JSON(w, http.StatusOK, res)
}
