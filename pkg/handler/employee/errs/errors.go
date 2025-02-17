package errs

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/dwarvesf/fortress-api/pkg/controller/employee"
	"github.com/dwarvesf/fortress-api/pkg/view"
)

var (
	ErrInvalidEmployeeID          = errors.New("invalid employee ID")
	ErrInvalidEmployeeStatus      = errors.New("invalid value for employee status")
	ErrInvalidJoinedDate          = errors.New("invalid join date")
	ErrInvalidPositionCode        = errors.New("invalid position code")
	ErrInvalidStackCode           = errors.New("invalid stack code")
	ErrInvalidProjectCode         = errors.New("invalid project code")
	ErrInvalidChapterCode         = errors.New("invalid chapter code")
	ErrInvalidSeniorityCode       = errors.New("invalid seniority code")
	ErrInvalidOrganizationCode    = errors.New("invalid organization code")
	ErrInvalidEmailDomain         = errors.New("invalid email domain")
	ErrRoleCannotBeEmpty          = errors.New("role cannot be empty")
	ErrCountryNotFound            = errors.New("country not found")
	ErrCityDoesNotBelongToCountry = errors.New("city does not belong to country")
)

func ConvertControllerErr(c *gin.Context, err error) {
	if err == nil {
		return
	}

	var status int

	switch err {
	case employee.ErrEmployeeNotFound,
		employee.ErrLineManagerNotFound,
		employee.ErrRoleNotfound,
		employee.ErrSeniorityNotfound,
		employee.ErrReferrerNotFound,
		employee.ErrOrganizationNotFound,
		employee.ErrStackNotFound,
		employee.ErrPositionNotFound:
		status = http.StatusNotFound

	case employee.ErrInvalidJoinedDate,
		employee.ErrInvalidLeftDate,
		employee.ErrLeftDateBeforeJoinedDate,
		employee.ErrEmployeeExisted,
		employee.ErrInvalidCountryOrCity,
		employee.ErrInvalidFileExtension,
		employee.ErrInvalidFileSize,
		employee.ErrInvalidAccountRole,
		employee.ErrEmailExisted,
		employee.ErrTeamEmailExisted,
		employee.ErrPersonalEmailExisted:
		status = http.StatusBadRequest

	default:
		status = http.StatusInternalServerError
	}

	c.JSON(status, view.CreateResponse[any](nil, nil, err, nil, ""))
}
