package cbc

import(
	"fmt"
	"path/filepath"
	"os"
	"strings"
	tpl "text/template"

	ent "github.com/pigfall/ent_utils"
)

type ModelGenerator struct{
	generator *Generator
}


func (g *ModelGenerator) Generate() error{
	for _,node := range g.generator.graph.GetNodes(){
		if err := g.generateModel(node);err!=nil{
			return err
		}
	}

	return nil
}


func (g *ModelGenerator) generateModel(node *ent.Type) error{
	tplIns,err := tpl.New("model.tmpl").ParseFS(templates,"tpls/model.tmpl")
	if err != nil{
		return err
	}
	if err :=os.MkdirAll(g.outputBasePath(),os.ModePerm);err != nil{
		return err
	}

	file,err := os.Create(filepath.Join(g.outputBasePath(),fmt.Sprintf("%s.js",strings.ToLower(node.Name()))))
	if err != nil{
		return err
	}
	defer file.Close()

	return tplIns.Execute(file,node)
}

func (g *ModelGenerator) outputBasePath() string{
		return filepath.Join(g.generator.outputPath,"model")
}
