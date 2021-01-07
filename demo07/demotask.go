package main

import (
	"log"
	"time"
)

func run(job *Job) {
	log.Printf("Starting job %s\n", job.id)

	executeStepIfValid(job, step1, job.id)
	executeStepIfValid(job, step2, job.id)
	executeStepIfValid(job, step3, job.id)

	cleanup(job.id)
}

func executeStepIfValid(job *Job, f func(p string), p string) {
	if job.result.err == errCancelled {
		log.Printf("Job %s has been cancelled. Bypass method execution.", job.id)
		return
	}

	f(p)
}

func step1(jobId string) {
	log.Printf("Starting step1 for Job %s\n", jobId)
	time.Sleep(10 * time.Second)
	log.Println("Ending step1")
}

func step2(jobId string) {
	log.Printf("Starting step2 for Job %s\n", jobId)
	time.Sleep(10 * time.Second)
	log.Println("Ending step2")
}

func step3(jobId string) {
	log.Printf("Starting step3 for Job %s\n", jobId)
	time.Sleep(10 * time.Second)
	log.Println("Ending step3")
}

func cleanup(jobId string) {
	log.Printf("Execute cleanup for Job %s\n", jobId)
}
