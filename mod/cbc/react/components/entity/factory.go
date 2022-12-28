package entity

import (
	ent "github.com/pigfall/ent_utils"
)

type Factory struct{}

func (Factory) NewEntityListViewGenerator(node *ent.Type) *EntityListViewGenerator {
	return &EntityListViewGenerator{
		EntNode: node,
	}
}
