package view

import (
	"math"
	"strings"

	"github.com/dwarvesf/fortress-api/pkg/model"
)

type ProjectSizeResponse struct {
	Data []*model.ProjectSize `json:"data"`
}

type WorkSurveyResponse struct {
	Data WorkSurveysData `json:"data"`
}

type Trend struct {
	Workload float64 `json:"workload"`
	Deadline float64 `json:"deadline"`
	Learning float64 `json:"learning"`
}

type WorkSurvey struct {
	EndDate  string  `json:"endDate"`
	Workload float64 `json:"workload"`
	Deadline float64 `json:"deadline"`
	Learning float64 `json:"learning"`
	Trend    *Trend  `json:"trend"`
}

type ActionItemTrend struct {
	High   float64 `json:"high"`
	Medium float64 `json:"medium"`
	Low    float64 `json:"low"`
}

type AuditActionItemReport struct {
	Quarter string           `json:"quarter"`
	High    int64            `json:"high"`
	Medium  int64            `json:"medium"`
	Low     int64            `json:"low"`
	Trend   *ActionItemTrend `json:"trend"`
}

type WorkSurveysData struct {
	Project     *BasicProjectInfo `json:"project"`
	WorkSurveys []*WorkSurvey     `json:"workSurveys"`
}

type ActionItemReportResponse struct {
	AuditActionItemReports []*AuditActionItemReport `json:"data"`
}

func ToWorkSurveyData(project *model.Project, workSurveys []*model.WorkSurvey) *WorkSurveysData {
	rs := &WorkSurveysData{}

	for _, ws := range workSurveys {
		rs.WorkSurveys = append(rs.WorkSurveys, &WorkSurvey{
			EndDate:  ws.EndDate.Format("02/01"),
			Workload: ws.Workload,
			Deadline: ws.Deadline,
			Learning: ws.Learning,
		})
	}

	if project != nil {
		rs.Project = toBasicProjectInfo(*project)
	}

	if workSurveys != nil && len(workSurveys) > 1 {
		for i := 1; i < len(workSurveys); i++ {
			rs.WorkSurveys[i].Trend = calculateTrend(workSurveys[i-1], workSurveys[i])
		}
	}

	return rs
}

func ToActionItemReportData(actionItemReports []*model.ActionItemReport) []*AuditActionItemReport {
	var rs []*AuditActionItemReport
	// reverse the order to correct timeline
	for i, j := 0, len(actionItemReports)-1; i < j; i, j = i+1, j-1 {
		actionItemReports[i], actionItemReports[j] = actionItemReports[j], actionItemReports[i]
	}
	for _, ws := range actionItemReports {
		rs = append(rs, &AuditActionItemReport{
			Quarter: strings.Split(ws.Quarter, "/")[1] + "/" + strings.Split(ws.Quarter, "/")[0],
			High:    ws.High,
			Medium:  ws.Low,
			Low:     ws.Low,
		})
	}

	if actionItemReports != nil && len(actionItemReports) > 1 {
		for i := 1; i < len(actionItemReports); i++ {
			rs[i].Trend = calculateActionItemReportTrend(actionItemReports[i-1], actionItemReports[i])
		}
	}

	return rs
}

// calculateTrend calculate the trend for work survey
func calculateTrend(previous *model.WorkSurvey, current *model.WorkSurvey) *Trend {
	rs := &Trend{}

	// if previous or current value = 0 trend = 0
	if previous.Workload == 0 || current.Workload == 0 {
		rs.Workload = 0
	} else {
		rs.Workload = (current.Workload - previous.Workload) / previous.Workload * 100
	}

	if previous.Deadline == 0 || current.Deadline == 0 {
		rs.Deadline = 0
	} else {
		rs.Deadline = (current.Deadline - previous.Deadline) / previous.Deadline * 100
	}

	if previous.Learning == 0 || current.Learning == 0 {
		rs.Learning = 0
	} else {
		rs.Learning = (current.Learning - previous.Learning) / previous.Learning * 100
	}

	return rs
}

// calculateTrend calculate the trend for action item report
func calculateActionItemReportTrend(previous *model.ActionItemReport, current *model.ActionItemReport) *ActionItemTrend {
	rs := &ActionItemTrend{}

	// if previous or current value = 0 trend = 0
	if previous.High == 0 || current.High == 0 {
		rs.High = 0
	} else {
		rs.High = float64(float64(current.High-previous.High) / float64(previous.High) * 100)
	}

	if previous.Medium == 0 || current.Medium == 0 {
		rs.Medium = 0
	} else {
		rs.Medium = float64(float64(current.Medium-previous.Medium) / float64(previous.Medium) * 100)
	}

	if previous.Low == 0 || current.Low == 0 {
		rs.Low = 0
	} else {
		rs.Low = float64(float64(current.Low-previous.Low) / float64(previous.Low) * 100)
	}

	return rs
}

