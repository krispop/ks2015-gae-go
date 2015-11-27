package guests

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/sign", sign)
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
他の人からのメッセージ:<p style="box-shadow: 0 2px 4px 0 rgba(0,0,0,0.16),0 2px 10px 0 rgba(0,0,0,0.12)" align=center>
{{.Content}}
</p>
{{end}}
</td>
</body>
`))

type Message struct {
	Time    time.Time
	Content string
}

func root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	msg := Message{
		Time:    time.Now(),
		Content: "test",
	}
	if err := rootTemplate.Execute(w, msg); err != nil {
		fmt.Fprint(w, "Error :(")
		log.Printf("templating error: %v", err)
	}
}

func sign(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, "ありがとうございます！<a href='/'>戻ります</a>")
}
