package client

import(
	tpl "text/template"
	"path/filepath"
	"os"
	"log"

	ent "github.com/pigfall/ent_utils"
)

type ClientGenerator struct{
	Nodes []*ent.Type
}

func (g *ClientGenerator) Generate(outputDir string)error{
	tplIns, err := tpl.New("client.go.tmpl").ParseFS(templates, "tpls/client.go.tmpl")
	if err != nil {
		log.Println(err)
		return err
	}

	outputDir = filepath.Join(outputDir, "client")
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		log.Println(err)
		return err
	}
	outputFile, err := os.Create(filepath.Join(outputDir, "Client.js"))
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
