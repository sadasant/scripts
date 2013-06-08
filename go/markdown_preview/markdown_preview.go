package main

import (
	"bytes"
	"github.com/knieriem/markdown"
	"log"
	"net"
	"net/http"
	"os"
	"os/exec"
)

func main() {
	// Args
	path := os.Args[1]
	port := ":8080"
	if len(os.Args) > 2 {
		port = ":" + os.Args[2]
	}

	// Open Markdown File
	mkd_file, err := os.Open(path)
	if err != nil {
		log.Fatal("Open File Error: ", err.Error())
	}
	defer mkd_file.Close()

	// Markdown Parser
	p := markdown.NewParser(nil)
	var markdown_buf bytes.Buffer
	p.Markdown(mkd_file, markdown.ToHTML(&markdown_buf))

	// Launch the HTTP server
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`
		<head><style>
a {
	color              : #666;
	outline            : 0;
	text-decoration    : none;
	transition         : color .5s ease 0s;
	-moz-transition    : color .5s ease 0s;
	-webkit-transition : color .5s ease 0s;
}
a:hover {
	font-style : italic;
}
strong {
	font-weight : bold;
}
hr {
	border : 1px solid black;
	margin : 30px 0;
}
::selection {
	background : #c4c4c4;
	color      : black;
}
::-moz-selection {
	background : #777;
	color      : black;
}
body {
	font-family : sans-serif;
	background  : black;
	color       : #ddd;
	width       : 77%;
	max-width   : 666px;
	padding     : 0px 20px 30px 0;
	font-size   : 17px;
}
h1, h2, h3 {
	font-weight : bold;
	margin-top  : 30px;
}
h1 {
	font-size : 1.8em;
}
h2 {
	font-size : 1.4em;
}
h3 {
	font-size  : 1em;
	margin-top : 20px;
}
p {
	margin : 16px 0;
}
a:hover {
	color : white;
}
em {
	font-style : italic;
}
code {
	font-size     : 13px;
	color         : #888;
	background    : black;
	padding       : 3px;
	border        : 1px solid #171717;
	border-top    : 1px solid #111111;
	border-bottom : 1px solid #222222;
	box-shadow    : 0 0 3px #070707;
}
pre, p code {
	display               : block;
	font-size             : 13px;
	color                 : #888;
	background            : black;
	padding               : 10px;
	margin                : 10px 0;
	overflow-x            : hidden;
	white-space           : pre;
	border                : 1px solid #171717;
	border-top            : 1px solid #111111;
	border-bottom         : 1px solid #222222;
	border-radius         : 3px;
	-moz-border-radius    : 3px;
	-webkit-border-radius : 3px;
	box-shadow            : inset 10px 0 30px #0A0A0A;
	                      , inset 30px 0 100px #070707;
}
pre code {
	border     : 0;
	box-shadow : 0;
}
pre:hover, p code:hover {
	overflow-x : scroll;
}
pre code {
	padding    : 0;
	border     : 0;
	background : none;
}
blockquote {
	padding    : 10px 10px 10px 15px;
	background : #070707;
	color      : #777;
	font-size  : 16px;
}
blockquote a {
	text-decoration : underline;
}
ul {
	margin : 8px 0 4px 0;
}
li {
	margin           : 0 0 8px 20px;
	padding-left     : 4px;
	list-style-image : url(http://sadasant.com/images/whitebullets.png);
}
img {
	max-width : 77%;
	border    : 1px solid #111;
}
		</style></head>
		<body>`))
		w.Write(markdown_buf.Bytes())
		w.Write([]byte("</body>"))
	})

	// Serving the doc
	server := &http.Server{
		Addr: port,
	}
	server_listener, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err.Error())
	}
	go server.Serve(server_listener)

	// Surf the url
	url := "http://localhost" + port
	println("Listening on: " + url + "/")
	cmd := exec.Command("surf", url)
	err = cmd.Run()
	if err != nil {
		log.Fatal(err.Error())
	}

	// Bye bye!
	server_listener.Close()
}
