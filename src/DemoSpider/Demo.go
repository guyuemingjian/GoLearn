package main

import (
	"database/sql"
	"fmt"
	_ "github.com/mattn/go-sqlite3" // sqlite3 dirver
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const BASE_URL string = "http://xh.5156edu.com/"
const BASE_HTMl string = "bs.html"
const SQL_INSERT string = "insert into font_ch(hanzi,pinyin,bushou,bihua,jieshi) values(?, ?, ?, ?, ?);"
const SQL_SELECT string = "select * from font_ch;"

//日志
var fileName = "D:/00work/GoDemo/build/Err_First_" + time.Now().Format("20060102_15040599") + ".log"
var logFile, _ = os.Create(fileName)
var errLog = log.New(logFile, "[err]", log.Llongfile)

type font struct {
	id     int32  //标识
	hanzi  string //汉字
	pinyin string //拼音
	bushou string //部首
	bihua  int    //笔画
	jieshi string //解释
}

func main() {

	//数据库连接
	db := DbConnect()

	//所有偏旁检索
	herf1 := fistGet()

	var count = 0
	var errcount = 0
	//当前偏旁对应汉字
	for i := 0; i < len(herf1); i++ {
		//for i := 0; i < 1; i++ {

		//详细汉字连接取得
		herf2 := HanZiInBsGet(herf1[i][1])

		for j := 0; j < len(herf2); j++ {
			//for j := 0; j < 1; j++ {
			count = count + 1
			fmt.Println(count)

			fmt.Println(time.Now().Format("2006-01-02 15:04:05.999999999") + " sub " + herf2[j][1])
			res, err := HanZiGet(herf2[j][1])

			if "" != err {
				errLog.Println(err)
				fmt.Println(err)
				errcount = errcount + 1
				fmt.Println(errcount)
				continue
			}

			//保存到数据库
			InsertFontCh(&res, db)
		}
	}

	db.Close()
	fmt.Println("爬取结束")
}

//得到所有偏旁连接
func fistGet() [][]string {

	resp, err := http.Get(BASE_URL + BASE_HTMl)
	checkErr(err, "部首取得失败")

	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err, "部首读取失败")

	var decodeBytes, _ = simplifiedchinese.GBK.NewDecoder().Bytes(body)
	var html = string(decodeBytes)
	//fmt.Println(html)

	commentCount := `<a class='fontbox' href="(.*?)">`
	rep := regexp.MustCompile(commentCount)
	herf1 := rep.FindAllStringSubmatch(html, -1)

	//for i := 0; i < len(herf1); i++ {
	//	fmt.Println(herf1[i][1])
	//}
	time.Sleep(time.Second)

	return herf1
}

//当前偏旁包含的汉字链接
func HanZiInBsGet(subUrl string) [][]string {
	resp, err := http.Get(BASE_URL + subUrl)
	checkErr(err, "汉字取得失败")

	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err, "部首下汉字读取失败")

	var decodeBytes, _ = simplifiedchinese.GBK.NewDecoder().Bytes(body)
	var html = string(decodeBytes)
	//fmt.Println(html)

	commentCount := `<a class='fontbox' href='(.*?)'>`
	rep := regexp.MustCompile(commentCount)
	herf2 := rep.FindAllStringSubmatch(html, -1)

	//for i := 0; i < len(herf2); i++ {
	//	fmt.Println(herf2[i])
	//}

	time.Sleep(time.Second)
	return herf2
}

