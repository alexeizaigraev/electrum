// electrum project main.go
package main

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"strings"

	_ "github.com/lib/pq"
)

type site struct {
	department string
	edrpou     string
	address    string
	register   string
}

type department struct {
	department      string
	region          string
	district_region string
	district_city   string
	city_type       string
	city            string
	street          string
	street_type     string
	hous            string
	post_index      string
	partner         string
	status          string
	register        string
	edrpou          string
	address         string
	partner_name    string
	id_terminal     string
	koatu           string
	tax_id          string
	koatu2          string
}

type terminal struct {
	department       string
	termial          string
	model            string
	serial_number    string
	date_manufacture string
	soft             string
	producer         string
	rne_rro          string
	sealing          string
	fiscal_number    string
	oro_serial       string
	oro_number       string
	ticket_serial    string
	ticket_1sheet    string
	ticket_number    string
	sending          string
	books_arhiv      string
	tickets_arhiv    string
	to_rro           string
	owner_rro        string
	register         string
	finish           string
}

var (
	ConnStr = "user=postgres password=postgres dbname=drm_go sslmode=disable"
)

func InsertOtborOneRecord(vec []string) error {
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		return err
		//panic(err)
	}
	defer db.Close()

	sqlStatement := `
INSERT INTO otbor (term, dep)
VALUES ($1, $2)
RETURNING term`

	term := ""
	err = db.QueryRow(sqlStatement, vec[0], vec[1]).Scan(&term)
	if err != nil {
		//panic(err)
		return err
	}
	//fmt.Println("New record term is:", term)
	return nil
}

func InsertDepartmentsOneRecord(vec []string) error {
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		return err
		//panic(err)
	}
	defer db.Close()

	sqlStatement := `
INSERT INTO departments (department, region, district_region, district_city, city_type, city, street, street_type, hous, post_index, partner, status, register, edrpou, address, partner_name, id_terminal, koatu, tax_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)
RETURNING department`

	department := ""
	err = db.QueryRow(sqlStatement, vec[0], vec[1], vec[2], vec[3], vec[4], vec[5], vec[6], vec[7], vec[8], vec[9], vec[10], vec[11], vec[12], vec[13], vec[14], vec[15], vec[16], vec[17], vec[18]).Scan(&department)
	if err != nil {
		//panic(err)
		return err
	}
	fmt.Println(department)
	return nil
}

func InsertTerminalsOneRecord(vec []string) error {
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		fmt.Println(err)
		//return err
		panic(err)
	}
	defer db.Close()

	sqlStatement := `
INSERT INTO terminals (department, termial, model, serial_number, date_manufacture, soft, producer, rne_rro, sealing, fiscal_number, oro_serial, oro_number, ticket_serial, ticket_1sheet, ticket_number, sending, books_arhiv, tickets_arhiv, to_rro, owner_rro, register, finish)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20)
RETURNING termial`

	termial := ""
	//--
	err = db.QueryRow(sqlStatement, vec[0], vec[1], vec[2], vec[3], vec[4], vec[5], vec[6], vec[7], vec[8], vec[9], vec[10], vec[11], vec[12], vec[13], vec[14], vec[15], vec[16], vec[17], vec[18], vec[19], vec[20], "").Scan(&termial)
	//err = db.QueryRow(sqlStatement, vec[0], vec[1]).Scan(&termial)
	if err != nil {
		//panic(err)
		return err
	}

	//fmt.Println(termial)
	return nil
}

func InsertDepartmentsFromFile() string {
	info := ""
	count := 0
	countErr := 0
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		fmt.Println(err)
		//return err
		panic(err)
	}
	defer db.Close()

	sqlStatement := `
INSERT INTO departments (department, region, district_region, district_city, city_type, city, street, street_type, hous, post_index, partner, status, register, edrpou, address, partner_name, id_terminal, koatu, tax_id)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19)
RETURNING department`

	department := ""

	//fmt.Println("InsertDepartmentsFromFile...")
	arr := FileToArr(DataInDir + "departments.csv")[1:]
	sizeVec := len(arr[0])
	for _, vec := range arr {
		if len(vec) == sizeVec {
			err = db.QueryRow(sqlStatement, vec[0], vec[1], vec[2], vec[3], vec[4], vec[5], vec[6], vec[7], vec[8], vec[9], vec[10], vec[11], vec[12], vec[13], vec[14], vec[15], vec[16], vec[17], vec[18]).Scan(&department)
			if err != nil {
				countErr += 1
				fmt.Println("err", vec[0])
				continue
			} else {
				count += 1
			}
			//fmt.Println(department)
		} else {
			countErr += 1
			fmt.Println("bed vec size", vec[1])
		}
	}
	if countErr == 0 {
		info += fmt.Sprintf("success dep %d", count)
	} else {
		info += fmt.Sprintf(">> dep err: %d ok: %d", countErr, count)
	}
	return info
}

