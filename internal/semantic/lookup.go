package semantic

func (p *Package) Type(name string) *Type {

	p.buildIndex()

	// return p.index.typesByName[name]
	return nil
}

func (p *Package) Function(name string) *Function {

	p.buildIndex()

	// return p.index.functionsByName[name]
	return nil
}

func (p *Package) Constant(name string) *Constant {

	p.buildIndex()

	// return p.index.constantsByName[name]
	return nil
}