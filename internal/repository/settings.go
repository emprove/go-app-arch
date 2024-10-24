package repository

import (
	"context"
	"encoding/json"
	"errors"
	"strings"

	"go-app-arch/internal/dto"
	"go-app-arch/internal/entity"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Settings struct {
	db *pgxpool.Pool
}

func NewSettingsRepository(db *pgxpool.Pool) *Settings {
	return &Settings{db: db}
}

func (repo *Settings) GetAll() ([]entity.Settings, error) {
	query := `
	SELECT 
		id,
		key, 
		meta
	FROM settings`

	rows, err := repo.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	dbmodels, err := pgx.CollectRows(rows, pgx.RowToStructByName[entity.Settings])
	if err != nil {
		return nil, err
	}

	return dbmodels, nil
}

func (repo *Settings) FindProductCategories(args *dto.ProductCategoryFindArgs, locale string) ([]entity.ProductCategory, error) {
	type dbrow struct {
		Title    string
		TitleEn  string
		Code     string
		Position int
	}

	tmpl := `
	select 
		pc.title, 
		pc.title_en, 
		pc.code,
		pc.position
	from product_categories pc
	{{where}}
	order by pc.position asc`

	conditions := []string{"where true"}
	namedArgs := pgx.NamedArgs{}

	if args.IsAvailable != nil {
		conditions = append(conditions, "pc.is_available = @is_available")
		namedArgs["is_available"] = args.IsAvailable
	}

	conditionsStr := strings.Join(conditions, " AND ")
	query := strings.Replace(tmpl, "{{where}}", conditionsStr, 1)

	rows, err := repo.db.Query(context.Background(), query, namedArgs)
	if err != nil {
		return nil, err
	}

	dbmodels, err := pgx.CollectRows(rows, pgx.RowToStructByName[dbrow])
	if err != nil {
		return nil, err
	}

	var res []entity.ProductCategory
	for _, m := range dbmodels {
		var pc entity.ProductCategory
		pc.Code = m.Code
		pc.Position = m.Position
		switch locale {
		case "en":
			pc.Title = m.TitleEn
		case "ru":
			pc.Title = m.Title
		default:
			return nil, errors.New("locale unknown")
		}
		res = append(res, pc)
	}

	return res, nil
}

func (repo *Settings) FindProductOptions(locale string) ([]entity.ProductOption, error) {
	type dbrow struct {
		Title   string
		TitleEn string
		Code    string
		Typeof  string
		Data    string
	}

	query := `
	select 
		po.title, 
		po.title_en, 
		po.code,
		po.typeof,
		po.data
	from product_options po
	order by po.id asc`

	rows, err := repo.db.Query(context.Background(), query)
	if err != nil {
		return nil, err
	}

	dbmodels, err := pgx.CollectRows(rows, pgx.RowToStructByName[dbrow])
	if err != nil {
		return nil, err
	}

	var res []entity.ProductOption
	for _, m := range dbmodels {
		var po entity.ProductOption
		po.Code = m.Code
		po.Typeof = m.Typeof
		switch locale {
		case "en":
			po.Title = m.TitleEn
		case "ru":
			po.Title = m.Title
		default:
			return nil, errors.New("locale unknown")
		}
		if err := json.Unmarshal([]byte(m.Data), &po.Data); err != nil {
			return nil, err
		}
		res = append(res, po)
	}

	return res, nil
}
