package baidu

import (
	"github.com/UlinoyaPed/BaiduTranslate"
	"log"
	"subtitle-translator-ai/utils/xlog"
)

type baiduTranslator struct {
	btr *BaiduTranslate.BaiduInfo
}

func NewClient(appid, secret string) *baiduTranslator {
	btr := BaiduTranslate.BaiduInfo{AppID: appid, SecretKey: secret}
	return &baiduTranslator{
		btr: &btr,
	}
}

func (that *baiduTranslator) Dial() error {
	return nil
}

func (that *baiduTranslator) Translate(message, parentMessageId, systemPrompt string) (messageId, content string, err error) {
	// 传入：(原文, 原文语言, 译文语言)

	// 完整实例
	s1 := that.btr.NormalTr(message, "en", "zh") // 对原文进行了url encode，原文可带空格
	if s1.Err() != nil {
		log.Println(xlog.Error(s1.Err()))
		return "", "", s1.Err()
	} else {
		return "", s1.Dst, nil
	}
}
