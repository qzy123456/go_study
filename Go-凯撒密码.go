package main

import (
	"fmt"
	"strings"
)

const STEP byte = '3'

// 凯撒密码加密
func caesarEn(strRaw string) string {
	//1. 将文本转为小写
	str_raw := strings.ToLower(strRaw)
	//2. 定义步长
	step_move := STEP
	//3. 将字符串转换为明文字符切片
	str_slice_src := []byte(str_raw)
	fmt.Println("Clear text character slice:", str_slice_src)

	//4. 创建一个密文字符切片
	str_slice_dst := str_slice_src

	//5.循环处理文本切片
	for i := 0; i < len(str_slice_src); i++ {
		//6.如果当前周期的明文特征在位移范围内，请直接添加位移步骤以保存密文字符切片
		if str_slice_src[i] < 123-step_move {
			str_slice_dst[i] = str_slice_src[i] + step_move
		} else { //7. 如果明文字符超出范围，则加上位移后的步长减去 26
			str_slice_dst[i] = str_slice_src[i] + step_move - 26
		}
	}
	//8. 输出结果
	fmt.Println("The encryption result is:", step_move, str_slice_dst, string(str_slice_dst))
	return string(str_slice_dst)
}

//2. 凯撒密码解密
func caesarDe(strCipher string) string {
	//1. 将文本转为小写
	str_cipher := strings.ToLower(strCipher)
	//2. 替代步长
	step_move := STEP
	//3. 将字符串转换为明文字符切片
	str_slice_src := []byte(str_cipher)
	fmt.Println("Ciphertext character slice:", str_slice_src)

	//4. 创建一个密文字符切片
	str_slice_dst := str_slice_src

	//5. 循环处理字符文本切片
	for i := 0; i < len(str_slice_src); i++ {
		//6. 如果当前周期的明文特征在位移范围内，请直接添加位移步骤以保存密文字符切片
		if str_slice_src[i] >= 97+step_move {
			str_slice_dst[i] = str_slice_src[i] - step_move
		} else { //7. 如果明文字符超出范围，则加上 26 减去位移后的步长
			str_slice_dst[i] = str_slice_src[i] + 26 - step_move
		}
	}
	//8. Output results
	fmt.Println("The decryption result is:", step_move, str_slice_dst, string(str_slice_dst))
	return string(str_slice_dst)
}

func main() {
	a := "abc"
	str := caesarEn(a)
	fmt.Println(str)
	fmt.Println(caesarDe(str))
}
