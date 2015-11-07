package lib

import (
	. "github.com/r7kamura/gospel"
	"path"
	"runtime"
	"testing"
)

func TestInterpretor(t *testing.T) {
	var (
		text     string
		expected []byte
		actual   []byte
		sut2     *Interpreter
		err      error
	)
	Describe(t, "初期化", func() {
		Context("インスタンス生成を生成するとき", func() {
			Before(func() {
				text = "東北ずんだ太ももの3拍子"
				sut2 = NewInterpreter(text)
			})
			It("nilにならないこと", func() {
				Expect(sut2).To(Exist)
			})
			It("originに値が設定されていること", func() {
				Expect(sut2.origin).To(Equal, text)
			})
		})
	})
	Describe(t, "変換処理", func() {
		Context("東北を置き換えるとき", func() {
			Before(func() {
				data := "東北"
				expected = []byte{'Z'}
				sut2 = NewInterpreter(data)
				actual, err = sut2.toChar()
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("Zが返ってくること", func() {
				Expect(actual).To(Equal, expected)
			})
		})
		Context("ずんだを置き換えるとき", func() {
			Before(func() {
				data := "ずんだ"
				expected = []byte{'N'}
				sut2 = NewInterpreter(data)
				actual, err = sut2.toChar()
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("Nが返ってくること", func() {
				Expect(actual).To(Equal, expected)
			})
		})
		Context("太ももを置き換えるとき", func() {
			Before(func() {
				data := "太もも"
				expected = []byte{'K'}
				sut2 = NewInterpreter(data)
				actual, err = sut2.toChar()
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("Kが返ってくること", func() {
				Expect(actual).To(Equal, expected)
			})
		})
		Context("変換対象に不要な文字が含まれているとき", func() {
			Before(func() {
				data := "東北ずんだ太ももの3拍子"
				expected = []byte{'Z', 'N', 'K'}
				sut2 = NewInterpreter(data)
				actual, err = sut2.toChar()
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("ZNKが返ってくること", func() {
				Expect(actual).To(Equal, expected)
			})
		})
	})
	Describe(t, "コードを生成する", func() {
		Context("スタックに0x41をpushするコマンドを生成するとき", func() {
			Before(func() {
				data := "東北東北東北東北太もも東北東北東北東北東北太ももずんだ"
				sut2 = NewInterpreter(data)
				err = sut2.toCode()
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("コマンドが１つだけ作成されること", func() {
				Expect(sut2.commands.Len()).To(Equal, 1)
			})
			It("stack push 0x41 コマンドになっていること", func() {
				Expect(sut2.commands.Get(1)).To(Equal, newSubCommandWithParam("stack", "push", 0x41))
			})
		})
		Context("スタックに0x43をpushするコマンドを生成するとき", func() {
			Before(func() {
				data := "東北東北東北太もも東北東北東北東北太もも太ももずんだ"
				sut2 = NewInterpreter(data)
				err = sut2.toCode()
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("コマンドが１つだけ作成れること", func() {
				Expect(sut2.commands.Len()).To(Equal, 1)
			})
			It("stack push 0x43 コマンドになっていること", func() {
				Expect(sut2.commands.Get(1)).To(Equal, newSubCommandWithParam("stack", "push", 0x43))
			})
		})
		Context("スタックに0x41をpushするコマンドを２つ生成するとき", func() {
			Before(func() {
				data := "東北, 東北, 東北, 太もも, 東北, 東北, 東北, 東北, 東北, 太もも, ずんだ, 東北, 東北, 東北, 太もも, 東北, 東北, 東北, 東北, 東北, 太もも, ずんだ"
				sut2 = NewInterpreter(data)
				err = sut2.toCode()
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("コマンドが2つ作成れること", func() {
				Expect(sut2.commands.Len()).To(Equal, 2)
			})
			It("１つ目のコマンドがコマンドがstack push 0x41 コマンドになっていること", func() {
				Expect(sut2.commands.Get(1)).To(Equal, newSubCommandWithParam("stack", "push", 0x41))
			})
			It("2つ目のコマンドがコマンドがstack push 0x41 コマンドになっていること", func() {
				Expect(sut2.commands.Get(2)).To(Equal, newSubCommandWithParam("stack", "push", 0x41))
			})
		})
	})
	Describe(t, "不要な文字をフィルタリングする", func() {
		Context("不要な文字AとBと\rが含まれるとき", func() {
			Before(func() {
				data := "'A', 東北, 'B', 太もも, '\r', ずんだ"
				sut2 = NewInterpreter(data)
				sut2.filter()
			})
			It("AとBと\rが排除されていること", func() {
				Expect(sut2.source).To(Equal, []byte{'Z', 'K', 'N'})
			})
		})
	})
}

func current_dir() string {
	_, fpath, _, _ := runtime.Caller(0)
	return path.Dir(fpath)
}
