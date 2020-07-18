package bilive

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func SendDanmaku(text string) error {
	url := "https://api.live.bilibili.com/msg/send"
	method := "POST"

	s := "color=16777215&fontsize=25&mode=1&msg=%s&rnd=1594867830&roomid=4153177&bubble=0&csrf_token=6d97ce7224f5f416dc1b05f95b4b3604&csrf=6d97ce7224f5f416dc1b05f95b4b3604"
	payload := strings.NewReader(fmt.Sprintf(s, text))

	client := &http.Client {
	}
	req, err := http.NewRequest(method, url, payload)

	if err != nil {
		fmt.Println(err)
	}
	req.Header.Add("authority", "api.live.bilibili.com")
	req.Header.Add("accept", "application/json, text/javascript, */*; q=0.01")
	req.Header.Add("user-agent", "Mozilla/5.0 (Macintosh; Intel Mac OS X 10_15_5) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/83.0.4103.116 Safari/537.36")
	req.Header.Add("content-type", "application/x-www-form-urlencoded; charset=UTF-8")
	req.Header.Add("origin", "https://live.bilibili.com")
	req.Header.Add("sec-fetch-site", "same-site")
	req.Header.Add("sec-fetch-mode", "cors")
	req.Header.Add("sec-fetch-dest", "empty")
	req.Header.Add("referer", "https://live.bilibili.com/4153177")
	req.Header.Add("accept-language", "zh-CN,zh;q=0.9,en;q=0.8,ja;q=0.7,zh-TW;q=0.6")
	req.Header.Add("cookie", "LIVE_BUVID=AUTO9815874837583600; _uuid=C2C60502-97B0-B722-2192-29FCE65367FA83031infoc; buvid3=6E5E19A7-F1F8-4A63-A4E1-3A4F6D03E87D155836infoc; sid=9z43k3gy; DedeUserID=2459271; DedeUserID__ckMd5=3c272ac53b78f4fc; SESSDATA=cdcc18ed%2C1605057673%2C38bbd*51; bili_jct=6d97ce7224f5f416dc1b05f95b4b3604; CURRENT_FNVAL=16; Hm_lvt_8a6e55dbd2870f0f5bc9194cddf32a02=1593479317,1593566663; rpdid=|(mmmYYRuRY0J'ulmRY|k|lR; bsource=search_google; _dfcaptcha=3f405c2b7f92f15cf356c8d4ccf8427e; Hm_lpvt_8a6e55dbd2870f0f5bc9194cddf32a02=1594867676; Hm_lvt_8a6d461cf92ec46bd14513876885e489=1594867746; Hm_lpvt_8a6d461cf92ec46bd14513876885e489=1594867746; PVID=14")

	res, err := client.Do(req)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	_, err = ioutil.ReadAll(res.Body)
	if err != nil {
		return err
	}
	return nil
}