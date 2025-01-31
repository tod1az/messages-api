package main

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type Message struct {
	Body string `json:"body"`
	Id   string `json:"id"`
}

var messages map[string]Message
var users []string

func main() {
	messages = make(map[string]Message)
	e := echo.New()
	e.Use(middleware.CORS())
	e.GET("/messages", homeHandler)
	e.POST("/messages", postMessage)
	e.GET("/all", getAllMessages)
	e.Logger.Fatal(e.Start(":3000"))
}

func homeHandler(c echo.Context) error {
	id := c.QueryParam("id")
	message, err := getRandomMessage(id)
	if err != nil {
		message.Id = "fakeid"
		message.Body = "this should be an ai message hehe"
	}
	return c.JSON(http.StatusOK, message)

}

func postMessage(c echo.Context) error {
	var newMessage Message
	err := json.NewDecoder(c.Request().Body).Decode(&newMessage)
	if err != nil {
		return err
	}
	users = append(users, newMessage.Id)
	messages[newMessage.Id] = newMessage

	return c.JSON(http.StatusOK, newMessage)
}

func getRandomMessage(currentId string) (*Message, error) {
	for id, message := range messages {
		if id != currentId {
			return &message, nil
		}
	}
	return &Message{}, errors.New("Error, there are no messages registered")
}
func getAllMessages(c echo.Context) error {
	var allMessages []Message
	for _, value := range messages {
		allMessages = append(allMessages, value)
	}
	return c.JSON(http.StatusOK, allMessages)
}
