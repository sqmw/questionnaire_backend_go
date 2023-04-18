package entity

// ResInfo 表示的是返回的信息
type ResInfo = struct {
	status  int8
	message string
	data    string
}

// User 用户的实体类，admin也是其中的一部分，根据account是否包含使用admin作为前缀来进行区分
type User = struct {
	Email    string
	Account  string
	Nickname string
	Password string
	Status   int8
}

// Questionnaire 问卷的实体类
type Questionnaire = struct {
	ID        string `db:"id"`
	Content   string `db:"content"`
	Account   string `db:"account"`
	Status    int8   `db:"status"`
	StartTime string `db:"start_time"`
	StopTime  string `db:"stop_time"`
}

// Template 模板的实体类
type Template = struct {
	ID      string `bd:"id"`
	Account string `db:"template"`
	TimeUse uint   `db:"time_use"`
	Content string `db:"content"`
	Detail  string `db:"detail"`
	Img     string `db:"img"`
}

// AnswerResult 答卷结果的实体类
type AnswerResult = struct {
	ID              string `db:"id"`
	Account         string `db:"account"`
	QuestionnaireId string `db:"questionnaire_id"`
	StartTime       string `db:"start_time"`
	StopTime        string `db:"stop_time"`
	Content         string `db:"content"`

	// JSON序列化以及反序列化的时候，这种套娃怎么处理
	AnswererInfo struct {
		IP    string
		Other string
	} `db:"answerer_info"`
}
