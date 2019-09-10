package sheetsapi

import (
	"fmt"
	"google.golang.org/api/sheets/v4"
	"japapp/models"
	"log"
	"strconv"
)

const spreadSheetId = "1-AakxLC0AO_eYlaXN0NReDDrkaxu6-D644wHdJ6Cijk"
const vocabCountCell = "A1"
const vocabOffset = 3

func GetAllNewQuestions() []models.Question {
	// Get number of available questions
	resp, err := srv.Spreadsheets.Values.Get(spreadSheetId, "General Vocab!" + vocabCountCell + ":" + vocabCountCell).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet. %v", err)
	}

	// Retrieve all questions
	lastRow := fmt.Sprintf("%v", resp.Values[0][0])
	readRange := "General Vocab!A3:L" + lastRow
	resp, err = srv.Spreadsheets.Values.Get(spreadSheetId, readRange).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet. %v", err)
	}
	log.Printf("Sheets: GetAll successful")
	return models.ConvertValuesToQuestions(resp.Values, vocabOffset)
}

func IncrNewQuestion(row string, isKanaCorrect bool, isKanjiCorrect bool, isGrpCorrect bool) {
	// retrieve new question
	resp, err := srv.Spreadsheets.Values.Get(spreadSheetId, "General Vocab!A" + row + ":L" + row).Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet. %v", err)
	}
	offset, _ := strconv.Atoi(row)

	question := models.ConvertValuesToQuestions(resp.Values, offset)[0]

	// incr only corr answer if all corr
	if isKanaCorrect && isKanjiCorrect && isGrpCorrect {
		question.CorrKana += 1
		question.CorrKanji += 1
		question.CorrGrp += 1
		question.TotalTest += 1
	} else {
		if !isKanaCorrect {question.WrongKana +=1}
		if !isKanjiCorrect {question.WrongKanji +=1}
		if !isGrpCorrect {question.WrongGrp +=1}
	}

	// update accordingly
	updateNewQuestion(&question)
}

func CheckUnderstoodNewQuestions() {
	questions := GetAllNewQuestions()
	for i := len(questions)-1; i >= 0; i-- {
		question := questions[i]
		if checkNewQuestionUnderstood(&question) {
			moveNewQuestion(&question)
		}
	}
}

func checkNewQuestionUnderstood (question *models.Question) bool{
	if question.TotalTest >= 8 &&
		float32(question.CorrKana) / float32(question.TotalTest) > 0.7 &&
		float32(question.CorrKanji) / float32(question.TotalTest) > 0.7 &&
		float32(question.CorrGrp) / float32(question.TotalTest) > 0.7 {
		return true
	} else {
		return false
	}
}

func moveNewQuestion(question *models.Question) {
	rb := &sheets.ValueRange{
		Values: models.ConvertQuestionToValues(question),
	}
	_, err := srv.Spreadsheets.Values.Append(spreadSheetId, "Understood Vocab!A3:K1000", rb).ValueInputOption("USER_ENTERED").InsertDataOption("OVERWRITE").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet. %v", err)
	}
	log.Printf("Sheets: Append successful")

	deleteRequest := sheets.DeleteDimensionRequest{
		Range: &sheets.DimensionRange{
			Dimension:       "ROWS",
			EndIndex:        question.Row,
			SheetId:         0,
			StartIndex:      question.Row - 1, // First row is index 0
		},
	}
	requests := []*sheets.Request{
		{DeleteDimension: &deleteRequest},
	}
	rb3 := &sheets.BatchUpdateSpreadsheetRequest{
		Requests: requests,
	}
	_, err = srv.Spreadsheets.BatchUpdate(spreadSheetId, rb3).Do()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Sheets: Delete Row successful")
}

func updateNewQuestion(question *models.Question) {
	rb := &sheets.ValueRange{
		Values: models.ConvertQuestionToValues(question),
	}
	rowStr := strconv.FormatInt(question.Row, 10)
	_, err := srv.Spreadsheets.Values.Update(spreadSheetId, "General Vocab!A" + rowStr + ":K" + rowStr, rb).ValueInputOption("USER_ENTERED").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve data from sheet. %v", err)
	}
	log.Printf("Sheets: UpdatedNewQuestion call successful")
}

