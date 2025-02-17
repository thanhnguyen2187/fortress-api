// Package notion please edit this file only with approval from hnh
package notion

import (
	"net/http"
	"strconv"
	"time"

	"github.com/dstotijn/go-notion"
	"github.com/gin-gonic/gin"

	"github.com/dwarvesf/fortress-api/pkg/model"
	"github.com/dwarvesf/fortress-api/pkg/view"
)

// ListEvents godoc
// @Summary Get list events from DF Dwarves Community Events
// @Description Get list events from DF Dwarves Community Events
// @Tags Notion
// @Accept  json
// @Produce  json
// @Success 200 {object} view.MessageResponse
// @Failure 400 {object} view.ErrorResponse
// @Router /notion/events [get]
func (h *handler) ListEvents(c *gin.Context) {
	filter := &notion.DatabaseQueryFilter{}

	nextDays := 7
	if c.Query("d") != "" {
		d, ok := c.GetQuery("d")
		if !ok {
			d = "7"
		}
		var err error
		nextDays, err = strconv.Atoi(d)
		if err != nil {
			nextDays = 7
		}
	}

	from := time.Now()
	to := from.Add(24 * time.Hour * time.Duration(nextDays))
	filter.And = append(filter.And, notion.DatabaseQueryFilter{
		Property: "Date",
		DatabaseQueryPropertyFilter: notion.DatabaseQueryPropertyFilter{
			Date: &notion.DatePropertyFilter{
				OnOrAfter: &from,
			},
		},
	})
	filter.And = append(filter.And, notion.DatabaseQueryFilter{
		Property: "Date",
		DatabaseQueryPropertyFilter: notion.DatabaseQueryPropertyFilter{
			Date: &notion.DatePropertyFilter{
				OnOrBefore: &to,
			},
		},
	})

	resp, err := h.service.Notion.GetDatabase(h.config.Notion.Databases.Event, filter, []notion.DatabaseQuerySort{
		{
			Property:  "Date",
			Direction: notion.SortDirAsc,
		},
	}, 0)
	if err != nil {
		c.JSON(http.StatusBadRequest, view.CreateResponse[any](nil, nil, err, nil, "can't get events from notion"))
		return
	}

	var events []model.NotionEvent

	for _, r := range resp.Results {
		props := r.Properties.(notion.DatabasePageProperties)

		name := props["Name"].Title[0].Text.Content
		if r.Icon != nil && r.Icon.Emoji != nil {
			name = *r.Icon.Emoji + " " + props["Name"].Title[0].Text.Content
		}

		activityType := ""
		if props["Activity Type"].Select != nil {
			activityType = props["Activity Type"].Select.Name
		}

		var date model.DateTime
		if props["Date"].Date != nil {
			date.Time = props["Date"].Date.Start.Time
			date.HasTime = props["Date"].Date.Start.HasTime()
		}

		events = append(events, model.NotionEvent{
			ID:           r.ID,
			Name:         name,
			ActivityType: activityType,
			Date:         date,
			CreatedAt:    r.CreatedTime,
		})
	}

	c.JSON(http.StatusOK, view.CreateResponse[any](events, nil, nil, nil, "get list events successfully"))
}
