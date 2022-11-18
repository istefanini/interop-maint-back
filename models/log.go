package models

type Log struct {
	EventID    string  `json:"EventID"`
	SysFechaC  string  `json:"sysFechaC"`
	Estado     float64 `json:"Estado"`
	LogProceso string  `json:"LogProceso"`
	MsgFinal   string  `json:"MsgFinal"`
}
