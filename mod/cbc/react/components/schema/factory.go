package schema

import (
	ent "github.com/pigfall/ent_utils"
)

type Factory struct{}

func (Factory) SchemaViewGenerator(entNode *ent.Type) *SchemaViewGenerator {
	return &SchemaViewGenerator{EntNode: entNode}
}
