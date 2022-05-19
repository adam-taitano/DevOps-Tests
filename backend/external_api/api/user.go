package api

import (
	"external_api/db"
	"external_api/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

func user_endpoints(rg *gin.RouterGroup, db *db.DbConfig) {
	rg.GET("/users", func(c *gin.Context) {
		get_users(c, db)
	})

	rg.GET("/user/:id", func(c *gin.Context) {
		if id, e := strconv.Atoi(c.Param("id")); e == nil {
			get_user(c, db, id)
		} else {
			fmt.Printf("[XAPI WARN] Query Error: %s", e)
		}
	})
}

func get_users(c *gin.Context, db *db.DbConfig) {
	rows := db.Query(`SELECT * FROM users`)
	fmt.Println(rows)

	var users []models.User
	for rows.Next() {
		var user models.User
		rows.Scan(&user.Id, &user.Username, &user.Password)
		users = append(users, user)
	}
	defer rows.Close()
	c.JSON(200, gin.H{
		"users": users,
	})
}

func get_user(c *gin.Context, db *db.DbConfig, id int) {
	var user models.User
	err := db.QueryOne("SELECT * FROM users WHERE id=$1", id).Scan(&user.Id, &user.Username, &user.Password)
	if err == nil {
		c.JSON(200, gin.H{
			"user": user,
		})
	} else {
		c.JSON(400, gin.H{
			"error": "user not found",
		})
	}
}
