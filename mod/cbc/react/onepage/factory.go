package onepage

import (
	ent "github.com/pigfall/ent_utils"
)

type Factory struct{}

func (Factory) SchemasViewerGenerator(nodes []*ent.Type) *SchemasViewerGenerator {
	return &SchemasViewerGenerator{
		Nodes: nodes,
	}
}

func (Factory) EntitiesViewerGenerator(nodes []*ent.Type) *EntitiesViewerGenerator {
	return &EntitiesViewerGenerator{
		Nodes: nodes,
	}
}
