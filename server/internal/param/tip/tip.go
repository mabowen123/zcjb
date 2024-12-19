package tip

type List struct {
	Id           uint   `json:"id"`
	Url          string `json:"url" dc:"爬取链接"`
	IntervalTime uint   `json:"interval_time" dc:"爬取间隔时间"`
	NextTime     uint   `json:"next_time" dc:"下次爬取时间"`
	Remark       uint   `json:"remark" dc:"备注"`
}
