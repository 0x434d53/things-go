package dal

// BSSyncronyMetadata struct
type BSSyncronyMetadata struct {
	UUID  string `db:"uuid"`
	Value []byte `db:"value"`
}

// Meta struct
type Meta struct {
	Key   string `db:"key"`
	Value string `db:"value"`
}

// TMArea struct
type TMArea struct {
	UUID         string `db:"uuid"`
	Title        string `db:"title"`
	Visible      int    `db:"visible"`
	Index        int    `db:"index"`
	CachedTags   []byte `db:"cachedTags"`
	Experimental []byte `db:"experimental"`
}

// TMAreaTag struct
type TMAreaTag struct {
	Areas string `db:"areas"`
	Tags  string `db:"tags"`
}

// TMChecklistItem struct
type TMChecklistItem struct {
	UUID                 string  `db:"uuid"`
	UserModificationDate float64 `db:"userModificationDate"`
	CreationDate         float64 `db:"creationDate"`
	Title                string  `db:"title"`
	Status               int     `db:"status"`
	StopDate             float64 `db:"stopDate"`
	Index                int     `db:"index"`
	Task                 string  `db:"task"`
	LeavesTombstone      int     `db:"leavesTombstone"`
	Experimental         []byte  `db:"experimental"`
}

// TMCommand struct
type TMCommand struct {
	UUID         string  `db:"uuid"`
	CreationDate float64 `db:"creationDate"`
	Type         int     `db:"type"`
	Info         []byte  `db:"info"`
}

// TMContact struct
type TMContact struct {
	UUID            string `db:"uuid"`
	DisplayName     string `db:"displayName"`
	FirstName       string `db:"firstName"`
	LastName        string `db:"lastName"`
	Emails          string `db:"emails"`
	AppleAddresssID string `db:"appleAddressBookId"`
	Index           int    `db:"index"`
}

// TMMetaItem struct
type TMMetaItem struct {
	UUID  string `db:"uuid"`
	Value []byte `db:"value"`
}

// TMSettings struct
type TMSettings struct {
	UUID               string  `db:"uuid"`
	LogInterval        int     `db:"logInterval"`
	ManualLogDate      float64 `db:"manualLogDate"`
	GroupTodayByParent int     `db:"groupTodayByParent"`
	URISchemeAuthToken string  `db:"uriSchemeAuthenticationToken"`
	Experimental       []byte  `db:"experimental"`
}

// TMTag struct
type TMTag struct {
	UUID         string       `db:"uuid"`
	Title        string       `db:"title"`
	Shortcut     *string      `db:"shortcut"`
	UsedDate     NullRealTime `db:"usedDate"`
	Parent       *string      `db:"parent"`
	Index        int          `db:"index"`
	Experimental []byte       `db:"experimental"`
}

// TMTask struct
type TMTask struct {
	UUID                 string       `db:"uuid"`
	LeavesTombstone      int          `db:"leavesTombstone"`
	CreationDate         NullRealTime `db:"creationDate"`
	UserModificationDate NullRealTime `db:"userModificationDate"`
	Type                 int          `db:"type"`
	Status               int          `db:"status"`
	StopDate             NullRealTime `db:"stopDate"`
	Trashed              int          `db:"trashed"`
	Title                string       `db:"title"`
	Notes                string       `db:"notes"`
	NotesSync            int          `db:"notesSync"`
	CachedTags           []byte       `db:"cachedTags"`

	Start                           int          `db:"start"`
	StartDate                       NullIntTime  `db:"startDate"`
	StartBucket                     int          `db:"startBucket"`
	ReminderTime                    NullIntTime  `db:"reminderTime"`
	LastReminderInteractionDate     NullRealTime `db:"lastReminderInteractionDate"`
	Deadline                        NullIntTime  `db:"deadline"`
	DeadlineSuppressionDate         NullIntTime  `db:"deadlineSuppressionDate"`
	T2DeadlineOffset                int          `db:"t2_deadlineOffset"`
	Index                           int          `db:"index"`
	TodayIndex                      int          `db:"todayIndex"`
	TodayIndexReferenceDate         NullIntTime  `db:"todayIndexReferenceDate"`
	Area                            *string      `db:"area"`
	Project                         *string      `db:"project"`
	Heading                         *string      `db:"heading"`
	Contact                         *string      `db:"contact"`
	UntrashedLeafActionsCount       int          `db:"untrashedLeafActionsCount"`
	OpenUntrashedLeafActionsCount   int          `db:"openUntrashedLeafActionsCount"`
	ChecklistItemsCount             int          `db:"checklistItemsCount"`
	OpenChecklistItemsCount         int          `db:"openChecklistItemsCount"`
	RT1RepeatingTemplate            *string      `db:"rt1_repeatingTemplate"`
	RT1RecurrenceRule               []byte       `db:"rt1_recurrenceRule"`
	RT1InstanceCreationStartDate    NullIntTime  `db:"rt1_instanceCreationStartDate"`
	RT1InstanceCreationPaused       *int         `db:"rt1_instanceCreationPaused"`
	RT1InstanceCreationCount        *int         `db:"rt1_instanceCreationCount"`
	RT1AfterCompletionReferenceDate NullIntTime  `db:"rt1_afterCompletionReferenceDate"`
	RT1NextInstanceStartDate        NullIntTime  `db:"rt1_nextInstanceStartDate"`
	Experimental                    []byte       `db:"experimental"`
	Repeater                        []byte       `db:"repeater"`
	RepeaterMigrationDate           NullRealTime `db:"repeaterMigrationDate"`
	Tags                            []TMTag      `db:"-"`
}

// TMTaskTag struct
type TMTaskTag struct {
	Tasks string `db:"tasks"`
	Tags  string `db:"tags"`
}

// TMTombstone struct
type TMTombstone struct {
	UUID              string  `db:"uuid"`
	DeletionDate      float64 `db:"deletionDate"`
	DeletedObjectUUID string  `db:"deletedObjectUUID"`
}

// ThingsTouch_ExtensionCommandStore_Commands struct
type ThingsTouchExtensionCommandStoreCommands struct {
	ID   int    `db:"id"`
	Type string `db:"type"`
	Body []byte `db:"body"`
}

// ThingsTouch_ExtensionCommandStore_Meta struct
type ThingsTouchExtensionCommandStoreMeta struct {
	Key   string `db:"key"`
	Value []byte `db:"value"`
}
