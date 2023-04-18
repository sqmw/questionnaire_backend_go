package dao

import (
	"database/sql"
	"fmt"
	"github.com/google/uuid"
	"questionnaireGo/src/entity"
)

// AddTemplate 添加模板
func AddTemplate(template entity.Template) bool {
	result, err := GetDBConnect().Exec("insert into template set id = ?, account = ?, content = ?, detail = ?, img = ?, time_use = ?",
		uuid.NewString(),
		template.Account,
		template.Content,
		template.Detail,
		template.Img,
		template.TimeUse)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return false
	}
	return true
}

// DelTemplateById 通过问卷的ID在数据库删除问卷
func DelTemplateById(id string) bool {
	result, err := GetDBConnect().Exec("delete from template where id = ?", id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if rowsAffected, _ := result.RowsAffected(); rowsAffected == 0 {
		return false
	}
	return true
}

// ModifyTemplateDetail 修改模板的详情
func ModifyTemplateDetail(id string, detail string) bool {
	var result sql.Result
	var err error
	var rowAffect int64
	result, err = GetDBConnect().Exec("update template set detail = ? where id = ?", detail, id)
	if err != nil {
		return false
	}
	if rowAffect, _ = result.RowsAffected(); rowAffect == 0 {
		return false
	}
	return true
}

// ModifyTemplateImg 修改模板的图片
func ModifyTemplateImg(id string, img string) bool {
	var result sql.Result
	var err error
	var rowAffect int64
	result, err = GetDBConnect().Exec("update template set img = ? where id = ?", img, id)
	if err != nil {
		return false
	}
	if rowAffect, _ = result.RowsAffected(); rowAffect == 0 {
		return false
	}
	return true
}

// ModifyTemplateContent 修改模板的content
func ModifyTemplateContent(id string, content string) bool {
	var result sql.Result
	var err error
	var rowAffect int64
	result, err = GetDBConnect().Exec("update template set content = ? where id = ?", content, id)
	if err != nil {
		fmt.Println(err)
		return false
	}
	if rowAffect, _ = result.RowsAffected(); rowAffect == 0 {
		return false
	}
	return true
}

// GetTemplateContent 获取模板的内容
func GetTemplateContent(id string) string {
	var err error
	// 查询字符串稍微有点不同
	var content = sql.NullString{}
	err = GetDBConnect().QueryRow("select content from template where id = ?", id).Scan(&content)
	if err != nil {
		return ""
	}
	return content.String
}
