package config

import "sync"

type Locale struct {
	Title    string `json:"title"`
	Iso      string `json:"iso"`
	Position int    `json:"position"`
}

type DBCfg struct {
	Dsn string
}

type Cfg struct {
	dbConfig       *DBCfg
	appURL         string
	appLumURL      string
	urlShop        string
	urlAdmin       string
	httpPort       int
	locales        []Locale
	allowedOrigins []string
}

// maybe it's better to use context.WithValue()
type DynamicState struct {
	mu     sync.RWMutex
	locale string
}

func (ds *DynamicState) SetLocale(locale string) {
	ds.mu.Lock()
	defer ds.mu.Unlock()
	ds.locale = locale
}

func (ds *DynamicState) GetLocale() string {
	ds.mu.RLock()
	defer ds.mu.RUnlock()
	return ds.locale
}

func NewConfig(dbCfg *DBCfg, appUrl, appLumUrl, urlShop, urlAdmin string, httpPort int, locales []Locale, allowedOrigins []string) *Cfg {
	return &Cfg{
		dbConfig:       dbCfg,
		appURL:         appUrl,
		appLumURL:      appLumUrl,
		urlShop:        urlShop,
		urlAdmin:       urlAdmin,
		httpPort:       httpPort,
		locales:        locales,
		allowedOrigins: allowedOrigins,
	}
}

func (c *Cfg) AvailableLocalesIso() []string {
	availableIso := []string{}
	for _, v := range c.locales {
		availableIso = append(availableIso, v.Iso)
	}
	return availableIso
}

func (c *Cfg) GetDBConfig() *DBCfg {
	return c.dbConfig
}

func (c *Cfg) GetAppURL() string {
	return c.appURL
}

func (c *Cfg) GetAppLumURL() string {
	return c.appLumURL
}

func (c *Cfg) GetUrlShop() string {
	return c.urlShop
}

func (c *Cfg) GetUrlAdmin() string {
	return c.urlAdmin
}

func (c *Cfg) GetHttpPort() int {
	return c.httpPort
}

func (c *Cfg) GetLocales() []Locale {
	return c.locales
}

func (c *Cfg) GetAllowedOrigins() []string {
	return c.allowedOrigins
}
