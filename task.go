package task

import (
	"fmt"
	"errors"
)

type Job struct {
	ID int
	jobTime int
	childJobIDs []int
}

func getJobByID(id int, jobs []Job) (Job, error){
	for i:=0; i<len(jobs); i++ {
		if jobs[i].ID == id {
			return jobs[i], nil
		}
	}
	return Job{}, errors.New("No such ID")
}

func countDuration(id int, jobs []Job) (int, error) {
	duration := 0

	job, err := getJobByID(id, jobs)
	if err != nil {
		return 0, err
	}

	for j := 0;  j<len(job.childJobIDs); j++ {
		currentDuration, err := countDuration(job.childJobIDs[j], jobs)
		if err != nil {
			fmt.Println(err)
			return 0, err
		}
		duration += currentDuration
	}
	return duration + job.jobTime, nil
}