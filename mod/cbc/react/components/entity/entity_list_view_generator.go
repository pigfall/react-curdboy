package entity

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	tpl "text/template"

	ent "github.com/pigfall/ent_utils"

	"github.com/pigfall/react-curdboy/mod/cbc/react/components"
)

// Generate a component which will used to show the list of entities
// Entity is the instance of a schema. Actually is a data record in database
type EntityListViewGenerator struct {
	EntNode *ent.Type
}

func (g *EntityListViewGenerator) Generate(outputDirBase string) error {
	tplIns, err := tpl.New("entity_list_view.tmpl").ParseFS(templates, "tpls/entity_list_view.tmpl")
	if err != nil {
		log.Println(err)
		return err
	}

	outputDir := filepath.Join(outputDirBase, components.BaseDir, "entities")
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		log.Println(err)
		return err
	}

	outputFile, err := os.Create(filepath.Join(outputDir, fmt.Sprintf("%sList.js", g.EntNode.Name())))
	if err != nil {
		log.Println(err)
		return err
	}

	if err := tplIns.Execute(outputFile, g); err != nil {
		log.Println(err)
		return err
	}

	return nil
}
