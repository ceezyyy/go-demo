package main

import (
	"bufio"
	"git.inke.cn/BackendPlatform/go-tools/cast"
	"os"
)

func main() {
	file, err := os.Open("tep.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sc := bufio.NewScanner(file)

	counter := 0
	users := make([]int64, 0, 20000)
	IDs := make([]int, 0, 20000)

	for sc.Scan() {
		item := sc.Text()

		if counter%2 == 0 {
			users = append(users, cast.ToInt64(item))
		} else {
			IDs = append(IDs, cast.ToInt(item))
		}

		counter++
	}

	if err = sc.Err(); err != nil {
		panic(err)
	}

}
