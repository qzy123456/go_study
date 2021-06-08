package main

import "fmt"

func countPrimes(n int) int {
	primes := []int{}
	isPrime := make([]bool, n)
	for i := range isPrime {
		isPrime[i] = true
	}

	for i := 2; i < n; i++ {
		if isPrime[i] {
			//fmt.Println(i) //2 3 5
			primes = append(primes, i)
		}
		for _, p := range primes {
			//fmt.Println(i,p) //2 2 , 3 2, 4 2, 5 2
			if i*p >= n {
				break
			}
			isPrime[i*p] = false
			if i%p == 0 {
				break
			}
		}
	}
	return len(primes)
}
func countPrimes2(n int) int {
	if n <= 2 {
		return 0
	}
	isPrime := make([]bool, n)
	isPrime[0], isPrime[1] = true, true
	for i := 2; i < n; i++ {
		if !isPrime[i] {
			for j := 2; ; {
				summation := j * i
				if summation < n {
					isPrime[summation] = true
					j++
					continue
				}
				break
			}
		}
	}
	count := 0
	for _, value := range isPrime {
		if !value {
			count++
		}
	}
	return count
}
func countPrimes3(n int) int {
	if n<2{
		return 0
	}
	var isPrime []bool
	for i:=0;i<n;i++{
		isPrime=append(isPrime, true)
	}
	for i:=2;i*i<n;i++{
		for j:=i*i;j<n;j=j+i{
			isPrime[j]=false
		}
	}
	cnt:=0
	for _,v:=range isPrime{
		if v==true{
			cnt++
		}
	}
	return cnt-2
}

func countPrimes4(n int) int {
	notPrime := make([]bool,n)
	count := 0
    //fmt.Println(notPrime) [false false false false false false]
	for i := 2; i < n; i++ {
		if notPrime[i] == false {
			fmt.Println("i=",i)
			count++
		}
		for j := i * i; j < n; j += i {
			fmt.Println("j=",j)
			notPrime[j] = true
		}
	}
	return count
}

func main()  {
	//fmt.Println(countPrimes(6))
	//fmt.Println(countPrimes2(6))
	fmt.Println(countPrimes3(10))
	fmt.Println(countPrimes4(10))
}
