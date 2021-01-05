package main

import (
	"log"
	"net/http"
	"sync"
	"time"
)

type jobsManager struct {
	m sync.Mutex

	jobList      map[string]job
	jobChannel   chan *job
	doneChannel  chan *job
	cancelChanel chan *job
}

type job struct {
	id     string
	cancel chan struct{}
	done   chan struct{}
	err    error
}

var jobs = jobsManager{
	jobList:      make(map[string]job),
	jobChannel:   make(chan *job),
	doneChannel:  make(chan *job),
	cancelChanel: make(chan *job),
}

func startTask(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")

	jobs.m.Lock()
	defer jobs.m.Unlock()
	jobs.jobList[id] = job{
		id:     id,
		cancel: make(chan struct{}),
		done:   make(chan struct{}),
	}

	newJob, _ := jobs.jobList[id]
	jobs.jobChannel <- &newJob
	log.Printf("Request with id %s has ended\n", id)
}

func stopTask(w http.ResponseWriter, r *http.Request) {

}

func status(w http.ResponseWriter, r *http.Request) {

}

func process(jobChannel chan *job) {
	for {
		newJob := <-jobChannel

		startJob(newJob)
	}

}

func startJob(newJob *job) {
	log.Printf("Starting job %s\n", newJob.id)
	time.Sleep(5 * time.Second)
	log.Printf("Ending job %s\n", newJob.id)
	jobs.doneChannel <- newJob
}

func doneJobs(doneChannel chan *job) {
	for {
		doneJob := <-doneChannel

		log.Printf("Job %s is done\n", doneJob.id)
	}
}

func main() {
	go process(jobs.jobChannel)
	go doneJobs(jobs.doneChannel)

	http.HandleFunc("/start", startTask)
	http.HandleFunc("/stop", stopTask)
	http.HandleFunc("/status", status)
	_ = http.ListenAndServe(":8000", nil)
}
