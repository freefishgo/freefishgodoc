package controllers

import (
	"bytes"
	"freefishgodoc/tools"
	"image/gif"
	"image/jpeg"
	"image/png"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/freefishgo/freefishgo/middlewares/mvc"
	"github.com/nfnt/resize"
)

type mediaController struct {
	mvc.Controller
}

func init() {
	mvc.AddHandlers(&mediaController{})
}

func (media *mediaController) UploadFilePost() {
	resp := &struct {
		Link string `json:"link"`
	}{}
	f, fh, err := media.Request.FormFile("file")
	defer f.Close()
	if err != nil {
		resp.Link = ""
		return
	}
	ty := filepath.Ext(fh.Filename)
	path := "/uploadfile/" + strings.ToLower(strings.Trim(ty, ".")) + "/" + tools.GetGuid() + ty
	resp.Link = path
	media.Response.WriteJson(resp)
	path = filepath.Join("static", path)
	if media.Query["type"] == "img" {
		// decode jpeg into image.Image
		b, err := ioutil.ReadAll(f)
		buf := new(bytes.Buffer)
		buf.Write(b)
		img, err := jpeg.Decode(buf)
		if err != nil {
			buf.Reset()
			buf.Write(b)
			img, err = png.Decode(buf)
			if err != nil {
				buf.Reset()
				buf.Write(b)
				img, err = gif.Decode(buf)
			}
		}
		// resize to width 1000 using Lanczos resampling
		// and preserve aspect ratio
		m := resize.Resize(1000, 0, img, resize.Lanczos3)
		os.MkdirAll(filepath.Dir(path), 0644)
		out, err := os.Create(path)
		if err != nil {
			log.Println(err)
		}
		defer out.Close()

		// write new image to file
		jpeg.Encode(out, m, nil)
		return
	}
	tools.WriteFileFromReader(path, f)
}

// 控制器执行前调用
func (media *mediaController) Prepare() {
	userinfo, _ := media.Response.GetSession("userinfo")
	userip, _ := media.Response.GetSession("userip")
	if userinfo != nil && userip != nil {
		if userip != media.Request.Host {
			media.Response.Write([]byte("违规操作"))
			media.SkipController()
		}
	} else {
		media.Response.Write([]byte("登录过期了"))
		media.SkipController()
	}
}
