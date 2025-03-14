package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	for i := 0; i < 5; i++ {
		// 出題する記号の列挙
		PERTS := [...]string{
			"?", "!", "#", "<", ">",
			"|", "@", "&", "*", "+", "'",
			"^", "\"",
			"\\", "%", "/", ";", ":",
			"(", ")", "{", "}", "[", "]", "_", "`",
			"$",
			"0", "1", "2", "3", "4", "5", "6", "7", "8", "9",
		}
		// 出題の選出
		const Q_LEN int = 5
		var ques string
		for i := 0; i < Q_LEN; i++ {
			rand.Seed(time.Now().UnixNano())
			ques += PERTS[rand.Intn(len(PERTS))]
		}
		// 出題
		fmt.Println(ques)
		var ans string
		fmt.Scan(&ans)
		// 正誤判定
		if ques != ans {
			fmt.Println("!!!!!!!!!!!! WRONG !!!!!!!!!!!!!!")
		}
	}
}
