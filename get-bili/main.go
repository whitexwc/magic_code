package main

import (
	"magic_code/get-bili/downloader"
	myfmt "magic_code/get-bili/fmt"
)

func main() {

	request := downloader.InfoRequest{Bvids: []string{"BV1Be4y1j78j", "BV1Nv4y197m7"}}
	response, err := downloader.BatchDownloadVideoInfo(request)
	if err != nil {
		panic(err)
	}

	for _, info := range response.Infos {
		myfmt.Logger.Printf("title:%s \n description:%s \n ", info.Data.Title, info.Data.Desc)
	}
}
