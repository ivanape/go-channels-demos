package main

import (
	"encoding/json"
	"net/http"
)

var jobsManager *JobsManager

func startTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	job := NewJob(id)
	_ = job.Do(run, job)
	job, _ = jobsManager.RunJob(job)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(job)
}

func stopTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	job, _ := jobsManager.StopJob(id)

	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(job)
}

func status(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(jobsManager.GetJobs())
}

func main2() {
	jobsManager = NewJobManager()
	jobsManager.StartManager()

	http.HandleFunc("/start", startTask)
	http.HandleFunc("/stop", stopTask)
	http.HandleFunc("/status", status)
	_ = http.ListenAndServe(":8000", nil)
}

func main() {
	jobsManager = NewJobManager()
	jobsManager.StartManager()

	/*job := NewJob("1")
	_ = job.Do(run, job)
	_, _ = jobsManager.RunJob(job)

	<- job.done*/

	job1 := NewJob("1")
	_ = job1.Do(run, job1)
	job2 := NewJob("2")
	_ = job2.Do(run, job2)
	_ = jobsManager.RunJobsInSequence(job1, job2)

}
