package helper

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
)

// EncodePNGToByte 图片编码二进制，PNG格式
func EncodePNGToByte(img image.Image) (ret []byte) {
	var buf bytes.Buffer
	if err := png.Encode(&buf, img); err != nil {
		panic(err.Error())
	}
	ret = buf.Bytes()
	buf.Reset()
	return
}

// EncodeJPEGToByte 图片编码二进制，IMAGE格式
func EncodeJPEGToByte(img image.Image, quality int) (ret []byte) {
	var buf bytes.Buffer
	if err := jpeg.Encode(&buf, img, &jpeg.Options{Quality: quality}); err != nil {
		panic(err.Error())
	}
	ret = buf.Bytes()
	buf.Reset()
	return
}

// DecodeByteToJpeg 字节编码图片
func DecodeByteToJpeg(b []byte) (img image.Image, err error) {
	var buf bytes.Buffer
	buf.Write(b)
	img, err = jpeg.Decode(&buf)
	buf.Reset()
	return
}

// DecodeByteToPng 字节编码图片
func DecodeByteToPng(b []byte) (img image.Image, err error) {
	var buf bytes.Buffer
	buf.Write(b)
	img, err = png.Decode(&buf)
	buf.Reset()
	return
}

// EncodePNGToBase64 base64编码
func EncodePNGToBase64(img image.Image) string {
	return fmt.Sprintf("data:%s;base64,%s", "image/png", base64.StdEncoding.EncodeToString(EncodePNGToByte(img)))
}

// EncodeJPEGToBase64 base64编码
func EncodeJPEGToBase64(img image.Image, quality int) string {
	return fmt.Sprintf("data:%s;base64,%s", "image/jpeg", base64.StdEncoding.EncodeToString(EncodeJPEGToByte(img, quality)))
}
