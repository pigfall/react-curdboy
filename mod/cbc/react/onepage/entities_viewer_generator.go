package onepage

import (
	"log"
	"os"
	"path/filepath"
	tpl "text/template"

	ent "github.com/pigfall/ent_utils"
	iter "github.com/pigfall/goiter"

	entity "github.com/pigfall/react-curdboy/mod/cbc/react/components/entity"
)

type EntitiesViewerGenerator struct {
	Nodes []*ent.Type
}

func (g *EntitiesViewerGenerator) Generate(outputDir string) error {
	if err := g.generateDependencies(outputDir); err != nil {
		return err
	}

	tplIns, err := tpl.New("entities_viewer.tmpl").ParseFS(templates, "tpls/entities_viewer.tmpl")
	if err != nil {
		log.Println(err)
		return err
	}

	if err := os.MkdirAll(filepath.Join(outputDir, "onepage"), os.ModePerm); err != nil {
		log.Println(err)
		return err
	}
	outputFile, err := os.Create(filepath.Join(outputDir, "onepage", "EntitiesViewer.js"))
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

func (g *EntitiesViewerGenerator) generateDependencies(outputDir string) error {
	entityFactory := (&entity.Factory{})

	if err := iter.ForEach(
		iter.Slice(g.Nodes),
		func(node *ent.Type) error {
			if err := entityFactory.NewEntityListViewGenerator(node).Generate(outputDir); err != nil {
				return err
			}

			return nil
		},
	); err != nil {
		return err
	}

	return nil
}
