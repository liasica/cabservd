// Copyright (C) liasica. 2023-present.
//
// Created at 2023-05-15
// Based on cabservd by liasica, magicrolan@qq.com.

package xlls

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"strconv"
	"strings"
	"time"

	"github.com/auroraride/adapter"
	"github.com/auroraride/adapter/log"
	"github.com/go-resty/resty/v2"
	jsoniter "github.com/json-iterator/go"
	"go.uber.org/zap"
)

// Request 请求体
type Request struct {
	Version   string `json:"version,omitempty"`
	Timestamp int64  `json:"timestamp,omitempty"`
	RequestId string `json:"requestId,omitempty"`
	AppId     string `json:"appId,omitempty"`
	Sign      string `json:"sign,omitempty"`
	Biz       string `json:"biz,omitempty"`
}

func (r *Request) String() string {
	b, _ := jsoniter.Marshal(r)
	return adapter.ConvertBytes2String(b)
}

// 生成请求体
func newRequest(biz any) (args *Request) {
	args = &Request{
		Version:   version,
		Timestamp: time.Now().UnixMilli(),
		RequestId: generateRequestID(),
		AppId:     appID,
	}
	switch v := biz.(type) {
	case string:
		args.Biz = v
	default:
		args.Biz, _ = jsoniter.MarshalToString(biz)
	}

	// 生成签名
	var s strings.Builder
	s.WriteString("appId=")
	s.WriteString(args.AppId)
	s.WriteString(",biz=")
	s.WriteString(args.Biz)
	s.WriteString(",timestamp=")
	s.WriteString(strconv.FormatInt(args.Timestamp, 10))

	// 计算HmacSHA1
	h := hmac.New(sha1.New, appSecret)
	h.Write([]byte(s.String()))
	sum := h.Sum(nil)
	args.Sign = base64.URLEncoding.EncodeToString(sum)
	return
}

func doRequest[T any](path string, biz any) (data T, err error) {
	var result ApiResponse[T]
	result, err = request[ApiResponse[T]](path, biz)
	data = result.Data
	return
}

func request[T any](path string, biz any) (result T, err error) {
	args := newRequest(biz)
	b, _ := jsoniter.Marshal(args)

	client := resty.New()
	var resp *resty.Response
	resp, err = client.R().
		EnableTrace().
		SetHeader("Content-Type", "application/json; charset=utf-8").
		SetBody(b).
		SetResult(&result).
		Post(baseURL + path)
	if err != nil {
		zap.L().Error("[api] 请求失败", zap.Error(err), zap.String("path", path), log.ResponseBody(resp.Body()), zap.ByteString("body", b))
		return
	}

	// TODO 日志记录增加elk索引
	zap.L().Info("[api] 请求成功", zap.String("path", path), log.ResponseBody(resp.Body()), zap.ByteString("body", b))
	return
}
