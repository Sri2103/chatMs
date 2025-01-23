import React, { useState } from 'react';
import ChatMessage from './message';
interface Message {
  time: string;
  sender: string;
  content: string;
  isUser: boolean;
}


const Chat: React.FC = () => {
  const [messages, setMessages] = useState<Message[]>([
    {
      time: '11:42 AM',
      sender: 'Abhiraj',
      content:
        "Hi there! Welcome to my portfolio chat. Feel free to explore my work at <a href='https://abhirajk.vercel.app' target='_blank'>abhirajk.vercel.app</a>",
      isUser: false,
    },
    {
      time: '11:43 AM',
      sender: 'User',
      content: "Amazing work! I'm impressed with your projects and technical skills.",
      isUser: true,
    },
    {
      time: '11:44 AM',
      sender: 'Abhiraj',
      content:
        "Thank you! I specialize in full-stack development using modern technologies. Would you like to discuss a potential collaboration?",
      isUser: false,
    },
  ]);

  const [typing, setTyping] = useState(true);

  const handleSendMessage = (message: string) => {
    setMessages([
      ...messages,
      { time: '11:45 AM', sender: 'User', content: message, isUser: true },
    ]);
    // Simulate a delay for the bot's response
    setTimeout(() => {
      setMessages([
        ...messages,
        {
          time: '11:46 AM',
          sender: 'Abhiraj',
          content: "I'd love to! Let's schedule a quick call to discuss further.",
          isUser: false,
        },
      ]);
      setTyping(false);
    }, 1000);
  };

  return (
    <div className="w-full bg-gray-100 h-screen flex flex-col">
      {/* Chat Header */}
      <div className="bg-blue-600 text-white p-4 shadow-lg">
        <div className="max-w-4xl mx-auto flex items-center justify-between">
          <div className="flex items-center space-x-4">
            <img
              src="https://res.cloudinary.com/djv4xa6wu/image/upload/v1735722165/AbhirajK/Abhirajk.webp"
              alt="Profile"
              className="w-10 h-10 rounded-full object-cover border-2 border-white"
            />
            <div>
              <h1 className="font-bold">Abhiraj K</h1>
              <a
                href="https://abhirajk.vercel.app"
                target="_blank"
                className="text-sm text-blue-100 hover:text-white transition"
              >
                Visit Portfolio →
              </a>
            </div>
          </div>
          <div className="flex items-center space-x-4">
            {/* ... (Header buttons) */}
          </div>
        </div>
      </div>

      {/* Chat Messages */}
      <div className="flex-1 overflow-y-auto p-4 chat-container">
        <div className="max-w-4xl mx-auto space-y-4">
          {messages.map((message, index) => (
            <ChatMessage key={index} {...message} />
          ))}

          {typing && (
            <div className="flex items-start space-x-2">
              <img
                src="https://res.cloudinary.com/djv4xa6wu/image/upload/v1735722165/AbhirajK/Abhirajk.webp"
                alt="Abhiraj"
                className="w-8 h-8 rounded-full object-cover"
              />
              <div className="bg-white rounded-lg rounded-tl-none p-3 shadow-md">
                <div className="typing-indicator">
                  <span>•</span>
                  <span>•</span>
                  <span>•</span>
                </div>
              </div>
            </div>
          )}
        </div>
      </div>

      {/* Chat Input */}
      <div className="bg-white border-t p-4">
        <div className="max-w-4xl mx-auto flex items-center space-x-4">
          {/* ... (Input buttons) */}
          <input
            type="text"
            placeholder="Type your message..."
            className="flex-1 p-2 border rounded-full focus:outline-none focus:border-blue-500"
            onKeyPress={(e) => {
              if (e.key === 'Enter') {
                handleSendMessage(e.target.value);
                e.target.value = ''; // Clear input field
              }
            }}
          />
          <button className="p-2 text-white bg-blue-600 rounded-full hover:bg-blue-700 transition">
            <svg className="w-6 h-6" fill="none" stroke="currentColor" viewBox="0 0 24 24">
              <path
                stroke-linecap="round"
                stroke-linejoin="round"
                stroke-width="2"
                d="M12 19l9 2-9-18-9 18 9-2zm0 0v-8"
              />
            </svg>
          </button>
        </div>
      </div>
    </div>
  );
};

export default Chat;
