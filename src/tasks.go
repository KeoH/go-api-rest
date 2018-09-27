package main

import (
	"log"
	"net/http"
	"fmt"
	"encoding/json"
	mgo "gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Task struct {
	ID 			bson.ObjectId 	`bson:"_id" json:"id"`
	Name		string			`bson:"name" json:"name"`
	Description	string			`bson:"description" json:"description"`
}

type TaskDAO struct {
	Server 		string
	Database	string
}

var dao = TaskDAO{}

var db *mgo.Database

const (
	COLLECTION = "tasks"
)

func (m *TaskDAO) Connect() {
	session, err := mgo.Dial(m.Server)
	if err != nil {
		log.Fatal(err)
	}
	db = session.DB(m.Database)
}

func (m *TaskDAO) FindAll() ([]Task, error) {
	var tasks []Task
	err := db.C(COLLECTION).Find(bson.M{}).All(&tasks)
	return tasks, err
}

func (m *TaskDAO) Insert(task Task) error {
	err := db.C(COLLECTION).Insert(&task)
	return err
}

func AllTaskEndPoint(w http.ResponseWriter, r *http.Request){
	fmt.Println("AllTaskEndPoint")	
	
	tasks, err := dao.FindAll()
	if err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusOK, tasks)
}

func GetTaskEndpoint(w http.ResponseWriter, r *http.Request){
	fmt.Println("GetTaskEndpoint")		
	fmt.Fprintln(w, "Not implemented yet!")
}

func CreateTaskEndpoint(w http.ResponseWriter, r *http.Request){
	fmt.Println("CreateTaskEndpoint")

	defer r.Body.Close()
	var task Task
	if err := json.NewDecoder(r.Body).Decode(&task); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	task.ID = bson.NewObjectId()

	if err := dao.Insert(task); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}

	respondWithJson(w, http.StatusCreated, task)
}

func UpdateTaskEndpoint(w http.ResponseWriter, r *http.Request){
	fmt.Println("UpdateTaskEndpoint")		
	fmt.Fprintln(w, "Not implemented yet!")
}

func DeleteTaskEndpoint(w http.ResponseWriter, r *http.Request){
	fmt.Println("DeleteTaskEndpoint")		
	fmt.Fprintln(w, "Not implemented yet!")
}

func respondWithError(w http.ResponseWriter, code int, msg string){
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}){
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func init(){
	dao.Server = "172.17.0.2"
	dao.Database = "go_microservice"
	dao.Connect()
}