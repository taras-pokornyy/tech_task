package task

import "testing"


func TestCountDurationValid(t *testing.T) {

	var jobs = []Job{
		Job{1, 30, []int{2, 4}},
		Job{2, 10, []int{3}},
		Job{4, 60, []int{}},
		Job{3, 20, []int{}},
	}

    duration, _ := countDuration(1, jobs)
    if duration != 120 {
		t.Errorf("Duration is incorrect. Expected %v, Actual: %v.", duration, 120)
    }
}

func TestCountDurationInValid(t *testing.T) {

	var jobs = []Job{
		Job{1, 30, []int{2, 4}},
		Job{2, 10, []int{3}},
		Job{4, 60, []int{}},
		Job{3, 20, []int{}},
	}

    _, err := countDuration(5, jobs)
    if err == nil {
		t.Errorf("Error is incorrect. Expected %v, Actual: %v.", "No such ID", err)
    }
}

func TestCountDurationInValidRecursive(t *testing.T) {

	var jobs = []Job{
		Job{1, 30, []int{2, 4, 6}},
		Job{2, 10, []int{3}},
		Job{4, 60, []int{}},
		Job{3, 20, []int{}},
	}

    _, err := countDuration(1, jobs)
    if err == nil {
		t.Errorf("Error is incorrect. Expected %v, Actual: %v.", "No such ID", err)

    }
}


func TestGetJobByIDValid(t *testing.T) {
	var jobs = []Job{
		Job{1, 30, []int{2, 4}},
		Job{2, 10, []int{3}},
		Job{4, 60, []int{}},
		Job{3, 20, []int{}},
	}

	job, _ := getJobByID(1, jobs)
	if job.ID != jobs[0].ID {
		t.Errorf("Job is incorrect. Expected %v, Actual: %v.", jobs[0] job)
	}
}

func TestGetJobByIDInValid(t *testing.T) {
	var jobs = []Job{
		Job{1, 30, []int{2, 4}},
		Job{2, 10, []int{3}},
		Job{4, 60, []int{}},
		Job{3, 20, []int{}},
	}

	_, err := getJobByID(5, jobs)
	if err == nil {
		t.Errorf("Error is incorrect. Expected %v, Actual: %v.", "No such ID", err)
	}
}
