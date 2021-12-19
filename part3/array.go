var a int[4]
a[0] = 1
i:=a[0]

b := [2]string{"jay", "kevin"}

b := [...]string{"Penn", "Teller"}

letters := []string{"a", "b", "c"}

func make([]T, len, cap) []T

var s []byte
s = make([]byte, 5, 5)

s := make(x, 5)

len(s) == 5
cap(s) == 5

b := []byte{'g', 'o', 'l', 'a', 'n', 'g'}
b[1:4] == []byte{'o', 'l', 'a'}

var x = [4]int{1, 2, 3 4}