type EngineeringHealthResponse struct {
	Data EngineeringHealthData `json:"data"`
}

type EngineeringHealthData struct {
	Average []*EngineeringHealth      `json:"average"`
	Groups  []*GroupEngineeringHealth `json:"groups"`
}

type EngineeringHealth struct {
	Quarter string  `json:"quarter"`
	Value   float64 `json:"avg"`
	Trend   float64 `json:"trend"`
}

type GroupEngineeringHealth struct {
	Quarter       string                 `json:"quarter"`
	Delivery      float64                `json:"delivery"`
	Quality       float64                `json:"quality"`
	Collaboration float64                `json:"collaboration"`
	Feedback      float64                `json:"feedback"`
	Trend         EngineeringHealthTrend `json:"trend"`
}

type EngineeringHealthTrend struct {
	Delivery      float64 `json:"delivery"`
	Quality       float64 `json:"quality"`
	Collaboration float64 `json:"collaboration"`
	Feedback      float64 `json:"feedback"`
}

func ToEngineeringHealthData(average []*model.AverageEngineeringHealth, groups []*model.GroupEngineeringHealth) *EngineeringHealthData {
	rs := &EngineeringHealthData{}

	// Reverse quarter order
	for i, j := 0, len(average)-1; i < j; i, j = i+1, j-1 {
		average[i], average[j] = average[j], average[i]
	}

	for i, j := 0, len(groups)-1; i < j; i, j = i+1, j-1 {
		groups[i], groups[j] = groups[j], groups[i]
	}

	for _, a := range average {
		rs.Average = append(rs.Average, &EngineeringHealth{
			Quarter: strings.Split(a.Quarter, "/")[1] + "/" + strings.Split(a.Quarter, "/")[0],
			Value:   a.Avg,
		})
	}

	calculateTrendForEngineeringHealthList(rs.Average)

	rs.Groups = toGroupEngineeringHealth(groups)
	calculateEngineeringHealthGroupTrend(rs.Groups)

	return rs
}

func toGroupEngineeringHealth(groups []*model.GroupEngineeringHealth) []*GroupEngineeringHealth {
	var rs []*GroupEngineeringHealth
	count := 0
	quarter := ""
	i := 0

	for i < len(groups) {
		if quarter != groups[i].Quarter {
			count++
			quarter = groups[i].Quarter

			if count > 4 {
				break
			}
		}

		rs = append(rs, &GroupEngineeringHealth{
			Quarter: strings.Split(groups[i].Quarter, "/")[1] + "/" + strings.Split(groups[i].Quarter, "/")[0],
		})

		for quarter == groups[i].Quarter {
			switch groups[i].Area {
			case model.AuditItemAreaDelivery:
				rs[count-1].Delivery = groups[i].Avg
			case model.AuditItemAreaQuality:
				rs[count-1].Quality = groups[i].Avg
			case model.AuditItemAreaCollaborating:
				rs[count-1].Collaboration = groups[i].Avg
			case model.AuditItemAreaFeedback:
				rs[count-1].Feedback = groups[i].Avg
			}

			i++
			if i >= len(groups) {
				break
			}
		}

	}

	return rs
}

func calculateTrendForEngineeringHealthList(healths []*EngineeringHealth) {
	for i := 1; i < len(healths); i++ {
		healths[i].Trend = calculateEngineeringHealthTrend(healths[i-1], healths[i])
	}
}

func calculateEngineeringHealthGroupTrend(groups []*GroupEngineeringHealth) {
	for i := 1; i < len(groups); i++ {
		groups[i].Trend.Delivery = calculateEngineeringHealthTrend(&EngineeringHealth{Value: groups[i-1].Delivery}, &EngineeringHealth{Value: groups[i].Delivery})
		groups[i].Trend.Quality = calculateEngineeringHealthTrend(&EngineeringHealth{Value: groups[i-1].Quality}, &EngineeringHealth{Value: groups[i].Quality})
		groups[i].Trend.Collaboration = calculateEngineeringHealthTrend(&EngineeringHealth{Value: groups[i-1].Collaboration}, &EngineeringHealth{Value: groups[i].Collaboration})
		groups[i].Trend.Feedback = calculateEngineeringHealthTrend(&EngineeringHealth{Value: groups[i-1].Feedback}, &EngineeringHealth{Value: groups[i].Feedback})
	}
}

