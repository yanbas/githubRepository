package App

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-acme/lego/v3/log"
)

type App struct {
}

func (a *App) getData(w http.ResponseWriter, r *http.Request) {

	w.Header().Set("Content-Type", "application/json")
	response := Response{}

	if r.Method != "GET" {
		response.Success = false
		response.Message = "Invalid HTTP Method"

		data, err := json.Marshal(response)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(data)
		return
	}

	//get parameter user in url
	param := r.URL.Path

	if len(param) < 8 {
		response.Success = false
		response.Message = "Invalid param user, please fill in"
		data, err := json.Marshal(response)
		if err != nil {
			log.Println(err.Error())
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		w.Write(data)
		return
	}

	fmt.Println(param[7:])

	a.CallHTTPServiceGithub(param[7:])

}

func (a *App) CallHTTPServiceGithub(user string) {

	client := &http.Client{}

	url := fmt.Sprintf("https://api.github.com/users/%s/repos", user)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/vnd.github.nebula-preview+json")
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Errored when sending request to the server")
		return
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	fmt.Println(resp.Status)
	fmt.Println(string(resp_body))

}
