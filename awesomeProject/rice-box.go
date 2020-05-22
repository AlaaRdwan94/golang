package main

import (
	"time"

	"github.com/GeertJohan/go.rice/embedded"
)

func init() {

	// define files
	file2 := &embedded.EmbeddedFile{
		Filename:    "Adminwelcome.html",
		FileModTime: time.Unix(1590091581, 0),

		Content: string("\n<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>Title</title>\n    <meta charset=\"utf-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n    <link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/bootstrap/3.4.1/css/bootstrap.min.css\">\n    <script src=\"https://ajax.googleapis.com/ajax/libs/jquery/3.5.1/jquery.min.js\"></script>\n    <script src=\"https://maxcdn.bootstrapcdn.com/bootstrap/3.4.1/js/bootstrap.min.js\"></script>\n    <script src=\"http://ajax.googleapis.com/ajax/libs/jquery/1.7.1/jquery.min.js\" type=\"text/javascript\">\n    </script>\n    <script src=\"/js/welcomepage.js\" type=\"text/javascript\"></script>\n</head>\n<body>\n<h1>welcome  {{.UserName}} </h1>\n<input type=\"hidden\" id=\"username\" value=\"{{.UserName}}\">\n<div class=\"container p-3 my-3 bg-dark text-white\">\n    <div class=\"row\">\n        <div class=\"col-sm-6\">\n            <h3>send message :: </h3>\n            <input class=\"form-control form-control-lg\" type=\"text\" placeholder=\".form-control-lg\" name=\"message\" id=\"message\">\n        </div>\n\n    </div>\n\n    <div class=\"row\">\n        <div class=\"col-sm-4\">\n            <h3></h3>\n            <button type=\"button\" class=\"btn btn-secondary\" onclick=\"send()\">Send</button>\n        </div></div>\n    <div class=\"row\" id=\"chathistory\">\n\n    </div>\n    <div class=\"row\">\n        <div class=\"col-sm-4\">\n\n        </div>\n    </div>\n</div>\n</body>\n</html>\n"),
	}
	file4 := &embedded.EmbeddedFile{
		Filename:    "home.go.html",
		FileModTime: time.Unix(1590007217, 0),

		Content: string("<head>\n{{/*    <link rel=\"stylesheet\" href=\"./login.css\">*/}}\n</head>\n<form action=\"/Login\" method=\"POST\" novalidate>\n\n    <div class=\"container\">\n        <label for=\"uname\" ><b>Username</b></label>\n        <input type=\"text\" placeholder=\"Enter Username\" name=\"userame\" required>\n\n        <label for=\"psw\" name=\"password\"><b>Password</b></label>\n        <input type=\"password\" placeholder=\"Enter Password\" name=\"password\" required>\n\n        <button type=\"submit\">Login</button>\n\n    </div>\n\n    <div class=\"container\" style=\"background-color:#f1f1f1\">\n        <button type=\"button\" class=\"cancelbtn\">Cancel</button>\n        <span class=\"psw\">create  <a href=\"register\">account?</a></span>\n    </div>\n</form>\n<style>\n    /* Bordered form */\n    form {\n        border: 3px solid #f1f1f1;\n    }\n\n    /* Full-width inputs */\n    input[type=\"text\"],\n    input[type=\"password\"] {\n        width: 100%;\n        padding: 12px 20px;\n        margin: 8px 0;\n        display: inline-block;\n        border: 1px solid #ccc;\n        box-sizing: border-box;\n    }\n\n    /* Set a style for all buttons */\n    button {\n        background-color: #4caf50;\n        color: white;\n        padding: 14px 20px;\n        margin: 8px 0;\n        border: none;\n        cursor: pointer;\n        width: 100%;\n    }\n\n    /* Add a hover effect for buttons */\n    button:hover {\n        opacity: 0.8;\n    }\n\n    /* Extra style for the cancel button (red) */\n    .cancelbtn {\n        width: auto;\n        padding: 10px 18px;\n        background-color: #f44336;\n    }\n\n    /* Center the avatar image inside this container */\n    .imgcontainer {\n        text-align: center;\n        margin: 24px 0 12px 0;\n    }\n\n    /* Avatar image */\n    img.avatar {\n        width: 40%;\n        border-radius: 50%;\n    }\n\n    /* Add padding to containers */\n    .container {\n        padding: 16px;\n    }\n\n    /* The \"Forgot password\" text */\n    span.psw {\n        float: right;\n        padding-top: 16px;\n    }\n\n    /* Change styles for span and cancel button on extra small screens */\n    @media screen and (max-width: 300px) {\n        span.psw {\n            display: block;\n            float: none;\n        }\n        .cancelbtn {\n            width: 100%;\n        }\n    }\n\n</style>"),
	}
	file5 := &embedded.EmbeddedFile{
		Filename:    "homeAdmin.html",
		FileModTime: time.Unix(1590091581, 0),

		Content: string("<head>\n    {{/*    <link rel=\"stylesheet\" href=\"./login.css\">*/}}\n</head>\n<form action=\"/Login\" method=\"POST\" novalidate>\n\n    <div class=\"container\">\n        <label for=\"uname\" ><b>Username</b></label>\n        <input type=\"text\" placeholder=\"Enter Username\" name=\"userame\" required>\n\n        <label for=\"psw\" name=\"password\"><b>Password</b></label>\n        <input type=\"password\" placeholder=\"Enter Password\" name=\"password\" required>\n\n        <button type=\"submit\">Login</button>\n\n    </div>\n\n    <div class=\"container\" style=\"background-color:#f1f1f1\">\n        <button type=\"button\" class=\"cancelbtn\">Cancel</button>\n        <span class=\"psw\">create  <a href=\"register\">Admin account?</a></span>\n    </div>\n</form>\n<style>\n    /* Bordered form */\n    form {\n        border: 3px solid #f1f1f1;\n    }\n\n    /* Full-width inputs */\n    input[type=\"text\"],\n    input[type=\"password\"] {\n        width: 100%;\n        padding: 12px 20px;\n        margin: 8px 0;\n        display: inline-block;\n        border: 1px solid #ccc;\n        box-sizing: border-box;\n    }\n\n    /* Set a style for all buttons */\n    button {\n        background-color: #4caf50;\n        color: white;\n        padding: 14px 20px;\n        margin: 8px 0;\n        border: none;\n        cursor: pointer;\n        width: 100%;\n    }\n\n    /* Add a hover effect for buttons */\n    button:hover {\n        opacity: 0.8;\n    }\n\n    /* Extra style for the cancel button (red) */\n    .cancelbtn {\n        width: auto;\n        padding: 10px 18px;\n        background-color: #f44336;\n    }\n\n    /* Center the avatar image inside this container */\n    .imgcontainer {\n        text-align: center;\n        margin: 24px 0 12px 0;\n    }\n\n    /* Avatar image */\n    img.avatar {\n        width: 40%;\n        border-radius: 50%;\n    }\n\n    /* Add padding to containers */\n    .container {\n        padding: 16px;\n    }\n\n    /* The \"Forgot password\" text */\n    span.psw {\n        float: right;\n        padding-top: 16px;\n    }\n\n    /* Change styles for span and cancel button on extra small screens */\n    @media screen and (max-width: 300px) {\n        span.psw {\n            display: block;\n            float: none;\n        }\n        .cancelbtn {\n            width: 100%;\n        }\n    }\n\n</style>"),
	}
	file7 := &embedded.EmbeddedFile{
		Filename:    "js/welcomepage.js",
		FileModTime: time.Unix(1590090195, 0),

		Content: string("function f() {\n    console.log(\"ggggggggggggggggggg\");\n\n}\n$(document).ready(f());\n\n\n\nfunction send() {\n    var msg = document.getElementById(\"message\").value\n    var sender = document.getElementById(\"username\").value\n    console.log(\"username\"+ sender+ \"and message\"+ msg)\n    let message = {\n        Message :msg,\n        From : sender\n    }\n    var ws = new WebSocket(\"ws://\" + window.location.host + \"/Socket\");\n    ws.onopen = function() {\n        ws.send(JSON.stringify(message))\n    }\n    $(\"#chathistory\").append(`<div class=\"col-sm-12 \" >\n            <div class=\"card\" style=\"width:400px\">\n                <div class=\"card-body\">\n                    <h4 class=\"card-title\">`+sender+`</h4>\n                    <p class=\"card-text\">`+ msg+`</p>\n            </div>\n           </div>\n`);\n    ws.onmessage = function(event) {\n        var m = JSON.parse(event.data);\n        console.log(m);\n        console.debug(\"Received message\", m.Message);\n        $(\"#chathistory\").append(`<div class=\"col-sm-12 \"  style=\"color: darkblue\">\n            <div class=\"card\" style=\"width:400px\">\n                <div class=\"card-body\">\n                    <h4 class=\"card-title\">`+m.From+`</h4>\n                    <p class=\"card-text\">`+ m.Message+`</p>\n            </div>\n           </div>\n`);\n    }\n    ws.onerror = function(event) {\n        console.debug(event)\n    }\n}\n"),
	}
	file8 := &embedded.EmbeddedFile{
		Filename:    "login.css",
		FileModTime: time.Unix(1590001722, 0),

		Content: string("/* Bordered form */\nform {\n    border: 3px solid #f1f1f1;\n}\n\n/* Full-width inputs */\ninput[type=\"text\"],\ninput[type=\"password\"] {\n    width: 100%;\n    padding: 12px 20px;\n    margin: 8px 0;\n    display: inline-block;\n    border: 1px solid #ccc;\n    box-sizing: border-box;\n}\n\n/* Set a style for all buttons */\nbutton {\n    background-color: #4caf50;\n    color: white;\n    padding: 14px 20px;\n    margin: 8px 0;\n    border: none;\n    cursor: pointer;\n    width: 100%;\n}\n\n/* Add a hover effect for buttons */\nbutton:hover {\n    opacity: 0.8;\n}\n\n/* Extra style for the cancel button (red) */\n.cancelbtn {\n    width: auto;\n    padding: 10px 18px;\n    background-color: #f44336;\n}\n\n/* Center the avatar image inside this container */\n.imgcontainer {\n    text-align: center;\n    margin: 24px 0 12px 0;\n}\n\n/* Avatar image */\nimg.avatar {\n    width: 40%;\n    border-radius: 50%;\n}\n\n/* Add padding to containers */\n.container {\n    padding: 16px;\n}\n\n/* The \"Forgot password\" text */\nspan.psw {\n    float: right;\n    padding-top: 16px;\n}\n\n/* Change styles for span and cancel button on extra small screens */\n@media screen and (max-width: 300px) {\n    span.psw {\n        display: block;\n        float: none;\n    }\n    .cancelbtn {\n        width: 100%;\n    }\n}\n"),
	}
	file9 := &embedded.EmbeddedFile{
		Filename:    "register.html",
		FileModTime: time.Unix(1590007982, 0),

		Content: string("<head>\n    {{/*    <link rel=\"stylesheet\" href=\"./login.css\">*/}}\n</head>\n<form action=\"/Register\" method=\"POST\" novalidate>\n\n    <div class=\"container\">\n        <label for=\"uname\" ><b>Username</b></label>\n        <input type=\"text\" placeholder=\"Enter Username\" name=\"userame\" required>\n\n        <label for=\"psw\" name=\"password\"><b>Password</b></label>\n        <input type=\"password\" placeholder=\"Enter Password\" name=\"password\" required>\n\n        <button type=\"submit\">Register</button>\n\n    </div>\n\n    <div class=\"container\" style=\"background-color:#f1f1f1\">\n        <button type=\"button\" class=\"cancelbtn\">Cancel</button>\n        <span class=\"psw\">have  <a href=\"home\">account?</a></span>\n    </div>\n</form>\n<style>\n    /* Bordered form */\n    form {\n        border: 3px solid #f1f1f1;\n    }\n\n    /* Full-width inputs */\n    input[type=\"text\"],\n    input[type=\"password\"] {\n        width: 100%;\n        padding: 12px 20px;\n        margin: 8px 0;\n        display: inline-block;\n        border: 1px solid #ccc;\n        box-sizing: border-box;\n    }\n\n    /* Set a style for all buttons */\n    button {\n        background-color: #4caf50;\n        color: white;\n        padding: 14px 20px;\n        margin: 8px 0;\n        border: none;\n        cursor: pointer;\n        width: 100%;\n    }\n\n    /* Add a hover effect for buttons */\n    button:hover {\n        opacity: 0.8;\n    }\n\n    /* Extra style for the cancel button (red) */\n    .cancelbtn {\n        width: auto;\n        padding: 10px 18px;\n        background-color: #f44336;\n    }\n\n    /* Center the avatar image inside this container */\n    .imgcontainer {\n        text-align: center;\n        margin: 24px 0 12px 0;\n    }\n\n    /* Avatar image */\n    img.avatar {\n        width: 40%;\n        border-radius: 50%;\n    }\n\n    /* Add padding to containers */\n    .container {\n        padding: 16px;\n    }\n\n    /* The \"Forgot password\" text */\n    span.psw {\n        float: right;\n        padding-top: 16px;\n    }\n\n    /* Change styles for span and cancel button on extra small screens */\n    @media screen and (max-width: 300px) {\n        span.psw {\n            display: block;\n            float: none;\n        }\n        .cancelbtn {\n            width: 100%;\n        }\n    }\n\n</style>"),
	}
	filea := &embedded.EmbeddedFile{
		Filename:    "registerAdmin.html",
		FileModTime: time.Unix(1590091581, 0),

		Content: string("<head>\n    {{/*    <link rel=\"stylesheet\" href=\"./login.css\">*/}}\n</head>\n<form action=\"/Register\" method=\"POST\" novalidate>\n\n    <div class=\"container\">\n        <label for=\"uname\" ><b>Username</b></label>\n        <input type=\"text\" placeholder=\"Enter Username\" name=\"userame\" required>\n\n        <label for=\"psw\" name=\"password\"><b>Password</b></label>\n        <input type=\"password\" placeholder=\"Enter Password\" name=\"password\" required>\n\n        <button type=\"submit\">Register</button>\n\n    </div>\n\n    <div class=\"container\" style=\"background-color:#f1f1f1\">\n        <button type=\"button\" class=\"cancelbtn\">Cancel</button>\n        <span class=\"psw\">have  <a href=\"home\">Admin account?</a></span>\n    </div>\n</form>\n<style>\n    /* Bordered form */\n    form {\n        border: 3px solid #f1f1f1;\n    }\n\n    /* Full-width inputs */\n    input[type=\"text\"],\n    input[type=\"password\"] {\n        width: 100%;\n        padding: 12px 20px;\n        margin: 8px 0;\n        display: inline-block;\n        border: 1px solid #ccc;\n        box-sizing: border-box;\n    }\n\n    /* Set a style for all buttons */\n    button {\n        background-color: #4caf50;\n        color: white;\n        padding: 14px 20px;\n        margin: 8px 0;\n        border: none;\n        cursor: pointer;\n        width: 100%;\n    }\n\n    /* Add a hover effect for buttons */\n    button:hover {\n        opacity: 0.8;\n    }\n\n    /* Extra style for the cancel button (red) */\n    .cancelbtn {\n        width: auto;\n        padding: 10px 18px;\n        background-color: #f44336;\n    }\n\n    /* Center the avatar image inside this container */\n    .imgcontainer {\n        text-align: center;\n        margin: 24px 0 12px 0;\n    }\n\n    /* Avatar image */\n    img.avatar {\n        width: 40%;\n        border-radius: 50%;\n    }\n\n    /* Add padding to containers */\n    .container {\n        padding: 16px;\n    }\n\n    /* The \"Forgot password\" text */\n    span.psw {\n        float: right;\n        padding-top: 16px;\n    }\n\n    /* Change styles for span and cancel button on extra small screens */\n    @media screen and (max-width: 300px) {\n        span.psw {\n            display: block;\n            float: none;\n        }\n        .cancelbtn {\n            width: 100%;\n        }\n    }\n\n</style>"),
	}
	fileb := &embedded.EmbeddedFile{
		Filename:    "welcome.html",
		FileModTime: time.Unix(1590090676, 0),

		Content: string("\n<!DOCTYPE html>\n<html lang=\"en\">\n<head>\n    <meta charset=\"UTF-8\">\n    <title>Title</title>\n    <meta charset=\"utf-8\">\n    <meta name=\"viewport\" content=\"width=device-width, initial-scale=1\">\n    <link rel=\"stylesheet\" href=\"https://maxcdn.bootstrapcdn.com/bootstrap/3.4.1/css/bootstrap.min.css\">\n    <script src=\"https://ajax.googleapis.com/ajax/libs/jquery/3.5.1/jquery.min.js\"></script>\n    <script src=\"https://maxcdn.bootstrapcdn.com/bootstrap/3.4.1/js/bootstrap.min.js\"></script>\n   <script src=\"http://ajax.googleapis.com/ajax/libs/jquery/1.7.1/jquery.min.js\" type=\"text/javascript\">\n</script>\n<script src=\"/js/welcomepage.js\" type=\"text/javascript\"></script>\n</head>\n<body>\n<h1>welcome  {{.UserName}} </h1>\n<input type=\"hidden\" id=\"username\" value=\"{{.UserName}}\">\n<div class=\"container p-3 my-3 bg-dark text-white\">\n    <div class=\"row\">\n        <div class=\"col-sm-6\">\n            <h3>send message :: </h3>\n            <input class=\"form-control form-control-lg\" type=\"text\" placeholder=\".form-control-lg\" name=\"message\" id=\"message\">\n        </div>\n\n    </div>\n\n        <div class=\"row\">\n            <div class=\"col-sm-4\">\n                <h3></h3>\n                <button type=\"button\" class=\"btn btn-secondary\" onclick=\"send()\">Send</button>\n            </div></div>\n        <div class=\"row\" id=\"chathistory\">\n\n        </div>\n        <div class=\"row\">\n            <div class=\"col-sm-4\">\n\n            </div>\n        </div>\n    </div>\n</body>\n</html>\n"),
	}

	// define dirs
	dir1 := &embedded.EmbeddedDir{
		Filename:   "",
		DirModTime: time.Unix(1590091581, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file2, // "Adminwelcome.html"
			file4, // "home.go.html"
			file5, // "homeAdmin.html"
			file8, // "login.css"
			file9, // "register.html"
			filea, // "registerAdmin.html"
			fileb, // "welcome.html"

		},
	}
	dir3 := &embedded.EmbeddedDir{
		Filename:   "css",
		DirModTime: time.Unix(1590028090, 0),
		ChildFiles: []*embedded.EmbeddedFile{},
	}
	dir6 := &embedded.EmbeddedDir{
		Filename:   "js",
		DirModTime: time.Unix(1590091337, 0),
		ChildFiles: []*embedded.EmbeddedFile{
			file7, // "js/welcomepage.js"

		},
	}

	// link ChildDirs
	dir1.ChildDirs = []*embedded.EmbeddedDir{
		dir3, // "css"
		dir6, // "js"

	}
	dir3.ChildDirs = []*embedded.EmbeddedDir{}
	dir6.ChildDirs = []*embedded.EmbeddedDir{}

	// register embeddedBox
	embedded.RegisterEmbeddedBox(`templates`, &embedded.EmbeddedBox{
		Name: `templates`,
		Time: time.Unix(1590091581, 0),
		Dirs: map[string]*embedded.EmbeddedDir{
			"":    dir1,
			"css": dir3,
			"js":  dir6,
		},
		Files: map[string]*embedded.EmbeddedFile{
			"Adminwelcome.html":  file2,
			"home.go.html":       file4,
			"homeAdmin.html":     file5,
			"js/welcomepage.js":  file7,
			"login.css":          file8,
			"register.html":      file9,
			"registerAdmin.html": filea,
			"welcome.html":       fileb,
		},
	})
}