func InsertTerminalsFromFile() string {
	info := ""
	count := 0
	count_err := 0
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		fmt.Println(err)
		//return err
		panic(err)
	}
	defer db.Close()

	sqlStatement := `
INSERT INTO terminals (department, termial, model, serial_number, date_manufacture, soft, producer, rne_rro, sealing, fiscal_number, oro_serial, oro_number, ticket_serial, ticket_1sheet, ticket_number, sending, books_arhiv, tickets_arhiv, to_rro, owner_rro, register, finish)
VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9, $10, $11, $12, $13, $14, $15, $16, $17, $18, $19, $20, $21, $22)
RETURNING termial`

	termial := ""

	//fmt.Println("InsertTerminalsFromFile...")
	arr := FileToArr(DataInDir + "terminals.csv")[1:]
	sizeVec := len(arr[0])
	for _, vec := range arr {
		if true || len(vec) == sizeVec {
			err = db.QueryRow(sqlStatement, vec[0], vec[1], vec[2], vec[3], vec[4], vec[5], vec[6], vec[7], vec[8], vec[9], vec[10], vec[11], vec[12], vec[13], vec[14], vec[15], vec[16], vec[17], vec[18], vec[19], "", "").Scan(&termial)

			if err != nil {
				count_err += 1
				//fmt.Println("err", vec[1])
				continue
			} else {
				count += 1
			}
			//fmt.Println(termial)
		} else {
			//countErr += 1
			//fmt.Println("bed vec size", vec[1])
		}
	}
	if count_err == 0 {
		info += fmt.Sprintf("success term %d", count)
	} else {
		info += fmt.Sprintf(">> term err: %d ok: %d", count_err, count)
	}
	return info
}

func InsertTerminalsFromFile0() {
	ok := true
	fmt.Println("InsertTerminalsFromFile...")
	arr := FileToArr(DataInDir + "terminals.csv")[1:]
	for _, vec := range arr {
		if len(vec) == 20 {
			err := InsertTerminalsOneRecord(vec)
			if err != nil {
				ok = false
				fmt.Println(">> err insert terminals", vec[1])
			}
		} else {
			fmt.Println("bed vec size", vec[1])
		}
	}
	if ok {
		fmt.Println("success")
	} else {
		fmt.Println("errors")
	}
}

func InsertOtborFromFile() string {
	info := ""
	count := 0
	count_err := 0
	ok := true
	fmt.Println("\n\tInsertOtborFromFile...")
	arr := FileToArr(DataInDir + "otbor.csv")[1:]
	for _, vec := range arr {
		err := InsertOtborOneRecord(vec)
		count += 1
		if err != nil {
			ok = false
			count_err += 1
			fmt.Println(">> err insert otbor", vec)
		}
	}
	if ok {
		info += fmt.Sprintf("success otbor %d", count)
		//fmt.Println("success")
	} else {
		info += fmt.Sprintf(">> otbor err: %d, ok: %d", count_err, count)
		//fmt.Println("errors")
	}
	return info
}

func ClearTableOtbor() error {
	db, err := sql.Open("postgres", ConnStr)
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

func ClearTableDepartments() {
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		//return err
		panic(err)
	}
	defer db.Close()
	_, err = db.Query("DELETE FROM departments;")
	if err != nil {
		panic(err)
		//return err
	}
	fmt.Println("clear table departments")
	//return nil
}

func ClearTableTerminals() {
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		//return err
		panic(err)
	}
	defer db.Close()
	_, err = db.Query("DELETE FROM terminals;")
	if err != nil {
		panic(err)
		//return err
	}
	fmt.Println("clear table terminals")
	//return nil
}

