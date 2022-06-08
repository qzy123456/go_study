package fiveinarow

import(
	"testing"
	"fmt"
	"math/rand"
	"time"
)

func Benchmark_AddAllinat(b *testing.B){
	for k:=0;k<b.N;k++{
		c := Allinat{}
	    chess := Coordinat{}
		for i:=0; i<10; i++{
			if i % 2 == 0{
				chess.Color = "white"
			}else{
				chess.Color = "black"
			}
			rand.Seed(time.Now().UnixNano())
			chess.X = rand.Intn(5)
			chess.Y = rand.Intn(5)
			xx,x := c.AddCoordinat(chess)
			fmt.Println(c)
			if xx == true{
				fmt.Printf("%#v\n\n",x)
				break
			}
			
		}
	}
}


func TestAddAllinat(t *testing.T){
	 c := Allinat{}
	 chess := Coordinat{}
	 cch := make(chan int,40)
	 for i:=0; i<5000000; i++{
		if i % 2 == 0{
			chess.Color = "white"
		}else{
			chess.Color = "black"
		}
		cch <- i
		go func(ccs chan int){
			rand.Seed(time.Now().UnixNano())
			chess.X = rand.Intn(330)
			chess.Y = rand.Intn(330)
			xx,x := c.AddCoordinat(chess)
			if xx == true{
				fmt.Printf("%#v========\n\n",x)
			}
			fmt.Println(<-ccs)
		}(cch)
		
	 }

}

func TestSlope(t *testing.T){
	chess := Coordinat{3,1,""}
	Slope(nil,chess)
}

func TestCenter(t *testing.T){
	a := new(Allinat)
	chess := Coordinat{3,1,""}
	a.AddCoordinat(chess)
	a.Center()
	fmt.Println(a.Key)
}