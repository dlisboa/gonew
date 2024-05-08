package static

import "embed"

//go:embed css/* js/* img/*
var FS embed.FS