func SelectAllDepartments() {
	//fmt.Println("SelectAllDepartments...")
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from departments order by department")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	//departments := []department{}
	outText := "department;region;district_region;district_city;city_type;city;street;street_type;hous;post_index;partner;status;register;edrpou;address;partner_name;id_terminal;koatu;tax_id;koatu2\n"
	for rows.Next() {
		p := department{}
		//err := rows.Scan(&p.department, &p.termial, &p.model, &p.serial_number, &p.date_manufacture, &p.soft, &p.producer, &p.rne_rro, &p.sealing, &p.fiscal_number, &p.oro_serial, &p.oro_number, &p.ticket_serial, &p.ticket_1sheet, &p.ticket_number, &p.sending, &p.books_arhiv, &p.tickets_arhiv, &p.to_rro, &p.owner_rro)
		err := rows.Scan(&p.department, &p.region, &p.district_region, &p.district_city, &p.city_type, &p.city, &p.street, &p.street_type, &p.hous, &p.post_index, &p.partner, &p.status, &p.register, &p.edrpou, &p.address, &p.partner_name, &p.id_terminal, &p.koatu, &p.tax_id, &p.koatu2)
		if err != nil {
			fmt.Println(err)
			continue
		}
		//departments = append(departments, p)
		//outText += p.department + ";" + p.termial + ";" + p.model + ";" + p.serial_number + ";" + p.date_manufacture + ";" + p.soft + ";" + p.producer + ";" + p.rne_rro + ";" + p.sealing + ";" + p.fiscal_number + ";" + p.oro_serial + ";" + p.oro_number + ";" + p.ticket_serial + ";" + p.ticket_1sheet + ";" + p.ticket_number + ";" + p.sending + ";" + p.books_arhiv + ";" + p.tickets_arhiv + ";" + p.to_rro + ";" + p.owner_rro
		outText += p.department + ";" + p.region + ";" + p.district_region + ";" + p.district_city + ";" + p.city_type + ";" + p.city + ";" + p.street + ";" + p.street_type + ";" + p.hous + ";" + p.post_index + ";" + p.partner + ";" + p.status + ";" + p.register + ";" + p.edrpou + ";" + p.address + ";" + p.partner_name + ";" + p.id_terminal + ";" + p.koatu + ";" + p.tax_id + ";" + p.koatu2 + "\n"
		//fmt.Println(p.department)
	}
	TextToFile(outText, DataInDir+"pg_departments.csv")
}

