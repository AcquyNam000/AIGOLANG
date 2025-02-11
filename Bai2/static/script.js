let chatHistory = [];

function sendMessage(person) {
    let messageInput = document.getElementById("messageInput");
    let message = messageInput.value.trim();
    if (message === "") return;

    chatHistory.push({ person, text: message });
    updateChatUI();
    messageInput.value = "";
}

function updateChatUI() {
    let chatBox = document.getElementById("chatBox");
    chatBox.innerHTML = "";
    chatHistory.forEach((chat, index) => {
        let msgElement = document.createElement("div");
        msgElement.classList.add("chat-message", chat.person === 'A' ? "chat-a" : "chat-b");
        msgElement.innerHTML = `<strong>${chat.person}:</strong> <span id="msg-${index}">${chat.text}</span>
        <button onclick="editMessage(${index})">✏️</button>`;
        chatBox.appendChild(msgElement);
    });
    chatBox.scrollTop = chatBox.scrollHeight;
}

function editMessage(index) {
    let newText = prompt("Chỉnh sửa tin nhắn:", chatHistory[index].text);
    if (newText !== null) {
        chatHistory[index].text = newText;
        updateChatUI();
    }
}

function generateSSML() {
    let voiceA = document.getElementById("voiceA").value;
    let voiceB = document.getElementById("voiceB").value;

    let ssml = '<speak xml:lang="vi-VN">\n';
    chatHistory.forEach(chat => {
        let voice = chat.person === 'A' ? voiceA : voiceB;
        ssml += `<voice name="${voice}">${chat.person}: ${chat.text}</voice>\n`;
    });
    ssml += '</speak>';

    document.getElementById("ssml").innerText = ssml;
}

function copySSML() {
    let ssmlText = document.getElementById("ssml").innerText;
    navigator.clipboard.writeText(ssmlText);
    alert("Đã sao chép SSML!");
}

function validateVoices() {
    let voiceA = document.getElementById("voiceA");
    let voiceB = document.getElementById("voiceB");

    if (voiceA.value === voiceB.value) {
        alert("Hai giọng nói không thể giống nhau!");
        voiceB.selectedIndex = (voiceB.selectedIndex + 1) % voiceB.options.length;
    }
}
