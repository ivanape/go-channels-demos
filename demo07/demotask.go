package main

import (
	"log"
	"time"
)

var (
	defaultDuration = 3
)

func run(job *Job) {
	log.Printf("Starting job %s\n", job.Id)

	executeStepIfValid(job, step1, job.Id)
	executeStepIfValid(job, step2, job.Id)
	executeStepIfValid(job, step3, job.Id)

	cleanup(job.Id)
}

func executeStepIfValid(job *Job, f func(p string), p string) {
	if job.result.err == errCancelled {
		log.Printf("Job %s has been cancelled. Bypass method execution.", job.Id)
		return
	}

	f(p)
}

func step1(jobId string) {
	log.Printf("Starting step1 for Job %s\n", jobId)
	time.Sleep(time.Duration(defaultDuration) * time.Second)
	log.Println("Ending step1")
}

func step2(jobId string) {
	log.Printf("Starting step2 for Job %s\n", jobId)
	time.Sleep(time.Duration(defaultDuration) * time.Second)
	log.Println("Ending step2")
}

func step3(jobId string) {
	log.Printf("Starting step3 for Job %s\n", jobId)
	time.Sleep(time.Duration(defaultDuration) * time.Second)
	log.Println("Ending step3")
}

func cleanup(jobId string) {
	log.Printf("Execute cleanup for Job %s\n", jobId)
}
