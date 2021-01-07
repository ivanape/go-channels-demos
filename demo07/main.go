package main

import (
	"encoding/json"
	"net/http"
)

var jobsManager *JobsManager

func startTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	job, _ := jobsManager.AddJob(id)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(job)
}

func stopTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	job, _ := jobsManager.RemoveJob(id)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(job)
}

func status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(jobsManager.GetJobs())
}

func main() {
	jobsManager = NewJobManager()
	jobsManager.StartManager()

	http.HandleFunc("/start", startTask)
	http.HandleFunc("/stop", stopTask)
	http.HandleFunc("/status", status)
	_ = http.ListenAndServe(":8000", nil)
}
