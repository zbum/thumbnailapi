package thumbnailapi

import (
	"bytes"
	"fmt"
	"gopkg.in/gographics/imagick.v3/imagick"
	"io"
	"net/http"
	"path/filepath"
	"strconv"
	"strings"
)

func Resize(res http.ResponseWriter, req *http.Request){

	mw := imagick.NewMagickWand()
	defer mw.Destroy()

	file, fileHeader, _ := req.FormFile("file")
	defer file.Close()

	fmt.Println(filepath.Ext(fileHeader.Filename))

	buf := bytes.NewBuffer(nil)
	if _, err := io.Copy(buf, file); err != nil {
		return
	}

	err := mw.ReadImageBlob(buf.Bytes())
	if err != nil {
		return
	}

	sizes := req.FormValue("sizes")
	split := strings.Split(sizes, ",")
	var doorayHeaders []string
	var imageBytesList [][]byte
	for _,size := range split {
		colsRows := strings.Split(size, "-")
		cols,_ := strconv.ParseUint(colsRows[0],10, 64)
		rows,_ := strconv.ParseUint(colsRows[1],10, 64)

		mw.ResizeImage(uint(cols), uint(rows), imagick.FILTER_BOX)
		mw.SetImageFormat("png")

		image := mw.GetImageBlob()
		imageBytesList = append(imageBytesList, image)
		doorayHeaders = append(doorayHeaders, size + ":" + strconv.Itoa(len(image)))
	}
	res.Header().Set("X-Content-Lengths", strings.Join(doorayHeaders, ","))

	for _,imageBytes := range imageBytesList {
		res.Write(imageBytes)
	}

}