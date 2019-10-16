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

func countDuration(id int, jobs []Job) (int, error) {
	duration := 0

	for i := 0; i<len(jobs); i++ {
		if jobs[i].ID == id {
			for j := 0;  j<len(jobs[i].childJobIDs); j++ {
				dur, err := countDuration(jobs[i].childJobIDs[j], jobs)
				if err != nil {
					fmt.Println(err)
					return 0, err
				}
				duration += dur
			}
			return duration + jobs[i].jobTime, nil
		}
	}
	return 0, errors.New("No such ID")
}