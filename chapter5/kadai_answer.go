package chapter5

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

// Cut is ...
// go-cutコマンドを実装しよう
func Cut() {
	flag.Parse()

	// チェック
	if err := Validation(flag.NArg(), *fields); err != nil {
		log.Fatal(err)
	}

	file, err := os.Open(flag.Args()[0])
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// cut実行
	if err := CutExecute(file, os.Stdout, *fields, *delimiter); err != nil {
		log.Fatal(err)
	}
}

// Validation is ...
// チェック
func Validation(arg, fieldNumber int) error {
	if arg == 0 {
		return fmt.Errorf("ファイルパスを指定してください。")
	}
	if fieldNumber <= 0 {
		return fmt.Errorf("-f は1以上である必要があります。")
	}
	return nil
}

// CutExecute is ...
// cut実行
func CutExecute(ioReader io.Reader, ioWriter io.Writer, fieldNumber int, delimiter string) error {
	scanner := bufio.NewScanner(ioReader)
	writer := bufio.NewWriter(ioWriter)
	for scanner.Scan() {
		sb := strings.Split(scanner.Text(), delimiter)
		if len(sb) <= fieldNumber-1 {
			return fmt.Errorf("-fの値に該当するデータがありません")
		}
		fmt.Fprintf(writer, sb[fieldNumber-1])
	}
	writer.Flush()
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}
