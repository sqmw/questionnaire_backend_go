package main

import (
	"fmt"
	"questionnaireGo/src/dao"
)

func main() {
	fmt.Println(dao.DelQuestionnaireById("10"))
}
