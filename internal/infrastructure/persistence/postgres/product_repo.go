package postgres

import (
	"context"
	"encoding/json"
	"errors"
	"strings"
	"time"

	"go-app-arch/internal/app/dto"
	"go-app-arch/internal/domain/entity"
	"go-app-arch/internal/domain/repository"
	"go-app-arch/internal/infrastructure/database"
	"go-app-arch/internal/infrastructure/mapper"
	"go-app-arch/internal/typefmt"
	"go-app-arch/internal/utils"

	"github.com/jackc/pgx/pgtype"
	"github.com/jackc/pgx/v5"
)

type productRepo struct {
	db            database.DB
	productMapper *mapper.ProductMapper
	fileMapper    *mapper.FileMapper
}

func NewProductRepository(db database.DB, pm *mapper.ProductMapper, fm *mapper.FileMapper) repository.Product {
	return &productRepo{db: db, productMapper: pm, fileMapper: fm}
}

func (repo *productRepo) FindOneAdm(args *dto.ProductFindOneAdmArgs) (*dto.ProductFindOneRowAdm, error) {
	type dbrow struct {
		ID              int
		CurrencyIso     string
		Category        string
		Name            string
		Annotation      *string
		Description     *string
		Price           float64
		Position        *int
		Options         string
		CreatedAt       *time.Time
		UpdatedAt       *time.Time
		VideoPath       *string
		Slug            string
		MakingInDaysMin *int
		MakingInDaysMax *int
		NameEn          *string
		AnnotationEn    *string
		DescriptionEn   *string
		SlugEn          string
		IsAvailable     bool
		IsPublished     bool
		Files           string
		Supplies        string
		RelatedIDs      string
	}
	query := `
	SELECT p.id AS "ID",
		p.category AS "Category",
		p.name AS "Name",
		p.annotation AS "Annotation",
		p.description AS "Description",
		p.price AS "Price",
		p.position AS "Position",
		COALESCE(p.options, '[]') AS "Options",
		p.created_at AS "CreatedAt",
		p.updated_at AS "UpdatedAt",
		p.video_path AS "VideoPath",
		p.slug AS "Slug",
		p.making_in_days_min AS "MakingInDaysMin",
		p.making_in_days_max AS "MakingInDaysMax",
		p.name_en AS "NameEn",
		p.annotation_en AS "AnnotationEn",
		p.description_en AS "DescriptionEn",
		p.slug_en AS "SlugEn",
		p.is_available AS "IsAvailable",
		p.is_published AS "IsPublished",
		c.iso_alfa AS "currency_iso",
		coalesce(jsonb_agg(distinct jsonb_build_object('id', phf.file_id, 'position', phf.position, 'name', f.name, 'path', f.path, 'path_thumb', f.path_thumb)) FILTER (WHERE phf.file_id IS NOT NULL), '[]') as Files,
		coalesce(jsonb_agg(distinct jsonb_build_object('id', ps.id, 'options', ps.options, 'quantity', ps.quantity)) FILTER (WHERE ps.id IS NOT NULL), '[]') as Supplies,
		coalesce(jsonb_agg(distinct r.related_id) FILTER (WHERE r.related_id IS NOT NULL), '[]') as RelatedIDs 
	FROM products AS p
		LEFT JOIN product_has_files AS phf ON (p.id = phf.product_id)
		LEFT JOIN files AS f ON (phf.file_id = f.id)
		LEFT JOIN product_supplies AS ps ON (p.id = ps.product_id)
		LEFT JOIN product_has_related AS r ON (p.id = r.product_id)
		LEFT JOIN currencies AS c ON (c.id = p.currency_id)
	where p.id = @id
	group by p.id, c.iso_alfa;`

	namedArgs := pgx.NamedArgs{"id": args.ID}
	rows, err := repo.db.Query(context.TODO(), query, namedArgs)
	if err != nil {
		return nil, err
	}

	dbmodel, err := pgx.CollectExactlyOneRow(rows.(pgx.Rows), pgx.RowToStructByName[dbrow])
	if err != nil {
		return nil, err
	}

	res := dto.ProductFindOneRowAdm{
		ID:              dbmodel.ID,
		CurrencyIso:     dbmodel.CurrencyIso,
		Category:        dbmodel.Category,
		Name:            dbmodel.Name,
		Annotation:      dbmodel.Annotation,
		Description:     dbmodel.Description,
		Price:           dbmodel.Price,
		Position:        dbmodel.Position,
		CreatedAt:       dbmodel.CreatedAt,
		UpdatedAt:       dbmodel.UpdatedAt,
		VideoPath:       dbmodel.VideoPath,
		Slug:            dbmodel.Slug,
		MakingInDaysMin: dbmodel.MakingInDaysMin,
		MakingInDaysMax: dbmodel.MakingInDaysMax,
		NameEn:          dbmodel.NameEn,
		AnnotationEn:    dbmodel.AnnotationEn,
		DescriptionEn:   dbmodel.DescriptionEn,
		SlugEn:          dbmodel.SlugEn,
		IsAvailable:     dbmodel.IsAvailable,
		IsPublished:     dbmodel.IsPublished,
		RelatedIDs:      typefmt.JsonArrayToIntSlice(dbmodel.RelatedIDs),
	}
	if err := json.Unmarshal([]byte(dbmodel.Options), &res.Options); err != nil {
		return nil, err
	}

	// files
	var filesJson []entity.FileJson
	if err := json.Unmarshal([]byte(dbmodel.Files), &filesJson); err != nil {
		return nil, err
	}
	if len(filesJson) > 0 {
		res.Files = repo.fileMapper.JsonFilesToFiles(filesJson)
	} else {
		res.Files = make([]entity.File, 0)
	}

	// supplies
	var supplies []entity.Supply
	if err := json.Unmarshal([]byte(dbmodel.Supplies), &supplies); err != nil {
		return nil, err
	}
	res.Supplies = supplies

	return &res, nil
}

