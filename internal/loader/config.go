package loader

import (
	"golang.org/x/tools/go/packages"
)

type Config struct {
	Dir     string
	Pattern string
}

func DefaultConfig(dir string) Config {
	return Config{
		Dir:     dir,
		Pattern: "./...",
	}
}

func (c Config) PackageConfig() *packages.Config {

	return &packages.Config{
		Dir: c.Dir,

		Mode: packages.NeedName |
			packages.NeedModule |
			packages.NeedFiles |
			packages.NeedCompiledGoFiles |
			packages.NeedImports |
			packages.NeedSyntax |
			packages.NeedTypes |
			packages.NeedTypesInfo,
	}
}
