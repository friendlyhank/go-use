package time

import (
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
