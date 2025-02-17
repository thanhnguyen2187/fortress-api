package employee

import (
	"errors"
	"strings"
	"time"

	"gorm.io/gorm"

	"github.com/dwarvesf/fortress-api/pkg/logger"
	"github.com/dwarvesf/fortress-api/pkg/model"
)

type UpdateEmployeeGeneralInfoInput struct {
	FullName           string
	Email              string
	Phone              string
	LineManagerID      model.UUID
	DisplayName        string
	GithubID           string
	NotionID           string
	NotionName         string
	NotionEmail        string
	DiscordName        string
	LinkedInName       string
	LeftDate           string
	JoinedDate         string
	OrganizationIDs    []model.UUID
	ReferredBy         model.UUID
	WiseRecipientID    string
	WiseRecipientEmail string
	WiseRecipientName  string
	WiseAccountNumber  string
	WiseCurrency       string
}

func (r *controller) UpdateGeneralInfo(employeeID string, body UpdateEmployeeGeneralInfoInput) (*model.Employee, error) {
	l := r.logger.Fields(logger.Fields{
		"controller": "employee",
		"method":     "UpdateGeneralInfo",
	})

	tx, done := r.repo.NewTransaction()

	// check line manager existence
	if !body.LineManagerID.IsZero() {
		exist, err := r.store.Employee.IsExist(tx.DB(), body.LineManagerID.String())
		if err != nil {
			l.Errorf(err, "failed to check line manager existence")
			return nil, done(err)
		}

		if !exist {
			return nil, done(ErrLineManagerNotFound)
		}
	}

	// check referrer existence
	if !body.ReferredBy.IsZero() {
		exist, err := r.store.Employee.IsExist(tx.DB(), body.ReferredBy.String())
		if err != nil {
			l.Errorf(err, "failed to check referer existence")
			return nil, done(err)
		}

		if !exist {
			return nil, done(ErrReferrerNotFound)
		}

		if employeeID == body.ReferredBy.String() {
			return nil, done(ErrCannotSelfReferral)
		}
	}

	emp, err := r.store.Employee.One(tx.DB(), employeeID, true)
	if err != nil {
		l.Errorf(err, "failed to get employee")
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, done(ErrEmployeeNotFound)
		}
		return nil, done(err)
	}

	if emp.TeamEmail != "" && emp.TeamEmail != body.Email {
		_, err = r.store.Employee.OneByEmail(r.repo.DB(), body.Email)
		if !errors.Is(err, gorm.ErrRecordNotFound) {
			if err == nil {
				return nil, ErrEmailExisted
			}
			return nil, err
		}
	}

	// 3. update information and return nil, done(err)

	if strings.TrimSpace(body.FullName) != "" {
		emp.FullName = body.FullName
	}

	if strings.TrimSpace(body.Email) != "" {
		emp.TeamEmail = body.Email
	}

	if strings.TrimSpace(body.Phone) != "" {
		emp.PhoneNumber = body.Phone
	}

	if strings.TrimSpace(body.DisplayName) != "" {
		emp.DisplayName = body.DisplayName
	}

	if strings.TrimSpace(body.JoinedDate) != "" {
		joinedDate, err := time.Parse("2006-01-02", body.JoinedDate)
		if err != nil {
			return nil, done(ErrInvalidJoinedDate)
		}
		emp.JoinedDate = &joinedDate
	}

	if strings.TrimSpace(body.LeftDate) != "" {
		leftDate, err := time.Parse("2006-01-02", body.LeftDate)
		if err != nil {
			return nil, done(ErrInvalidLeftDate)
		}
		emp.LeftDate = &leftDate
	}

	if emp.JoinedDate != nil && emp.LeftDate != nil {
		if emp.LeftDate.Before(*emp.JoinedDate) {
			return nil, done(ErrLeftDateBeforeJoinedDate)
		}
	}

	emp.LineManagerID = body.LineManagerID
	emp.ReferredBy = body.ReferredBy
	if strings.TrimSpace(body.WiseRecipientID) != "" {
		emp.WiseRecipientID = body.WiseRecipientID
	}

	if strings.TrimSpace(body.WiseAccountNumber) != "" {
		emp.WiseAccountNumber = body.WiseAccountNumber
	}

	if strings.TrimSpace(body.WiseRecipientEmail) != "" {
		emp.WiseRecipientEmail = body.WiseRecipientEmail
	}

	if strings.TrimSpace(body.WiseRecipientName) != "" {
		emp.WiseRecipientName = body.WiseRecipientName
	}

	if strings.TrimSpace(body.WiseCurrency) != "" {
		emp.WiseCurrency = body.WiseCurrency
	}

	discordID := ""
	if strings.TrimSpace(body.DiscordName) != "" {
		// Get discord info
		discordMember, err := r.service.Discord.GetMemberByUsername(body.DiscordName)
		if err != nil {
			l.Errorf(err, "failed to discord member member by discord name", "discordName", body.DiscordName)
			return nil, done(err)
		}

		if discordMember == nil {
			return nil, done(ErrCouldNotFoundDiscordMemberInGuild)
		}

		discordID = discordMember.User.ID
	}

	accountInUsed := false
	tmpE, err := r.store.Employee.GetByDiscordID(tx.DB(), discordID)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		l.Errorf(err, "failed to get employee by discord id", "discordID", discordID)
		return nil, done(err)
	}

	if !errors.Is(err, gorm.ErrRecordNotFound) {
		if tmpE.ID != emp.ID {
			accountInUsed = true
		}
	}

	if accountInUsed {
		return nil, done(ErrDiscordAccountAlreadyUsedByAnotherEmployee)
	}

	discordAccountInput := &model.DiscordAccount{
		DiscordID: discordID,
		Username:  body.DiscordName,
	}

	discordAccount, err := r.store.DiscordAccount.Upsert(tx.DB(), discordAccountInput)
	if err != nil {
		l.Errorf(err, "failed to upsert discord account", "discordAccount", discordAccountInput)
		return nil, done(err)
	}

	emp.DiscordAccountID = discordAccount.ID

	// Update social accounts
	saInput := model.SocialAccountInput{
		GithubID:     body.GithubID,
		NotionID:     body.NotionID,
		NotionName:   body.NotionName,
		NotionEmail:  body.NotionEmail,
		LinkedInName: body.LinkedInName,
	}

	if err := r.updateSocialAccounts(tx.DB(), saInput, emp.ID); err != nil {
		l.Errorf(err, "failed to update social account", "socialAccount", saInput)
		return nil, done(err)
	}

	_, err = r.store.Employee.UpdateSelectedFieldsByID(tx.DB(), employeeID, *emp,
		"full_name",
		"team_email",
		"phone_number",
		"line_manager_id",
		"display_name",
		"joined_date",
		"left_date",
		"referred_by",
		"wise_recipient_id",
		"wise_account_number",
		"wise_recipient_email",
		"wise_recipient_name",
		"wise_currency",
		"discord_account_id",
	)
	if err != nil {
		return nil, done(err)
	}

	if len(body.OrganizationIDs) > 0 {
		// Check organizations existence
		organizations, err := r.store.Organization.All(tx.DB())
		if err != nil {
			return nil, done(err)
		}

		orgMaps := model.ToOrganizationMap(organizations)
		for _, sID := range body.OrganizationIDs {
			_, ok := orgMaps[sID]
			if !ok {
				l.Errorf(ErrOrganizationNotFound, "organization not found with id: ", sID.String())
				return nil, done(ErrOrganizationNotFound)
			}
		}

		// Delete all exist employee organizations
		if err := r.store.EmployeeOrganization.DeleteByEmployeeID(tx.DB(), employeeID); err != nil {
			return nil, done(err)
		}

		// Create new employee position
		for _, orgID := range body.OrganizationIDs {
			_, err := r.store.EmployeeOrganization.Create(tx.DB(), &model.EmployeeOrganization{
				EmployeeID:     model.MustGetUUIDFromString(employeeID),
				OrganizationID: orgID,
			})
			if err != nil {
				return nil, done(err)
			}
		}
	}

	emp, err = r.store.Employee.One(tx.DB(), employeeID, true)
	if err != nil {
		return nil, done(err)
	}

	return emp, done(nil)
}

