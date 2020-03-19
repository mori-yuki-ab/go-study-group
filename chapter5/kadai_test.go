package chapter5

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidation(t *testing.T) {
	t.Run("正常: チェックが正常に完了", func(t *testing.T) {
		t.Parallel()
		if err := Validation(1, 1); err != nil {
			t.Fail()
		}
	})

	t.Run("異常: ファイルパス未指定", func(t *testing.T) {
		t.Parallel()
		err := Validation(0, 1)
		assert.EqualError(t, err, "ファイルパスを指定してください。")
	})

	t.Run("異常: -f未指定", func(t *testing.T) {
		t.Parallel()
		err := Validation(1, 0)
		assert.EqualError(t, err, "-f は1以上である必要があります。")
	})
}

func TestCutExecute(t *testing.T) {
	t.Run("正常: 正常にcut実行できる", func(t *testing.T) {
		t.Parallel()
		bufferString := bytes.NewBufferString("aaa,bbb")
		buffer := new(bytes.Buffer)
		err := CutExecute(bufferString, buffer, 1, ",")
		assert.NoError(t, err)
		assert.Equal(t, "aaa", string(buffer.Bytes()))
	})

	t.Run("異常: fieldNumberに該当するデータなし", func(t *testing.T) {
		t.Parallel()
		bufferString := bytes.NewBufferString("aaa,bbb")
		buffer := new(bytes.Buffer)
		err := CutExecute(bufferString, buffer, 3, ",")
		assert.EqualError(t, err, "-fの値に該当するデータがありません")
	})
}

func BenchmarkCutExecute(b *testing.B) {
	b.ResetTimer()
	bufferString := bytes.NewBufferString("aaa,bbb")
	buffer := new(bytes.Buffer)
	for i := 0; i < b.N; i++ {
		CutExecute(bufferString, buffer, 1, ",")
	}
}
