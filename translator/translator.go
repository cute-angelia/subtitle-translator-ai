package translator

import (
	"log"
	"subtitle-translator-ai/translator/baidu"
	"subtitle-translator-ai/translator/openai"
	"subtitle-translator-ai/translator/tencent"
	"subtitle-translator-ai/utils/xlog"
)

type TranslatorClient interface {
	Dial() error
	Translate(message, parentMessageId, systemPrompt string) (messageId, content string, err error)
}

type Translator struct {
	Engine string
	Client TranslatorClient
}

func NewBaiduTranslator(appid, secret string) (*Translator, error) {
	client := baidu.NewClient(appid, secret)
	return &Translator{
		Engine: "OpenAI",
		Client: client,
	}, nil
}

// NewTencentTranslator 腾讯翻译
// https://cloud.tencent.com/document/api/551/15619#SDK
func NewTencentTranslator(appid, secret string) (*Translator, error) {
	if client, err := tencent.NewClient(appid, secret); err != nil {
		return nil, err
	} else {
		return &Translator{
			Engine: "Tencent",
			Client: client,
		}, nil
	}
}

func MustNewOpenAITranslator(openaiKey, proxy string) *Translator {
	translator, err := NewOpenAITranslator(openaiKey, proxy)
	if err != nil {
		log.Fatal(xlog.Fatal(err))
	}
	return translator
}

func NewOpenAITranslator(openaiKey, proxy string) (*Translator, error) {
	client := openai.NewClient(openaiKey, proxy)
	if err := client.Dial(); err != nil {
		return nil, err
	}
	return &Translator{
		Engine: "OpenAI",
		Client: client,
	}, nil
}
