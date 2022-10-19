package schema

import(
	"log"
	"os"
	tpl "text/template"
	"fmt"
	"path/filepath"

	ent "github.com/pigfall/ent_utils"
	iter "github.com/pigfall/goiter"
)


type SchemaViewGenerator struct{
	EntNode *ent.Type
}

func (g *SchemaViewGenerator) Generate(outputDirBase string)error{
	tplIns,err := tpl.New("schema_view.tmpl").ParseFS(templates,"tpls/schema_view.tmpl")
	if err != nil{
		log.Println(err)
		return err
	}

	outputDir :=filepath.Join(outputDirBase,"components","schemas")
	if err:=os.MkdirAll(outputDir,os.ModePerm);err != nil{
		log.Println(err)
		return err
	}

	outputFile,err := os.Create(filepath.Join(outputDir,fmt.Sprintf("%s.js",g.EntNode.Name())))
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

func (g *SchemaViewGenerator) Generated_ComponentName() string{
	return g.EntNode.Name()
}

func GenerateSchemaViews(nodes []*ent.Type,outputDir string)error{
	factory := Factory{}
	return iter.ForEach(
		iter.Slice(nodes),
		func(node *ent.Type)error{
			return factory.SchemaViewGenerator(node).Generate(filepath.Join(outputDir,"components","schema"))
		},
	)
}


