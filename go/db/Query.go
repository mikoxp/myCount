package db

const (
	GetAllLogsSelect = "SELECT id, info FROM public.info_log;"
	GetAllStepsSelect = "SELECT day, steps_number FROM numeration.steps;"
)
