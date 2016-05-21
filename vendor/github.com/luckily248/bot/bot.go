package bot

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"unicode/utf8"

	"github.com/luckily248/bot/handler"
	"github.com/luckily248/bot/models"
)

var whitelist []string
var token string

func isWhiteList(id string) bool {
	for _, wid := range whitelist {
		if wid == id {
			return true
		}
	}
	return false
}
func init() {
	whitelist = []string{"135326188", "125657347", "72121797", "72125169", "125253862", "140087332", "92363924", "93138188"} //law in luckbot patt tom law luck shab morecowbell cata
	token = "376a2810f86a0133f8cb675ee1cd23ec"
}

//main
func Run() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})
	http.HandleFunc("/bot", WarDataController)
	http.ListenAndServe(":"+os.Getenv("PORT"), nil)
}
func WarDataController(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		fmt.Printf("otherMethod:%s\n", r.Method)
		return
	}
	var rec models.GMrecModel
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	if err := json.Unmarshal(body, &rec); err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Printf("rec:%v\n", rec)
	if rec.Text == "" {
		fmt.Printf("is empty\n")
		return
	}
	if strings.HasPrefix(rec.Text, "!") {
		go handle(rec)
	}
	if strings.Contains(rec.Text, "removed") && rec.System {
		go checkremove(rec)
	}
	return
}
func handle(rec models.GMrecModel) {
	reptext, err := handler.HandlecocText(rec)
	fmt.Printf("reptextlen:%d\n", utf8.RuneCountInString(reptext))
	rep := &models.GMrepModel{}
	rep.InitbyGID(rec.Group_id)
	if err != nil {
		rep.SetText(err.Error())
		fmt.Printf("err:%s\n", err.Error())
	} else {
		rep.SetText(reptext)
		fmt.Printf("ob:%v\n", rep)
	}
	buff, err := json.Marshal(rep)
	if err != nil {
		fmt.Printf("err:%s\n", err.Error())
		return
	}
	fmt.Println(string(buff))
	httpPost(buff)
}
func httpPost(rep []byte) {
	resp, err := http.Post("https://api.groupme.com/v3/bots/post",
		"application/x-www-form-urlencoded",
		bytes.NewReader(rep))
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
}
func checkremove(rec models.GMrecModel) {
	name := strings.Split(rec.Text, " removed ")[0]
	group, err := httpPostGetGroup(rec.Group_id)
	if err != nil {
		fmt.Println(err)
		return
	}
	var remover_id string
	for _, member := range group.Response.Members {
		if member.Nickname == name {
			remover_id = member.Id
		}
	}
	if remover_id == "" {
		fmt.Println("remover not found")
		return
	}
	if isWhiteList(remover_id) {
		fmt.Println("in whitelist")
		return
	}
	httpPostRemove(rec.Group_id, remover_id)
	return
}
func httpPostRemove(group_id string, membership_id string) {
	url := fmt.Sprintf("https://api.groupme.com/v3/groups/%s/members/%s/remove?token=%s", group_id, membership_id, token)
	resp, err := http.Post(url,
		"application/x-www-form-urlencoded",
		nil)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(body))
	return
}
func httpPostGetGroup(group_id string) (group models.GMrecGroupModel, err error) {
	url := fmt.Sprintf("https://api.groupme.com/v3/groups/%s?token=%s", group_id, token)
	fmt.Println(url)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return
	}
	err = json.Unmarshal(body, &group)
	if err != nil {
		return
	}
	fmt.Printf("group:%v\n", group)
	return
}
