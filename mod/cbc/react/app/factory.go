package app

import (
	ent "github.com/pigfall/ent_utils"
)

type Factory struct{}

func (Factory) SimpleAppGenerator(nodes []*ent.Type) *SimpleAppGenerator {
	return &SimpleAppGenerator{
		nodes: nodes,
	}
}
