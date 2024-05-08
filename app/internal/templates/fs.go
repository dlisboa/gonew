package templates

import "embed"

//go:embed layout/* pages/* partial/*
var FS embed.FS
