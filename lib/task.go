package lib

var Description string

type Status int

const (
	CLOSE Status = iota
	OPEN
)

type Task struct {
	Id          int
	Description string
	Status      Status
}

func CreateTask(description string, status Status) Task {
	return Task{Description: description, Status: status}
}

func (tk Task) GetStatusString() string {
	switch tk.Status {
	case CLOSE:
		return "CLOSE"
	default:
		return "OPEN"
	}
}

func GetStatusFromString(status string) Status {
	switch status {
	case "OPEN":
		return OPEN
	default:
		return CLOSE
	}
}