//汉字信息取得
func HanZiGet(subUrl string) (font, string) {

	var result font

	resp, err := http.Get(BASE_URL + subUrl)
	checkErr(err, "汉字详细取得失败"+subUrl)

	body, err := ioutil.ReadAll(resp.Body)
	checkErr(err, "汉字详细读取失败"+subUrl)

	var decodeBytes, _ = simplifiedchinese.GBK.NewDecoder().Bytes(body)
	var html = string(decodeBytes)
	html = strings.Replace(html, "\r", "", -1)
	html = strings.Replace(html, "\n", "", -1)
	html = strings.Replace(html, " ", "", -1)
	html = strings.Replace(html, "\t", "", -1)
	html = strings.Replace(html, "&nbsp;", "", -1)
	html = strings.Replace(html, "<BR>", "<br>", -1)
	html = strings.Replace(html, "<P>", "", -1)
	html = strings.Replace(html, "</P>", "</p>", -1)
	//fmt.Println(html)

	commentCount := `基本解释(.*)详细解释`
	rep := regexp.MustCompile(commentCount)
	temp := rep.FindAllStringSubmatch(html, -1)
	if len(temp) < 1 {
		errLog.Println(html)
		return result, "基本解释没有" + subUrl
	}
	xxTemp := temp[0][1]
	//fmt.Println(xxTemp)

	//tempF.hanzi = "丨"
	commentCount = `<hrclass=hr1>(.*)<br>`
	rep = regexp.MustCompile(commentCount)
	temp1 := rep.FindAllStringSubmatch(xxTemp, -1)
	if len(temp1) < 1 {
		errLog.Println(xxTemp)
		return result, "汉字没有" + subUrl
	}
	hanzi := temp1[0][1]
	ind := strings.Index(hanzi, "<br>")
	if ind > 0 {
		hanzi = hanzi[0:ind]
	}
	result.hanzi = hanzi
	//fmt.Println(result.hanzi)

	//tempF.pinyin = "g"
	commentCount = `<spanclass=font_15>(.*)</span><br>`
	rep = regexp.MustCompile(commentCount)
	temp1 = rep.FindAllStringSubmatch(xxTemp, -1)
	if len(temp1) < 1 {
		errLog.Println(xxTemp)
		return result, "拼音没有" + subUrl
	}
	pinyin := temp1[0][1]
	ind1 := strings.Index(pinyin, "</span>")
	if ind1 > 0 {
		pinyin = pinyin[0:ind1]
	}
	result.pinyin = pinyin
	//fmt.Println("拼音 " + result.pinyin)

	//tempF.bushou = "鱼"
	commentCount = `部首：(.*)；`
	rep = regexp.MustCompile(commentCount)
	temp1 = rep.FindAllStringSubmatch(xxTemp, -1)
	if len(temp1) < 1 {
		errLog.Println(xxTemp)
		return result, "部首没有" + subUrl
	}
	result.bushou = temp1[0][1]
	//fmt.Println("部首 " + result.bushou)

	//tempF.bihua = 3
	commentCount = `笔画数：(.*)；<br>部首`
	rep = regexp.MustCompile(commentCount)
	temp1 = rep.FindAllStringSubmatch(xxTemp, -1)
	if len(temp1) < 1 {
		errLog.Println(xxTemp)
		return result, "笔画没有" + subUrl
	}
	result.bihua, _ = strconv.Atoi(temp1[0][1])
	//fmt.Printf("笔画 %d \n", result.bihua)

	//tempF.jieshi = "瞎说"
	commentCount = `</span><br>(.*)笔画数`
	rep = regexp.MustCompile(commentCount)
	temp1 = rep.FindAllStringSubmatch(xxTemp, -1)
	if len(temp1) < 1 {
		errLog.Println(xxTemp)
		return result, "解释没有" + subUrl
	}
	jieshi := temp1[0][1]
	result.jieshi = strings.Replace(jieshi, "<br>", "", -1)
	//fmt.Println("解释 " + result.jieshi)

	time.Sleep(100 * time.Millisecond)
	return result, ""
}

func checkErr(err error, msg string) {
	if err != nil {
		fmt.Println(err, " \n", msg)
		panic(1)
	}
}

func DbConnect() *sql.DB {
	db, err := sql.Open("sqlite3", "D:/SQLiteStudio/DB/test.db")
	checkErr(err, "连接数据库失败")

	return db
}

func SelectFontCH(wkDB *sql.DB) {

	stmt, err := wkDB.Prepare(SQL_SELECT)
	checkErr(err, "查询预处理失败")

	rows, err := stmt.Query()
	checkErr(err, "查询失败")

	for rows.Next() {
		var tempF font
		rows.Scan(&tempF.id, &tempF.hanzi, &tempF.pinyin, &tempF.bushou, &tempF.bihua, &tempF.jieshi)
		fmt.Printf("标识:%d  汉字:%s  拼音:%s  部首:%s  笔画:%d  解释:%s  \n", tempF.id, tempF.hanzi, tempF.pinyin, tempF.bushou, tempF.bihua, tempF.jieshi)
	}
}

func InsertFontCh(wkFont *font, wkDB *sql.DB) {

	//var tempF font
	//tempF.hanzi = "丨"
	//tempF.pinyin = "g"
	//tempF.bushou = "鱼"
	//tempF.bihua = 3
	//tempF.jieshi = "瞎说"
	//InsertFontCh(&tempF, db)

	stmt, err := wkDB.Prepare(SQL_INSERT)
	checkErr(err, "插入预处理失败")

	_, err = stmt.Exec(wkFont.hanzi, wkFont.pinyin, wkFont.bushou, wkFont.bihua, wkFont.jieshi)
	checkErr(err, "插入失败")
}
