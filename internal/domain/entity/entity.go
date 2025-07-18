package entity

import (
	"fmt"
	"time"

	"go-app-arch/internal/domain/valueobject"
)

type Settings struct {
	ID   int    `json:"id"`
	Key  string `json:"key"`
	Meta any    `json:"meta"`
}

type Supply struct {
	ID       int `json:"id"`
	Options  any `json:"options"`
	Quantity int `json:"quantity"`
}

type User struct {
	ID          int               `json:"id"`
	Name        string            `json:"name"`
	Email       valueobject.Email `json:"email"`
	Permissions []string          `json:"permissions"`
}

type FileJson struct {
	ID        int     `json:"id"`
	Position  *int    `json:"position"`
	Name      *string `json:"name"`
	Path      string  `json:"path"`
	PathThumb string  `json:"path_thumb"`
}

type File struct {
	ID        int     `json:"id"`
	Name      *string `json:"name"`
	Path      string  `json:"path"`
	PathThumb string  `json:"path_thumb"`
	Position  *int    `json:"position"`
}

type ProductCategory struct {
	Title    string `json:"title"`
	Code     string `json:"code"`
	Position int    `json:"position"`
}

type ProductOption struct {
	Title  string `json:"title"`
	Code   string `json:"code"`
	Typeof string `json:"typeof"`
	Data   any    `json:"data"`
}

type TimeYMD struct {
	*time.Time
}

func (t *TimeYMD) MarshalJSON() ([]byte, error) {
	if t.Time == nil {
		return []byte("null"), nil
	}

	return fmt.Appendf(nil, "\"%s\"", t.Format("2006-01-02")), nil
}
