package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)
 
func getSQLStatement() {
	file, err := os.Open("test.sql")
 
	if err != nil {
		log.Fatalf("failed opening file: %s", err)
	}
 
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	var ksqlStatement []string
 
	for scanner.Scan() {
		ksqlStatement = append(ksqlStatement, scanner.Text())
	}
 
	file.Close()
 
	for _, eachline := range ksqlStatement {
		SendSQL(eachline)
	}
}

// SendSQL - send sql statement
func SendSQL(sqlStatement string)  {
	requestBody, err := json.Marshal(map[string]string{
		"ksql": sqlStatement,
	})

	if err != nil{
		log.Fatalln("request ERR",err)
	}

	resp, err := http.Post("http://localhost:8088/ksql","application/vnd.ksql.v1+json",bytes.NewBuffer(requestBody))
	if err != nil{
		log.Fatalln("resp ERR:",err)
	}

	defer resp.Body.Close()

	body,err := ioutil.ReadAll(resp.Body)
	if err != nil{
		log.Fatalln("body ERR:",err)
	}

	// 去除Body回傳的前後大括號，在轉換型態時比較方便
	newBody := strings.Trim(string(body),"[")
	newBody = strings.Trim(newBody,"]")

	respContent := strings.Split(string(body),"\"")
	DetermineType(respContent[3],newBody)
}

// DetermineType - 判斷回傳的型態(@type)，針對不同型態給予不同的type
func DetermineType(whichType string, body string)  {
	body = strings.Replace(body,"@type","currentType",1)

	if whichType == "currentStatus" {
		// 創建Table或刪除
		data := CurrentStatus{}
		err := json.Unmarshal([]byte(body), &data)
		if err != nil {log.Println(err)}

		log.Println("使用者執行指令:",data.StatementText)
		log.Println("指令執行狀態:",data.CommandStatus.Status)
		log.Println("KSQL Server回傳訊息:",data.CommandStatus.Message)
		log.Println("---------------")


	}else if  whichType == "queries"{
		// 正在執行的queries執行
		data := Queries{}
		err := json.Unmarshal([]byte(body), &data)
		if err != nil {log.Println("err1:",err)}

		log.Println("使用者執行指令:",data.StatementText)

		getID := strings.Split(body,"\"id\":")
		getID = strings.Split(getID[1],",\"state\"")
		ID := strings.Trim(getID[0],"\"")
		log.Println("ID為：",ID)
		// log.Println("刪除Query中…")
		// SendSQL("TERMINATE "+ID+";")
		log.Println()
	}
}

// CurrentStatus - CurrentStatus
type CurrentStatus struct{
	CurrentType string `json:"currentType"`
	StatementText string `json:"statementText"`
	CommandID string `json:"commandID"`
	CommandStatus CommandStatus `json:"commandStatus"`
	CommandSequenceNumber int64 `json:"commandSequenceNumber"`
	Warnings []string `json:"warnings"`
}

// CommandStatus - CommandStatus
type CommandStatus struct{
	Status string `json:"status"`
	Message string `json:"message"`
}

// Queries - Queries
type Queries struct{
	CurrentType string `json:"currentType"`
	StatementText string `json:"statementText"`
	SubQueries SubQuery `json:"SubQueries"`
	Warnings []string `json:"warnings"`
}

// SubQuery - SubQuery
type SubQuery struct{
	QueryString string `json:"QueryString"`
	Sinks []string `json:"Sinks"`
	SinkKafkaTopics []string `json:"SinkKafkaTopics"`
	ID string `json:"ID"`
	State string `json:"State"`
}