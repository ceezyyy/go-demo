package main

import (
	"fmt"
	"time"
)

func main() {

	now := time.Now()
	fmt.Println(now)

	then := time.Date(2021, 7, 20, 9, 0, 60, 651387237, time.UTC)
	fmt.Println(then)

	fmt.Println(then.Weekday())

	// Compare two times
	fmt.Println(then.Before(now))
	fmt.Println(then.After(now))
	fmt.Println(then.Equal(now))

	diff1 := now.Sub(then)
	diff2 := then.Sub(now)
	// Duration 带正负
	fmt.Println(diff1.Hours())
	fmt.Println(diff2.Hours())

	fmt.Println(time.Now().Format("2006-01-02 15:04:05.999"))

	str := "2020-03-08 11:33:44"
	t, err := time.Parse(time.RFC3339, str)

	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(t)

	tomorrow := time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	fmt.Println(tomorrow)

	fmt.Println(time.Now().AddDate(0, 0, 2).Format("2006-01-02 15:04:05"))

	nowTS := time.Now().Unix()
	timeT := time.Unix(nowTS, 0) // timestamp to time.Time
	fmt.Println(timeT.Hour())

	rounded := time.Date(now.Year()+1, 1, 1, 0, 0, 0, 0, t.Location())
	fmt.Println(rounded)

	start := time.Now().Format("200601") + "01"
	fmt.Println(start)

	tmp := 1 * 24 * time.Hour.Seconds()
	fmt.Println(int64(tmp))

	end := time.Now().AddDate(0, 0, 7).Unix()
	fmt.Println(end)

	fmt.Println(time.Now().Format("01021504"))

}
