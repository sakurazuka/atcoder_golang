package beginners

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func WelcomeToAtcoder(r io.Reader) {
	scanner := bufio.NewScanner(r)
	var a, b, c int
	var arr []string
	var s string

	scanner.Scan()
	a, _ = strconv.Atoi(scanner.Text())

	scanner.Scan()
	arr = strings.Split(scanner.Text(), " ")
	b, _ = strconv.Atoi(arr[0])
	c, _ = strconv.Atoi(arr[1])

	scanner.Scan()
	s = scanner.Text()

	fmt.Printf("%d %s\n", a+b+c, s)
}

func Product(r io.Reader) {
	scanner := bufio.NewScanner(r)
	var a, b int
	var arr []string
	var s string

	scanner.Scan()
	arr = strings.Split(scanner.Text(), " ")
	a, _ = strconv.Atoi(arr[0])
	b, _ = strconv.Atoi(arr[1])

	if (a*b)%2 == 0 {
		s = "Even"
	} else {
		s = "Odd"
	}
	fmt.Printf("%s\n", s)
}

func PlacingMarbles(r io.Reader) {
	scanner := bufio.NewScanner(r)
	var a, b, c int
	var arr []string

	scanner.Scan()
	arr = strings.Split(scanner.Text(), "")
	a, _ = strconv.Atoi(arr[0])
	b, _ = strconv.Atoi(arr[1])
	c, _ = strconv.Atoi(arr[2])

	fmt.Printf("%d\n", a+b+c)
}

func ShiftOnly(r io.Reader) {
	// 操作通りに処理をするパターン
	scanner := bufio.NewScanner(r)
	var a, b int
	var m []int

	scanner.Scan()
	_, _ = strconv.Atoi(scanner.Text())

	scanner.Scan()
	for _, x := range strings.Split(scanner.Text(), " ") {
		a, _ = strconv.Atoi(x)
		m = append(m, a)
	}

	var flag bool
	for {
		flag = false
		for _, y := range m {
			flag = (y%2 == 1)
			if flag {
				break
			}
		}
		if flag {
			break
		}

		for i, z := range m {
			m[i] = (z / 2)
		}

		b += 1
	}

	fmt.Printf("%d\n", b)

	// 線形探索で解くパターン

	// ２進数を使って解くパターン

	// fmt.Printf("%d\n", x)
}