func (repo *productRepo) FindListAdm(args *dto.ProductFindListAdmArgs) (*dto.ProductFindListAdm, error) {
	type dbrow struct {
		TotalCount  int
		ID          int
		CurrencyIso string
		Category    string
		Name        string
		Price       float64
		Position    *int
		IsPublished bool
		Files       string
	}
	tmpl := `
	WITH cte AS (
		select p.id as ids,
		COUNT(*) OVER() AS total_count
		from products p
		{{where}}
		ORDER BY p.position, p.id ASC NULLS last
		LIMIT @limit
		OFFSET @offset
	)
	SELECT (SELECT total_count FROM cte LIMIT 1) AS TotalCount,
		p.id AS "id",
		p.category AS "category",
		p.name AS "name",
		p.price AS "price",
		p.position AS "position",
		p.is_published AS "is_published",
		c.iso_alfa AS "currency_iso",
		coalesce(jsonb_agg(distinct jsonb_build_object('id', phf.file_id, 'position', phf.position, 'name', f.name, 'path', f.path, 'path_thumb', f.path_thumb)) FILTER (WHERE phf.file_id IS NOT NULL), '[]') as files
	FROM products AS p
		LEFT JOIN product_has_files AS phf ON (p.id = phf.product_id)
		LEFT JOIN files AS f ON (phf.file_id = f.id)
		LEFT JOIN currencies AS c ON (c.id = p.currency_id)
	where p.id in (SELECT ids FROM cte)
	group by p.id, c.iso_alfa
	ORDER BY p.position, p.id ASC NULLS last;`

	conditions := []string{"where true"}
	namedArgs := pgx.NamedArgs{}

	if len(args.ID) > 0 {
		ids := &pgtype.Int4Array{}
		ids.Set(args.ID)
		conditions = append(conditions, "p.id = ANY(@ids)")
		namedArgs["ids"] = ids
	}
	if len(args.Categories) > 0 {
		conditions = append(conditions, "p.category = ANY(@categories)")
		namedArgs["categories"] = args.Categories
	}
	if args.Name != "" {
		conditions = append(conditions, "p.name like @name")
		namedArgs["name"] = "%" + args.Name + "%"
	}
	if args.Price != nil {
		conditions = append(conditions, "p.price = @price")
		namedArgs["price"] = args.Price
	}
	if args.IsPublished != nil {
		conditions = append(conditions, "p.is_published = @is_published")
		namedArgs["is_published"] = args.IsPublished
	}

	namedArgs["limit"] = args.PerPage
	namedArgs["offset"] = utils.GetOffset(args.Page, args.PerPage)

	conditionsStr := strings.Join(conditions, " AND ")
	query := strings.Replace(tmpl, "{{where}}", conditionsStr, 1)

	rows, err := repo.db.Query(context.Background(), query, namedArgs)
	if err != nil {
		return nil, err
	}

	dbmodels, err := pgx.CollectRows(rows.(pgx.Rows), pgx.RowToStructByName[dbrow])
	if err != nil {
		return nil, err
	}

	var res dto.ProductFindListAdm
	if len(dbmodels) == 0 {
		res.Products = make([]dto.ProductFindListProductAdm, 0)
		return &res, nil
	}

	res.TotalCount = dbmodels[0].TotalCount
	for _, m := range dbmodels {
		var p dto.ProductFindListProductAdm
		p.ID = m.ID
		p.CurrencyIso = m.CurrencyIso
		p.Category = m.Category
		p.Name = m.Name
		p.Price = m.Price
		p.Position = m.Position
		p.IsPublished = m.IsPublished

		var jsonFiles []entity.FileJson
		err := json.Unmarshal([]byte(m.Files), &jsonFiles)
		if err != nil {
			return nil, err
		}
		p.Files = repo.fileMapper.JsonFilesToFiles(jsonFiles)

		res.Products = append(res.Products, p)
	}

	return &res, nil
}

