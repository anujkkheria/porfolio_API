package projects

import (
	"context"
	"fmt"
	"os"

	"github.com/chromedp/chromedp"
	"github.com/gofiber/fiber/v2"
)

type NewProject struct{
	Link string `json:"link" bson:"link"`
	Title string `json:"title" bson:"title"`
	// Image string `json:"image" bson:"image"`
	Category string `json:"category" bson:"category"`
}

func CreateNewProject(c *fiber.Ctx)error{

	 project := new(NewProject)
		if err := c.BodyParser(project); err != nil{
			return c.Status(400).JSON(fiber.Map{
				"message":"invalid Body please check the body structure",
			})
		}
	err := getScreenShotfromLink(project.Link, project.Title)
	if err !=nil{
		return c.SendStatus(500)
	}
		return c.SendStatus(201)
}

	func getScreenShotfromLink(url string, title string) error{
		ctx,cancel:=	chromedp.NewContext(context.Background())
		defer cancel()
		
		fmt.Print(ctx,url)
var buf []byte
 
err:= chromedp.Run(ctx, chromedp.Tasks{
	chromedp.Navigate(url),
	chromedp.FullScreenshot(&buf,90),
})

if err !=nil{
	return err
}

err = SaveImage(title, buf)

if err != nil{
	return err
}

return nil

}

func SaveImage(title string, buf []byte)error{
	filename := fmt.Sprintf("%s.%s",title,".png")
	err := os.WriteFile(filename, buf, 0644)
if err != nil{ 
	return err
}
	return nil
} 