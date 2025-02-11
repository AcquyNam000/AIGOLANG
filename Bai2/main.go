package main

import (
	"github.com/kataras/iris/v12"
)

func main() {
	app := iris.New()

	app.Handle("GET", "/", func(ctx iris.Context) {
		ctx.HTML(`
		<!DOCTYPE html>
		<html lang="vi">
		<head>
		    <meta charset="UTF-8">
		    <meta name="viewport" content="width=device-width, initial-scale=1.0">
		    <title>SSML Generator</title>
		    <script src="/static/script.js"></script>
		    <link rel="stylesheet" href="/static/style.css">
		</head>
		<body>
		    <h2>Nhập hội thoại</h2>

		    <div class="voice-selection">
		        <label>Voice A:</label>
		        <select id="voiceA" onchange="validateVoices()">
        <option value="en-US-AndrewMultilingualNeural">Andrew (EN)</option>
        <option value="en-US-ChristopherNeural">Christopher (EN)</option>
        <option value="en-US-EricNeural">Eric (EN)</option>
        <option value="vi-VN-HoaiMyNeural">Hoài My (VN)</option>
        <option value="vi-VN-NamMinhNeural">Nam Minh (VN)</option>
		        </select>

		        <label>Voice B:</label>
		        <select id="voiceB" onchange="validateVoices()">
        <option value="en-US-AndrewMultilingualNeural">Andrew (EN)</option>
        <option value="en-US-ChristopherNeural">Christopher (EN)</option>
        <option value="en-US-EricNeural">Eric (EN)</option>
        <option value="vi-VN-HoaiMyNeural">Hoài My (VN)</option>
        <option value="vi-VN-NamMinhNeural">Nam Minh (VN)</option>
		        </select>
		    </div>

		    <h3>Hội thoại</h3>
		    <div id="chatBox" class="chat-box"></div>

		    <div class="input-container">
		        <input type="text" id="messageInput" placeholder="Nhập tin nhắn...">
		        <button onclick="sendMessage('A')">Gửi (A)</button>
		        <button onclick="sendMessage('B')">Gửi (B)</button>
		        <button onclick="generateSSML()">Xuất SSML</button>
		    </div>

		    <h3>SSML Output:</h3>
		    <div class="ssml-box">
		        <button class="copy-btn" onclick="copySSML()">📋</button>
		        <pre id="ssml"></pre>
		    </div>
		</body>
		</html>
		`)
	})

	app.HandleDir("/static", iris.Dir("./static"))
	app.Listen(":8080")
}
