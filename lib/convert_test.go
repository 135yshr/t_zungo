package lib

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestConerter(t *testing.T) {
	var (
		sut      func([]byte) (*Command, int, error)
		cmd      *Command
		seek     int
		expected []byte
		err      error
	)
	Describe(t, "コンバーターのインスタンスを作成する", func() {
		Context("東北を渡したとき", func() {
			Before(func() {
				sut, err = createFunction(byte('Z'))
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("インスタンスが返ってくること", func() {
				Expect(sut).To(Exist)
			})
		})
		Context("ずんだを渡したとき", func() {
			Before(func() {
				sut, err = createFunction(byte('N'))
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("インスタンスが返ってくること", func() {
				Expect(sut).To(Exist)
			})
		})
		Context("太ももを渡したとき", func() {
			Before(func() {
				sut, err = createFunction(byte('K'))
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("インスタンスが返ってくること", func() {
				Expect(sut).To(Exist)
			})
		})
		Context("東北ずんだ太もも以外の文字が渡されたとき", func() {
			Before(func() {
				sut, err = createFunction(byte('A'))
			})
			It("インスタンスがnilなること", func() {
				Expect(sut).To(NotExist)
			})
			It("エラーが発生すること", func() {
				Expect(err).To(Exist)
			})
		})
	})
	Describe(t, "スタックを操作する", func() {
		Context("SSKNを渡したとき", func() {
			Before(func() {
				expected = []byte{'Z', 'Z', 'K', 'N'}
				cmd, seek, err = stackManipulation(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("stack push 1 コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newSubCommandWithParam("stack", "push", 1))
			})
		})
		Context("ZZKZNを渡したとき", func() {
			Before(func() {
				expected = []byte{'Z', 'Z', 'K', 'Z', 'N'}
				cmd, seek, err = stackManipulation(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("stack push 2 コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newSubCommandWithParam("stack", "push", 2))
			})
		})
		Context("ZZKZZNを渡したとき", func() {
			Before(func() {
				expected = []byte{'Z', 'Z', 'K', 'Z', 'Z', 'N'}
				cmd, seek, err = stackManipulation(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("stack push 4 コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newSubCommandWithParam("stack", "push", 4))
			})
		})
		Context("ZKKNを渡したとき", func() {
			Before(func() {
				expected = []byte{'Z', 'K', 'K', 'N'}
				cmd, seek, err = stackManipulation(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("stack push -1 コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newSubCommandWithParam("stack", "push", -1))
			})
		})
		Context("LSを渡したとき", func() {
			Before(func() {
				expected = []byte{'N', 'Z'}
				cmd, seek, err = stackManipulation(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("stack copy コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newSubCommand("stack", "copy"))
			})
		})
		Context("NKを渡したとき", func() {
			Before(func() {
				expected = []byte{'N', 'K'}
				cmd, seek, err = stackManipulation(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("stack swap コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newSubCommand("stack", "swap"))
			})
		})
		Context("NNを渡したとき", func() {
			Before(func() {
				expected = []byte{'N', 'N'}
				cmd, seek, err = stackManipulation(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("stack remove コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newSubCommand("stack", "remove"))
			})
		})
		Context("KZZKNを渡したとき", func() {
			Before(func() {
				expected = []byte{'K', 'Z', 'Z', 'K', 'N'}
				cmd, seek, err = stackManipulation(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("stack ncopy 1 コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newSubCommandWithParam("stack", "ncopy", 1))
			})
		})
		Context("KNZKNを渡したとき", func() {
			Before(func() {
				expected = []byte{'K', 'N', 'Z', 'K', 'N'}
				cmd, seek, err = stackManipulation(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("stack move 1 コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newSubCommandWithParam("stack", "move", 1))
			})
		})
		Context("KKを渡したとき", func() {
			Before(func() {
				expected = []byte{'K', 'K'}
				cmd, seek, err = stackManipulation(expected)
			})
			It("コマンドが存在しないこと", func() {
				Expect(cmd).To(NotExist)
			})
			It("エラーが発生すること", func() {
				Expect(err).To(Exist)
			})
			It("読み込んだ文字数が0であること", func() {
				Expect(seek).To(Equal, 0)
			})
		})
	})
	Describe(t, "制御に関する命令", func() {
		Context("ZZKZZKNを渡したとき", func() {
			Before(func() {
				expected = []byte{'Z', 'Z', 'K', 'Z', 'Z', 'K', 'N'}
				cmd, seek, err = flowControl(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("label 1001 コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newSubCommand("label", "1001"))
			})
		})
		Context("ZKKZZKNを渡したとき", func() {
			Before(func() {
				expected = []byte{'Z', 'K', 'K', 'Z', 'Z', 'K', 'N'}
				cmd, seek, err = flowControl(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("call 1001 コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newSubCommand("call", "1001"))
			})
		})
		Context("ZNKZZKNを渡したとき", func() {
			Before(func() {
				expected = []byte{'Z', 'N', 'K', 'Z', 'Z', 'K', 'N'}
				cmd, seek, err = flowControl(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("goto 1001 コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newSubCommand("goto", "1001"))
			})
		})
		Context("KZKZZKNを渡したとき", func() {
			Before(func() {
				expected = []byte{'K', 'Z', 'K', 'Z', 'Z', 'K', 'N'}
				cmd, seek, err = flowControl(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("if 0 goto 1001 コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newSubCommand("if stack==0 then goto", "1001"))
			})
		})
		Context("KKKZZKNを渡したとき", func() {
			Before(func() {
				expected = []byte{'K', 'K', 'K', 'Z', 'Z', 'K', 'N'}
				cmd, seek, err = flowControl(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("if <0 goto 1001 コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newSubCommand("if stack<0 then goto", "1001"))
			})
		})
		Context("KNを渡したとき", func() {
			Before(func() {
				expected = []byte{'K', 'N'}
				cmd, seek, err = flowControl(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("return コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newCommand("return"))
			})
		})
		Context("NNを渡したとき", func() {
			Before(func() {
				expected = []byte{'N', 'N'}
				cmd, seek, err = flowControl(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("exit コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newCommand("exit"))
			})
		})
		Context("NKを渡したとき", func() {
			Before(func() {
				expected = []byte{'N', 'K'}
				cmd, seek, err = flowControl(expected)
			})
			It("コマンドが存在しないこと", func() {
				Expect(cmd).To(NotExist)
			})
			It("エラーが発生すること", func() {
				Expect(err).To(Exist)
			})
			It("読み込んだ文字数が0であること", func() {
				Expect(seek).To(Equal, 0)
			})
		})
	})
	Describe(t, "演算命令の関数", func() {
		Context("ZZを渡したとき", func() {
			Before(func() {
				expected = []byte{'Z', 'Z'}
				cmd, seek, err = arithmetic(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("add コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newCommand("add"))
			})
		})
		Context("ZKを渡したとき", func() {
			Before(func() {
				expected = []byte{'Z', 'K'}
				cmd, seek, err = arithmetic(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("sub コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newCommand("sub"))
			})
		})
		Context("ZNを渡したとき", func() {
			Before(func() {
				expected = []byte{'Z', 'N'}
				cmd, seek, err = arithmetic(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("mul コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newCommand("mul"))
			})
		})
		Context("KZを渡したとき", func() {
			Before(func() {
				expected = []byte{'K', 'Z'}
				cmd, seek, err = arithmetic(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("div コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newCommand("div"))
			})
		})
		Context("KKを渡したとき", func() {
			Before(func() {
				expected = []byte{'K', 'K'}
				cmd, seek, err = arithmetic(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("mod コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newCommand("mod"))
			})
		})
		Context("KNを渡したとき", func() {
			Before(func() {
				expected = []byte{'K', 'N'}
				cmd, seek, err = arithmetic(expected)
			})
			It("コマンドが存在しないこと", func() {
				Expect(cmd).To(NotExist)
			})
			It("エラーが発生すること", func() {
				Expect(err).To(Exist)
			})
			It("読み込んだ文字数が0であること", func() {
				Expect(seek).To(Equal, 0)
			})
		})
	})
	Describe(t, "ヒープ領域を操作する", func() {
		Context("Sを渡したとき", func() {
			Before(func() {
				expected = []byte{'Z'}
				cmd, seek, err = heapAccess(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("heap push コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newSubCommand("heap", "push"))
			})
		})
		Context("Kを渡したとき", func() {
			Before(func() {
				expected = []byte{'K'}
				cmd, seek, err = heapAccess(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("heap pop コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newSubCommand("heap", "pop"))
			})
		})
		Context("Nを渡したとき", func() {
			Before(func() {
				expected = []byte{'N'}
				cmd, seek, err = heapAccess(expected)
			})
			It("コマンドが存在しないこと", func() {
				Expect(cmd).To(NotExist)
			})
			It("エラーが発生すること", func() {
				Expect(err).To(Exist)
			})
			It("読み込んだ文字数が0であること", func() {
				Expect(seek).To(Equal, 0)
			})
		})
	})
	Describe(t, "I/O操作の関数", func() {
		Context("SSを渡したとき", func() {
			Before(func() {
				expected = []byte{'Z', 'Z'}
				cmd, seek, err = i_o(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("putc コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newCommand("putc"))
			})
		})
		Context("SKを渡したとき", func() {
			Before(func() {
				expected = []byte{'Z', 'K'}
				cmd, seek, err = i_o(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("putn コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newCommand("putn"))
			})
		})
		Context("KSを渡したとき", func() {
			Before(func() {
				expected = []byte{'K', 'Z'}
				cmd, seek, err = i_o(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("gutc コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newCommand("getc"))
			})
		})
		Context("KKを渡したとき", func() {
			Before(func() {
				expected = []byte{'K', 'K'}
				cmd, seek, err = i_o(expected)
			})
			It("エラーが発生しないこと", func() {
				Expect(err).To(NotExist)
			})
			It("読み込んだ文字数がパラメータの配列数と同じであること", func() {
				Expect(seek).To(Equal, len(expected))
			})
			It("コマンドが存在すること", func() {
				Expect(cmd).To(Exist)
			})
			It("gutn コマンドが作成されること", func() {
				Expect(cmd).To(Equal, newCommand("getn"))
			})
		})
		Context("KNを渡したとき", func() {
			Before(func() {
				expected = []byte{'K', 'N'}
				cmd, seek, err = i_o(expected)
			})
			It("コマンドが存在しないこと", func() {
				Expect(cmd).To(NotExist)
			})
			It("エラーが発生すること", func() {
				Expect(err).To(Exist)
			})
			It("読み込んだ文字数が0であること", func() {
				Expect(seek).To(Equal, 0)
			})
		})
	})
}
