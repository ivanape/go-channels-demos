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
	jobChannel    chan *Job
	doneChannel   chan *Job
	cancelChannel chan *Job
	workerSize    int
}

func NewJobManager() *JobsManager {
	return &JobsManager{
		jobList:       make(map[string]*Job),
		jobChannel:    make(chan *Job),
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

func (j *JobsManager) AddJob(id string) (*Job, error) {
	j.m.Lock()
	defer j.m.Unlock()
	newJob := &Job{
		id:      id,
		status:  Pending,
		funcs:   make(map[string]interface{}),
		fparams: make(map[string][]interface{}),
	}
	_ = newJob.SetTask(run, newJob)

	j.jobList[id] = newJob
	j.jobChannel <- newJob

	return newJob, nil
}

func (j *JobsManager) RemoveJob(id string) (*Job, error) {
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
		case job := <-j.jobChannel:
			job.status = Running
			_, _ = job.Run()

			if job.result.err != errCancelled {
				jobsManager.doneChannel <- job
			}

		case doneJob := <-j.doneChannel:
			doneJob.status = Done
			log.Printf("Job %s is done\n", doneJob.id)

		case cancelledJob := <-j.cancelChannel:
			cancelledJob.status = Cancelled
			cancelledJob.result = JobResult{
				err: errCancelled,
			}
			log.Printf("Job %s is cancelled\n", cancelledJob.id)
		}
	}
}