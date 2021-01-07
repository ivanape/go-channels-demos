package main

import (
	"encoding/json"
	"fmt"
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

	job := NewJob("1")
	_ = job.Do(run, job)
	_, _ = jobsManager.RunJob(job)
	job.Wait()

	job1 := NewJob("3")
	_ = job1.Do(run, job1)
	job2 := NewJob("4")
	_ = job2.Do(run, job2)
	_ = jobsManager.RunJobsInSequence(job1, job2)

	job3 := NewJob("5")
	_ = job3.Do(run, job3)
	job4 := NewJob("hello")
	_ = job4.Do(task1, "ivan")
	_ = jobsManager.RunJobsInParallel(job3, job4)

	fmt.Printf("%v\n", jobsManager.GetJobs())
}
