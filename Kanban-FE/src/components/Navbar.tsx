import { useContext, useEffect, useState } from "react";
import { IoMdNotifications } from "react-icons/io";
import { NavLink } from "react-router";
import { AuthContext } from "../context/auth";
import { axiosInstance } from "../utils/axios";

const Navbar = () => {
  const { auth } = useContext(AuthContext) || { auth: null };
  const [notiCount, setNotiCount] = useState(0);
  const [openNoti, setOpenNoti] = useState(false);
  const [notifications, setNotifications] = useState([]);
  const handleNotiClick = () => {
    setOpenNoti(!openNoti);
  };
  const handleLogout = async () => {
    try {
      const response = await axiosInstance.post("/users/logout");
      if (response.status === 200) {
        window.location.href = "/login";
      }
    } catch (error) {
      console.error("Error during logout:", error);
    }
  };
  const fetchNotificationCount = async () => {
    try {
      const response = await axiosInstance.get("/notifications/unread/count");
      if (response.status === 200) {
        setNotiCount(response.data.data);
      }
    } catch (error) {
      console.error("Error fetching notifications:", error);
    }
  };
  const fetchNotifications = async () => {
    try {
      const response = await axiosInstance.get("/notifications/my");
      if (response.status === 200) {
        setNotifications(response.data.data);
      }
    } catch (error) {
      console.error("Error fetching notifications:", error);
    }
  };
  const handleNotificationClick = (id) => {
    try {
      const response = axiosInstance.put(`/notifications/mark-as-read/${id}`);
      if (response.status === 200) {
        setNotifications((prev) => prev.filter((item) => item.ID !== id));
        setNotiCount((prev) => prev - 1);
      }
    } catch (error) {
      console.error("Error handling notification click:", error);
    }
  };
  useEffect(() => {
    const interval = setInterval(() => {
      fetchNotificationCount();
      fetchNotifications();
    }, 1000);
    return () => clearInterval(interval);
  }, []);
  return (
    <nav className="w-full h-fit p-4 bg-slate-400 m-0 grid grid-cols-2 items-center">
      <NavLink to="/" className="">
        Clicknext Kanban
      </NavLink>
      <div
        className="flex justify-end items-center gap-4"
        onClick={handleNotiClick}
      >
        <div>
          <div className="relative">
            <IoMdNotifications className="text-4xl" />
            {notiCount > 0 && (
              <span className="absolute top-0 right-0 bg-red-500 text-white rounded-full w-4 h-4 flex items-center justify-center text-xs">
                {notiCount}
              </span>
            )}
          </div>
          {openNoti && (
            <div className="relative">
              <div className="w-64 absolute top-0 right-0 bg-white shadow-lg rounded-lg p-4">
                <h1 className="text-lg font-bold">Notifications</h1>
                <div className="flex flex-col gap-2 mt-2">
                  <div className="flex flex-col gap-4 items-center justify-between">
                    <h1 className="text-sm font-semibold"></h1>
                    {notifications.length === 0
                      ? "No notifications"
                      : notifications.map((item) => (
                          <div
                            onClick={handleNotificationClick.bind(
                              this,
                              item.ID
                            )}
                            key={item.ID}
                            className={`flex flex-col rounded p-2 gap-2 w-full ${
                              item.read ? "bg-slate-300" : "bg-slate-50"
                            }`}
                          >
                            <h1 className="text-left">{item.title}</h1>
                            <p className="text-sm text-gray-500">
                              {item.message}
                            </p>
                            <span className="text-xs text-gray-500">
                              {new Date(item.CreatedAt).toLocaleString()}
                            </span>
                          </div>
                        ))}
                  </div>
                </div>
              </div>
            </div>
          )}
        </div>

        <div>
          <div className="grid grid-cols-2 items-center divide-x-2 divide-slate-600">
            <NavLink to="/profile/" end className={"text-center px-2"}>
              {auth?.firstName} {auth?.lastName}
            </NavLink>
            <button className="text-center px-2" onClick={handleLogout}>
              logout
            </button>
          </div>
        </div>
      </div>
    </nav>
  );
};

export default Navbar;