func (r *controller) updateSocialAccounts(db *gorm.DB, input model.SocialAccountInput, employeeID model.UUID) error {
	l := r.logger.Fields(logger.Fields{
		"handler":    "employee",
		"method":     "updateSocialAccounts",
		"input":      input,
		"employeeID": employeeID,
	})

	accounts, err := r.store.SocialAccount.GetByEmployeeID(db, employeeID.String())
	if err != nil {
		l.Error(err, "failed to get social accounts by employeeID")
		return err
	}

	accountsInput := map[model.SocialAccountType]model.SocialAccount{
		model.SocialAccountTypeGitHub: {
			Type:       model.SocialAccountTypeGitHub,
			EmployeeID: employeeID,
			AccountID:  input.GithubID,
			Name:       input.GithubID,
		},
		model.SocialAccountTypeNotion: {
			Type:       model.SocialAccountTypeNotion,
			EmployeeID: employeeID,
			AccountID:  input.NotionID,
			Name:       input.NotionName,
			Email:      input.NotionEmail,
		},
		model.SocialAccountTypeLinkedIn: {
			Type:       model.SocialAccountTypeLinkedIn,
			EmployeeID: employeeID,
			AccountID:  input.LinkedInName,
			Name:       input.LinkedInName,
		},
	}

	for _, account := range accounts {
		delete(accountsInput, account.Type)

		switch account.Type {
		case model.SocialAccountTypeGitHub:
			account.AccountID = input.GithubID
			account.Name = input.GithubID
		case model.SocialAccountTypeNotion:
			account.Name = input.NotionName
			account.Email = input.NotionEmail
		case model.SocialAccountTypeLinkedIn:
			account.AccountID = input.LinkedInName
			account.Name = input.LinkedInName
		default:
			continue
		}

		if _, err := r.store.SocialAccount.UpdateSelectedFieldsByID(db, account.ID.String(), *account, "account_id", "name", "email"); err != nil {
			l.Errorf(err, "failed to update social account %s", account.ID)
			return err
		}
	}

	for _, account := range accountsInput {
		if _, err := r.store.SocialAccount.Create(db, &account); err != nil {
			l.AddField("account", account).Error(err, "failed to create social account")
			return err
		}
	}

	return nil
}
