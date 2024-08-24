package blogs

import (
	"github.com/gofiber/fiber/v2"
	blogsHandler "github.com/portfolio_API/Handlers/blogs"
)

func Blogs(v *fiber.App){
	blogs := v.Group("/blogs")
	blogs.Get("/", func (c *fiber.Ctx) error{
		return c.Status(200).SendString("you visited blogs")
	})
	blogs.Post("/", blogsHandler.CreateBlogs)
	blogs.Get("/:id")
	blogs.Patch("/:id")
}

