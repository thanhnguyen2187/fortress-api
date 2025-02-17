package consts

// Define basecamp resource id
const (
	WoodlandScheduleID             int64 = 1346305137
	PlaygroundScheduleID           int64 = 1941398077
	WoodlandID                           = 9403032
	PlaygroundID                         = 12984857
	PlaygroundCampfireID                 = 1941398072
	HiringID                             = 13171568
	HiringScheduleID                     = 1974612191
	PlaygroundTodoID                     = 1941398075
	PlaygroundDynamicsTodoID             = 2235833582
	PlaygroundExpenseTodoID              = 2436015405
	FoundationID                         = 9405283
	FoundationTodoID                     = 1346687281
	FoundationCampfireID                 = 1346687275
	WoodlandDynamicsTodoID               = 2566078344
	VenturesID                           = 13484002
	DesignID                             = 9404034
	OperationID                          = 9403043
	ManagementID                         = 15240223
	AccountingID                         = 15258324
	TechRadarID                          = 13307382
	ReadifyID                            = 16277845
	ReadifyCampfireID                    = 2502884934
	HookID                               = 393058
	WoodlandTodoID                       = 1346305133
	OperationTodoID                      = 1346306047
	ManagementTodoID                     = 2326974480
	AccountingTodoID                     = 2329633561
	ExpenseTodoID                        = 2353511928
	AutoBotID                            = 25727627
	WoodlandMessageBoardID               = 1346305130
	PlaygroundMessageBoardID             = 1941398073
	HiringMessageBoardID                 = 1974612189
	CompanyBasecampID                    = 4108948
	ProjectManagementID                  = 1970322248
	TechRadarProjectID                   = 13307382
	TechRadarMessageBoardID              = 1998256729
	TechRadarTodoSetID                   = 1998256730
	TechRadarAssessTodoListID            = 1998259939 // temporary hardcode
	HiringTodoSetID                      = 1974612190
	HRTodoID                             = 2040601263
	TechRadarCampfireID                  = 1998256728
	OperationCampfireID                  = 1346306044
	WoodlandCampfireID                   = 1346305126
	FortressCampfireID                   = 1347531395
	HiringCampfireID                     = 1974612188
	FortressID                           = 9410372
	OnleaveID                            = 5669914253
	OnleavePlaygroundID                  = 2243342506
	WarehouseID                          = 15921521
	PlaygroundHRTodoID                   = 2475678340
	BirthdayGift2020                     = 2576503375
	BirthdayGift2021                     = 3338845535
	PaperTrailTodoListID                 = 2685205937
	PlaygroundPaperTrailTodoListID       = 2685503124
	ShareholderID                        = 16944388
	ShareholderCampfireID                = 2635581864
	SudoID                               = 16473245
	SudoCampfireID                       = 2538063314
	OpsExpenseTodoID                     = 4665885355
	PlaygroundBirthdayTodoID             = 3942871393
	BirthdayToDoListID                   = 3941578970

	// People
	LyBasecampID           = 21564173
	HanBasecampID          = 21562923
	QuangBasecampID        = 22659105
	AnBasecampID           = 21562943
	HuyNguyenBasecampID    = 22658825
	HuyGiangBasecampID     = 22658816
	MinhTranBasecampID     = 21564151
	NamTranBasecampID      = 21675130
	DuyenBasecampID        = 26160403
	TrungPhanBasecampID    = 21574701
	PhuongTruongBasecampID = 21842626
	ThanhNguyenBasecampID  = 21572501
	VanNguyenBasecampID    = 26595807
	KhaiLeBasecampID       = 24006290
	KhanhTruongBasecampID  = 24419646
	GiangThanBasecampID    = 26160802
	HelenBasecampID        = 40439249
	ThuongBasecampID       = 38246363
	NamNguyenBasecampID    = 21581534

	// BucketName
	BucketNameWoodLand   = "Woodland"
	BucketNamePlayGround = "Fortress | Playground"

	// OrgChart
	ManagementLevel = 3

	// Basecamp kind list
	TodoCreate         = "todo_created"
	TodoComplete       = "todo_completed"
	TodoUncomplete     = "todo_uncompleted"
	CommentCreate      = "comment_created"
	MessageBoardCreate = "message_created"
	MessageBoardActive = "message_active"

	AutoBotSgID = "BAh7CEkiCGdpZAY6BkVUSSIpZ2lkOi8vYmMzL1BlcnNvbi8yNTcyNzYyNz9leHBpcmVzX2luBjsAVEkiDHB1cnBvc2UGOwBUSSIPYXR0YWNoYWJsZQY7AFRJIg9leHBpcmVzX2F0BjsAVDA=--5a1528460315bfd57bc41cf6bd3f899b1c346e7b"

	// Basecamp Comment Message
	CommentThankYouEmailSent                      = "Thank you email has been sent"
	CommentThankYouEmailFailed                    = "Unable to send thank you email invoice."
	CommentUpdateInvoiceSuccessfully              = "Invoice status has been set to paid"
	CommentUpdateInvoiceFailed                    = `Unable to update invoice status`
	CommentMoveInvoicePDFToPaidDirSuccessfully    = "GDrive file has been updated"
	CommentMoveInvoicePDFToPaidDirFailed          = "Unable to move invoice pdf to paid directory"
	CommentCantFindInvoice                        = "Invoice not found"
	CommentInvoiceFileMoved                       = "Invoice file has been moved to Paid folder in Google Drive"
	CommentUnableToUpdateGDLoc                    = "Unable to update Google Drive location"
	CommentMissingConfirmation                    = "Missing Confirmation"
	CommentInvalidOnLeaveFormat                   = "Invalid on leave format"
	CommentCreateScheduleSuccessfully             = "Create schedule successfully"
	CommentCreateScheduleFailed                   = "Unable to create schedule"
	CommentCreateExpenseSuccessfully              = "Create expense successfully"
	CommentCreateExpenseFailed                    = "Unable to create expense"
	CommentDeleteExpenseSuccessfully              = "Delete expense successfully"
	CommentDeleteExpenseFailed                    = "Delete expense failed"
	CommentStoreAccountingTransactionFailed       = "Store accounting transaction failed"
	CommentStoreAccountingTransactionSuccessfully = "Store accounting transaction successfully"
	CommentProbationReviewSuccessfully            = "Probation review succeed"
	CommentThankYouEmailSentSuccessfully          = "Thank you email sent successfully"
	CommentOfferEmailSentSuccessfully             = "Offer email sent successfully"
	CommentHiringNoActionTaken                    = "Neither Offered nor Fail, no action taken"
	CommentHiredCandidate                         = "Candidate hired"
	CommentRejectedCandidate                      = "Candidate rejected"
)
