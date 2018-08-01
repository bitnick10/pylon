package frd

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"gok/chart"
	"pylon/chartdb"
	"pylon/checker"
	"strings"
	"time"
)

func GetCorrectTime(t time.Time) time.Time {
	ret := t
	ret = ret.Add(5 * time.Minute)
	ret = ret.Add(8 * time.Hour)
	return ret
}

func GetCorrectCSVContent(id, src string) string {
	//src = strings.Replace(src, ",", "", -1)
	var buffer bytes.Buffer
	//src = RemoveSomeRows(src)
	r := csv.NewReader(strings.NewReader(src))
	r.Comma = ','
	records, err := r.ReadAll()
	if err != nil {
		panic(err)
	}
	buffer.WriteString("date\topen\thigh\tlow\tclose\tvolume\n")
	for i := 1 ;i< len(records) ;i++ {
		row := records[i]
		timeString := row[0]
		t, _ := time.Parse("2006-01-02 15:04:05", timeString)
		t = GetCorrectTime(t)
		buffer.WriteString(t.Format("2006-01-02 15:04") + "\t")
		buffer.WriteString(row[1] + "\t")
		buffer.WriteString(row[3] + "\t")
		buffer.WriteString(row[4] + "\t")
		buffer.WriteString(row[2] + "\t")
		buffer.WriteString("0" + "\n")
	}
	return buffer.String()
}

func Feed(id, csv string) {
	csvContent := GetCorrectCSVContent(id, csv)
	cc, err := chart.NewChartFromCSVString(id, id, chart.M5, csvContent)
	if err != nil {
		fmt.Println("error NewChartFromCSVString")
		fmt.Println(err)
	}
	// fmt.Println("cc.Len()", cc.Len())
	err = checker.IsSticksOKRule3(cc)
	if err != nil {
		fmt.Println("check sticks error")
		fmt.Println(err)
	} else {
		fmt.Println(cc.Id(), cc.Period().String(), "sticks check OK len", cc.Len())
		// return
		chartdb.CreateTable(cc)
		chartdb.ReplaceInto(cc)
		fmt.Println(cc.Id(), cc.Period().String(), "Replace Into finished")
	}
}
