package youzan_func

import (
	"context"
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/shatian114/verifysign"
)

// 对有赞请求结果进行验签
func (s *Client) VerifySign(ctx context.Context, resHeader *http.Header, resStr *string) (err error) {

	resTimestamp, err := strconv.ParseInt(resHeader.Get("Youzan-Timestamp"), 10, 64)
	if err != nil {
		err = fmt.Errorf("有赞头信息里的时间戳不正确: %d", resTimestamp)
		return
	}

	if time.Unix(int64(resTimestamp), 0).Before(time.Now().Add(-5 * time.Minute)) {
		err = fmt.Errorf("有赞响应时间戳超过5分钟，不可信")
		return
	}
	verifySignContent := fmt.Sprintf("%s\n%d\n%s\n", *resStr, resTimestamp, resHeader.Get("Youzan-Noncestr"))
	resSign := resHeader.Get("Youzan-Sign")
	err = verifysign.VerifySign(ctx, s.PubKey, resSign, verifySignContent)
	if err != nil {
		err = fmt.Errorf("有赞验签失败")
		return
	}
	return
}
