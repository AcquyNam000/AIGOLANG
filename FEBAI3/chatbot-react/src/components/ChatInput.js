import React, { useState } from "react";

function ChatInput({ onSubmit, mode }) {
  const [input, setInput] = useState("");

  const handleSend = () => {
    if (input.trim() === "") return;
    onSubmit(input);
    setInput("");
  };

  return (
    <div className="chat-input">
      <textarea
        value={input}
        onChange={(e) => setInput(e.target.value)}
        placeholder={
          mode === "process"
            ? "Nhập một đoạn hội thoại tiếng Việt..."
            : "Nhập từng câu thoại giữa hai người..."
        }
      />
      <button onClick={handleSend}>Gửi</button>
    </div>
  );
}

export default ChatInput;
