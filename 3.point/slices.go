package main
import "fmt"

func main(){
	p := []int{1,2,3,4,5,6,7}
	fmt.Println("p==", p)

	for i := 0; i < len(p); i++ {
		fmt.Println("p[%d] ==  %d\n", i, p[i])
	}
}

