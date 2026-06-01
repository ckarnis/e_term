package windows

import (
	"fmt"
	"os"
	"path/filepath"
	"strconv"
)

const PIDDir = "runtime/pids"

func pidFile(name string) string {
	return filepath.Join(PIDDir, name+".pid")
}

func WritePID(name string, pid int) error {

	err := os.MkdirAll(PIDDir, 0755)

	if err != nil {
		return err
	}

	return os.WriteFile(
		pidFile(name),
		[]byte(strconv.Itoa(pid)),
		0644,
	)
}

func ReadPID(name string) (int, error) {

	data, err := os.ReadFile(pidFile(name))

	if err != nil {
		return 0, err
	}

	pid, err := strconv.Atoi(string(data))

	if err != nil {
		return 0, err
	}

	return pid, nil
}

func RemovePID(name string) {
	_ = os.Remove(pidFile(name))
}

func DebugPID(name string) {
	pid, _ := ReadPID(name)
	fmt.Println(name, pid)
}
