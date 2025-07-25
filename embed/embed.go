package embed

import "embed"

//go:embed templates/*.tmpl templates/**/*
var Templates embed.FS