func calculateEngineeringHealthTrend(previous *EngineeringHealth, current *EngineeringHealth) float64 {
	// if previous or current value = 0 trend = 0
	if previous.Value == 0 || current.Value == 0 {
		return 0
	}

	// return the value fixed 2 decimal places
	return float64(math.Trunc((current.Value-previous.Value)/previous.Value*100*100)) / 100
}

type AuditResponse struct {
	Data AuditData `json:"data"`
}

type AuditData struct {
	Average []*Audit      `json:"average"`
	Groups  []*GroupAudit `json:"groups"`
}

type Audit struct {
	Quarter string  `json:"quarter"`
	Value   float64 `json:"avg"`
	Trend   float64 `json:"trend"`
}

type GroupAudit struct {
	Quarter    string          `json:"quarter"`
	Frontend   float64         `json:"frontend"`
	Backend    float64         `json:"backend"`
	System     float64         `json:"system"`
	Process    float64         `json:"process"`
	Mobile     float64         `json:"mobile"`
	Blockchain float64         `json:"blockchain"`
	Trend      GroupAuditTrend `json:"trend"`
}

type GroupAuditTrend struct {
	Frontend   float64 `json:"frontend"`
	Backend    float64 `json:"backend"`
	System     float64 `json:"system"`
	Process    float64 `json:"process"`
	Mobile     float64 `json:"mobile"`
	Blockchain float64 `json:"blockchain"`
}

func ToAuditData(average []*model.AverageAudit, groups []*model.GroupAudit) *AuditData {
	rs := &AuditData{}

	// Reverse quarter order
	for i, j := 0, len(average)-1; i < j; i, j = i+1, j-1 {
		average[i], average[j] = average[j], average[i]
	}

	for i, j := 0, len(groups)-1; i < j; i, j = i+1, j-1 {
		groups[i], groups[j] = groups[j], groups[i]
	}

	for _, a := range average {
		rs.Average = append(rs.Average, &Audit{
			Quarter: strings.Split(a.Quarter, "/")[1] + "/" + strings.Split(a.Quarter, "/")[0],
			Value:   a.Avg,
		})
	}

	calculateTrendForAuditList(rs.Average)

	rs.Groups = toGroupAudit(groups)
	calculateAuditGroupTrend(rs.Groups)

	return rs
}

func toGroupAudit(groups []*model.GroupAudit) []*GroupAudit {
	var rs []*GroupAudit

	for i := range groups {
		rs = append(rs, &GroupAudit{
			Quarter:    strings.Split(groups[i].Quarter, "/")[1] + "/" + strings.Split(groups[i].Quarter, "/")[0],
			Frontend:   groups[i].Frontend,
			Backend:    groups[i].Backend,
			System:     groups[i].System,
			Process:    groups[i].Process,
			Mobile:     groups[i].Mobile,
			Blockchain: groups[i].Blockchain,
		})
	}

	return rs
}

func calculateTrendForAuditList(healths []*Audit) {
	for i := 1; i < len(healths); i++ {
		healths[i].Trend = calculateAuditTrend(healths[i-1], healths[i])
	}
}

func calculateAuditGroupTrend(groups []*GroupAudit) {
	for i := 1; i < len(groups); i++ {
		groups[i].Trend.Frontend = calculateAuditTrend(&Audit{Value: groups[i-1].Frontend}, &Audit{Value: groups[i].Frontend})
		groups[i].Trend.Backend = calculateAuditTrend(&Audit{Value: groups[i-1].Backend}, &Audit{Value: groups[i].Backend})
		groups[i].Trend.System = calculateAuditTrend(&Audit{Value: groups[i-1].System}, &Audit{Value: groups[i].System})
		groups[i].Trend.Process = calculateAuditTrend(&Audit{Value: groups[i-1].Process}, &Audit{Value: groups[i].Process})
		groups[i].Trend.Mobile = calculateAuditTrend(&Audit{Value: groups[i-1].Mobile}, &Audit{Value: groups[i].Mobile})
		groups[i].Trend.Blockchain = calculateAuditTrend(&Audit{Value: groups[i-1].Blockchain}, &Audit{Value: groups[i].Blockchain})
	}
}

func calculateAuditTrend(previous *Audit, current *Audit) float64 {
	// if previous or current value = 0 trend = 0
	if previous.Value == 0 || current.Value == 0 {
		return 0
	}

	// return the value fixed 2 decimal places
	return float64(math.Trunc((current.Value-previous.Value)/previous.Value*100*100)) / 100
}