func (repo *productRepo) FindList(args *dto.ProductFindListArgs, locale string) (*dto.ProductFindList, error) {
	type dbrow struct {
		TotalCount    int
		ID            int
		CurrencyIso   string
		Category      string
		Name          string
		Annotation    *string
		Description   *string
		Price         float64
		Position      *int
		Slug          string
		NameEn        *string
		AnnotationEn  *string
		DescriptionEn *string
		SlugEn        string
		Files         string
	}
	tmpl := `
	WITH cte AS (
		select p.id as ids,
		COUNT(*) OVER() AS total_count
		from products p
		{{where}}
		ORDER BY p.position, p.id ASC NULLS last
		LIMIT @limit
		OFFSET @offset
	)
	SELECT (SELECT total_count FROM cte LIMIT 1) AS TotalCount,
		p.id AS "id",
		p.category AS "category",
		p.name AS "name",
		p.annotation AS "annotation",
		p.description AS "description",
		p.price AS "price",
		p.position AS "position",
		p.slug AS "slug",
		p.name_en AS "name_en",
		p.annotation_en AS "annotation_en",
		p.description_en AS "description_en",
		p.slug_en AS "slug_en",
		c.iso_alfa AS "currency_iso",
		coalesce(jsonb_agg(distinct jsonb_build_object('id', phf.file_id, 'position', phf.position, 'name', f.name, 'path', f.path, 'path_thumb', f.path_thumb)) FILTER (WHERE phf.file_id IS NOT NULL), '[]') as files
	FROM products AS p
		LEFT JOIN product_has_files phf ON (p.id = phf.product_id)
		LEFT JOIN files f ON (phf.file_id = f.id)
		LEFT JOIN currencies c ON (c.id = p.currency_id)
	where p.id in (SELECT ids FROM cte)
	group by p.id, c.iso_alfa
	ORDER BY p.position, p.id ASC NULLS last;`

	conditions := []string{"where true", "p.is_published = true"}
	namedArgs := pgx.NamedArgs{}
	if args.Category != "" {
		conditions = append(conditions, "p.category = @category")
		namedArgs["category"] = args.Category
	}
	if len(args.IDs) > 0 {
		ids := &pgtype.Int4Array{}
		ids.Set(args.IDs)
		conditions = append(conditions, "p.id = ANY(@ids)")
		namedArgs["ids"] = ids
	}

	namedArgs["limit"] = args.PerPage
	namedArgs["offset"] = utils.GetOffset(args.Page, args.PerPage)

	conditionsStr := strings.Join(conditions, " AND ")
	query := strings.Replace(tmpl, "{{where}}", conditionsStr, 1)

	rows, err := repo.db.Query(context.Background(), query, namedArgs)
	if err != nil {
		return nil, err
	}

	dbmodels, err := pgx.CollectRows(rows.(pgx.Rows), pgx.RowToStructByName[dbrow])
	if err != nil {
		return nil, err
	}

	var res dto.ProductFindList
	if len(dbmodels) == 0 {
		res.Products = make([]dto.ProductFindListProduct, 0)
		return &res, nil
	}

	res.TotalCount = dbmodels[0].TotalCount
	for _, m := range dbmodels {
		var p dto.ProductFindListProduct
		p.ID = m.ID
		p.CurrencyIso = m.CurrencyIso
		p.Category = m.Category
		p.Price = m.Price
		p.Position = m.Position
		p.Slug = m.Slug
		switch locale {
		case "en":
			if m.NameEn != nil {
				p.Name = *m.NameEn
			} else {
				p.Name = m.Name
			}
			p.Annotation = m.AnnotationEn
			p.Description = m.DescriptionEn
			p.Slug = m.SlugEn
		case "ru":
			p.Name = m.Name
			p.Annotation = m.Annotation
			p.Description = m.Description
			p.Slug = m.Slug
		default:
			return nil, errors.New("locale unknown")
		}

		var jsonFiles []entity.FileJson
		err := json.Unmarshal([]byte(m.Files), &jsonFiles)
		if err != nil {
			return nil, err
		}
		p.Files = repo.fileMapper.JsonFilesToFiles(jsonFiles)

		res.Products = append(res.Products, p)
	}

	return &res, nil
}

