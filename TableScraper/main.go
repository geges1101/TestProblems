package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/tanaikech/go-gdoctableapp"
	"net/http"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	url := "https://confluence.hflabs.ru/pages/viewpage.action?pageId=1181220999"

	response, err := http.Get(url)
	defer response.Body.Close()
	checkError(err)

	if response.StatusCode > 400 {
		fmt.Println("Status code:", response.StatusCode)
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	checkError(err)

	var table [][]interface{}
	var tableS [][]string
	doc.Find("div.table-wrap").Find("tr").Each(func(i int, item *goquery.Selection) {
		var row []interface{}
		var rowS []string
		item.Find("th").Each(func(j int, curr *goquery.Selection) {
			row = append(row, curr.Text())
			rowS = append(rowS, curr.Text())
		})
		table = append(table, row)
		tableS = append(tableS, rowS)
	})

	documentID := "1X5G41rdtekvhsELQe--xFSqEYgjqd7hPG0G3VpyWT_Y" // id документа
	tableIndex := 0

	client := OAuth2()
	g := gdoctableapp.New()

	res, err := g.Docs(documentID).TableIndex(tableIndex).GetValues().Do(client)

	same := true

	currTable := res.Values

	//проверяем что данные в таблице с сайта такие же что в документе
	for i := 0; i < len(currTable); i++ {
		if len(tableS) < len(currTable) {
			same = false
			break
		}
		for j := 0; j < len(currTable[i]); j++ {
			if len(tableS[i]) < len(currTable[i]) {
				same = false
				break
			}
			if currTable[i][j] != tableS[i][j] {
				same = false
				break
			}
		}
	}

	//иначе удаляем и создаем новую таблицу
	if !same {
		_, err1 := g.Docs(documentID).TableIndex(tableIndex).DeleteTable().Do(client)
		checkError(err1)
		obj := &gdoctableapp.CreateTableRequest{
			Rows:    int64(len(table)),
			Columns: 2,
			Append:  true,
			// Append:  true, // When this is used instead of "Index", new table is created to the end of Document.
			Values: table,
		}
		_, err2 := g.Docs(documentID).CreateTable(obj).Do(client)
		checkError(err2)
	}
}
