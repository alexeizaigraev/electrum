package control_otbor

import (
	"database/sql"
	"electrum/db"
	"electrum/module"
	"fmt"
	"strings"
)

func OtborTextTerm() module.Out5 {
	ClearTableOtbor()
	info := ""
	count := 0
	count_err := 0
	ok := true
	term := ""
	dep := ""
	db, err := sql.Open("postgres", db.ConnStr)
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	arr, _ := module.FileToVec(module.DataInDir + "otbor_term.csv")

	for _, term0 := range arr {

		if term0 == "" || term0 == "\n" {
			continue
		}
		term = strings.Trim(term0, " ")
		term = strings.Trim(term, "\n")
		dep = string([]rune(term)[:7])

		sqlStatement := `
INSERT INTO otbor (term, dep)
VALUES ($1, $2)
RETURNING term`

		myTerm := ""
		err = db.QueryRow(sqlStatement, term, dep).Scan(&myTerm)
		if err != nil {
			fmt.Println(err)
			ok = false
			count_err += 1
		}
		count += 1
		fmt.Println(term, dep)
	}

	if ok {
		info += fmt.Sprintf("success otbor %d", count)
		//fmt.Println("success")
	} else {
		info += fmt.Sprintf(">> otbor err: %d, ok: %d", count_err, count)
		//fmt.Println("errors")
	}
	return module.Out5{
		Kind: "Отбор терминалы",
		//Data:  strings.Split(Unfind+OutText, "/n"),
		//Fname: outFname,
		//Err:   "",
		Rez: info,
	}
}

func ClearTableOtbor() error {
	db, err := sql.Open("postgres", db.ConnStr)
	if err != nil {
		return err
		//panic(err)
	}
	defer db.Close()
	_, err = db.Query("DELETE FROM otbor;")
	if err != nil {
		return err
	}
	fmt.Println("clear table otbor")
	return nil
}
