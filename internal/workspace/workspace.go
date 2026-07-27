package workspace

// Workspace is the root object of the application.
type Workspace struct {
	Module *Module

	Packages []*Package
	
	index *Index
}

// PackageCount returns number of packages.
func (w *Workspace) PackageCount() int {
	return len(w.Packages)
}

// FileCount returns number of Go files.
func (w *Workspace) FileCount() int {

	n := 0

	for _, p := range w.Packages {
		n += len(p.Files)
	}

	return n
}

func (w *Workspace) PackageByPath(path string) *Package {

	w.buildIndex()

	return w.index.packagesByPath[path]
}

func (w *Workspace) PackagesByName(name string) []*Package {

	w.buildIndex()

	return w.index.packagesByName[name]
}

func (w *Workspace) FileByPath(path string) *File {

	w.buildIndex()

	return w.index.filesByPath[path]
}

func (w *Workspace) FilesByName(name string) []*File {

	w.buildIndex()

	return w.index.filesByName[name]
}
