import React from "react";

function ChatBox({ dialog }) {
  return (
    <div className="chat-box">
      {dialog ? <pre>{dialog.content}</pre> : <p>Nhập hội thoại để bắt đầu</p>}
    </div>
  );
}

export default ChatBox;
