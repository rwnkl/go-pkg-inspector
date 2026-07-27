package workspace

import (
	"path/filepath"
	"sync"
)

type Index struct {
	once sync.Once

	packagesByPath map[string]*Package
	packagesByName map[string][]*Package

	filesByPath map[string]*File
	filesByName map[string][]*File
}

func (w *Workspace) buildIndex() {

	if w.index == nil {
		w.index = &Index{}
	}

	w.index.once.Do(func() {

		w.index.packagesByPath = make(map[string]*Package)
		w.index.packagesByName = make(map[string][]*Package)

		w.index.filesByPath = make(map[string]*File)
		w.index.filesByName = make(map[string][]*File)

		for _, pkg := range w.Packages {

			w.index.packagesByPath[pkg.Path] = pkg

			w.index.packagesByName[pkg.Name] =
				append(w.index.packagesByName[pkg.Name], pkg)

			for _, file := range pkg.Files {

				w.index.filesByPath[file.Path] = file

				name := filepath.Base(file.Path)

				w.index.filesByName[name] =
					append(w.index.filesByName[name], file)
			}
		}
	})
}