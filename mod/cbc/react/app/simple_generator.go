package app

import (
	"log"
	"os"
	"path/filepath"
	tpl "text/template"

	ent "github.com/pigfall/ent_utils"

	onepage "github.com/pigfall/react-curdboy/mod/cbc/react/onepage"
)

type SimpleAppGenerator struct {
	nodes []*ent.Type
}

func (g *SimpleAppGenerator) Generate(outputDir string) error {
	if err := g.generateDependencies(outputDir); err != nil {
		return err
	}
	tplIns, err := tpl.New("simple_app.tmpl").ParseFS(templates, "tpls/simple_app.tmpl")
	if err != nil {
		log.Println(err)
		return err
	}

	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		log.Println(err)
		return err
	}

	outputFile, err := os.Create(filepath.Join(outputDir, "CDBSimpleApp.js"))
	if err != nil {
		log.Println(err)
		return err
	}
	defer outputFile.Close()

	if err := tplIns.Execute(outputFile, g); err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (g *SimpleAppGenerator) generateDependencies(outputDir string) error {
	factory := onepage.Factory{}
	if err := factory.SchemasViewerGenerator(g.nodes).Generate(outputDir); err != nil {
		return err
	}

	if err := factory.EntitiesViewerGenerator(g.nodes).Generate(outputDir); err != nil {
		return err
	}

	return nil
}
