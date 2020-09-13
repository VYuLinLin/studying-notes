package main

var keyChar = []byte("0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func genKey(n int) string {
	if n == 0 {
		return string(keyChar[0])
	}
	l := len(keyChar) // 62
	s := make([]byte, 20)
	i := len(s) // 20
	for n > 0 && i >= 0 {
		i--
		j := n % i      // 问题二：为什么要用i取余
		n = (n - j) / l // 问题三： 为什么要除以l
		s[i] = keyChar[j]
	}
	return string(s[i:])
}

func main() {

}
