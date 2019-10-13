package controllers

import (
	"broz3r.com/HelloGo_Rest-API/models"
	"broz3r.com/HelloGo_Rest-API/net"
	"net/http"
	"time"
)

// Get Init godoc
// @Summary Initialize App
// @Description Initialize Application with GDPR version and Update Object if app update is needed.
// @ID get-init
// @Tags Init
// @Accept  json
// @Produce  json
// @Param appVersion query string true "App Version"
// @Success 200 {object} models.InitResponse
// @Failure 400
// @Failure 501
// @Router /init [get]
func HelloWorld(w http.ResponseWriter, r *http.Request) {
	resp := models.HelloWorld {
		Value: "Hello World !",
		Date: time.Now().UTC(),
	}
	net.Respond(w, resp)
}
