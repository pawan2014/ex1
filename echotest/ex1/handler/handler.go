package handler

import "github.com/labstack/echo"

type Excuse struct {
	Error string `json:"error"`
	Id    string `json:"id"`
	Quote string `json:"quote"`
}

type Handler struct {
}

// Machines give all machines
func (h *Handler) GetMachines(c echo.Context) (err error) {
	return nil
}

// Machine add one machine with Tags
func (h *Handler) AddMachineWithTags(c echo.Context) (err error) {
	return nil
}

// Get one machine with tags /machine/tags
func (h *Handler) GetMachineWithTags(c echo.Context) (err error) {
	return nil
}

// /machine/{id}/tag POST & PUT
func (h *Handler) UpdateTagUnderMachine(c echo.Context) (err error) {
	return nil
}

// /machine/{}/start
func (h *Handler) StartMachineStream(c echo.Context) (err error) {
	return nil
}

// /machine/{}/stop
func (h *Handler) StopMachineStream(c echo.Context) (err error) {
	return nil
}

// /machine/{}/status
func (h *Handler) MachineStreamingStatus(c echo.Context) (err error) {
	return nil
}
