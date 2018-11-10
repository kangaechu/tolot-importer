package main

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
	"time"

	"github.com/chromedp/chromedp"
	"github.com/spf13/viper"
)

func main() {
	// settings.yamlから設定を読み込み
	viper.SetConfigName("settings")  // 設定ファイル名を拡張子抜きで指定する
	viper.AddConfigPath(".")         // 現在のワーキングディレクトリを探索することもできる
	vipererr := viper.ReadInConfig() // 設定ファイルを探索して読み取る
	if vipererr != nil {             // 設定ファイルの読み取りエラー対応
		panic(fmt.Errorf("設定ファイル読み込みエラー: %s", vipererr))
	}
	var userID = viper.GetString("userID")
	var password = viper.GetString("password")
	var inFileName = viper.GetString("addressFileName")

	// アドレス帳を開く
	file, err := os.Open(inFileName)
	if err != nil {
		panic(fmt.Errorf("アドレス帳がありません: %s", err))
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.LazyQuotes = true // ダブルクオートを厳密にチェックしない

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome instance
	c, err := chromedp.New(ctxt, chromedp.WithErrorf(log.Printf))
	if err != nil {
		log.Fatal(err)
	}

	// ログイン→TOLOT連絡帳への遷移
	var site, res string
	err = c.Run(ctxt, login(userID, password, &site, &res))
	if err != nil {
		log.Fatal(err)
	}

	for {
		record, err := reader.Read() // 1行読み出す
		if err == io.EOF {
			break
		} else {
			if err != nil {
				log.Fatal(err)
			}
			err = c.Run(ctxt, addNewContact(record, &site, &res))
			if err != nil {
				log.Fatal(err)
			}
		}
	}

	// shutdown chrome
	err = c.Shutdown(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	// wait for chrome to finish
	err = c.Wait()
	if err != nil {
		log.Fatal(err)
	}
}

func login(userID string, password string, site, res *string) chromedp.Tasks {
	tasks := chromedp.Tasks{
		// ログインページ
		chromedp.Navigate(`https://tolot.com/member/login`),
		chromedp.WaitVisible(`input[name="password"]`, chromedp.ByQuery),
		chromedp.SendKeys(`input[name="user_name"]`, userID, chromedp.ByQuery),
		chromedp.SendKeys(`input[name="password"]`, password, chromedp.ByQuery),
		chromedp.Click(`input[type="submit"]`, chromedp.ByQuery),
		// マイページ
		chromedp.WaitVisible(`.m-content-menu--contact`, chromedp.ByQuery),
		chromedp.Click(`.m-content-menu--contact`, chromedp.ByQuery),
		// TOLOT連絡帳
		chromedp.WaitVisible(`a[href="/member/contact/new"]`, chromedp.ByQuery),
	}
	return tasks
}

func addNewContact(record []string, site, res *string) chromedp.Tasks {
	println(record[0] + record[1])
	zip := strings.Replace(record[3], "-", "", 1)

	tasks := chromedp.Tasks{
		// TOLOT連絡帳に追加
		chromedp.Click(`a[href="/member/contact/new"]`, chromedp.ByQuery),
		chromedp.WaitVisible(`input.m-input--name-last`, chromedp.ByQuery),
		chromedp.SendKeys(`input.m-input--name-last`, record[0], chromedp.ByQuery),
		chromedp.SendKeys(`.m-input--name-first`, record[1], chromedp.ByQuery),
		chromedp.SendKeys(`.m-input--joint_name-first1`, record[2], chromedp.ByQuery),
		chromedp.SendKeys(`.m-input--zip`, zip, chromedp.ByQuery),
		chromedp.SendKeys(`.m-input--state`, record[4], chromedp.ByQuery),
		chromedp.SendKeys(`.m-input--address1`, record[5], chromedp.ByQuery),
		chromedp.SendKeys(`.m-input--address2`, record[6], chromedp.ByQuery),
		chromedp.Sleep(3 * time.Second),
		chromedp.Click(`input[type="submit"]`, chromedp.ByQuery),
		// OKボタン
		chromedp.WaitVisible(`a.resolve`, chromedp.ByQuery),
		chromedp.Click(`a.resolve`, chromedp.ByQuery),
		// TOLOT連絡帳
		chromedp.WaitVisible(`a[href="/member/contact/new"]`, chromedp.ByQuery),
	}
	return tasks
}
