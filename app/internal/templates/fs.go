package templates

import "embed"

//go:embed layout/* page/* partial/* error/*
var FS embed.FS
