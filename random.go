// 1	足立区
// 2	荒川区
// 3	板橋区
// 4	江戸川区
// 5	大田区
// 6	葛飾区
// 7	北区
// 8	江東区
// 9	品川区
// 10	渋谷区
// 11	新宿区
// 12	杉並区
// 13	墨田区
// 14	世田谷区
// 15	台東区
// 16	千代田区
// 17	中央区
// 18	豊島区
// 19	中野区
// 20	練馬区
// 21	文京区
// 22	港区
// 23	目黒区

package main

import (
	"fmt"
	"math/rand"
	"time"
)

type District struct {
	number int
	name   string
}

// func DistrictName(num int) string {
// 	districts := []District{{1, "足立区"}, {12, "杉並区"}}
// 	return "足立区"
// }

func random() District {
	districts := []District{{1, "足立区"}, {12, "杉並区"}}
	rand.Seed(time.Now().UnixNano())
	i := rand.Intn(len(districts))

	return districts[i]
}

func main() {
	district := random()
	fmt.Println(district)
}
