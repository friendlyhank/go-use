package time

import (
	"fmt"
	"testing"
	"time"
)

//相当于time.Now().Sub(t)
func TestTimeSince(t *testing.T) {
	elapsed := time.Since(time.Now())
	t.Logf("%v",elapsed)
}

//TestTimeGetYear- 获取年
func TestTimeGetYear(t *testing.T){
	year := time.Now().Year()
	t.Logf("%v",year)
}

//TestTimeGetMonth - 获取月
func TestTimeGetMonth(t *testing.T){
	month := time.Now().Month()
	t.Logf("%v",month.String())
	mNums := int(month)
	t.Logf("%v",mNums)
}

//TestTimeGetWeek -获取星期
func TestTimeGetWeek(t *testing.T){
	week := time.Now().Weekday()
	t.Logf("%v",week.String())
	wNums := int(week)
	t.Logf("%v",wNums)
}

//TestTimeGetDay- 获取天
func TestTimeGetDay(t *testing.T){
	day := time.Now().Day()
	t.Logf("%v",day)
}

//TestTimeGetDate -构造指定的时间,通常比如1天前,7天前
func TestTimeGetDate(t *testing.T){
	year,month,_ := time.Now().Date()
	currentTime := time.Date(year,month,15,0,0,0,0,time.Local)
	t.Logf("%v",currentTime)
	t.Logf("%v",currentTime.Unix())
}

//TestParseTime -格式化时间
func TestParseTime(t *testing.T){
	//格式化时间
	strTime :=time.Now().Format("2006-01-02 15:04:05")
	fmt.Println(strTime)
	//将格式化的时间再转化为时间
	time,_ := time.Parse("2006-01-02 15:04:05",strTime)
	fmt.Println(time)
}
