package main

import (
	"fmt"
	"github.com/xuri/glc"
	"sync"
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
	var workGroup sync.WaitGroup
	workGroup.Add(1)

	glc.NewGLC(glc.InitOption{
		Path:     "/root/projects/vivo_module/bin/x64/Release/logs/",
		Prefix:   "backend",
		Interval: time.Duration(time.Second * 30),
		Reserve:  time.Duration(time.Hour * 12),
	})
	workGroup.Wait()
}
