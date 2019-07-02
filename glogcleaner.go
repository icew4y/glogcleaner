package main

import (
	"fmt"
	"github.com/xuri/glc"
	"time"
)

func main() {
	//var s string;
	//s = "backend_log20190610-032609.33";
	//prefix := strings.HasPrefix(s, "backend");
	//if (prefix){
	//	str := strings.Split(s, `.`)
	//	if (len(str) == 7){
	//
	//	}
	//	fmt.Print("yes")
	//}
	var cstZone = time.FixedZone("CST", 8*3600)

	fmt.Print(time.Now().In(cstZone).Format("2006-01-02 15:04:05") + "\n")
	glc.NewGLC(glc.InitOption{
		Path:     "/root/projects/vivo_module/bin/x64/Release/logs/",
		Prefix:   "backend",
		Interval: time.Duration(time.Hour * 12),
		Reserve:  time.Duration(time.Hour * 12),
	})
	defer func() { select {} }()
}
