package guests

import (
	"fmt"
	"html/template"
	"math/rand"
	"net/http"
	"time"

	"appengine"
	"appengine/datastore"
)

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/sign", sign)
	time.Local, _ = time.LoadLocation("Asia/Tokyo")
}

var rootTemplate = template.Must(template.New("root").Parse(`
<head>
<link rel="stylesheet" href="pure/pure-min.css">
<title>Kawaii-Science Hackathon Guestbook</title>
</head>
<body>
<h1>理系女子のHackathonにメッセージを残そう！</h1>
<table width=100%>
<td width=50%>
<form action="sign">
<textarea name=message placeholder="世界に伝えたいこと" rows=10 cols=80></textarea>
<br><input type=submit value="送信">
</form>
</td>
<td width=50% valign=top>
{{if .Content}}
他の人からのメッセージ:<p style='border-style: solid; border-color: blue; border-width: 2px;' align=center>
{{.Content}}
</p>
@{{.Time}}
{{end}}
</td>
</body>
`))

type Message struct {
	// Timeは送信されたタイムスタンプ
	Time time.Time
	// Contentはメッセージの中身
	Content string
}

const kind = "Message"

func root(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	var msgs []*Message
	if _, err := datastore.NewQuery(kind).GetAll(c, &msgs); err != nil {
		fmt.Fprint(w, "DB Error :(")
		c.Errorf("db error: %v", err)
		return

	}
	msg := &Message{}
	c.Infof("messages: %d", len(msgs))
	if len(msgs) > 0 {
		msg = msgs[rand.Int31n(int32(len(msgs)))]
	}
	if err := rootTemplate.Execute(w, msg); err != nil {
		fmt.Fprint(w, "Error :(")
		c.Errorf("templating error: %v", err)
		return
	}
}

func sign(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	msg := r.FormValue(kind)
	if msg == "" {
		fmt.Fprint(w, "もう少し面白いのを入力してみよう！")
		return
	}

	if _, err := datastore.Put(c, datastore.NewIncompleteKey(c, kind, nil), &Message{
		Time:    time.Now(),
		Content: msg,
	}); err != nil {
		fmt.Fprint(w, "失敗しました。。。 :(")
		c.Errorf("datastore error: %v", err)
		return
	}
	fmt.Fprint(w, "ありがとうございます！<a href='/'>戻ります</a>")
}
