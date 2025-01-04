package mapper

import (
	"cmp"
	"slices"

	"go-app-arch/internal/app/config"
	"go-app-arch/internal/domain/entity"
	"go-app-arch/internal/utils"
)

type FileMapper struct {
	Cfg *config.Cfg
}

func NewFileMapper(cfg *config.Cfg) *FileMapper {
	return &FileMapper{Cfg: cfg}
}

func (m *FileMapper) JsonFilesToFiles(jsonfiles []entity.FileJson) []entity.File {
	var files []entity.File
	for _, jf := range jsonfiles {
		var f entity.File
		f.ID = jf.ID
		f.Name = jf.Name
		f.Path = utils.GetResourceStorageUrl(jf.Path, m.Cfg.GetAppLumURL())
		f.PathThumb = utils.GetResourceStorageUrl(jf.PathThumb, m.Cfg.GetAppLumURL())
		f.Position = jf.Position
		files = append(files, f)
	}
	slices.SortFunc(files, func(a, b entity.File) int {
		var av, bv int
		if a.Position != nil {
			av = *a.Position
		}
		if b.Position != nil {
			bv = *b.Position
		}
		return cmp.Compare(av, bv)
	})

	return files
}
