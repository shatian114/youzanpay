package youzan_func

import (
	"bytes"
	"context"
	"encoding/json/v2"
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"strconv"
	"strings"
	"time"

	"github.com/shatian114/verifysign"
	youzan "github.com/shatian114/youzanpay"
	youzan_model "github.com/shatian114/youzanpay/model"
)

type Client struct {
	PubKeyId string
	PrivKey  string
	PubKey   string
	MchId    string
	SerialNo string
}

func (s *Client) Request[T any](ctx context.Context, paramObj any, serverName string, apiVer string) (res youzan_model.YouZanResCommon[T], err error) {

	path := strings.ReplaceAll(serverName, "_", ".") + "/" + apiVer
	serverName = serverName + "_" + apiVer
	paramByte, err := json.Marshal(paramObj)
	if err != nil {
		err = fmt.Errorf("有赞请求，参数转json出错: %s", err.Error())
		return
	}
	bodyJson := string(paramByte)

	// 构造签名
	timestampStr := strconv.FormatInt(time.Now().Unix(), 10)
	noncestr := youzan.RandStr(32)
	signStr := fmt.Sprintf("%s\n%s\n%s\n%s\n%s\n", serverName, timestampStr, noncestr, bodyJson, s.MchId)

	sign, err := verifysign.BuildSign(ctx, signStr, s.PrivKey)

	httpReq, err := http.NewRequest(http.MethodPost, youzan.ApiUrl+path, bytes.NewReader(paramByte))
	if err != nil {
		err = fmt.Errorf("有赞请求，创建http request出错: %s", err)
		return
	}

	httpReq.Header.Set("Content-Type", "application/json")
	httpReq.Header.Set("signtype", "SHA256withRSA")
	httpReq.Header.Set("timestamp", timestampStr)
	httpReq.Header.Set("noncestr", noncestr)
	httpReq.Header.Set("sign", sign)
	httpReq.Header.Set("publickeyid", s.PubKeyId)
	httpReq.Header.Set("mchid", s.MchId)

	reqAllByte, err := httputil.DumpRequest(httpReq, true)
	if err != nil {

		youzan.Logger.Error(ctx, "获取有赞请求数据出错: ", err.Error())
	}
	youzan.Logger.Debug(ctx, "有赞请求: ", reqAllByte)

	httpClient := &http.Client{}
	httpRes, err := httpClient.Do(httpReq)

	if err != nil {
		err = fmt.Errorf("有赞接口请求出错: %s", err.Error())
		return
	}
	defer httpRes.Body.Close()

	resAllByte, err := httputil.DumpResponse(httpRes, true)
	if err != nil {
		youzan.Logger.Error(ctx, "获取有赞请响应据出错: ", err.Error())
	}
	youzan.Logger.Debug(ctx, "有赞响应: ", resAllByte)

	resByte, err := io.ReadAll(httpRes.Body)
	if err != nil {
		err = fmt.Errorf("有赞请求，读取响应byte出错: %s", err)
		return
	}
	resStr := string(resByte)

	// 判断是否为前置api错误
	preErr := youzan_model.PreErr{}
	err = json.Unmarshal(resByte, &preErr)
	if err == nil && preErr.GwErrResp.ErrMsg != "" {
		err = fmt.Errorf("有赞请求出现前置错误: %s", preErr.GwErrResp.ErrMsg)
		return
	}

	// 验签
	// 判断响应时间是否超过5分钟，超过则遗弃本次响应
	resHeader := httpRes.Header

	err = s.VerifySign(ctx, &resHeader, &resStr)
	if err != nil {
		err = fmt.Errorf("有赞验签失败")
		return
	}
	err = json.Unmarshal(resByte, &res)
	if err != nil {
		return
	}
	if !res.Success {
		err = fmt.Errorf("有赞提示出错: %s", res.Message)
		return
	}
	return
}
