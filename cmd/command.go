package cmd

type Executable interface {
	Execute()
}

type UICommand struct {
	Executable
}

type BackendCommand struct {
	Executable
}
