

import React from "react";

const ChatInput: React.FC = () => {
  const [message, setMessage] = React.useState("");

  const handleSend = () => {
    if (message.trim() !== "") {
      console.log("Message sent:", message);
      setMessage(""); // Clear input after sending
    }
  };

  return (
    <div className="flex items-center mt-4">
      <input
        type="text"
        placeholder="Type your message..."
        value={message}
        onChange={(e) => setMessage(e.target.value)}
        className="flex-1 p-2 border rounded-lg focus:outline-none"
      />
      <button
        onClick={handleSend}
        className="ml-2 px-4 py-2 bg-blue-500 text-white rounded-lg hover:bg-blue-600"
      >
        Send
      </button>
    </div>
  );
};

export default ChatInput;
