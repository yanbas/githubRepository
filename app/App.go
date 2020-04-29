package App

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/go-acme/lego/v3/log"
)

type App struct {
	ResponseData []ResponseGithub
}

func (a *App) Initialize() {
	log.Println(http.ListenAndServe(":8085", http.HandlerFunc(a.getData)))
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

	// Call REST Client for github API
	err := a.CallHTTPServiceGithub(param[7:])
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	response.Message = "Done"
	response.Success = true
	response.List = a.ResponseData

	data, err := json.Marshal(response)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Write(data)

}

func (a *App) CallHTTPServiceGithub(user string) error {

	client := &http.Client{}

	url := fmt.Sprintf("https://api.github.com/users/%s/repos", user)

	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Accept", "application/vnd.github.nebula-preview+json")
	resp, err := client.Do(req)

	if err != nil {
		log.Println("Errored when sending request to the server")
		return err
	}

	defer resp.Body.Close()
	resp_body, _ := ioutil.ReadAll(resp.Body)

	log.Println(fmt.Sprintf("Response Status : %s", resp.Status))
	log.Println(fmt.Sprintf("Response Body : %s", resp_body))

	err = json.Unmarshal(resp_body, &a.ResponseData)
	if err != nil {
		return err
	}

	return nil
}
