package animation

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/gif"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

import (
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
)

type CreateFrameParam struct {
	URI   string // URL or local path. support gif|jpeg|png
	Delay int    // Delay time, one per frame, in 100ths of a second. Set to gif.GIF.Delay
}

type CreateAnimationParam struct {
	Frames []CreateFrameParam
}

func Create(param CreateAnimationParam) ([]byte, error) {
	log.Println("start create animation")

	frameParams := param.Frames
	if len(frameParams) == 0 {
		return nil, errors.New("require CreateAnimationParam.Frames")
	}

	g := &gif.GIF{}
	for _, p := range frameParams {
		err := appendFrame(g, p)
		if err != nil {
			return nil, fmt.Errorf("failed appendFrame. %w", err)
		}
	}

	var all bytes.Buffer
	err := gif.EncodeAll(&all, g)
	if err != nil {
		return nil, fmt.Errorf("failed gif.EncodeAll. %w", err)
	}
	return all.Bytes(), nil
}

func appendFrame(g *gif.GIF, p CreateFrameParam) error {
	origin, err := fetchImage(p.URI)
	if err != nil {
		return fmt.Errorf("failed fetchImage. %w", err)
	}

	decodedImage, _, err := image.Decode(bytes.NewReader(origin))
	if err != nil {
		return fmt.Errorf("failed image.Decode. %w", err)
	}

	var encodedGIF bytes.Buffer
	err = gif.Encode(&encodedGIF, decodedImage, nil)
	if err != nil {
		return fmt.Errorf("failed gif.Encode. %w", err)
	}

	decodedGIF, err := gif.Decode(bytes.NewReader(encodedGIF.Bytes()))
	if err != nil {
		return fmt.Errorf("failed gif.Decode. %w", err)
	}

	g.Image = append(g.Image, decodedGIF.(*image.Paletted))
	g.Delay = append(g.Delay, 100)
	return nil
}

func fetchImage(uri string) ([]byte, error) {
	log.Printf("fetchImage uri: %s\n", uri)
	if isURL(uri) {
		return fetchImageFromURL(uri)
	}
	return ioutil.ReadFile(uri)
}

func isURL(uri string) bool {
	_, err := url.ParseRequestURI(uri)
	if err != nil {
		return false
	}
	u, err := url.Parse(uri)
	if err != nil || u.Scheme == "" || u.Host == "" {
		return false
	}
	return true
}

func fetchImageFromURL(url string) ([]byte, error) {
	response, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed http.Get. url %s : %w", url, err)
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("statusCode is not 200. url: %s", url)
	}
	defer response.Body.Close()
	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed ioutil.ReadAll. url %s : %w", url, err)
	}
	return data, nil
}
