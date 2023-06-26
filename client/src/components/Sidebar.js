// import React from 'react';
import {
  FiBell,
  FiMessageCircle,
  FiLogOut,
  FiArchive,
  FiTrash2,
} from "react-icons/fi";
import logo from '../assets/about.jpg'

const Sidebar = () => {
  const chatRooms = [
    { title: "Notification", icon: <FiBell className="text-gray-600 mr-2" /> },
    { title: "User-List", icon: <FiBell className="text-gray-600 mr-2" /> },
    { title: "Pinned", icon: null },
    { title: "All", icon: <FiMessageCircle className="text-gray-600 mr-2" /> },
    { title: "Archived", icon: <FiArchive className="text-gray-600 mr-2" /> },
    { title: "Trash", icon: <FiTrash2 className="text-gray-600 mr-2" /> },
    { title: "Settings", icon: null },
  ];

  return (
    <div className="flex flex-col w-64 justify-between bg-black text-white p-5 h-screen">
      <div className="flex flex-col">
        <h2 className="flex justify-center text-center font-bold text-3xl p-5 border-b-2 border-gray-900">
          DevTalks
        </h2>
        <ul className="flex flex-col mt-5">
          {chatRooms.map((room, index) => (
            <li
              key={index}
              className="flex items-center px-4 py-2 rounded-xl hover:bg-[#F3FC8A] hover:text-black cursor-pointer"
            >
              {room.icon}
              <span>{room.title}</span>
            </li>
          ))}
        </ul>
      </div>
      <div className="flex items-center flex-col justify-end">
        <img
          className="w-32 h-32 rounded-full mr-2"
          src={logo}
          alt="User Profile"
        />
        <h2 className="text-lg font-bold">Username</h2>
      </div>
    </div>
  );
};

export default Sidebar;
