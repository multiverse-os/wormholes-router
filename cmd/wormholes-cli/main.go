package main

import (
	"fmt"
)


type Command struct {
	*cmd.Command

	Path string

	Basename string
	Extension string
	Filename string


	ManageSubprocesses bool
	MemoryFD bool
	FD uintptr
	Output []string
	PID int // Will need file

	// Logging system
	// NestProccesses

	StartedAt time.Time
	StoppedAt time.Time
}

func Start() *Command {
	// Start subprocesses, ensure they dont die, store only parent in PID file
	// but keep a session folder in the user home path with all the subprocess
	// PID numbers and any ephemeral data only relating to the session and 
	// immediately useless upon deletion. 
	return &Command{}
}

func (self *Command) Restart() *Command { return self }
func (self *Command) Status() *Command { return self }
func (self *Command) Stop() *Command { return self } 
 

// TODO: These are experimental and possibly just unneccessarily replicating 
//       features that should be provided by the OS in my application.
//func (self *Command) Attach() *Command { return self }
//func (self *Command) Detach() *Command { return self }


func (self *Command) String() string {
	return filepath.Join(self.Path, self.Filename) 
}

func (self *Command) Where() string { return self.String() }

func (self *Command) InExecutableFilePath() bool { return true } 

func (self *Command) AppendToExecutableFilePath() (bool, error) {
	return false, nil
}
