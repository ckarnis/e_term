//go:build windows

package windows

import (
	"os"

	"golang.org/x/sys/windows"
)

func openConsoleFiles() (*os.File, *os.File, error) {
	// Open CONIN$ with the specific access flags Windows requires
	conin, err := windows.CreateFile(
		windows.StringToUTF16Ptr("CONIN$"),
		windows.GENERIC_READ|windows.GENERIC_WRITE,
		windows.FILE_SHARE_READ|windows.FILE_SHARE_WRITE,
		nil,
		windows.OPEN_EXISTING,
		0,
		0,
	)
	if err != nil {
		return nil, nil, err
	}

	conout, err := windows.CreateFile(
		windows.StringToUTF16Ptr("CONOUT$"),
		windows.GENERIC_READ|windows.GENERIC_WRITE,
		windows.FILE_SHARE_READ|windows.FILE_SHARE_WRITE,
		nil,
		windows.OPEN_EXISTING,
		0,
		0,
	)
	if err != nil {
		windows.CloseHandle(conin)
		return nil, nil, err
	}

	return os.NewFile(uintptr(conin), "CONIN$"),
		os.NewFile(uintptr(conout), "CONOUT$"),
		nil
}
