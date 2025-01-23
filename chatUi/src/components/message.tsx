import React from "react";
const ChatMessage: React.FC<{ time: string; sender: string; content: string; isUser: boolean }> = ({
  time,
  sender,
  content,
  isUser,
}) => (
  <div className={`flex items-start space-x-2 ${isUser ? 'justify-end' : ''}`}>
    {isUser && (
      <img
        src="https://res.cloudinary.com/djv4xa6wu/image/upload/v1735722163/AbhirajK/Abhirajk%20mykare.webp"
        alt="User"
        className="w-8 h-8 rounded-full object-cover"
      />
    )}
    <div>
      <div className={`bg-${isUser ? 'blue-600' : 'white'} rounded-lg rounded-tl-none p-3 shadow-md max-w-md`}>
        <p>{content}</p>
      </div>
      <span className="text-gray-500 text-xs message-time">{time}</span>
    </div>
    {!isUser && (
      <img
        src="https://res.cloudinary.com/djv4xa6wu/image/upload/v1735722165/AbhirajK/Abhirajk.webp"
        alt="Abhiraj"
        className="w-8 h-8 rounded-full object-cover"
      />
    )}
  </div>
);


export default ChatMessage
