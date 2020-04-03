package youdao

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/go-errors/errors"
	json "github.com/json-iterator/go"
	uuid "github.com/nu7hatch/gouuid"
	"github.com/qshuai/keyhot/fetcher"
)

type YouDaoFetcher struct {
	appId  string
	appKey string
}

func New(appId, appKey string, args ...string) fetcher.Fetcher {
	return &YouDaoFetcher{
		appId:  appId,
		appKey: appKey,
	}
}

type YouDaoResult struct {
	Query       string          `json:"query,omitempty"`
	Translation []string        `json:"translation,omitempty"`
	ErrorCode   string          `json:"errorCode"`
	Web         []ExplainDetail `json:"web,omitempty"`
}

type ExplainDetail struct {
	Key   string   `json:"key"`
	Value []string `json:"value,omitempty"`
}

func (y *YouDaoFetcher) Translate(word string) (*fetcher.Result, error) {
	uid, err := uuid.NewV4()
	if err != nil {
		return nil, err
	}

	timestamp := time.Now().Unix()
	sign := signature(y.appId, y.appKey, word, uid.String(), timestamp)

	params := fmt.Sprintf("from=en&to=zh-CHS&appKey=%s&salt=%s&sign=%s&signType=v3&curtime=%d&q=%s", y.appId, uid.String(), sign, timestamp, word)
	resp, err := http.Get("https://openapi.youdao.com/api?" + params)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var result YouDaoResult
	err = json.Unmarshal(body, &result)
	if err != nil {
		return nil, err
	}
	if result.ErrorCode != "0" {
		return nil, errors.New(fmt.Sprintf("response failed with non-zero code: %s", string(body)))
	}

	ret := &fetcher.Result{
		Origin: result.Query,
		Target: strings.Join(result.Translation, " "),
	}
	examples := make([]fetcher.Example, 0, len(result.Web))
	for _, item := range result.Web {
		examples = append(examples, fetcher.Example{
			Word:    item.Key,
			Explain: strings.Join(item.Value, " "),
		})
	}
	ret.Examples = examples

	return ret, nil
}

func signature(appId, appKey, word, salt string, timestamp int64) string {
	input := word
	if len(word) > 20 {
		input = word[0:10] + strconv.Itoa(len(word)) + word[len(word)-10:]
	}

	sum := sha256.Sum256([]byte(appId + input + salt + strconv.FormatInt(timestamp, 10) + appKey))
	return hex.EncodeToString(sum[:])
}
