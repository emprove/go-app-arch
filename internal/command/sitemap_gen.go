package command

import (
	"os"

	"go-app-arch/internal/database/postgres"
	"go-app-arch/internal/dto"
	"go-app-arch/internal/mapper"
)

func (c *SitemapGen) Run(args []string) *error {
	productMapper := mapper.NewProductMapper(c.app.Cfg)
	fileMapper := mapper.NewFileMapper(c.app.Cfg)
	productRepo := postgres.NewProductRepository(c.app.DB, productMapper, fileMapper)

	str := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n"
	str += "<urlset xmlns=\"http://www.sitemaps.org/schemas/sitemap/0.9\">\n"

	for _, locale := range c.app.Cfg.AvailableLocalesIso() {
		str += "<url>\n"
		str += "<loc>" + c.app.Cfg.GetUrlShop() + "/" + locale + "/" + "</loc>\n"
		str += "</url>\n"

		str += "<url>\n"
		str += "<loc>" + c.app.Cfg.GetUrlShop() + "/" + locale + "/service</loc>\n"
		str += "</url>\n"

		str += "<url>\n"
		str += "<loc>" + c.app.Cfg.GetUrlShop() + "/" + locale + "/products</loc>\n"
		str += "</url>\n"

		args := &dto.ProductFindListArgs{IDs: []int{}, Category: "", PerPage: 1000, Page: 1}
		products, err := productRepo.FindList(args, locale)
		if err != nil {
			return &err
		}

		for _, product := range products.Products {
			str += "<url>\n"
			str += "<loc>" + c.app.Cfg.GetUrlShop() + "/" + locale + "/products/" + product.Slug + "</loc>\n"
			str += "</url>\n"
		}
	}

	str += "</urlset>"

	os.WriteFile("sitemap.xml", []byte(str), 0644)

	return nil
}
