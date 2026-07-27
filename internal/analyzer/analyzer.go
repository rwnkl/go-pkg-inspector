package analyzer

import (
	"github.com/rwnkl/go-package-inspector/internal/semantic"
	"github.com/rwnkl/go-package-inspector/internal/workspace"
)

type Analyzer struct {
	collectors []Collector
}

func New() *Analyzer {

	return &Analyzer{
		collectors: []Collector{

    		&TypeCollector{},

    		&StructCollector{},
		},
	}
}

func (a *Analyzer) Analyze(
	ws *workspace.Workspace,
) (*semantic.Model, error) {

	model := &semantic.Model{}

	ctx := &Context{
		Workspace: ws,
		Model:     model,
	}

	for i, wp := range ws.Packages {

		ctx.Package = &PackageContext{
			Workspace: ws,

			WorkspacePackage: wp,

			SemanticPackage: model.Packages[i],
		}

		for _, c := range a.collectors {

			if err := c.Collect(ctx); err != nil {
				return nil, err
			}
		}
	}

	return model, nil
}