package handler

//------------------------------------------------REQUEST ENTITY-------------------------------------------------------
// Task ID

// ReqFetchTaskID Request entity get task id with count filter
type ReqFetchTaskID struct {
}

// Task in line

// ReqFetchBaseTasks Request entity get base info tasks with count filter
type ReqFetchBaseTasks struct {
}

// ReqFetchBaseTaskByID Request entity get base info task by ID
type ReqFetchBaseTaskByID struct {
}

// Task global edit

// ReqFetchTask Request entity for get Task with count filter
type ReqFetchTask struct {
}

// ReqFetchTaskByID Request entity get task by ID
type ReqFetchTaskByID struct {
}

// ReqAddTask Request entity for add new task
type ReqAddTask struct {
}

// ReqDeleteTask Request entity for delete tasks by ID
type ReqDeleteTask struct {
}

// ReqEditTask Request entity for edit task
type ReqEditTask struct {
}

// Status

// ReqGetAccessStatus Request entity get array access status for user
type ReqGetAccessStatus struct {
}

// ReqEditStatusInTask Request entity edit status in task
type ReqEditStatusInTask struct {
}

// Severity

// ReqEditSeverityInTask Request entity for edit severity in task
type ReqEditSeverityInTask struct {
}

// Tags

// ReqGetAccessTags Request entity get array access tag for user
type ReqGetAccessTags struct {
}

// ReqAddTagInTask Request entity for add new tag in task
type ReqAddTagInTask struct {
}

// ReqDeleteTagInTask Request entity for delete tag in task by ID
type ReqDeleteTagInTask struct {
}

// DayUse

// ReqEditDayUse Request entity for edit date use task
type ReqEditDayUse struct {
}

//------------------------------------------------RESPONSE ENTITY-------------------------------------------------------

// ResFetchTaskID Request entity get task id with count filter
type ResFetchTaskID struct {
}

// Task in line

// ResFetchBaseTasks Request entity get base info tasks with count filter
type ResFetchBaseTasks struct {
}

// ResFetchBaseTaskByID Request entity get base info task by ID
type ResFetchBaseTaskByID struct {
}

// Task global edit

// ResFetchTask Request entity for get Task with count filter
type ResFetchTask struct {
}

// ResFetchTaskByID Request entity get task by ID
type ResFetchTaskByID struct {
}

// ResAddTask Request entity for add new task
type ResAddTask struct {
}

// ResDeleteTask Request entity for delete tasks by ID
type ResDeleteTask struct {
}

// ResEditTask Request entity for edit task
type ResEditTask struct {
}

// Status

// ResGetAccessStatus Request entity get array access status for user
type ResGetAccessStatus struct {
}

// ResEditStatusInTask Request entity edit status in task
type ResEditStatusInTask struct {
}

// Severity

// ResEditSeverityInTask Request entity for edit severity in task
type ResEditSeverityInTask struct {
}

// Tags

// ResGetAccessTags Request entity get array access tag for user
type ResGetAccessTags struct {
}

// ResAddTagInTask Request entity for add new tag in task
type ResAddTagInTask struct {
}

// ResDeleteTagInTask Request entity for delete tag in task by ID
type ResDeleteTagInTask struct {
}

// DayUse

// ResEditDayUse Request entity for edit date use task
type ResEditDayUse struct {
}
