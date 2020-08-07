package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os/exec"
	"strconv"
	"strings"

	"github.com/amiraliio/gompeg"
	"github.com/gin-gonic/gin"
	un "shelltear.loli/unmarshal"
)

func main() {
	router := gin.Default()

	router.GET("getid/:name", func(c *gin.Context) {
		name := c.Param("name")

		id := FindAnimeID(name)

		c.String(http.StatusOK, "%s", id)
	})

	router.GET("getlinks/:id/:onlyinfo", func(c *gin.Context) {
		id := c.Param("id")
		onlyinfo := c.Param("onlyinfo")

		links := FindLinksOfID(id, onlyinfo)

		c.String(http.StatusOK, "%s", links)
	})

	router.Run(":8080")
}

func FindAnimeID(name string) string {

	req, err := http.NewRequest("GET", "https://animeflix.io/api/search?q="+url.QueryEscape(name), nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Authority", "animeflix.io")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Accept-Language", "fi-FI,fi;q=0.9,en-US;q=0.8,en;q=0.7")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	out, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	output2, err := un.UnmarshalAnimedown3(out)
	link1 := "https://animeflix.io/api/episodes?anime_id=" + strconv.FormatInt(output2.Data[0].ID, 10) + "&limit=30&sort=DESC"

	req, err = http.NewRequest("GET", link1, nil)

	if err != nil {
		fmt.Println(err)
	}

	resp, err = http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	out, err = ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	output, err := un.UnmarshalAnimedown(out)

	var send string

	for i := range output.Data {
		send += "Anime title: " + output.Data[i].Title + "\nEpisode: " + output.Data[i].EpisodeNum + "\nSub: " +
			strconv.Itoa(output.Data[i].Sub) + " Dub: " + strconv.Itoa(output.Data[i].Dub) +
			"\nID: " + strconv.Itoa(output.Data[i].ID) + "\n\n"
	}

	return send
}

func FindLinksOfID(id string, onlyinfo string) string {

	req, err := http.NewRequest("GET", "https://animeflix.io/api/videos?episode_id="+id, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	req.Header.Set("Authority", "animeflix.io")
	req.Header.Set("Accept", "application/json, text/plain, */*")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/84.0.4147.105 Safari/537.36")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Accept-Language", "fi-FI,fi;q=0.9,en-US;q=0.8,en;q=0.7")

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()

	out, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println(err)
		return
	}

	output, err := un.UnmarshalAnimedown2(out)

	if err != nil {
		fmt.Println(err)
		return
	}

	var links string

	if onlyinfo == "true" {
		for i := range output {
			links += "ID: " + output[i].ID + "\nFile: " + output[i].File +
				"\nType: " + output[i].Type + "\nLang: " + output[i].Lang +
				"\nThumb: " + output[i].Thumbnail + "\n\n"
		}
	} else {
		for i := range output {
			if output[i].Type == "hls" {
				links = output[i].File
				break
			}
		}
		links = DownloadAnime(links, id)
	}
	return links
}

func DownloadAnime(link string, id string) string {

	fix := strings.Split(link, "://")
	fix[0] = "http"
	command := "youtube-dl -f 1092-0 -g " + strings.Join(fix, "://")
	cmdString := strings.TrimSuffix(command, "\n")
	cmdString2 := strings.Fields(cmdString)

	link2, err := exec.Command(cmdString2[0], cmdString2[1:]...).Output()

	if err != nil {
		fmt.Println(err)
		return
	}

	reallink := string(link2[:])
	reallink = strings.TrimSuffix(reallink, "\n")

	stream := new(gompeg.Media)
	stream.SetInputPath(reallink)
	stream.SetOutputPath(id)
	stream.SetVideoCodec("libx264")
	stream.SetPreset("veryfast")
	stream.SetAudioCodec("aac")
	stream.SetQuality(23)
	stream.SetOutputFormat("flv")

	if err := stream.Build(); err != nil {
		fmt.Println(err)
	}

	return "done"
}
