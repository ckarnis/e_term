package stuff

import "os"

func FileExists(name string) bool {
	_, err := os.Stat(name)
	return err == nil
}

// if file exists err  is nil
// true is returned
