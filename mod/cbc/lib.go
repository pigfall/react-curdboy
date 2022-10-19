package cbc

import(
	"log"

	ent "github.com/pigfall/ent_utils"

	"github.com/pigfall/react-curdboy/mod/cbc/react/app"
)

func init(){
	log.SetFlags(log.Llongfile | log.LstdFlags)
}



type Generator struct{
	graph *ent.Graph
	entSchemaPath string
	outputPath string
}

func (g *Generator) Generate() error{
	factory := app.Factory{}
	return factory.SimpleAppGenerator(g.graph.GetNodes()).Generate(g.outputPath)
}
