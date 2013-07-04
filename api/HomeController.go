package api

import (
	. "github.com/pjvds/httpcallback.io/mvc"
	"net/http"
	"time"
)

type HomeController struct {
	StartTime time.Time
}

func NewHomeController() *HomeController {
	return &HomeController{
		StartTime: time.Now(),
	}
}

func (c *HomeController) HandleIndex(request *http.Request) ActionResult {
	return JsonResult(&JsonDocument{
		"message": "welcome!",
		"uptime":  time.Now().Sub(c.StartTime).String(),
	})
}

type PingResponse struct {
	Message string `json:"message"`
}

func (c *HomeController) HandlePing(req *http.Request) ActionResult {
	return JsonResult(&PingResponse{
		Message: "pong",
	})
}
