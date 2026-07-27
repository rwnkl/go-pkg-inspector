package semantic

// import "sync"

type Index struct {

    Symbols map[string]Symbol
}

func (p *Package) buildIndex() {

	if p.index == nil {
		p.index = &Index{
			Symbols: make(map[string]Symbol),
		}
	}

	// p.index.once.Do(func() {

	// 	p.index.typesByName =
	// 		make(map[string]*Type)

	// 	p.index.functionsByName =
	// 		make(map[string]*Function)

	// 	p.index.constantsByName =
	// 		make(map[string]*Constant)

	// 	for _, t := range p.Types {
	// 		p.index.typesByName[t.Name] = t
	// 	}

	// 	for _, f := range p.Functions {
	// 		p.index.functionsByName[f.Name] = f
	// 	}

	// 	for _, c := range p.Constants {
	// 		p.index.constantsByName[c.Name] = c
	// 	}
	// })
}