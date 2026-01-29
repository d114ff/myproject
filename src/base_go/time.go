package main

import (
	"fmt"
	"time"
)

const (
	TIME_FMT = "2006-01-02 15:04:05.000" //年-月-日 时:分:秒.毫秒
	DATE_FMT = "20060102"
)

func basicTime() {
	t0 := time.Now()
	fmt.Printf("秒%d, 毫秒%d, 微妙%d, 纳秒%d\n", t0.Unix(), t0.UnixMilli(), t0.UnixMicro(), t0.UnixNano())
	time.Sleep(2 * time.Second) // 休眠2秒
	t1 := time.Now()
	fmt.Println(t1)
	diff1 := t1.Sub(t0) //计算时间差，返回类型是time.Duration
	//diff1 := time.Since(t0)
	fmt.Println(diff1)                                                           //秒上相同
	diff2 := time.Since(t0)                                                      //等价于t1 := time.Now();t1.Sub(t0)
	fmt.Println(diff2)                                                           //秒上相同
	fmt.Printf("diff1=diff2 %t\n", int(diff1.Seconds()) == int(diff2.Seconds())) //在秒级别上diff1和diff2是相等的
	fmt.Printf("t1 > t0 %t\n", t1.After(t0))                                     //判断时间的先后顺序
	d := time.Duration(3 * time.Hour)                                            //Duration表示两个时刻之间的距离
	t2 := t1.Add(d)                                                              //加一段时间
	fmt.Printf("hour1 %d, hour2 %d \n", t1.Hour(), t2.Hour())                    //获取一个时刻的Hour
	fmt.Printf("week dat %s %d\n", t2.Weekday().String(), t2.Weekday())          //周日的Weekday()是0，周六的Weekday()是6
	fmt.Printf("year %d\n", t2.Year())
	fmt.Printf("month %s,%d,%d\n", t1.Month(), t1.Month(), int(t1.Month())) //在go语言内部Month就是int类型。type Month int
	fmt.Printf("day %d\n", t2.Day())                                        //属于一个月当中的第几天
	fmt.Printf("day in year %d\n", t2.YearDay())                            //属于一年当的第几天
	fmt.Printf("%d-%d-%d %d:%d:%d.%d\n", t1.Year(), t1.Month(), t1.Day(), t1.Hour(), t1.Minute(), t1.Second(), t1.UnixMilli()%t1.Unix())

}

// 时间格式化
func parse_format() {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	now := time.Now()
	ts := now.Format(TIME_FMT)
	fmt.Println("格式时间：", ts)

	t, _ := time.Parse(TIME_FMT, ts) //不建议使用Parse
	if t.Unix() != now.Unix() {
		fmt.Printf("Parse时间解析错误：%d != %d\n", t.Unix(), now.Unix())
	}

	t, _ = time.ParseInLocation(TIME_FMT, ts, loc)
	if t.Unix() != now.Unix() {
		fmt.Printf("ParseInLocation时间解析错误：%s\n", t.Format(TIME_FMT))
	} else {
		t2, _ := time.ParseInLocation(TIME_FMT, ts, loc)
		fmt.Println(t2)
	}

}

// 执行周期任务
func ticker() {
	tk := time.NewTicker(1 * time.Second)
	for i := 0; i < 10; i++ {
		<-tk.C // 阻塞1秒中
		fmt.Printf("现在时间%s\n", time.Now().Format(TIME_FMT))
	}
	tk.Stop() //用完记得调用stop
}

// 经过一段时间后，触发一次任务
func timer() {
	fmt.Printf("现在时间是%s\n", time.Now().Format(TIME_FMT))
	tm := time.NewTimer(3 * time.Second)
	<-tm.C //阻塞3秒钟，最多只能连续调用一次
	fmt.Printf("现在时间是%s\n", time.Now().Format(TIME_FMT))
	tm.Stop()
	//通过time.After也可以实现上述功能
	fmt.Println("------------------------------------------")
	fmt.Printf("现在时间是%s\n", time.Now().Format(TIME_FMT))
	<-time.After(3 * time.Second) // 阻塞3秒
	fmt.Printf("现在时间是%s\n", time.Now().Format(TIME_FMT))

}

func main20() {
	basicTime()
	//parse_format()
	//ticker()
	//timer()
}
