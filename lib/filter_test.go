package lib

import (
	. "github.com/r7kamura/gospel"
	"testing"
)

func TestFilter(t *testing.T) {
	Describe(t, "初期化", func() {
		Context("インスタンスを作成ししたとき", func() {
			It("インスタンスを取得できること", func() {
				sut := NewFilter([]string{""})
				Expect(sut).To(Exist)
			})
		})
	})
	Describe(t, "フィルタリングする", func() {
		var (
			sut *Filter
		)
		Before(func() {
			sut = NewFilter([]string{"東北", "ずんだ", "太もも"})
		})
		Context("東北サイコーを渡ししたとき", func() {
			It("東北が抽出できる", func() {
				actual := sut.Filter("東北が抽出できる")
				Expect(actual).To(Equal, map[int]string{0: "東北"})
			})
		})
		Context("東北のアイドル東北ずん子を渡ししたとき", func() {
			It("東北が２つ抽出できる", func() {
				actual := sut.Filter("東北のアイドル東北ずん子")
				Expect(actual).To(Equal, map[int]string{0: "東北", 21: "東北"})
			})
		})
		Context("東北ずん子とずんだ餅を食べようを渡ししたとき", func() {
			It("東北とずんだが抽出できる", func() {
				actual := sut.Filter("東北ずん子とずんだ餅を食べよう")
				Expect(actual).To(Equal, map[int]string{0: "東北", 18: "ずんだ"})
			})
		})
		Context("東北ずん子とずんだ餅を食べたい。ちなみにチャームポイントは太もも？を渡ししたとき", func() {
			It("東北ずんだ太ももが抽出できる", func() {
				actual := sut.Filter("東北ずん子とずんだ餅を食べたい。ちなみにチャームポイントは太もも？")
				Expect(actual).To(Equal, map[int]string{0: "東北", 18: "ずんだ", 87: "太もも"})
			})
		})
	})
}
