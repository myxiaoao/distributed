package grades

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func RegisterHandlers() {
	handler := new(studentsHandler)
	http.Handle("/students", handler)
	http.Handle("/students/", handler)
}

type studentsHandler struct{}

func (s studentsHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	pathSegments := strings.Split(r.URL.Path, "/")
	switch len(pathSegments) {
	case 2:
		s.getAll(w, r)
	case 3:
		id, err := strconv.Atoi(pathSegments[2])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		s.getOne(w, r, id)
	case 4:
		id, err := strconv.Atoi(pathSegments[2])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		s.addGrade(w, r, id)
	default:
		w.WriteHeader(http.StatusNotFound)
	}
}

func (s studentsHandler) getAll(w http.ResponseWriter, r *http.Request) {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()

	data, err := s.toJSON(students)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (s studentsHandler) getOne(w http.ResponseWriter, r *http.Request, id int) {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()

	student, err := students.GetById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println(err)
		return
	}

	data, err := s.toJSON(student)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (s studentsHandler) addGrade(w http.ResponseWriter, r *http.Request, id int) {
	studentsMutex.Lock()
	defer studentsMutex.Unlock()

	student, err := students.GetById(id)
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		log.Println(err)
		return
	}

	var g Grade
	dec := json.NewDecoder(r.Body)
	err = dec.Decode(&g)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	student.Grades = append(student.Grades, g)
	w.WriteHeader(http.StatusCreated)

	data, err := s.toJSON(g)
	if err != nil {
		log.Println(err)
		return
	}
	w.Header().Add("Content-Type", "application/json")
	_, err = w.Write(data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
}

func (s studentsHandler) toJSON(obj interface{}) ([]byte, error) {
	var b bytes.Buffer
	enc := json.NewEncoder(&b)
	err := enc.Encode(obj)
	if err != nil {
		return nil, fmt.Errorf("faild to serialize students: %q", err)
	}
	return b.Bytes(), nil
}
