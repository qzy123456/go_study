package main

import (
	"bytes"
	"fmt"
	"os"
	"unicode"
)

func main() {
	fmt.Println(bytes.Equal([]byte{}, []byte{}))            // true
	fmt.Println(bytes.Equal([]byte{'A', 'B'}, []byte{'a'})) // false
	fmt.Println(bytes.Equal([]byte{'a'}, []byte{'a'}))      // true
	fmt.Println(bytes.Equal([]byte{}, nil))                 // true
	fmt.Println(bytes.Equal([]byte{}, []byte{}))            // true
	fmt.Println(bytes.Equal([]byte{'A', 'B'}, []byte{'a'})) // false
	fmt.Println(bytes.Equal([]byte{'a'}, []byte{'a'}))      // true
	fmt.Println(bytes.Equal([]byte{}, nil))                 // true
	fmt.Println(bytes.EqualFold([]byte{},[]byte{})) // true
	fmt.Println(bytes.EqualFold([]byte{'A'},[]byte{'a'})) // true
	fmt.Println(bytes.EqualFold([]byte{'B'},[]byte{'a'})) // false
	fmt.Println(bytes.EqualFold([]byte{},nil)) // true
	fmt.Println(bytes.HasPrefix([]byte{1, 2, 3}, []byte{1})) // true
	fmt.Println(bytes.HasPrefix([]byte{1, 2, 3}, []byte{2})) // false

	fmt.Println(bytes.HasSuffix([]byte{1, 2, 3, 3}, []byte{3, 3}))  // true
	fmt.Println(bytes.HasSuffix([]byte{1, 2, 3, 3}, []byte{3, 4}))  // false

	fmt.Println(bytes.Contains([]byte{1, 2, 3}, []byte{1}))    // true
	fmt.Println(bytes.Contains([]byte{1, 2, 3}, []byte{1, 3})) // false

	fmt.Println(bytes.Index([]byte{1, 2, 3, 4, 5}, []byte{4, 5})) // 3
	fmt.Println(bytes.Index([]byte{1, 2, 3, 4, 5}, []byte{0, 1})) // -1

	fmt.Println(bytes.IndexByte([]byte{1, 2, 3}, 3)) // 2
	fmt.Println(bytes.IndexByte([]byte{1, 2, 3}, 0)) // -1

	fmt.Println(bytes.LastIndex([]byte("hi go"), []byte("go"))) // 3
	fmt.Println(bytes.LastIndex([]byte{1, 2, 3}, []byte{2, 3})) // 1

	fmt.Println(bytes.IndexAny([]byte("hi go"), "go")) // 3

	fmt.Println(bytes.Count([]byte("hi go go go go go go go go go"), []byte("go"))) // 9

	fmt.Println(bytes.IndexRune([]byte("你好吗,不太好啊,hi go go go go go go go go go"), '不')) // 9

	fmt.Println(bytes.IndexFunc([]byte("hi go"), func(r rune) bool {
		return r == 'g'
	})) // 3
	fmt.Println(string(bytes.Title([]byte("AAA")))) // AAA
	fmt.Println(string(bytes.Title([]byte("aaa")))) // Aaa

	fmt.Println(string(bytes.ToTitle([]byte("Aaa")))) // AAA

	fmt.Println(string(bytes.ToUpper([]byte("Aaa")))) // AAA
	fmt.Println(string(bytes.ToUpperSpecial(unicode.SpecialCase{}, []byte("Aaa")))) // AAA

	fmt.Println(string(bytes.ToLower([]byte("aAA")))) // aaa
	fmt.Println(string(bytes.ToLowerSpecial(unicode.SpecialCase{}, []byte("aAA")))) // aaa
	fmt.Println(bytes.Repeat([]byte{1,2},3)) // [1 2 1 2 1 2]

	fmt.Println(bytes.Replace([]byte{1,2,1,2,3,1,2,1,2}, []byte{1,2}, []byte{0,0},3)) // [0 0 0 0 3 0 0 1 2]

	fmt.Println(string(bytes.Map(func(r rune) rune {
		return r + 1 // 将每一个字符都+1
	},[]byte("abc")))) // bcd

	fmt.Println(string(bytes.Trim([]byte("hello my"), "my")))  // hello

	fmt.Println(string(bytes.TrimSpace([]byte(" hello my bro "))))   // hello my bro

	fmt.Println(string(bytes.TrimLeft([]byte("hi hi go"), "hi"))) //  hi go
	fmt.Println(string(bytes.TrimPrefix([]byte("hi hi go"),[]byte("hi")))) //  hi go
	fmt.Println(string(bytes.TrimSuffix([]byte("hi go go"),[]byte("go")))) // hi go
	r := bytes.NewReader([]byte("ABCDEFGHIJKLMN IIIIII LLLLLLLL SSSSSS"))

	fmt.Println(r.Len()) // 37
	fmt.Println(r.Size()) // 37

	tmp := make([]byte,5)
	n,_ := r.Read(tmp)
	fmt.Println(string(tmp[:n])) // ABCDE
	fmt.Println(r.Len(),r.Size()) // 32 37

	fmt.Println(r.ReadByte()) // 70 <nil> // F
	fmt.Println(r.ReadRune()) // 71 1 <nil>
	fmt.Println(r.Len(),r.Size()) // 30 37

	b := []byte("III")     // cap 3
	n,_ = r.ReadAt(b,1)
	fmt.Println(string(b),n)   // BCD 3

	r.Reset([]byte("Hi,My god"))
	fmt.Println(r.Len(),r.Size()) // 9 9

	r.WriteTo(os.Stdout)  // Hi,My god
	bb := bytes.NewBufferString("ABCDEFGH")

	fmt.Println(bb.String()) // ABCDEFGH

	fmt.Println(bb.Len()) // 8

	fmt.Println(string(bb.Next(2))) // AB
	tmp = make([]byte,2)

	nn,_ := bb.Read(tmp)
	fmt.Println(string(tmp[:nn])) //CD

	nextByte,_ := bb.ReadByte()
	fmt.Println(string(nextByte)) // E

	line,_ := bb.ReadString('G')
	fmt.Println(line) // FG
	fmt.Println(bb.String()) // H
	bb = bytes.NewBufferString("abcdefgh")

	line2,_ := bb.ReadBytes('b')
	fmt.Println(string(line2)) // ab
	fmt.Println(bb.String()) // cdefgh

	fmt.Println(string(bb.Bytes())) // defgh

	b1 := new(bytes.Buffer)

	b1.WriteByte('a')
	fmt.Println(b1.String()) // a

	b1.Write([]byte{98,99})
	fmt.Println(b1.String()) // abc

	b1.WriteString(" hello")
	fmt.Println(b1.String()) // abc hello

	b1.Truncate(3)
	fmt.Println(b1.String()) // abc

	n1,_ := b1.WriteTo(os.Stdout) // abc
	fmt.Println(n1) // 3

	b1.Reset()
	fmt.Println(b1.Len(),b1.String()) // 0

}
