package beginners

import (
	"bytes"
	"io"
	"os"
	"testing"
)

func assert(t *testing.T, fnc func(r io.Reader), input string) string {
	t.Helper()

	// 既存のStdoutを退避する
	orgStdout := os.Stdout
	defer func() {
		// 出力先を元に戻す
		os.Stdout = orgStdout
	}()
	// パイプの作成(r: Reader, w: Writer)
	r, w, _ := os.Pipe()
	// Stdoutの出力先をパイプのwriterに変更する
	os.Stdout = w
	// 標準入力に渡したい文字列をio.Readerインターフェイスに合うように変更
	ioReader := bytes.NewBufferString(input)
	// テスト対象の関数を実行する
	fnc(ioReader)
	// Writerをクローズする
	// Writerオブジェクトはクローズするまで処理をブロックするので注意
	w.Close()
	// Bufferに書き込こまれた内容を読み出す
	var buf bytes.Buffer
	if _, err := buf.ReadFrom(r); err != nil {
		t.Fatalf("failed to read buf: %v", err)
	}
	// 文字列を取得する
	return buf.String()
}

func TestWelcomeToAtcoder1(t *testing.T) {
	input := `1
2 3
test
`
	output := `6 test
`
	result := assert(t, WelcomeToAtcoder, input)
	if output != result {
		t.Errorf("actual %v\nwant %v", result, output)
	}
}

func TestWelcomeToAtcoder2(t *testing.T) {
	input := `72
128 256
myonmyon
`
	output := `456 myonmyon
`
	result := assert(t, WelcomeToAtcoder, input)
	if output != result {
		t.Errorf("actual %v\nwant %v", result, output)
	}
}

func TestProduct1(t *testing.T) {
	input := `3 4
`
	output := `Even
`
	result := assert(t, Product, input)
	if output != result {
		t.Errorf("actual %v\nwant %v", result, output)
	}
}

func TestProduct2(t *testing.T) {
	input := `1 21
`
	output := `Odd
`
	result := assert(t, Product, input)
	if output != result {
		t.Errorf("actual %v\nwant %v", result, output)
	}
}

func TestPlacingMarbles1(t *testing.T) {
	input := `101
`
	output := `2
`
	result := assert(t, PlacingMarbles, input)
	if output != result {
		t.Errorf("actual %v\nwant %v", result, output)
	}
}

func TestPlacingMarbles2(t *testing.T) {
	input := `000
`
	output := `0
`
	result := assert(t, PlacingMarbles, input)
	if output != result {
		t.Errorf("actual %v\nwant %v", result, output)
	}
}

func TestShiftOnly1(t *testing.T) {
	input := `3
8 12 40
`
	output := `2
`
	result := assert(t, ShiftOnly, input)
	if output != result {
		t.Errorf("actual %v\nwant %v", result, output)
	}
}

func TestShiftOnly2(t *testing.T) {
	input := `4
5 6 8 10
`
	output := `0
`
	result := assert(t, ShiftOnly, input)
	if output != result {
		t.Errorf("actual %v\nwant %v", result, output)
	}
}

func TestShiftOnly3(t *testing.T) {
	input := `6
382253568 723152896 37802240 379425024 404894720 471526144
`
	output := `8
`
	result := assert(t, ShiftOnly, input)
	if output != result {
		t.Errorf("actual %v\nwant %v", result, output)
	}
}
