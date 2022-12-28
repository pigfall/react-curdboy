package app

import (
	"embed"
)

//go:embed tpls/*
var templates embed.FS
