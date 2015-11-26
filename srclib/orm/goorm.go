package main
import (
	"fmt"
	
)
const promt = `Please enter number of operation:
1.Create new count
2.Show detail of count
3.Delete count
4.Exit`
func main(){
	fmt.Print("welcome bank of xorm!")
Exit:
	for {
		fmt.Println(promt)
		var num int
		fmt.Scanf("%d\n",&num)
		switch num{
			case 1:
				fmt.Println("Please enter <name> <balance>")
				var name string
				var balance float64
				fmt.Scanf("%s %f\n",&name,&balance)
				if err := newAccount(name,balance);err != nil{
					fmt.Print(err)
				}
			case 2: 
				fmt.Println("Please id int: ")
				var id int64
				fmt.Scanf("%d\n",&id)
				a,err := getAccount(id)
				if err != nil {
					fmt.Printf("%#v\n",a)
				}				
					fmt.Printf("%#v\n",a)
			case 9: 
				break Exit				
		}
	}
	
	
}

