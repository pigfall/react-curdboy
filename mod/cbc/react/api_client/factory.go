package client

import(
	ent "github.com/pigfall/ent_utils"
)

type Factory struct{

}

func (Factory) ClientGenerator(nodes []*ent.Type)*ClientGenerator{
	return &ClientGenerator{
		Nodes: nodes,
	}
}
