package types

type SplatoonSchedule struct {
	Data struct {
		TurfWar RegularSchedule `json:"regularSchedules"`
		Anarchy BankaraSchedule `json:"bankaraSchedules"`
		XBattle XSchedule       `json:"xSchedules"`
	} `json:"data"`
}

type RegularSchedule struct {
	Nodes []struct {
		Node
		Setting Setting `json:"regularMatchSetting"`
	} `json:"nodes"`
}

type BankaraSchedule struct {
	Nodes []struct {
		Node
		Setting []Setting `json:"bankaraMatchSettings"`
	} `json:"nodes"`
}

type XSchedule struct {
	Nodes []struct {
		Node
		Setting Setting `json:"xMatchSetting"`
	} `json:"nodes"`
}

type Node struct {
	StartTime string `json:"startTime"`
	EndTime   string `json:"endTime"`
}

type Setting struct {
	IsVsSetting string  `json:"__isVsSetting"`
	TypeName    string  `json:"__typename"`
	Stages      []Stage `json:"vsStages"`
	VsRule      Rule    `json:"vsRule"`
}

type Stage struct {
	StageId int    `json:"vsStageId"`
	Name    string `json:"name"`
	Image   Image  `json:"image"`
	Id      string `json:"id"`
}

type Image struct {
	Url string `json:"url"`
}

type Rule struct {
	Name string `json:"name"`
	Rule string `json:"rule"`
	Id   string `json:"id"`
}
