package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
)

func main() {
	// 引数を取得
	if len(os.Args) < 2 {
		log.Fatal("画像ファイルへのパスを指定してください")
	}
	imagePath := os.Args[1]

	// 画像ファイルを読み込む
	file, err := os.Open(imagePath)
	if err != nil {
		log.Fatalf("画像ファイルの読み込みに失敗しました: %v", err)
	}
	defer file.Close()

	// 一時ファイルを作成
	tempFile, err := os.CreateTemp("", "image-*.png")
	if err != nil {
		log.Fatalf("一時ファイルの作成に失敗しました: %v", err)
	}
	defer os.Remove(tempFile.Name()) // 一時ファイルを削除

	// 画像データを一時ファイルにコピー
	if _, err := io.Copy(tempFile, file); err != nil {
		log.Fatalf("一時ファイルへの書き込みに失敗しました: %v", err)
	}
	if err := tempFile.Close(); err != nil {
		log.Fatalf("一時ファイルのクローズに失敗しました: %v", err)
	}

	// pbcopyを使用して画像をクリップボードにコピー
	cmd := exec.Command("osascript", "-e", fmt.Sprintf(`
        set the clipboard to (read (POSIX file "%s") as JPEG picture)
    `, tempFile.Name()))

	if err := cmd.Run(); err != nil {
		log.Fatalf("クリップボードへのコピーに失敗しました: %v", err)
	}

	fmt.Println("画像がクリップボードにコピーされました")
}
