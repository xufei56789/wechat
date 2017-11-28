package function

//截取字符串 sinfo:需要截取的字符串 start:截取字符串的开始位置 end:截取字符串的结束位置 (返回截取好的字符串)
func RegexpString(sinfo string , start int , end int) (string) {
	rtime := []rune(sinfo)
	tissan := string(rtime[start:end])
	return tissan
}