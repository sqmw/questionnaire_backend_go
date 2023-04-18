package dao

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"questionnaireGo/src/entity"
)

//	fmt.Println(dao.AddAnswerResult(entity.AnswerResult{
//		ID:              "test",
//		Account:         "test",
//		QuestionnaireId: "test",
//		StartTime:       "2022-1-1",
//		StopTime:        "2022-1-1",
//		Content:         "test",
//		AnswererInfo: struct {
//			IP    string
//			Other string
//		}{IP: "ip", Other: "other"},
//	}))
//
// AddAnswerResult 添加一个回答的结果
func AddAnswerResult(answerResult entity.AnswerResult) bool {
	answererInfoJson, _ := json.Marshal(answerResult.AnswererInfo)
	result, err := GetDBConnect().Exec("insert into answer_result set id = ?, account = ?, questionnaire_id = ?,content = ?, start_time = ?, stop_time = ?, answerer_info = ?",
		uuid.NewString(),
		answerResult.Account,
		answerResult.QuestionnaireId,
		answerResult.Content,
		answerResult.StartTime,
		answerResult.StopTime,
		string(answererInfoJson))
	if err != nil {
		fmt.Println(err)
		return false
	}
	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return false
	}
	return true
}
