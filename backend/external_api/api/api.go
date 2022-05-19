package api

import (
	"external_api/db"

	"github.com/gin-gonic/gin"
)

func GroupApi(r *gin.Engine, db *db.DbConfig) {
	api := r.Group("/api")
	user_endpoints(api, db)
}
