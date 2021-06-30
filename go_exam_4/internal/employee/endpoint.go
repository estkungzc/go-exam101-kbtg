package employee

import (
	"context"
	"github.com/labstack/echo/v4"
	"go-exam101-kbtg/go_exam_4/internal"
	"go-exam101-kbtg/go_exam_4/internal/models"
	"net/http"
)

type employeeService interface {
	GetEmployeeById(ctx context.Context, empId string) ([]models.Employee, error)
	GetEmployeeByFirstName(ctx context.Context, firstName string) ([]models.Employee, error)
}

type Endpoint struct {
	cv  *internal.Configs
	srv employeeService //Service Interface Tier
}

func NewEndpoint(cv *internal.Configs) *Endpoint {
	return &Endpoint{cv: cv, srv: NewService(cv)}
}

func (e Endpoint) GetEmployeeById(c echo.Context) error {
	var request struct {
		EmployeeId string `json:"employee_id"`
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: "mismatch json request format",
			Data:    nil,
		})
	}

	data, err := e.srv.GetEmployeeById(c.Request().Context(), request.EmployeeId)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Success: true,
		Code:    http.StatusOK,
		Message: "ok",
		Data:    data,
	})
}

func (e Endpoint) GetEmployeeByFirstName(c echo.Context) error {
	var request struct {
		FirstName string `json:"first_name"`
	}

	if err := c.Bind(&request); err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: "mismatch json request format",
			Data:    nil,
		})
	}

	data, err := e.srv.GetEmployeeByFirstName(c.Request().Context(), request.FirstName)
	if err != nil {
		return c.JSON(http.StatusBadRequest, models.Response{
			Success: false,
			Code:    http.StatusBadRequest,
			Message: err.Error(),
			Data:    nil,
		})
	}

	return c.JSON(http.StatusOK, models.Response{
		Success: true,
		Code:    http.StatusOK,
		Message: "ok",
		Data:    data,
	})
}
