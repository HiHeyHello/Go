package myStock

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"
)

const (
	webSide = "http://www.tse.com.tw" //網站位址
)

const (
	/* 每日收盤行情 */
	dailyClsPrcs = "/ch/trading/exchange/MI_INDEX/MI_INDEX.php" //網站API(每日收盤行情)

	typeListHead   = "<select name='selectType'>\n<option  selected  value='"
	typeListTail   = "</option>	</select>"
	typeListSepRow = "</option><option   value='"
	typeListSepCol = "'>"

	prmDate = "qdate"      //網站API(每日收盤行情),參數資料日期
	prmType = "selectType" //網站API(每日收盤行情),參數分類項目

	// startDate = "2004/02/11 00:00:00"
	startDate = "2017/04/10 00:00:00" //網站提供資訊啟始日期
)

func GetOneDayHist() {
	//Request
	resp, err := http.PostForm(webSide+dailyClsPrcs, url.Values{prmDate: {"106/04/14"}, prmType: {"MS"}})
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}

func GetAllHist() {
	firstDate, err := time.Parse("2006/01/02 15:04:05", startDate)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Printf("First Date:\t%s(%d)\n", firstDate.Format("2006/01/02 15:04:05"), firstDate.Unix())

	nowDate := time.Now()
	fmt.Printf("Now Date:\t%s(%d)\n", nowDate.Format("2006/01/02 15:04:05"), nowDate.Unix())

	fmt.Println()

	for ; firstDate.
		Unix() < nowDate.Unix(); nowDate = nowDate.AddDate(0, 0, -1) {
		fmt.Printf("%s(%d)\n", nowDate.Format("2006/01/02"), nowDate.Weekday())
	}
}

func GetTypeList() [][]string {
	//Request
	resp, err := http.PostForm(webSide+dailyClsPrcs, nil)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return nil
	}

	//Type List Head
	typeListHeadIndex := strings.Index(string(body), typeListHead)
	if typeListHeadIndex < 0 {
		fmt.Printf("找不到Type list head(%s)\n", typeListHead)
		return nil
	}
	typeListHeadIndex += len(typeListHead)

	typeListHead := string(body[typeListHeadIndex:])

	//Type List Tail
	typeListTailIndex := strings.Index(typeListHead, typeListTail)
	if typeListTailIndex < 0 {
		fmt.Printf("找不到Type list tail(%s)\n", typeListTail)
		return nil
	}

	//Type list string
	typeListStr := typeListHead[:typeListTailIndex]

	//Type list string separate by row
	typeListSliptRow := strings.Split(typeListStr, typeListSepRow)

	//Type list string separate by column
	typeList := [][]string{}

	for _, val := range typeListSliptRow {
		typeList = append(typeList, strings.Split(val, typeListSepCol))
	}

	return typeList
}

func IsExist(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return true, err
}

func UpdHist(path string) {

}
