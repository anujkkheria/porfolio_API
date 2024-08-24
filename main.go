package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/portfolio_API/common"
	blogs "github.com/portfolio_API/routers/blogs"
)



func main(){

	err := Server()

	if err != nil{
		panic(err)
	}

}

func Server() error{
	 err := common.LoadEnv()

		if err != nil {
			return err
		}

		err = common.InitDB()

		if err != nil {
			return err 
		}

		defer common.CloseDB()

		app := fiber.New()

	app.Use(logger.New())
	app.Use(cors.New())
	app.Use(recover.New())



	// v1 := app.Group("v1")
 // blogs := v1.Group("blogs")

	// blogs.Get("/:id", )
	// blogs.Post("/create", CreateBlogs())
	
	blogs.Blogs(app)


	app.Get("/", func(c *fiber.Ctx)error{
		fmt.Print("hi server is running")
		return c.SendString("hellow world !!")
	})



	app.Listen(":3000")
	return nil
}