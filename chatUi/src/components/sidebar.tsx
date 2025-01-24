import { Link } from "react-router-dom";

const Sidebar: React.FC = () => {
  return (
    <div className="w-1/4 bg-gray-800 text-white h-full p-4">
      <h2 className="text-xl font-bold mb-4">Channels</h2>
      <ul>
        <li className="mb-2 cursor-pointer hover:text-blue-400">General</li>
        <li className="mb-2 cursor-pointer hover:text-blue-400">Random</li>
        <li className="mb-2 cursor-pointer hover:text-blue-400">
          <Link to={"/chat"}>Chat</Link>
        </li>
        <li className="mb-2 cursor-pointer hover:text-blue-400">
          <Link to={"/room"}>Room</Link>
        </li>
      </ul>
    </div>
  );
};

export default Sidebar;
