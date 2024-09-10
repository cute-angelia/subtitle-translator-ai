package tencent

import (
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/profile"
	"github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/common/regions"
	tmt "github.com/tencentcloud/tencentcloud-sdk-go/tencentcloud/tmt/v20180321"
)

type tencentTranslator struct {
	client *tmt.Client
}

func NewClient(appid, secret string) (*tencentTranslator, error) {
	credential := common.NewCredential(appid, secret)
	client, err := tmt.NewClient(credential, regions.Guangzhou, profile.NewClientProfile())
	return &tencentTranslator{
		client: client,
	}, err
}

func (that *tencentTranslator) Dial() error {
	return nil
}

func (that *tencentTranslator) Translate(message, parentMessageId, systemPrompt string) (messageId, content string, err error) {
	id := int64(0)

	// 语种识别
	// 如果需要，可以把注释打开
	//languageRequest := tmt.NewLanguageDetectRequest()
	//languageRequest.Text = &message
	//languageRequest.ProjectId = &id
	//languageResponse, _ := that.client.LanguageDetect(languageRequest)
	//lang := *languageResponse.Response.Lang

	lang := "en"
	tar := "zh"

	request := tmt.NewTextTranslateRequest()
	// request.Source = &lang
	request.Source = &lang
	//if lang == "zh" {
	//	tar = "en"
	//} else {
	//	tar = "zh"
	//}
	request.SourceText = &message
	request.Target = &tar
	request.ProjectId = &id
	response, err := that.client.TextTranslate(request)

	return "", *response.Response.TargetText, err
}
