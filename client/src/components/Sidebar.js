import React, { useEffect, useState } from 'react';
import {
  FiBell,
  FiMessageCircle,
  FiLogOut,
  FiArchive,
  FiTrash2,
  FiSettings,
  FiUsers,
  FiBookmark,
} from 'react-icons/fi';
import logo from '../assets/about.jpg';

const Sidebar = ({ onOptionClick }) => {
  const [selectedOption, setSelectedOption] = useState('all');
  const [username, setUsername] = useState(''); 

    useEffect(() => {
    // Retrieve the user ID from localStorage
    const userID = localStorage.getItem('id');
    if (userID) {
      // Fetch the user's username
      fetchUserByID(userID);
    }
  }, []);

  const fetchUserByID = async (userID) => {
    try {
      const response = await fetch(`http://localhost:8000/users/${userID}`);
      if (response.ok) {
        const data = await response.json();
        console.log(data)
        setUsername(data.username);
      } else {
        console.error('Failed to fetch user:', response.status);
      }
    } catch (error) {
      console.error('Error:', error);
    }
  };

  const handleItemClick = (option) => {
    setSelectedOption(option);
    onOptionClick(option);
  };

  const chatRooms = [
    { title: 'Notification', icon: <FiBell className="text-gray-600 mr-2" />, option: 'notification' },
    { title: 'Users', icon: <FiUsers className="text-gray-600 mr-2" />, option: 'user-list' },
    { title: 'Chats', icon: <FiMessageCircle className="text-gray-600 mr-2" />, option: 'all' },
    { title: 'Pinned', icon: <FiBookmark className="text-gray-600 mr-2" />, option: 'pinned' },
    { title: 'Archived', icon: <FiArchive className="text-gray-600 mr-2" />, option: 'archived' },
    { title: 'Trash', icon: <FiTrash2 className="text-gray-600 mr-2" />, option: 'trash' },
    { title: 'Settings',icon: <FiSettings className="text-gray-600 mr-2" />, option: 'settings' },
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
              className={`flex items-center px-4 py-2 rounded-xl cursor-pointer ${
                room.option === selectedOption ? 'bg-[#F3FC8A] text-black' : 'hover:bg-[#F3FC8A] hover:text-black'
              }`}
              onClick={() => handleItemClick(room.option)} // Handle item click
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
        <h2 className="text-lg font-bold">{username}</h2>
      </div>
    </div>
  );
};

export default Sidebar;
