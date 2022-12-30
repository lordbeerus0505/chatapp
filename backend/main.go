package main

import (
	"backend/src"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	cors "github.com/rs/cors/wrapper/gin"
)

func main() {
	r := gin.Default()
	r.Use(cors.Default())
	src.Login(r)
	src.Register(r)
	src.AddContact(r)
	src.GetChats(r)
	r.Run()

}
