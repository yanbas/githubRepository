package App

import (
	"net/http"

	"github.com/go-acme/lego/log"
)

func (a *App) Initialize() {

	log.Println(http.ListenAndServe(":8085", http.HandlerFunc(a.getData)))

}
