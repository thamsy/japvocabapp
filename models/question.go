package models

import (
	"fmt"
	"strconv"
)

type Question struct {
	Row        int64
	Eng        string
	Kana       string
	Kanji      string
	Chi        string
	Grp        string
	CorrKana   int
	WrongKana  int
	CorrKanji  int
	WrongKanji int
	CorrGrp    int
	WrongGrp   int
	TotalTest  int
}

func ConvertValuesToQuestions(values [][]interface{}, offset int) []Question {
	questions := make([]Question, 0)
	for idx, rowDataRaw := range values {
		rowData := make([]interface{}, 12, 12)
		copy(rowData, rowDataRaw)
		corrKana, _ := strconv.Atoi(fmt.Sprintf("%v", rowData[5]))
		wrongKana, _ := strconv.Atoi(fmt.Sprintf("%v", rowData[6]))
		corrKanji, _ := strconv.Atoi(fmt.Sprintf("%v", rowData[7]))
		wrongKanji, _ := strconv.Atoi(fmt.Sprintf("%v", rowData[8]))
		corrGrp, _ := strconv.Atoi(fmt.Sprintf("%v", rowData[9]))
		wrongGrp, _ := strconv.Atoi(fmt.Sprintf("%v", rowData[10]))
		totalTest, _ := strconv.Atoi(fmt.Sprintf("%v", rowData[11]))
		question := Question{
			Row:        int64(idx + offset),
			Eng:        fmt.Sprintf("%v", rowData[0]),
			Kana:       fmt.Sprintf("%v", rowData[1]),
			Kanji:      fmt.Sprintf("%v", rowData[2]),
			Chi:        fmt.Sprintf("%v", rowData[3]),
			Grp:        fmt.Sprintf("%v", rowData[4]),
			CorrKana:   corrKana,
			WrongKana:  wrongKana,
			CorrKanji:  corrKanji,
			WrongKanji: wrongKanji,
			CorrGrp:    corrGrp,
			WrongGrp:   wrongGrp,
			TotalTest:  totalTest,
		}
		questions = append(questions, question)
	}
	return questions
}

func ConvertQuestionToValues(question *Question) [][]interface{} {
	res := [][]interface{}{{question.Eng, question.Kana, question.Kanji, question.Chi, question.Grp,
		strconv.Itoa(question.CorrKana), strconv.Itoa(question.WrongKana), strconv.Itoa(question.CorrKanji),
		strconv.Itoa(question.WrongKanji), strconv.Itoa(question.CorrGrp), strconv.Itoa(question.WrongGrp)}}
	return res
}