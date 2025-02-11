document.addEventListener("DOMContentLoaded", function () {
    const sendButton = document.getElementById("sendButton");
    if (sendButton) {
        sendButton.addEventListener("click", sendPrompt);
    } else {
        console.error("❌ Không tìm thấy nút gửi!");
    }
});

async function sendPrompt() {
    const prompt = document.getElementById("prompt").value;
    const responseDiv = document.getElementById("response");

    if (!prompt.trim()) {
        responseDiv.innerHTML = "<p style='color:red'>❌ Vui lòng nhập prompt!</p>";
        return;
    }

    responseDiv.innerHTML = "⏳ Đang xử lý...";

    try {
        console.log("🔍 Gửi request đến API...");

        const response = await fetch("http://localhost:8080/api/groq", {  // ✅ Đảm bảo đường dẫn API chính xác
            method: "POST",  // ✅ Đúng method POST
            mode: "cors",  // ✅ Cho phép gửi request CORS
            headers: {
                "Content-Type": "application/json",
                "Accept": "application/json"
            },
            body: JSON.stringify({ prompt }) // ✅ Gửi đúng JSON format
        });

        if (!response.ok) {
            throw new Error(`Lỗi HTTP: ${response.status}`);
        }

        const data = await response.json();
        responseDiv.innerHTML = `<strong>Phản hồi:</strong> <br> ${data.response}`;
    } catch (error) {
        console.error("❌ Lỗi khi gọi API:", error);
        responseDiv.innerHTML = "<p style='color:red'>❌ Không thể kết nối API!</p>";
    }
}
