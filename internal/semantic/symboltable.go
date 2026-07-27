package semantic

import (
	"sort"
	"strings"
)

type SymbolTable struct {
	byID   map[ID]Symbol
	byName map[string][]Symbol
}

func NewSymbolTable() *SymbolTable {
	return &SymbolTable{
		byID:   make(map[ID]Symbol),
		byName: make(map[string][]Symbol),
	}
}

func (t *SymbolTable) Add(s Symbol) {
	t.byID[s.ID()] = s
	t.byName[s.Name()] = append(t.byName[s.Name()], s)
}

func (t *SymbolTable) Symbol(id ID) Symbol {
	return t.byID[id]
}

func (t *SymbolTable) Symbols(name string) []Symbol {
	return t.byName[name]
}

func (t *SymbolTable) All() []Symbol {
	result := make([]Symbol, 0, len(t.byID))
	for _, s := range t.byID {
		result = append(result, s)
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].PackagePath() != result[j].PackagePath() {
			return result[i].PackagePath() < result[j].PackagePath()
		}
		return result[i].Name() < result[j].Name()
	})

	return result
}

// Fuzzy search

// Because all symbols are in one table, implementing "Go to Symbol" becomes straightforward.
// The search can be performed by prefix matching on the symbol names, and the results can be sorted by package path and name.
// You can later replace this with a radix tree or fuzzy matcher without changing the UI.
func (t *SymbolTable) Prefix(prefix string) []Symbol {
	var result []Symbol

	for _, s := range t.byID {
		if strings.HasPrefix(s.Name(), prefix) {
			result = append(result, s)
		}
	}

	sort.Slice(result, func(i, j int) bool {
		return result[i].Name() < result[j].Name()
	})

	return result
}