package test

import (
	"fmt"
	"net/http"
	"time"
)

const serverPort = 3333

func timer() {
	ticker := time.NewTicker(2 * time.Second)
	i := 0
	//子协程
	go func() {
		for {
			i++
			fmt.Println(<-ticker.C)
			send()
			if i == 1000 {
				ticker.Stop()
			}
		}
	}()
	time.Sleep(time.Second * 30000000)
}

func send() {
	fmt.Println("test")
	resp, err := http.Get("https://baidu.com")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)
	//fmt.Println(string(body))
	fmt.Println(resp.StatusCode)
	if resp.StatusCode == 200 {
		fmt.Println("ok")
	} else {
		fmt.Println("not ok")
	}
}
