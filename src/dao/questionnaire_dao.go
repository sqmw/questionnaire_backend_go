package dao

import (
	"database/sql"
	"github.com/google/uuid"
)

// AddQuestionnaire 根据本问卷内容创建的account添加新的we暖
func AddQuestionnaire(account string, content string) bool {
	result, err := GetDBConnect().Exec("insert into questionnaire set id = ?, account = ?, content = ?, status = 1", uuid.NewString(), account, content)
	if err != nil {
		return false
	}
	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return false
	}
	return true
}

// DelQuestionnaireById 通过问卷的ID在数据库删除问卷
func DelQuestionnaireById(id string) bool {
	result, err := GetDBConnect().Exec("delete from questionnaire where id = ?", id)
	if err != nil {
		return false
	}
	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return false
	}
	// 删除对应的问卷收集的结果，这里外键对效率有影响，因此不是用户外键
	result, err = GetDBConnect().Exec("delete from answer_result where questionnaire_id = ?", id)
	if err != nil {
		return false
	}
	return true
}

// ModifyQuestionnaireStatus 修改问卷的状态
func ModifyQuestionnaireStatus(id string, status int8) bool {
	var result sql.Result
	var err error
	var rowAffect int64
	result, err = GetDBConnect().Exec("update questionnaire set status = ? where id = ?", status, id)
	if err != nil {
		return false
	}
	if rowAffect, _ = result.RowsAffected(); rowAffect == 0 {
		return false
	}
	return true
}

// ModifyQuestionnaireStartTime 修改问卷的开始时间
func ModifyQuestionnaireStartTime(id string, startTime string) bool {
	var result sql.Result
	var err error
	var rowAffect int64
	result, err = GetDBConnect().Exec("update questionnaire set start_time = ? where id = ?", startTime, id)
	if err != nil {
		return false
	}
	if rowAffect, _ = result.RowsAffected(); rowAffect == 0 {
		return false
	}
	return true
}

// ModifyQuestionnaireStopTime 修改问卷的开始时间
func ModifyQuestionnaireStopTime(id string, stopTime string) bool {
	var result sql.Result
	var err error
	var rowAffect int64
	result, err = GetDBConnect().Exec("update questionnaire set stop_time = ? where id = ?", stopTime, id)
	if err != nil {
		return false
	}
	if rowAffect, _ = result.RowsAffected(); rowAffect == 0 {
		return false
	}
	return true
}

// GetQuestionnaireContent 获取问卷的内容
func GetQuestionnaireContent(id string) string {
	var err error
	// 查询字符串稍微有点不同
	var content = sql.NullString{}
	err = GetDBConnect().QueryRow("select content from questionnaire where id = ?", id).Scan(&content)
	if err != nil {
		return ""
	}
	return content.String
}
