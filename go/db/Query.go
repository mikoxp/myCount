package db

const (
	GetAllLogsSelect  = "SELECT id, info FROM public.info_log;"
	GetAllStepsSelect = "SELECT to_char(day, 'YYYY-MM-dd'), steps_number FROM numeration.steps;"
)
