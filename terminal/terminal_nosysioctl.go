// +build windows plan9 solaris

package terminal

func getWinsize() (*winsize, error) {
	ws := new(winsize)

	ws.Col = 80
	ws.Row = 24

	return ws, nil
}
