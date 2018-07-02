package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"fmt"
	//"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"wunan/fabricportal/fabric"
)

func main() {
	router:=gin.Default()
	router.Static("/static", "./static")
	router.LoadHTMLGlob("templates/*")
	router.GET("/login.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
		})
	})
	router.GET("/sign_up.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "sign_up.html", gin.H{
		})
	})
	router.GET("/DASHBOARD.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "DASHBOARD.html", gin.H{
		})
	})
	router.GET("/network.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "network.html", gin.H{
		})
	})
	router.GET("/blocks.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "blocks.html", gin.H{
		})
	})
	router.GET("/transactions.html", func(c *gin.Context) {
		c.HTML(http.StatusOK, "transactions.html", gin.H{
		})
	})
	router.GET("/user/name/:username/:password", getuser)
	router.GET("/user/register/query_name:username", registeruser)
	router.GET("/user/register/do_register:username/:pasword", registeruserpassword)
	router.Run(":8080")
} 
func getuser(c *gin.Context) {
	name := c.Param("username")
	password:= c.Param("password") 
	fmt.Println("name:", name, "pass:", password)
	result := fabric.Fabric(name, password)
	c.String(http.StatusOK, result)
}

func registeruser(c *gin.Context) {
	rename := c.Param("username")
	password:= c.Param("password") 
	fmt.Println("name:", rename, "pass:", password)
	result := fabric.Register(rename, password)
	fmt.Println(result)
	c.String(http.StatusOK, result)
}

func registeruserpassword(c *gin.Context) {
	rename := c.Param("username")
	password:= c.Param("password") 
	fmt.Println("name:", rename, "pass:", password)
	result := fabric.Register(rename, password)
	fmt.Println(result)
	c.String(http.StatusOK, result)
}



