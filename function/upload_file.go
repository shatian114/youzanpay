package youzan_func

import (
	"bytes"
	"context"
	"encoding/json/v2"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"strconv"
	"time"

	"github.com/shatian114/verifysign"
	youzan "github.com/shatian114/youzanpay"
	youzan_model "github.com/shatian114/youzanpay/model"
)

// 上传文件
func (s *Client) UpFile(ctx context.Context, bArr []byte) (imgKey string, err error) {
	timestampStr := strconv.FormatInt(time.Now().Unix(), 10)
	noncestr := youzan.RandStr(32)
	signStr := fmt.Sprintf("%s\n%s\n%s\n\n%s\n", "media.photo.upload", timestampStr, noncestr, s.MchId)
	sign, err := verifysign.BuildSign(ctx, signStr, s.PrivKey)

	body := bytes.NewBuffer(nil)
	writer := multipart.NewWriter(body)
	// 创建文件字段
	part, _ := writer.CreateFormFile("file", "1gtgrt.jpg")
	// part.Write(bArr)
	io.Copy(part, bytes.NewReader(bArr))
	writer.Close()
	contentType := writer.FormDataContentType()

	req, err := http.NewRequest(http.MethodPost, youzan.UpApiUrl, body)
	if err != nil {
		err = fmt.Errorf("创建http请求出错: %s", err.Error())
		return
	}
	req.Header.Set("Content-Type", contentType)
	req.Header.Set("signtype", "SHA256withRSA")
	req.Header.Set("publickeyid", s.PubKeyId)
	req.Header.Set("timestamp", timestampStr)
	req.Header.Set("noncestr", noncestr)
	req.Header.Set("sign", sign)
	req.Header.Set("mchid", s.MchId)
	httpClient := http.Client{}
	res, err := httpClient.Do(req)
	if err != nil {
		err = fmt.Errorf("http 请求出错: %s", err.Error())
		return
	}
	defer res.Body.Close()
	resBArr, err := io.ReadAll(res.Body)
	if err != nil {
		err = fmt.Errorf("读取 http 结果出错: %s", err.Error())
		return
	}
	resCommon := youzan_model.YouZanResCommon[youzan_model.UploadFileRes]{}
	err = json.Unmarshal(resBArr, &resCommon)
	if err != nil {
		err = fmt.Errorf("转换有赞结果出错: %s", err)
		return
	}
	if !resCommon.Success {
		err = fmt.Errorf("有赞上传文件失败: %s", resCommon.Message)
		return
	}

	if err != nil {
		err = fmt.Errorf("解码有赞上传文件的imageKey出错: %s", err.Error())
		return
	}
	if resCommon.Data.ImageKey == "" {
		err = fmt.Errorf("有赞上传文件提示成功，但是没有返回imageKey")
		return
	}
	imgKey = resCommon.Data.ImageKey
	return
}
