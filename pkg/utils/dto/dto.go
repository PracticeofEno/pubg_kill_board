package dto

type PercentData struct {
	Count   int
	Percent int
}

type PercentReqeust struct {
	PercentData  []PercentData
	RandomString string
}

type UpdateUserData struct {
	RandomString string        `json:"random_string" binding:"required"`
	PercentData  []PercentData `json:"percent_data" binding:"required"`
	Nickname     string        `json:"nickname" binding:"required"`
	TargetKill   int           `json:"target_kill"`
	CurrentKill  int           `json:"current_kill"`
}