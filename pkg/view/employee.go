package view

import (
	"time"

	"github.com/dwarvesf/fortress-api/pkg/model"
)

// EmployeeData view for listing data
type EmployeeData struct {
	model.BaseModel

	// basic info
	FullName      string     `json:"fullName"`
	DisplayName   string     `json:"displayName"`
	TeamEmail     string     `json:"teamEmail"`
	PersonalEmail string     `json:"personalEmail"`
	Avatar        string     `json:"avatar"`
	PhoneNumber   string     `json:"phoneNumber"`
	Address       string     `json:"address"`
	MBTI          string     `json:"mbti"`
	Gender        string     `json:"gender"`
	Horoscope     string     `json:"horoscope"`
	DateOfBirth   *time.Time `json:"birthday"`
	DiscordID     string     `json:"discordID"`
	GithubID      string     `json:"githubID"`
	NotionID      string     `json:"notionID"`

	// working info
	WorkingStatus model.WorkingStatus `json:"status"`
	JoinedDate    *time.Time          `json:"joinedDate"`
	LeftDate      *time.Time          `json:"leftDate"`

	Seniority   *model.Seniority      `json:"seniority"`
	LineManager *BasicEmployeeInfo    `json:"lineManager"`
	Positions   []Position            `json:"positions"`
	Stacks      []Stack               `json:"stacks"`
	Roles       []Role                `json:"roles"`
	Projects    []EmployeeProjectData `json:"projects"`
	Chapters    []Chapter             `json:"chapters"`
}

type UpdateGeneralInfoEmployeeData struct {
	model.BaseModel

	// basic info
	FullName    string             `json:"fullName"`
	TeamEmail   string             `json:"teamEmail"`
	PhoneNumber string             `json:"phoneNumber"`
	DiscordID   string             `json:"discordID"`
	GithubID    string             `json:"githubID"`
	NotionID    string             `json:"notionID"`
	LineManager *BasicEmployeeInfo `json:"lineManager"`
}

type UpdateSkillEmployeeData struct {
	model.BaseModel

	Seniority *model.Seniority `json:"seniority"`
	Positions []model.Position `json:"positions"`
	Stacks    []model.Stack    `json:"stacks"`
	Chapters  []model.Chapter  `json:"chapters"`
}

type UpdatePersonalEmployeeData struct {
	model.BaseModel

	PersonalEmail string     `json:"personalEmail"`
	Address       string     `json:"address"`
	Gender        string     `json:"gender"`
	DateOfBirth   *time.Time `json:"birthday"`
}

type BasicEmployeeInfo struct {
	ID          string `json:"id"`
	FullName    string `json:"fullName"`
	DisplayName string `json:"displayName"`
	Avatar      string `json:"avatar"`
}

type UpdateEmployeeStatusResponse struct {
	Data EmployeeData `json:"data"`
}

type EmployeeListDataResponse struct {
	Data []EmployeeData `json:"data"`
}

type UpdataEmployeeStatusResponse struct {
	Data EmployeeData `json:"data"`
}

type UpdateSkillsEmployeeResponse struct {
	Data UpdateSkillEmployeeData `json:"data"`
}

type UpdatePersonalEmployeeResponse struct {
	Data UpdatePersonalEmployeeData `json:"data"`
}

type UpdateGeneralEmployeeResponse struct {
	Data UpdateGeneralInfoEmployeeData `json:"data"`
}

func ToUpdatePersonalEmployeeData(employee *model.Employee) *UpdatePersonalEmployeeData {
	return &UpdatePersonalEmployeeData{
		BaseModel: model.BaseModel{
			ID:        employee.ID,
			CreatedAt: employee.CreatedAt,
			UpdatedAt: employee.UpdatedAt,
		},
		DateOfBirth:   employee.DateOfBirth,
		Gender:        employee.Gender,
		Address:       employee.Address,
		PersonalEmail: employee.PersonalEmail,
	}
}

func ToUpdateSkillEmployeeData(employee *model.Employee) *UpdateSkillEmployeeData {
	positions := make([]model.Position, 0, len(employee.EmployeePositions))
	for _, v := range employee.EmployeePositions {
		positions = append(positions, v.Position)
	}

	stacks := make([]model.Stack, 0, len(employee.EmployeeStacks))
	for _, v := range employee.EmployeeStacks {
		stacks = append(stacks, v.Stack)
	}

	chapters := make([]model.Chapter, 0, len(employee.EmployeeChapters))
	for _, v := range employee.EmployeeChapters {
		chapters = append(chapters, v.Chapter)
	}

	return &UpdateSkillEmployeeData{
		BaseModel: model.BaseModel{
			ID:        employee.ID,
			CreatedAt: employee.CreatedAt,
			UpdatedAt: employee.UpdatedAt,
		},
		Seniority: employee.Seniority,
		Positions: positions,
		Stacks:    stacks,
		Chapters:  chapters,
	}
}

