// カレンダー v.0.1
// https://play.golang.org/p/cWxfngR2imA

package main

import "fmt"

var mmdd = [42]int{99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99, 99}

const bs_nen int = 2000
const bs_yobi int = 6


func main() {
	//fmt.Println("Start")
      var in_nen int
      var in_tsuki int
      fmt.Print("input year > ")
      fmt.Scan(&in_nen)
      fmt.Print("input month > ")
      fmt.Scan(&in_tsuki)
      var nisu int = 0
      var yobi int = 0
      if in_nen >= bs_nen {
	    nisu = nen_nisu(in_nen)          
	    nisu = nisu + 
                    tsuki_nisu(0, in_nen, in_tsuki) 
            yobi = yobi_edit(nisu)

       } else {
           nisu = nen_nisu2(in_nen)
           nisu = nisu +
                    tsuki_nisu2(0,in_nen,in_tsuki)
            yobi = yobi_edit2(nisu)

       }
	var max int = 
                           tsuki_nisu(1, in_nen, in_tsuki)
	cal_edit(yobi, max)                   
	cal_out(in_nen, in_tsuki)               

	//fmt.Println("End")
}

//　閏年チェック
func uruchk(nen int) int {
	var flg int = 0
	if nen%4 == 0 {
		if nen%100 == 0 {
			if nen%400 == 0 {
				flg = 1
			}
		} else {
			flg = 1
		}
	}
	return flg
}

// 2000年から対象年の前年までに日数算出
func nen_nisu(nen int) int {
	var w_nisu int = 0
	var w_nen int = nen - 1
	for w_nen >= bs_nen {
		if uruchk(w_nen) == 1 {
			w_nisu = w_nisu + 366
		} else {
			w_nisu = w_nisu + 365
		}
		w_nen = w_nen - 1
	}
	return w_nisu
}


// 対象年の翌年から1999までの日数算出
func nen_nisu2(nen int) int {
	var w_nisu int = 0
	var w_nen int = nen + 1
	for w_nen < bs_nen {
		if uruchk(w_nen) == 1 {
			w_nisu = w_nisu + 366
		} else {
			w_nisu = w_nisu + 365
		}
		w_nen = w_nen + 1
	}
	return w_nisu
}

// 対象年の1/1から対象月の前月までの日数算出
func tsuki_nisu(sw int, nen int, tsuki int) int {
	var w_nisu int = 0
	tsuki_dt := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if uruchk(nen) == 1 {
		tsuki_dt[1] = 29
	}
	if sw == 1 {
		return tsuki_dt[tsuki-1]
	}

	var w_tsuki int = 1
	for w_tsuki < tsuki {
		w_nisu = w_nisu + 
                      tsuki_dt[w_tsuki-1]
		w_tsuki = w_tsuki + 1
	}
	return w_nisu
}

// 対象月から対象年の12月までの日数算出
func tsuki_nisu2(sw int, nen int, tsuki int) int {
	var w_nisu int = 0
	tsuki_dt := [12]int{31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31}
	if uruchk(nen) == 1 {
		tsuki_dt[1] = 29
	}
	if sw == 1 {
		return tsuki_dt[tsuki-1]
	}

	var w_tsuki int = tsuki
	for w_tsuki <= 12 {
		w_nisu = w_nisu + 
                      tsuki_dt[w_tsuki - 1]
		w_tsuki = w_tsuki + 1
	}
	return w_nisu
}


// 対象月の1日の曜日算出　日:0 月:1 ...
func yobi_edit(nisu int) int {
	var w_sa int
	var yobi int = bs_yobi
	w_sa = nisu % 7
	yobi = yobi + w_sa
	if yobi > 6 {
		yobi = yobi % 7
	}
	return yobi
}


// 対象月の1日の曜日算出　日:0 月:1 ...
func yobi_edit2(nisu int) int {
	var w_sa int
	var yobi int = bs_yobi
	w_sa = nisu % 7
	yobi = yobi - w_sa
	return yobi
}

// カレンダー編集
func cal_edit(s1 int, mx int) {
	var hi int = 1
	for hi <= mx {
		mmdd[s1] = hi
		s1 = s1 + 1
		hi = hi + 1
	}
}

// カレンダー出力
func cal_out(nen int, tsuki int) {
	fmt.Printf("%5d", nen)
	fmt.Print("/")
	fmt.Printf("%2d", tsuki)

	fmt.Println()
	fmt.Println()
	fmt.Println("  S  M  T  W  T  F  S")

	var ss1 int = 0
	for ss1 < 42 {
		if mmdd[ss1] == 99 {
			fmt.Print("   ")
		} else {
			fmt.Printf("%3d", mmdd[ss1])
		}
		ss1 = ss1 + 1
		if ss1%7 == 0 {
			fmt.Println()
		}
	}
}

