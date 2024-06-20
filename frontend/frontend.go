package frontend

import (
	"embed"
	"io/fs"
)

//go:generate npm install
//go:generate npm run build
//go:embed build/*
var embFs embed.FS
var F fs.FS

func init() {
	var err error
	F, err = fs.Sub(embFs, "build")
	if err != nil {
		panic(err)
	}
}
