package main

import (
	"errors"
	"log"
	"sync"
)

var (
	errCancelled = errors.New("jobCancelled")
)

type JobsManager struct {
	m             sync.Mutex
	jobList       map[string]*Job
	workerChannel chan *Job
	doneChannel   chan *Job
	cancelChannel chan *Job
	workerSize    int
}

func NewJobManager() *JobsManager {
	return &JobsManager{
		jobList:       make(map[string]*Job),
		workerChannel: make(chan *Job),
		doneChannel:   make(chan *Job),
		cancelChannel: make(chan *Job),
		workerSize:    100, //By default allow 100 concurrent tasks
	}
}

func (j *JobsManager) StartManager() {
	for i := 0; i < j.workerSize; i++ {
		go j.registerWorker()
	}
}

func (j *JobsManager) RunJob(job *Job) (*Job, error) {
	j.m.Lock()
	defer j.m.Unlock()

	j.jobList[job.Id] = job
	j.workerChannel <- job

	return job, nil
}

func (j *JobsManager) RunJobsInSequence(jobs ...*Job) error {
	j.m.Lock()
	for _, job := range jobs {
		j.jobList[job.Id] = job
	}
	j.m.Unlock()

	for _, job := range jobs {
		j.workerChannel <- job
		<-job.done
	}

	return nil
}

func (j *JobsManager) RunJobsInParallel(jobs ...*Job) (chan struct{}, error) {
	j.m.Lock()
	defer j.m.Unlock()

	for _, job := range jobs {
		j.jobList[job.Id] = job
		j.workerChannel <- job
	}

	wait := make(chan struct{})

	return wait, nil
}

func (j *JobsManager) StopJob(id string) (*Job, error) {
	j.m.Lock()
	defer jobsManager.m.Unlock()
	job := j.jobList[id]

	j.cancelChannel <- job

	return job, nil
}

func (j *JobsManager) GetJobs() map[string]*Job {
	log.Printf("%v", j.jobList)
	return j.jobList
}

func (j *JobsManager) registerWorker() {
	for {
		select {
		case job := <-j.workerChannel:
			job.Status = Running
			_, _ = job.Run()

			if job.result.err != errCancelled {
				jobsManager.doneChannel <- job
				close(job.done)
			}

		case job := <-j.doneChannel:
			job.Status = Done
			log.Printf("Job %s is done\n", job.Id)

		case job := <-j.cancelChannel:
			job.Status = Cancelled
			job.result = JobResult{
				err: errCancelled,
			}
			close(job.done)
			log.Printf("Job %s is cancelled\n", job.Id)
		}
	}
}
