package main

import (
	"embed"

	"github.com/neel4os/warg/cmd"
	"github.com/neel4os/warg/internal/common/util"
	"github.com/spf13/cobra"
)

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen -config cfg.yaml ./openapi/api.yaml

//go:embed console/.output/public/*
var staticFiles embed.FS

func main() {
	util.NewStaticFileLocation(&staticFiles)
	command := cmd.New()
	cobra.CheckErr(command.Execute())
}