/*
func SelectAllDepartments() {
	fmt.Println("SelectAllDepartments...")
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from departments")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	//departments := []department{}
	outText := "department;region;district_region;district_city;city_type;city;street;street_type;hous;post_index;partner;status;register;edrpou;address;partner_name;id_terminal;koatu;tax_id\n"
	for rows.Next() {
		p := department{}
		//err := rows.Scan(&p.department, &p.termial, &p.model, &p.serial_number, &p.date_manufacture, &p.soft, &p.producer, &p.rne_rro, &p.sealing, &p.fiscal_number, &p.oro_serial, &p.oro_number, &p.ticket_serial, &p.ticket_1sheet, &p.ticket_number, &p.sending, &p.books_arhiv, &p.tickets_arhiv, &p.to_rro, &p.owner_rro)
		err := rows.Scan(&p.department, &p.region, &p.district_region, &p.district_city, &p.city_type, &p.city, &p.street, &p.street_type, &p.hous, &p.post_index, &p.partner, &p.status, &p.register, &p.edrpou, &p.address, &p.partner_name, &p.id_terminal, &p.koatu, &p.tax_id)
		if err != nil {
			fmt.Println(err)
			continue
		}
		//departments = append(departments, p)
		//outText += p.department + ";" + p.termial + ";" + p.model + ";" + p.serial_number + ";" + p.date_manufacture + ";" + p.soft + ";" + p.producer + ";" + p.rne_rro + ";" + p.sealing + ";" + p.fiscal_number + ";" + p.oro_serial + ";" + p.oro_number + ";" + p.ticket_serial + ";" + p.ticket_1sheet + ";" + p.ticket_number + ";" + p.sending + ";" + p.books_arhiv + ";" + p.tickets_arhiv + ";" + p.to_rro + ";" + p.owner_rro
		outText += p.department + ";" + p.region + ";" + p.district_region + ";" + p.district_city + ";" + p.city_type + ";" + p.city + ";" + p.street + ";" + p.street_type + ";" + p.hous + ";" + p.post_index + ";" + p.partner + ";" + p.status + ";" + p.register + ";" + p.edrpou + ";" + p.address + ";" + p.partner_name + ";" + p.id_terminal + ";" + p.koatu + ";" + p.tax_id + "\n"
		//fmt.Println(p.department)
	}
	TextToFile(outText, DataInDir+"pg_departments.csv")
}
*/
func DepsToArr() [][]string {
	var arr [][]string
	vec := make([]string, 19)
	fmt.Println("SelectAllDepartments...")
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from departments order by department")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	//departments := []department{}
	//outText := "department;region;district_region;district_city;city_type;city;street;street_type;hous;post_index;partner;status;register;edrpou;address;partner_name;id_terminal;koatu;tax_id\n"
	for rows.Next() {
		p := department{}
		//err := rows.Scan(&p.department, &p.termial, &p.model, &p.serial_number, &p.date_manufacture, &p.soft, &p.producer, &p.rne_rro, &p.sealing, &p.fiscal_number, &p.oro_serial, &p.oro_number, &p.ticket_serial, &p.ticket_1sheet, &p.ticket_number, &p.sending, &p.books_arhiv, &p.tickets_arhiv, &p.to_rro, &p.owner_rro)
		err := rows.Scan(&p.department, &p.region, &p.district_region, &p.district_city, &p.city_type, &p.city, &p.street, &p.street_type, &p.hous, &p.post_index, &p.partner, &p.status, &p.register, &p.edrpou, &p.address, &p.partner_name, &p.id_terminal, &p.koatu, &p.tax_id, &p.koatu2)
		if err != nil {
			fmt.Println(err)
			continue
		}
		if p.department == "1700999" {
			continue
		}
		//departments = append(departments, p)
		//outText += p.department + ";" + p.termial + ";" + p.model + ";" + p.serial_number + ";" + p.date_manufacture + ";" + p.soft + ";" + p.producer + ";" + p.rne_rro + ";" + p.sealing + ";" + p.fiscal_number + ";" + p.oro_serial + ";" + p.oro_number + ";" + p.ticket_serial + ";" + p.ticket_1sheet + ";" + p.ticket_number + ";" + p.sending + ";" + p.books_arhiv + ";" + p.tickets_arhiv + ";" + p.to_rro + ";" + p.owner_rro
		//outText += p.department + ";" + p.region + ";" + p.district_region + ";" + p.district_city + ";" + p.city_type + ";" + p.city + ";" + p.street + ";" + p.street_type + ";" + p.hous + ";" + p.post_index + ";" + p.partner + ";" + p.status + ";" + p.register + ";" + p.edrpou + ";" + p.address + ";" + p.partner_name + ";" + p.id_terminal + ";" + p.koatu + ";" + p.tax_id + "\n"
		//fmt.Println(p.department)
		vec[0] = p.department
		vec[1] = p.region
		vec[2] = p.district_region
		vec[3] = p.district_city
		vec[4] = p.city_type
		vec[5] = p.city
		vec[6] = p.street
		vec[7] = p.street_type
		vec[8] = p.hous
		vec[9] = p.post_index
		vec[10] = p.partner
		vec[11] = p.status
		vec[12] = p.register
		vec[13] = p.edrpou
		vec[14] = p.address
		vec[15] = p.partner_name
		vec[16] = p.id_terminal
		vec[17] = p.koatu
		vec[18] = p.tax_id
		vec[19] = p.koatu2

		arr = append(arr, vec)
		//fmt.Println(arr)
	}
	//TextToFile(outText, DataInDir+"pg_departments.csv")
	return arr
}

