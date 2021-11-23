package templatemethod

import "fmt"

// Printerインターフェース
type Printer interface {
	open()
	print()
	close()
}

// AbstractDisplay構造体
type AbstractDisplay struct {
	printer Printer
}

// AbstractDisplayにポインタ指定
func (a *AbstractDisplay) Display() {
	a.printer.open()
	for i := 0; i < 5; i++ {
		a.printer.print()
	}
	a.printer.close()
}

// CharDisplay講座体
type CharDisplay struct {
	*AbstractDisplay
	ch byte
}

// memo 関数からポインタを返すべきなのか
// 引数やレシーバを関数内で書き換える必要がある場合は使う
// 逆にそれ以外の場面ではわざわざ利用する必要はない
// https://zenn.dev/uji/articles/f6ab9a06320294146733
// https://qiita.com/momotaro98/items/5518210c1affc29a489d

func NewCharDisplay(ch byte) *CharDisplay {
	// &はアドレス演算子
	charDisplay := &CharDisplay{
		AbstractDisplay: &AbstractDisplay{},
		ch:              ch,
	}
	charDisplay.printer = charDisplay
	return charDisplay
}

func (c *CharDisplay) open() {
	fmt.Print("<<")
}
func (c *CharDisplay) print() {
	fmt.Print(string(c.ch))
}
func (c *CharDisplay) close() {
	fmt.Println(">>")
}

// StringDisplay構造体
type StringDisplay struct {
	*AbstractDisplay
	str   string
	width int
}

func NewStringDisplay(str string) *StringDisplay {
	stringDisplay := &StringDisplay{
		AbstractDisplay: &AbstractDisplay{},
		str:             str,
		width:           len(str),
	}
	stringDisplay.printer = stringDisplay
	return stringDisplay
}

func (s *StringDisplay) open() {
	s.printLine()
}
func (s *StringDisplay) print() {
	fmt.Printf("%s%s%s\n", "|", s.str, "|")
}
func (s *StringDisplay) close() {
	s.printLine()
}

func (s *StringDisplay) printLine() {
	fmt.Print("+")
	for i := 0; i < s.width; i++ {
		fmt.Print("-")
	}
	fmt.Println("+")
}
