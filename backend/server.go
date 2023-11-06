package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"packages/database"
	// "go.mongodb.org/mongo-driver/bson"
	"encoding/base64"
	"errors"
	"fmt"
	"io/ioutil"
	"strings"
	"github.com/google/uuid"
	"time"

)

var mimeToExtension = map[string]string{
	"image/png":  "png",
	"image/jpeg": "jpg",
	"image/gif":  "gif",
	"image/webp": "webp",
}


func saveDataURI(dataURI, outputDir string) (string, error) {
	// Split the data URI into its components.
	parts := strings.Split(dataURI, ",")
	if len(parts) != 2 {
		return "", errors.New("invalid data URI format")
	}

	// Extract MIME type from the data URI prefix.
	mimeType := strings.TrimPrefix(strings.Split(parts[0], ";")[0], "data:")

	// Map the MIME type to a file extension.
	extension, ok := mimeToExtension[mimeType]
	if !ok {
		return "", fmt.Errorf("unsupported MIME type: %s", mimeType)
	}

	// Decode the base64 data.
	decodedData, err := base64.StdEncoding.DecodeString(parts[1])
	if err != nil {
		return "", err
	}

	// Generate a UUID
	uuid := uuid.New()

	// Get the current timestamp in nanoseconds
	timestamp := time.Now().UnixNano()
 
	// Combine the UUID and timestamp to create a unique ID
	uniqueID := fmt.Sprintf("%s-%d", uuid, timestamp)

	filename := fmt.Sprintf("%s.%s", uniqueID,extension)

	// Save the decoded data to the desired directory.
	outputPath := fmt.Sprintf("%s/%s", outputDir, filename)
	err = ioutil.WriteFile(outputPath, decodedData, 0644)
	if err != nil {
		return "", err
	}

	return outputPath, nil
}



type Post struct {
    Title string `json:"title" bson:"title"`
	Description string `json:"description" bson:"description"`
	AttachedFilesBase64 []string `json:"attachedFiles" bson:"title"`
	Time int64 `json:"time" bson:"clientTime"`
}

func main() {
	database.InitializeDatabase("mongodb://localhost:27017", "production", "posts")
	app := fiber.New(fiber.Config{
		BodyLimit: 1024 * 1024 * 1024,
	})
	app.Use(logger.New())
	app.Use(cors.New())
	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Post("/api/v1/post", func(c *fiber.Ctx) error {
		post := new(Post)
		if err := c.BodyParser(post); err != nil {
            return err
        }
		
		for _, each := range post.AttachedFilesBase64 {
			_, err := saveDataURI(each, "files")
			fmt.Println(err)
		}

		return nil
	})

	app.Get("/api/v1/posts", func(c *fiber.Ctx) error {
		var posts []Post

		posts = append(posts, Post{Title: "Hello"})
		fmt.Println(posts)
		body := fiber.Map{
			"data": posts,
		}
		fmt.Println(body)
		return c.JSON(body)
	})


	app.Listen(":3001")
}
