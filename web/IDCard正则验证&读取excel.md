### IDCard正则验证
~~~
func IsIDCard(idCard string) bool {
	if idCard != "" {
		if isOk, _ := regexp.MatchString(`(^[1-9]\d{5}(18|19|([23]\d))\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}[0-9Xx]$)|(^[1-9]\d{5}\d{2}((0[1-9])|(10|11|12))(([0-2][1-9])|10|20|30|31)\d{3}$)`, idCard); isOk {
			return true
		}
	}

	return false
}

~~~


### 读取excel

~~~
package main

import (
	"bytes"
	"fmt"
	mrand "math/rand"
	"time"

	"github.com/tealeg/xlsx"
)

func main() {

	var buffer bytes.Buffer
	retUserID := readExcelForAction()
	for _, v := range retUserID {
		buffer.WriteString(v + ",")
	}
	fmt.Println(buffer.String())
}

var userIDList []string

func readExcelForAction() []string {
	excelFileName := "/Users/tom/Desktop/users.xlsx"
	xlFile, err := xlsx.OpenFile(excelFileName)
	if err != nil {
		fmt.Println("err:", err)
	}

	for _, sheet := range xlFile.Sheets {

		if sheet.Name == "Sheet1" {
			continue
		}
		for _, row := range sheet.Rows {

			for idx, cell := range row.Cells {

				if idx == 0 && cell.Value == "" {
					continue
				}

				userIDList = append(userIDList, cell.String())

			}

		}
	}
	return userIDList
}

~~~