func SelectAllTerminals() {
	//fmt.Println("SelectAllTerminals...")
	db, err := sql.Open("postgres", ConnStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	rows, err := db.Query("select * from terminals order by termial")
	if err != nil {
		panic(err)
	}
	defer rows.Close()
	//terminals := []department{}
	outText := "department;termial;model;serial_number;date_manufacture;soft;producer;rne_rro;sealing;fiscal_number;oro_serial;oro_number;ticket_serial;ticket_1sheet;ticket_number;sending;books_arhiv;tickets_arhiv;to_rro;owner_rro\n"
	for rows.Next() {
		p := terminal{}
		err := rows.Scan(&p.department, &p.termial, &p.model, &p.serial_number, &p.date_manufacture, &p.soft, &p.producer, &p.rne_rro, &p.sealing, &p.fiscal_number, &p.oro_serial, &p.oro_number, &p.ticket_serial, &p.ticket_1sheet, &p.ticket_number, &p.sending, &p.books_arhiv, &p.tickets_arhiv, &p.to_rro, &p.owner_rro, &p.register, &p.finish)
		//err := rows.Scan(&p.department, &p.region, &p.district_region, &p.district_city, &p.city_type, &p.city, &p.street, &p.street_type, &p.hous, &p.post_index, &p.partner, &p.status, &p.register, &p.edrpou, &p.address, &p.partner_name, &p.id_terminal, &p.koatu, &p.tax_id)
		if err != nil {
			fmt.Println(err)
			continue
		}
		//departments = append(departments, p)
		outText += p.department + ";" + p.termial + ";" + p.model + ";" + p.serial_number + ";" + p.date_manufacture + ";" + p.soft + ";" + p.producer + ";" + p.rne_rro + ";" + p.sealing + ";" + p.fiscal_number + ";" + p.oro_serial + ";" + p.oro_number + ";" + p.ticket_serial + ";" + p.ticket_1sheet + ";" + p.ticket_number + ";" + p.sending + ";" + p.books_arhiv + ";" + p.tickets_arhiv + ";" + p.to_rro + ";" + p.owner_rro + ";" + p.register + ";" + p.finish + "\n"
		//outText += p.department + ";" + p.region + ";" + p.district_region + ";" + p.district_city + ";" + p.city_type + ";" + p.city + ";" + p.street + ";" + p.street_type + ";" + p.hous + ";" + p.post_index + ";" + p.partner + ";" + p.status + ";" + p.register + ";" + p.edrpou + ";" + p.address + ";" + p.partner_name + ";" + p.id_terminal + ";" + p.koatu + ";" + p.tax_id + "\n"
		//fmt.Println(p.department)
	}
	TextToFile(outText, DataInDir+"pg_terminals.csv")
}

func MkVsyoZapros() {
	dede := FileToArr(DataInDir + "pg_departments.csv")
	deps := dede[1:]
	depsHead := "departments." + strings.Join(dede[0], ";")

	tete := FileToArr(DataInDir + "pg_terminals.csv")
	terms := tete[1:]
	termsHead := "terminals." + strings.Join(tete[0], ";")
	//outText := "terminals.department;termial;model;serial_number;date_manufacture;soft;producer;rne_rro;sealing;fiscal_number;oro_serial;oro_number;ticket_serial;ticket_1sheet;ticket_number;sending;books_arhiv;tickets_arhiv;to_rro;owner_rro;departments.department;region;district_region;city_type;city;district_city;street_type;street;hous;post_index;partner;status;register;edrpou;address;partner name;id_terminal;koatu;tax_id\n"
	outText := termsHead + ";" + depsHead + "\n"
	for _, term := range terms {
		termLine := strings.Join(term, ";")
		for _, dep := range deps {
			if term[0] == dep[0] {
				depLine := strings.Join(dep, ";")
				outText += termLine + ";" + depLine + "\n"
			}
		}
	}
	TextToFile(outText, DataInDir+"pg_vsyo_zapros.csv")
}

func MkOtbor() {
	fmt.Print("From -> ")
	inputStart := ""
	fmt.Fscan(os.Stdin, &inputStart)
	//fmt.Println()

	fmt.Print("Qual -> ")
	inputQual := ""
	fmt.Fscan(os.Stdin, &inputQual)

	start, _ := strconv.Atoi(inputStart)
	qual, _ := strconv.Atoi(inputQual)

	otborFunc(start, qual)

}

func otborFunc(start int, qual int) {
	head := "term;dep\n"
	outText := head
	for x := start; x < start+qual; x++ {
		dep := strconv.Itoa(x)
		term := strconv.Itoa(x*10 + 1)
		outText += term + ";" + dep + "\n"
		fmt.Println(term, dep)
	}
	TextToFile(outText, DataInDir+"otbor.csv")
}

func ZaprosOtbor() {
	obobo := FileToArr(DataInDir + "otbor.csv")
	otbor := obobo[1:]
	vsyooo, _ := FileToVec(DataInDir + "pg_vsyo_zapros.csv")
	vsyo := vsyooo[1:]

	outText := vsyooo[0] + "\n"
	for _, otborLine := range otbor {
		for _, vsyoLine := range vsyo {
			term := strings.Split(vsyoLine, ";")[1]
			if otborLine[0] == term {
				outText += vsyoLine + "\n"
			}
		}
	}
	TextToFile(outText, DataInDir+"pg_vsyo_zapros_vnesh_otbor.csv")
}

func RefreshAll() {

	ClearTableDepartments()
	InsertDepartmentsFromFile()

	ClearTableTerminals()
	InsertTerminalsFromFile()

	ClearTableOtbor()
	InsertOtborFromFile()

	//SelectAllDepartments()
	//SelectAllTerminals()
	//MkVsyoZapros()
	fmt.Println("\n\tall success refresged")
}