func (repo *productRepo) FindOne(args *dto.ProductFindOneArgs, locale string) (*dto.ProductFindOneRow, error) {
	type dbrow struct {
		ID              int
		CurrencyIso     string
		Category        string
		Name            string
		Annotation      *string
		Description     *string
		Price           float64
		Options         string
		VideoPath       *string
		Slug            string
		MakingInDaysMin *int
		MakingInDaysMax *int
		IsAvailable     bool
		NameEn          *string
		AnnotationEn    *string
		DescriptionEn   *string
		SlugEn          string
		Files           string
		Supplies        string
		RelatedIDs      string
	}
	query := `
	SELECT p.id AS "ID",
		p.category AS "Category",
		p.name AS "Name",
		p.annotation AS "Annotation",
		p.description AS "Description",
		p.price AS "Price",
		COALESCE(p.options, '[]') AS "Options",
		p.video_path AS "VideoPath",
		p.slug AS "Slug",
		p.making_in_days_min AS "MakingInDaysMin",
		p.making_in_days_max AS "MakingInDaysMax",
		p.is_available AS "IsAvailable",
		p.name_en AS "NameEn",
		p.annotation_en AS "AnnotationEn",
		p.description_en AS "DescriptionEn",
		p.slug_en AS "SlugEn",
		c.iso_alfa AS "currency_iso",
		coalesce(jsonb_agg(distinct jsonb_build_object('id', phf.file_id, 'position', phf.position, 'name', f.name, 'path', f.path, 'path_thumb', f.path_thumb)) FILTER (WHERE phf.file_id IS NOT NULL), '[]') as Files,
		coalesce(jsonb_agg(distinct jsonb_build_object('id', ps.id, 'options', ps.options, 'quantity', ps.quantity)) FILTER (WHERE ps.id IS NOT NULL), '[]') as Supplies,
		coalesce(jsonb_agg(distinct r.related_id) FILTER (WHERE r.related_id IS NOT NULL), '[]') as RelatedIDs 
	FROM products AS p
		LEFT JOIN product_has_files AS phf ON (p.id = phf.product_id)
		LEFT JOIN files AS f ON (phf.file_id = f.id)
		LEFT JOIN product_supplies AS ps ON (p.id = ps.product_id)
		LEFT JOIN product_has_related AS r ON (p.id = r.product_id)
		LEFT JOIN currencies AS c ON (c.id = p.currency_id)
	where p.slug = @slug 
		OR p.slug_en = @slug
	group by p.id, c.iso_alfa;`

	namedArgs := pgx.NamedArgs{"slug": args.Slug}
	rows, err := repo.db.Query(context.TODO(), query, namedArgs)
	if err != nil {
		return nil, err
	}

	dbmodel, err := pgx.CollectExactlyOneRow(rows.(pgx.Rows), pgx.RowToStructByName[dbrow])
	if err != nil {
		return nil, err
	}

	res := dto.ProductFindOneRow{
		ID:              dbmodel.ID,
		CurrencyIso:     dbmodel.CurrencyIso,
		Category:        dbmodel.Category,
		Name:            dbmodel.Name,
		Annotation:      dbmodel.Annotation,
		Description:     dbmodel.Description,
		Price:           dbmodel.Price,
		Slug:            dbmodel.Slug,
		MakingInDaysMin: dbmodel.MakingInDaysMin,
		MakingInDaysMax: dbmodel.MakingInDaysMax,
		IsAvailable:     dbmodel.IsAvailable,
		RelatedIDs:      typefmt.JsonArrayToIntSlice(dbmodel.RelatedIDs),
	}
	if dbmodel.VideoPath != nil {
		res.VideoPath = utils.GetResourceStorageUrl(*dbmodel.VideoPath, repo.fileMapper.Cfg.GetAppLumURL())
	} else {
		res.VideoPath = ""
	}
	if err := json.Unmarshal([]byte(dbmodel.Options), &res.Options); err != nil {
		return nil, err
	}

	switch locale {
	case "en":
		if dbmodel.NameEn != nil {
			res.Name = *dbmodel.NameEn
		} else {
			res.Name = dbmodel.Name
		}
		res.Annotation = dbmodel.AnnotationEn
		res.Description = dbmodel.DescriptionEn
		res.Slug = dbmodel.SlugEn
	case "ru":
		res.Name = dbmodel.Name
		res.Annotation = dbmodel.Annotation
		res.Description = dbmodel.Description
		res.Slug = dbmodel.Slug
	default:
		return nil, errors.New("locale unknown")
	}

	// files
	var filesJson []entity.FileJson
	if err := json.Unmarshal([]byte(dbmodel.Files), &filesJson); err != nil {
		return nil, err
	}
	res.Files = repo.fileMapper.JsonFilesToFiles(filesJson)

	// supplies
	var supplies []entity.Supply
	if err := json.Unmarshal([]byte(dbmodel.Supplies), &supplies); err != nil {
		return nil, err
	}
	res.Supplies = supplies

	return &res, nil
}
