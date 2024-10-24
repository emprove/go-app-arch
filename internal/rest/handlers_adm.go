package rest

import (
	"errors"
	"net/http"

	"go-app-arch/internal/app"
	"go-app-arch/internal/dto"
	"go-app-arch/internal/typefmt"
)

func (h *ProductHandlerAdm) FindList(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	args := &dto.ProductFindListAdmArgs{
		ID:          typefmt.StrToIntSlice(r.Form["id"]),
		Categories:  r.Form["category[]"],
		Name:        r.Form.Get("name"),
		Price:       typefmt.StrToNilInt(r.Form.Get("price")),
		IsPublished: typefmt.StrToNilBool(r.Form.Get("is_published")),
		PerPage:     10,
		Page:        1,
	}
	if perPage, ok := typefmt.StrToInt(r.Form.Get("per_page")); ok {
		args.PerPage = perPage
	}
	if page, ok := typefmt.StrToInt(r.Form.Get("page")); ok {
		args.Page = page
	}

	res, err := h.sProduct.FindListAdm(args)
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

func (h *ProductHandlerAdm) FindOne(w http.ResponseWriter, r *http.Request) {
	id, ok := typefmt.StrToInt(r.PathValue("id"))
	if !ok {
		BadRequest(w, r, errors.New("id is required"))
		return
	}

	args := &dto.ProductFindOneAdmArgs{
		ID: id,
	}
	product, err := h.sProduct.FindOneAdm(args)
	if err != nil {
		var vErr *app.ValidationError
		if errors.As(err, &vErr) {
			FailedValidation(w, r, vErr.Validator)
			return
		}
		ServerError(w, r, err)
		return
	}

	JSON(w, http.StatusOK, product)
}
