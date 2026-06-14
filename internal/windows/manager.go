package windows

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

type WindowManager struct {
	windows map[string]*exec.Cmd
}

var Manager = NewManager()

func NewManager() *WindowManager {
	return &WindowManager{
		windows: map[string]*exec.Cmd{},
	}
}

func (wm *WindowManager) Open(name string) error {

	if _, err := ReadPID(name); err == nil {
		return fmt.Errorf("already open")
	}
	/*if _, exists := wm.windows[name]; exists {
		return fmt.Errorf("already open")
	}*/

	exe, err := os.Executable()

	if err != nil {
		return err
	}

	var cmd *exec.Cmd

	switch runtime.GOOS {

	case "darwin":

		cmd = exec.Command(
			"osascript",
			"-e",
			fmt.Sprintf(
				`tell application "Terminal" to do script "%s child %s"`,
				exe,
				name,
			),
		)

	case "linux":

		cmd = exec.Command(
			"x-terminal-emulator",
			"-e",
			exe,
			"child",
			name,
		)

	case "windows":

		cmd = exec.Command(
			"cmd",
			"/C",
			"start",
			exe,
			"child",
			name,
		)

	default:
		return fmt.Errorf("unsupported OS")
	}

	err = cmd.Start()

	if err != nil {
		return err
	}

	wm.windows[name] = cmd

	return nil
}

func (wm *WindowManager) Close(name string) error {

	pid, err := ReadPID(name)

	if err != nil {
		return fmt.Errorf("window not found")
	}

	process, err := os.FindProcess(pid)

	if err != nil {
		return err
	}

	err = process.Kill()

	if err != nil {
		return err
	}

	RemovePID(name)

	delete(wm.windows, name)

	return nil
}
