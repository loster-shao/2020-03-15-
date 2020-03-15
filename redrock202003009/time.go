package main

import (
	"fmt"
	"time"
)

func main()  {
	sss := time.Now().Unix()
	tm2 := time.Now().Format("2006-01-02")
	ss := time.Unix(tm2, 0)
	fmt.Println(sss, tm2, ss)
}
