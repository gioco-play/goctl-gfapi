package errorx

type Err struct {
	status  int
	message string
}

func New(status int, msgs ...string) error {
	e := &Err{
		status: status,
	}
	if len(msgs) > 0 {
		e.message = msgs[0]
	}
	return e
}

func (e *Err) Error() string {
	return e.message
}

func (e *Err) GetStatus() int {
	return e.status
}
