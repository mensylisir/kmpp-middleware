package entity

type PostgresLog struct {
	OperatorName      string `json:"operator_name"`
	OperatorNamespace string `json:"operator_namespace"`
	Instance
}

type LogFormat struct {
	Time  string `json:"time"`
	Level string `json:"level"`
	Msg   string `json:"msg"`
}
