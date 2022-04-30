package registry

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
)

const ServerPort = ":3000"
const ServersURL = "http://localhost" + ServerPort + "/services"

type registry struct {
	registrations []Registration
	// 互斥锁，多线程操作，并发访问线程安全
	mutex *sync.Mutex
}

func (r *registry) add(registration Registration) error {
	r.mutex.Lock()
	r.registrations = append(r.registrations, registration)
	r.mutex.Unlock()
	return nil
}

func (r *registry) remove(url string) error {
	for i := range reg.registrations {
		if reg.registrations[i].ServiceURL == url {
			r.mutex.Lock()
			reg.registrations = append(reg.registrations[:i], reg.registrations[i+1:]...)
			r.mutex.Unlock()
			return nil
		}
	}
	return fmt.Errorf("service at URL %s not found", url)
}

var reg = registry{
	registrations: make([]Registration, 0),
	mutex:         new(sync.Mutex),
}

type Service struct{}

func (s Service) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	log.Println("Request received.")
	switch r.Method {
	case http.MethodPost:
		dec := json.NewDecoder(r.Body)
		var r Registration
		err := dec.Decode(&r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		log.Printf("Add services: %v with URL: %s\n", r.ServiceName, r.ServiceURL)
		err = reg.add(r)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusBadRequest)
			return
		}
	case http.MethodDelete:
		all, err := ioutil.ReadAll(r.Body)
		if err != nil {
			log.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		url := string(all)
		log.Printf("Removing service at URL: %s", url)
		err = reg.remove(url)
		return
	default:
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
