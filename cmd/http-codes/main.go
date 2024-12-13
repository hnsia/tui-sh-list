package main

import (
	"net/http"
	"time"

	tea "github.com/charmbracelet/bubbletea"
)

const url = "https://charm.sh"

type model struct {
	status int
	err    error
}

// make a request to the server, and return the result as a Msg
func checkServer() tea.Msg {
	c := &http.Client{Timeout: 10 * time.Second}

	res, err := c.Get(url)
	if err != nil {
		return errMsg{err}
	}

	return statusMsg(res.StatusCode)
}

type (
	statusMsg int
	errMsg    struct{ err error }
)

// Error implements the error interface
func (e errMsg) Error() string { return e.err.Error() }
