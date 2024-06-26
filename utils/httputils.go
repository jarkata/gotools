package utils

import (
	"encoding/json"
	"io"
	"log/slog"
	"net/http"
)

type responseDTO struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data"`
}

func newResponseDTO(code int, msg string, data any) responseDTO {
	return responseDTO{Code: code, Msg: msg, Data: data}
}

func newJsonResponse(code int, msg string, data any) []byte {
	b, err := json.Marshal(newResponseDTO(code, msg, data))
	if err != nil {
		slog.Error("转换为JSON格式数据发生异常", err)
	}
	return b
}

/*
读取请求报文，并转换为结构体
*/
func ReadBody(request io.Reader, target any) error {
	body, err := io.ReadAll(request)
	if err != nil {
		slog.Error("读取报文发生异常:", err)
		return err
	}
	err = json.Unmarshal(body, target)
	if err != nil {
		slog.Error("解析报文发生异常:", err)
		return err
	}
	return nil
}

/*
读取请求报文，并转换为结构体
*/
func ReadBodyWithLog(request io.Reader, target any) error {
	body, err := io.ReadAll(request)
	if err != nil {
		slog.Error("读取报文发生异常:", err)
		return err
	}
	slog.Info("请求报文", "Body", string(body))
	err = json.Unmarshal(body, target)
	if err != nil {
		slog.Error("解析报文发生异常:", err)
		return err
	}
	return nil
}

/*
*
发送成功消息
code:200,
msg:OK,
data:null
*/
func ResponseOK(w http.ResponseWriter) {
	sendJsonResponse(w, http.StatusOK, newJsonResponse(http.StatusOK, "OK", nil))
}

/*
返回指定数据,

	code:200
	msg:success
	data:any

*
*/
func WriteJsonResponse(w http.ResponseWriter, msg any) {
	b := newJsonResponse(http.StatusOK, "success", msg)
	sendJsonResponse(w, http.StatusOK, b)
}

/*
*
发送指定错误码及错误信息
消息结构如下:
code:200
msg:errMsg
data:nil
*/
func WriteErrorResponse(w http.ResponseWriter, statusCode int, errMsg string) {
	b := newJsonResponse(statusCode, errMsg, nil)
	sendJsonResponse(w, http.StatusOK, b)
}

func sendJsonResponse(w http.ResponseWriter, statusCode int, msg []byte) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")
	w.WriteHeader(statusCode)
	_, err := w.Write(msg)
	if err != nil {
		slog.Error("WriteResposneFail", "Error", err)
		return
	}
}
