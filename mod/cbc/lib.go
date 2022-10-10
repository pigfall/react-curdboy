package cbc

import(
	"embed"
	ent "github.com/pigfall/ent_utils"
)


//go:embed tpls/*
var templates embed.FS

type Generator struct{
	graph *ent.Graph
	entSchemaPath string
	outputPath string
}

func (g *Generator) Generate() error{
	factory := Factory{}
	if err :=factory.NewModelGenerator(g).Generate();err != nil{
		return err
	}



	panic("TODO")
}
