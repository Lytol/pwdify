package pwdify

type Engine struct {
	Files    []string
	Password string
}

func New(files []string, password string) *Engine {
	return &Engine{
		Files:    files,
		Password: password,
	}
}

func (e *Engine) Run() error {
	// TODO
	return nil
}

func (e *Engine) Status() map[string]bool {
	status := make(map[string]bool)

	// TODO

	return status
}

func (e *Engine) Completed() bool {
	// TODO
	return false
}
