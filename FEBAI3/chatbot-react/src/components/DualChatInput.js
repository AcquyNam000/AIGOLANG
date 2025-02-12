import React, { useState, useRef, useEffect } from "react";

function DualChatInput({ onSubmit }) {
  const [messages, setMessages] = useState([]);
  const messageEndRef = useRef(null);

  useEffect(() => {
    messageEndRef.current?.scrollIntoView({ behavior: "smooth" });
  }, [messages]);

  const handleSend = (sender) => {
    const inputField = document.getElementById("dual-chat-input");
    if (!inputField.value.trim()) return;

    const newMessage = { sender, text: inputField.value };
    setMessages([...messages, newMessage]);
    inputField.value = "";
  };

  const handleEdit = (index) => {
    const newText = prompt("Chỉnh sửa tin nhắn:", messages[index].text);
    if (newText !== null) {
      const updatedMessages = [...messages];
      updatedMessages[index].text = newText;
      setMessages(updatedMessages);
    }
  };

  const handleSubmit = () => {
    if (messages.length === 0) return;
    onSubmit(messages.map((msg) => `${msg.sender}: ${msg.text}`).join("\n"));
    setMessages([]);
  };

  return (
    <div className="dual-chat-input">
      <div className="chat-box">
        {messages.map((msg, index) => (
          <div
            key={index}
            className={`message ${msg.sender === "A" ? "sent" : "received"}`}
          >
            <span className="sender">{msg.sender}:</span>
            {msg.text}
            <button className="edit-button" onClick={() => handleEdit(index)}>
              ✏️
            </button>
          </div>
        ))}
        <div ref={messageEndRef}></div>
      </div>
      <textarea id="dual-chat-input" placeholder="Nhập tin nhắn..." />
      <div className="button-group">
        <button onClick={() => handleSend("A")}>Gửi A</button>
        <button onClick={() => handleSend("B")}>Gửi B</button>
      </div>
      <button onClick={handleSubmit} style={{ marginTop: "10px" }}>
        Gửi Hội Thoại
      </button>
    </div>
  );
}

export default DualChatInput;
