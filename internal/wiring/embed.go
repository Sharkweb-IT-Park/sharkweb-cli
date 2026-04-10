package wiring

import "embed"

// 🔥 NO wildcard (safer)
//go:embed templates/wiring/backend/**
var WiringTemplates embed.FS
