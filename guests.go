package guests

import (
	"fmt"
	"net/http"
)

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/send", root)
}

func root(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, `
<head>
<link rel="stylesheet" href="pure/pure-min.css">
<title>Kawaii-Science Hackathon Guestbook</title>
</head>
<body>
<h1>理系のHackathonにメッセージを残そう！</h1>
<table width=100%>
<td width=50%>
<form action="send">
<textarea name=message placeholder="世界に伝えたいこと" rows=10 cols=80></textarea>
</form>
</td>
<td width=50% valign=top>他の人からのメッセージ
...
</td>
</body>
`)
}
