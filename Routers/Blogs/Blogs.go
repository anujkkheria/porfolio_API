package blogs

import (
	"github.com/gofiber/fiber/v2"
	blogsHandler "github.com/portfolio_API/handlers/blogs"
)

func DemoHandler(c *fiber.Ctx) error{
	return c.Status(200).SendString("you visited blogs")
}

func Blogs(v *fiber.App){
	blogs := v.Group("/blogs")
	blogs.Get("/", DemoHandler)
	blogs.Get("/getall", blogsHandler.GetBlogs)
	blogs.Post("/createBlog", blogsHandler.CreateBlogs)
	blogs.Patch("/:id", blogsHandler.UpdateBlog)
}

