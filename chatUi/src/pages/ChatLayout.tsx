
import Sidebar from "../components/sidebar";
import ChatWindow from "../components/chatwindow";

const ChatLayout: React.FC = () => {
  return (
    <div className="flex h-screen">
      <Sidebar />
      <ChatWindow />
    </div>
  );
};

export default ChatLayout;