func ToUpdateGeneralInfoEmployeeData(employee *model.Employee) *UpdateGeneralInfoEmployeeData {
	rs := &UpdateGeneralInfoEmployeeData{
		BaseModel: model.BaseModel{
			ID:        employee.ID,
			CreatedAt: employee.CreatedAt,
			UpdatedAt: employee.UpdatedAt,
		},
		FullName:    employee.FullName,
		TeamEmail:   employee.TeamEmail,
		PhoneNumber: employee.PhoneNumber,
		DiscordID:   employee.DiscordID,
		GithubID:    employee.GithubID,
		NotionID:    employee.NotionID,
	}

	if employee.LineManager != nil {
		rs.LineManager = &BasicEmployeeInfo{
			ID:       employee.LineManager.ID.String(),
			FullName: employee.LineManager.FullName,
			Avatar:   employee.LineManager.Avatar,
		}
	}

	return rs
}

// ToEmployeeData parse employee date to response data
func ToEmployeeData(employee *model.Employee) *EmployeeData {
	employeeProjects := make([]EmployeeProjectData, 0, len(employee.ProjectMembers))
	for _, v := range employee.ProjectMembers {
		employeeProjects = append(employeeProjects, ToEmployeeProjectData(&v))
	}
	var lineManager *BasicEmployeeInfo
	if employee.LineManager != nil {
		lineManager = &BasicEmployeeInfo{
			ID:          employee.LineManager.BaseModel.ID.String(),
			FullName:    employee.LineManager.FullName,
			DisplayName: employee.LineManager.DisplayName,
			Avatar:      employee.LineManager.Avatar,
		}
	}

	rs := &EmployeeData{
		BaseModel: model.BaseModel{
			ID:        employee.ID,
			CreatedAt: employee.CreatedAt,
			UpdatedAt: employee.UpdatedAt,
		},
		FullName:      employee.FullName,
		DisplayName:   employee.DisplayName,
		TeamEmail:     employee.TeamEmail,
		PersonalEmail: employee.PersonalEmail,
		Avatar:        employee.Avatar,
		PhoneNumber:   employee.PhoneNumber,
		Address:       employee.Address,
		MBTI:          employee.MBTI,
		Gender:        employee.Gender,
		Horoscope:     employee.Horoscope,
		DateOfBirth:   employee.DateOfBirth,
		GithubID:      employee.GithubID,
		DiscordID:     employee.DiscordID,
		NotionID:      employee.NotionID,
		WorkingStatus: employee.WorkingStatus,
		Seniority:     employee.Seniority,
		JoinedDate:    employee.JoinedDate,
		LeftDate:      employee.LeftDate,
		Projects:      employeeProjects,
		LineManager:   lineManager,
		Roles:         ToRoles(employee.EmployeeRoles),
		Positions:     ToPositions(employee.EmployeePositions),
		Stacks:        ToEmployeeStacks(employee.EmployeeStacks),
		Chapters:      ToChapters(employee.EmployeeChapters),
	}

	if employee.Seniority != nil {
		rs.Seniority = employee.Seniority
	}

	if employee.LineManager != nil {
		rs.LineManager = &BasicEmployeeInfo{
			ID:       employee.LineManager.ID.String(),
			FullName: employee.LineManager.FullName,
			Avatar:   employee.LineManager.Avatar,
		}
	}

	return rs
}

func ToEmployeeListData(employees []*model.Employee) []EmployeeData {
	rs := make([]EmployeeData, 0, len(employees))
	for _, emp := range employees {
		empRes := ToEmployeeData(emp)
		rs = append(rs, *empRes)
	}
	return rs
}

type EmployeeContentData struct {
	Url string `json:"url"`
}

type EmployeeContentDataResponse struct {
	Data *EmployeeContentData `json:"data"`
}

func ToContentData(url string) *EmployeeContentData {
	return &EmployeeContentData{
		Url: url,
	}
}

func toBasicEmployeeInfo(employee model.Employee) *BasicEmployeeInfo {
	return &BasicEmployeeInfo{
		ID:          employee.ID.String(),
		FullName:    employee.FullName,
		DisplayName: employee.DisplayName,
		Avatar:      employee.Avatar,
	}
}
