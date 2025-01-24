import RoomWindow from "../components/roomwindow"
import Sidebar from "../components/sidebar"

const RoomLayout = () => {
  return (
    <div className="flex h-screen">
      <Sidebar />
      <RoomWindow />
    </div>
  )
}

export default RoomLayout
