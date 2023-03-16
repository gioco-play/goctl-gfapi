package errorx

type Err struct {
	Status  int
	Message string
}

func New(status int, msgs ...string) error {
	e := &Err{
		Status: status,
	}
	if len(msgs) > 0 {
		e.Message = msgs[0]
	}
	return e
}

func NewStatic(err *Err) error {
	return err
}

func (e *Err) Error() string {
	return e.Message
}

func (e *Err) GetStatus() int {
	return e.Status
}
