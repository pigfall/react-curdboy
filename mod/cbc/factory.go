package cbc

import(
	ent "github.com/pigfall/ent_utils"
)

type Factory struct{}

func (Factory) NewGenerator(entSchemaPath,outputPath string)(*Generator,error){
	g,err :=  ent.LoadGraph(entSchemaPath)
	if err != nil{
		return nil,err
	}
	return &Generator{
		graph:g,
		entSchemaPath:entSchemaPath,
		outputPath:outputPath,
	},nil
}

func (Factory) NewModelGenerator(generator *Generator)*ModelGenerator{
	return &ModelGenerator{
		generator:generator,
	}
}
