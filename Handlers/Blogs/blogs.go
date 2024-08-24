package blogs

import (
	"github.com/gofiber/fiber/v2"
	"github.com/portfolio_API/common"
)

	type Blog struct{
		Title string `json:"title" bson:"title"`
		Type string `json:"type" bson:"type"`
		Link string `json:"link" bson:"link"`
	}

	func CreateBlogs(c *fiber.Ctx) error{

	body := new(Blog)
	if err := c.BodyParser(body); err != nil{
		return c.Status(400).JSON(fiber.Map{
			"error":"invalid Body",
		})
	}

	coll := common.GetDBCollection("blogs")

	result,err := coll.InsertOne(c.Context(), body)

	if err != nil{
		return c.Status(500).JSON(fiber.Map{
			"message" :"failed to create book",
			"error": err.Error(),
		})
	}

return c.Status(fiber.StatusCreated).JSON(fiber.Map{
	"result":result,
})
}