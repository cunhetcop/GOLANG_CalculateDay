// TÍNH THỜI GIAN TỪ NGÀY HIỆN TẠI ĐẾN NGÀY BẤT KÌ
package methods

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/fatih/color"
)

func Time() {
	//set thời gian chạy function
	timeOut := time.AfterFunc(10*time.Second, func() {
		color.Red("Thời gian chờ đã kết thúc. Tự động thoát chương trình.")
		// os.Exit(0)
		panic("timeout")
	})
	defer timeOut.Stop()

	//nhập dữ liệu bằng package bufio/scanner
	var input string
	scanner := bufio.NewScanner(os.Stdin)
	count := 0
	for count < 3 {
		color.Yellow("CHƯƠNG TRÌNH TÍNH THỜI GIAN TỪ NGÀY HIỆN TẠI ĐẾN NGÀY BẤT KỲ")
		fmt.Println("Nhập ngày tháng theo định dạng dd/mm/yyyy: ")

		//check time valid
		if scanner.Scan() {
			input = scanner.Text()
			input = strings.TrimSpace(input)
			layout := "02/01/2006"
			t, err := time.ParseInLocation(layout, input, time.Local)
			//xét điều kiện để chạy lại vòng lặp, nếu vượt quá 3 lần thì thoát vòng lặp
			if err != nil {
				fmt.Printf("Nhập sai %d lần \n", count+1)
				count++
				continue
			}
			timeOut.Stop()
			color.Yellow("Thời gian nhập vào: %s\n", time.Now().Format("02/01/2006 15:04:05"))
			now := time.Now()
			diff := t.Sub(now)
			days := int(diff.Hours() / 24)
			hours := int(diff.Hours()) % 24
			minutes := int(diff.Minutes()) % 60

			//xét điều kiện thời gian quá khứ hay tương lai
			if diff < 0 {
				color.Yellow("Thời gian đã trôi qua tính giờ 0h ngày hôm nay: %v ngày %v giờ %v phút", days, hours, minutes)
			} else {
				color.Yellow("Thời gian tới ngày hôm đó tính từ 0h hôm nay còn: %v ngày %v giờ %v phút", days, hours, minutes)
			}
			return
		}
	}
	//đoạn if này có thể ko cần thiết
	if err := scanner.Err(); err != nil {
		fmt.Println("Lỗi nhập dữ liệu:", err)
	}
	color.Red("Bạn đã nhập sai quá 3 lần. Game over.")
}
