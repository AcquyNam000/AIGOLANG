document.addEventListener("DOMContentLoaded", function () {
    const sendButton = document.getElementById("sendButton");
    sendButton.addEventListener("click", sendPrompt);
});

async function sendPrompt() {
    const promptInput = document.getElementById("prompt");
    const chatBox = document.getElementById("chatBox");
    const userMessage = promptInput.value.trim();

    if (userMessage === "") return;

    // Hiển thị câu hỏi của User
    chatBox.innerHTML += `<div class="user-message">${userMessage}</div>`;
    chatBox.scrollTop = chatBox.scrollHeight;

    // Gửi API
    try {
        const response = await fetch("http://localhost:8080/api/groq", {
            method: "POST",
            headers: { "Content-Type": "application/json" },
            body: JSON.stringify({ prompt: userMessage })
        });

        const data = await response.json();
        const botMessage = data.response || "❌ Lỗi API!";

        // Hiển thị phản hồi từ Bot
        chatBox.innerHTML += `<div class="bot-message">${botMessage}</div>`;
        chatBox.scrollTop = chatBox.scrollHeight;
    } catch (error) {
        console.error("❌ Lỗi API:", error);
        chatBox.innerHTML += `<div class="bot-message">Lỗi khi gửi yêu cầu!</div>`;
    }

    promptInput.value = ""; // Xóa ô nhập sau khi gửi
}
