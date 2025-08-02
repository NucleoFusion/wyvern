package main

import (
	"wyvern-server/internal/dbctx"
	"wyvern-server/internal/dbctx/mgrctx"
	"wyvern-server/internal/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	dbctx.SetCtx()
	mgrctx.SetCtx(dbctx.GetCtx().Pg)

	r := gin.Default()

	// r.Use(middleware.Inject(ctx)) // Injecting AppContext
	routes.AddRoutes(r)

	r.Run(":3000")
}
