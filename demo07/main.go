package main

import (
	"log"
	"net/http"
)

var jobsManager *JobsManager

func startTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	_, _ = jobsManager.AddJob(id)

	log.Printf("Request with id %s has ended\n", id)
}

func stopTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	_, _ = jobsManager.RemoveJob(id)
}

func status(w http.ResponseWriter, r *http.Request) {
	jobsManager.GetJobs()
}

func main() {
	jobsManager = NewJobManager()
	jobsManager.StartManager()

	http.HandleFunc("/start", startTask)
	http.HandleFunc("/stop", stopTask)
	http.HandleFunc("/status", status)
	_ = http.ListenAndServe(":8000", nil)
}
