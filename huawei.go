package main
import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main(){
	reader := bufio.NewReader(os.Stdin)
	text,err:=reader.ReadString('\n')
	for ;err==nil && len(text) >0 && text != "exit" ;text,err=reader.ReadString('\n') {
		oo := strings.Fields(text)
		var a= []int{}
		var sum int
		for _,v := range oo{
			j,err := strconv.Atoi(v)
			if err != nil{
				panic(err)
			}
		sum = sum+j
		a = append(a, j)
		}
		fmt.Println("Array: ",a,"\nsum:",sum)
		
		half := getMinimumDiff(a,sum/2)
		fmt.Printf("First half:%d\n",half)
		ot := sum -half
		fmt.Printf("Minimum difference is:%d\n",ot -half)
		text = ""
		fmt.Println("Please input an integer array: \n")
	}
	fmt.Println("Bye bye~~")
}

func getMinimumDiff(str []int, half int) int{
	if(half ==0){
		return 0;
	} 
	if(len(str)<=0){
		return -1;
	}
	 if(half <0){
                  return -1
          }
          
	for ; half>0;half--{
		if(getHalf(str, half)){
			return half
		}
	}
	return half
}

func getHalf(a []int, va int) bool{
	if va ==0 {
		return true
	}
	if len(a) <=0 {
		return false
	}
	for i,v:= range(a){
		if getHalf(a[i+1:], va - v) {
			return true
		}
	}
	return false
}



