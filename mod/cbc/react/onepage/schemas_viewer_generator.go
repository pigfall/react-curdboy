package onepage

import(
	"path/filepath"
	tpl "text/template"
	"os"
	"log"

	ent "github.com/pigfall/ent_utils"
	iter "github.com/pigfall/goiter"

	schema "github.com/pigfall/react-curdboy/mod/cbc/react/components/schema"
)

type SchemasViewerGenerator struct{
	Nodes []*ent.Type
}


func (g *SchemasViewerGenerator) Generate(outputDir string) error{
	nodes := g.Nodes
	if err := g.generateDependencies(nodes,outputDir);err != nil{
		return err
	}
	tplIns,err := tpl.New("schemas_viewer.tmpl").ParseFS(templates,"tpls/schemas_viewer.tmpl")
	if err != nil{
		log.Println(err)
		return err
	}

	if err := os.MkdirAll(filepath.Join(outputDir,"onepage"),os.ModePerm);err != nil{
		log.Println(err)
		return err
	}
	outputFile,err := os.Create(filepath.Join(outputDir,"onepage","SchemasViewer.js"))
	if err != nil{
		log.Println(err)
		return err
	}
	defer outputFile.Close()

	if err := tplIns.Execute(outputFile,g);err != nil{
		log.Println(err)
		return err
	}

	return nil
}

func (g *SchemasViewerGenerator) generateDependencies(nodes []*ent.Type,outputDir string)error{
	schemaFactory :=(&schema.Factory{})

	if err := iter.ForEach(
		iter.Slice(nodes),
		func(node *ent.Type)error{
				return schemaFactory.SchemaViewGenerator(node).Generate(outputDir)
		},
	);err != nil{
		return err
	}

	return nil
}

func (g *SchemasViewerGenerator) Generated_ComponentName()string{
	return "Schemas"
}
