package handlers

import (
	"net/http"

	"log"

	"github.com/CloudyKit/jet/v6"

	"github.com/gorilla/websocket"
)

var views = jet.NewSet(
	jet.NewOSFileSystemLoader("../../html"),
	jet.InDevelopmentMode(),
)

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func Home(w http.ResponseWriter, r *http.Request) {

	RenderPage(w, "home.jet", nil)

}

// defines reponse sent back from websocket

type WsJsonResponse struct {
	Action  string `json:"action"`
	Message string `json:"message"`
	MsgType string `json:"msg_type"`
}

type WsPayload struct {
	Action   string `json:"action"`
	Message  string `json:"message"`
	Username string `json:"username"`
}

func WsEndPoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)

	}

	var response WsJsonResponse
	response.Message = `<em><small>Connected to server </small></em>`

	ws.WriteJSON(response)

}
func RenderPage(w http.ResponseWriter, tmpl string, data jet.VarMap) error {
	view, err := views.GetTemplate(tmpl)
	if err != nil {
		log.Println("no such template:", err)
		return err
	}

	if err = view.Execute(w, data, nil); err != nil {
		log.Println("cant execute template", err)
		return err
	}
	return nil

}
