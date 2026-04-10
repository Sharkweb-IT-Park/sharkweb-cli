package generator

import "embed"

//go:embed templates/module/full/**
var templateFS embed.FS
