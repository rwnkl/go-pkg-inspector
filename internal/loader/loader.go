package loader

import (
	"context"
	"fmt"
	"os"
	"path/filepath"

	"github.com/rwnkl/go-package-inspector/internal/workspace"
	"golang.org/x/tools/go/packages"
)

type Loader struct {
	config Config
	observer Observer
}


func New(config Config, observer Observer) *Loader {
	return &Loader{
		config: config,
		observer: observer,
	}
}

func (l *Loader) emit(e Event) {
	if l.observer != nil {
		l.observer.OnEvent(e)
	}
}

func (l *Loader) Load(ctx context.Context, observer Observer) (*Result, error) {

	if l.config.Pattern == "" {
		return nil, ErrEmptyPattern
	}

	l.emit(Event{
		Kind:    EventMessage,
		Message: "Checking Go module...",
	})

	if err := verifyModule(l.config.Dir); err != nil {
		return nil, err
	}

	l.emit(Event{
		Kind: EventModule,
		Name: l.config.Dir,
	})

	cfg := l.config.PackageConfig()
	cfg.Context = ctx

	l.emit(Event{
		Kind:    EventMessage,
		Message: "Loading packages...",
	})
	
	pkgs, err := packages.Load(cfg, l.config.Pattern)
	if err != nil {
		return nil, err
	}

	builder := workspace.NewBuilder()
	// result := &Result{
    // 	Workspace: ws,
	// }
	result := &Result{}

	for i, pkg := range pkgs {

		l.emit(Event{
			Kind:    EventPackage,
			Name:    pkg.PkgPath,
			Current: i + 1,
			Total:   len(pkgs),
		})

		builder.AddPackage(pkg)

		// wp := &workspace.Package{
		// 	Name: pkg.Name,
		// 	Path: pkg.PkgPath,
		// 	Pkg:  pkg,
		// }

		// if pkg.Types != nil {
		// 	wp.Types = pkg.Types
		// }

		// if pkg.Fset != nil {
		// 	wp.FileSet = pkg.Fset
		// }

		// if len(pkg.Syntax) > 0 {
		// 	wp.AST = pkg.Syntax
		// }

		// for _, file := range pkg.GoFiles {

		// 	l.emit(Event{
		// 		Kind: EventFile,
		// 		Name: file,
		// 	})

		// 	wp.Files = append(
		// 		wp.Files,
		// 		&workspace.File{
		// 			Name: filepath.Base(file),
		// 			Path: file,
		// 		},
		// 	)
		// }

		for _, err := range pkg.Errors {

			d := Diagnostic{
				Severity: SeverityError,
				Package:  pkg.PkgPath,
				Message:  err.Msg,
			}

			result.Diagnostics = append(result.Diagnostics, d)

			l.emit(Event{
				Kind:       EventDiagnostic,
				Diagnostic: &d,
			})
		}

		// ws.Packages = append(ws.Packages, wp)

		// if ws.Module == nil && pkg.Module != nil {

		// 	ws.Module = &workspace.Module{
		// 		Name: filepath.Base(pkg.Module.Path),
		// 		Path: pkg.Module.Path,
		// 	}
		// }

		if pkg.Module != nil {
			
			builder.SetModule(
				filepath.Base(pkg.Module.Path),
				pkg.Module.Path,
			)
		}

		// builder.AddPackage(pkg)
	}

	result.Workspace = builder.Workspace()

	return result, nil
}

func verifyModule(dir string) error {

	name := filepath.Join(dir, "go.mod")

	info, err := os.Stat(name)

	if err != nil {
		return ErrNoGoMod
	}

	if info.IsDir() {
		return fmt.Errorf("%s is a directory", name)
	}

	return nil
}