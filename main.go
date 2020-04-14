package main

import (
	"errors"
	"fmt"
	"log"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/go-vgo/robotgo"
	"github.com/qshuai/keyhot/fetcher"
	"github.com/qshuai/keyhot/fetcher/youdao"
	"github.com/qshuai/keyhot/gui"
	"github.com/qshuai/keyhot/storage"
	"github.com/qshuai/keyhot/storage/sqlite"
	hook "github.com/robotn/gohook"
	"github.com/spf13/viper"
)

type app struct {
	fetcher fetcher.Fetcher
	storage storage.Storage
}

func main() {
	viper.SetConfigName("app")  // name of config file (without extension)
	viper.SetConfigType("yaml") // REQUIRED if the config file does not have the extension in the name
	viper.AddConfigPath(".")    // optionally look for config in the working directory
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	key := "c"
	arr := []string{"cmd"}
	s := hook.Start()
	defer hook.End()

	fetch := youdao.New(viper.GetString("youdao.appid"), viper.GetString("youdao.appkey"))
	//store, err := mysql.New(viper.GetString("mysql.host"), viper.GetInt("mysql.port"), viper.GetString("mysql.user"), viper.GetString("mysql.passwd"), viper.GetString("mysql.database"))
	store, err := sqlite.New()
	if err != nil {
		log.Printf("connect to storage error: %s", err)
	}

	ct := false
	k := 0
	for {
		e := <-s

		l := len(arr)
		if l > 0 {
			for i := 0; i < l; i++ {
				ukey := robotgo.Keycode[arr[i]]

				if e.Kind == hook.KeyHold && e.Keycode == ukey {
					k++
				}

				if k == l {
					ct = true
				}

				if e.Kind == hook.KeyUp && e.Keycode == ukey {
					if k > 0 {
						k--
					}
					// time.Sleep(10 * time.Microsecond)
					ct = false
				}
			}
		} else {
			ct = true
		}

		if ct && e.Kind == hook.KeyUp && e.Keycode == robotgo.Keycode[key] {
			// 获取鼠标点击的词语
			text, err := clipboard.ReadAll()
			if err != nil {
				log.Printf("read clipboard error: %s", err)
				continue
			}

			cleanWord, err := filterWord(text)
			if err != nil {
				log.Printf("invalid word: %s", err)
			}
			result, err := fetch.Translate(cleanWord)
			if err != nil {
				log.Printf("fetch tanslation for word(%s) error: %s", text, err)
				continue
			}

			// 展示组件
			gui.ShowTranslation(result.Origin, result.Target)

			// 存储组件 异步化
			go func() {
				err = store.Create(result.Origin, result.Target)
				if err != nil {
					log.Printf("create a new record error: %s", err)
				}
			}()
		}
	}
}

func filterWord(word string) (string, error) {
	if len(word) <= 0 {
		return "", errors.New("empty string")
	}

	if strings.ContainsAny(word, "/。,><《》、\\|]}[{\":;；：！!`~@#￥%……&*（）+-_") {
		return "", errors.New("not valid word")
	}

	return strings.Trim(word, " .,\"'"), nil
}
