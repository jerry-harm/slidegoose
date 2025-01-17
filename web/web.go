package web

import "embed"

//go:generate pnpm build --base=/web
//go:embed dist
var WebFS embed.FS
