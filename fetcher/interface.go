package fetcher

type Fetcher interface {
	Translate(word string) (*Result, error)
}

type Result struct {
	Origin   string    `json:"origin"`             // 被翻译的词语
	Target   string    `json:"target"`             // 翻译成的词语
	Examples []Example `json:"examples,omitempty"` // 例子列表
}

type Example struct {
	Word    string `json:"word"`              // 词语、例句
	Explain string `json:"explain,omitempty"` // 释义
}
