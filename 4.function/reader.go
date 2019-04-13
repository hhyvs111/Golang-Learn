package main
import(
	"fmt"
	"io"
	"strings"
)

func main() {
	r := strings.NewReader("hello, world!")

	//这里每次循环读取8个字节，然后输出
	b := make([]byte, 16)
	for {
		n, err := r.Read(b)
		fmt.Printf("n = %v err = %v b = %v\n", n, err, b)
		//这里的是输出0到n的长度的字符数组
		fmt.Printf("b[:n] = %q\n", b[:n])
		if err == io.EOF {
			break
		}
	}
}


		

