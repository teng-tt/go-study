package main

import (
	"fmt"
	"math/rand"
)

func main() {

	//生成随机数,猜字游戏
	for {
		number := rand.Intn(100)
		fmt.Println("这是一个猜数字游戏，猜测数字范围为0-100，你有5次猜测机会")
		for count := 4; count >= 0; count-- {
			var guess int
			fmt.Print("请输入你猜测的数字:")
			fmt.Scan(&guess)

			if guess >= number {
				fmt.Printf("你输入的数字：%d,太大了，你还有%d重试次机会\n", guess, count)
			} else if guess <= number {
				fmt.Printf("你输入的数字：%d,太小了，你还有%d重试次机会\n", guess, count)
			} else {
				fmt.Printf("恭喜你猜对了，你太聪明了，答案为%d\n", guess)
				break
			}

		}
		fmt.Println("Game Over! 游戏结束,你太笨了！")

	}

}
