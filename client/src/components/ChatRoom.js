import React from 'react';
import { FiPaperclip, FiSmile, FiSend } from 'react-icons/fi';
import profile from "../assets/about.jpg"

const ChatRoom = ({ chat }) => {
    const handleSendMessage = (message) => {
        // Handle sending the message logic here
      };

  return (
    <div className="flex flex-col h-full">
      {/* Upper bar */}
      <div className="flex items-center justify-between bg-black p-2">
        <div className="flex items-center">
          <img
            className="w-8 h-8 rounded-full mr-2"
            src={profile}
            alt="Profile Picture"
          />
          <div>
            <h2 className="font-bold text-white">Username</h2>
            <p className="text-gray-400 text-sm">Last seen: 8:30pm</p>
          </div>
        </div>
        <div className="flex items-center bg-[#1D1E24]">
          {/* Include your search bar component here */}
        </div>
      </div>

{/* Main chat content */}
<div className="flex-grow bg-[#1D1E24] flex flex-col-reverse">
  {/* Render the chat messages here */}
  <div className="p-4">
    <div className="flex items-center mb-2">
      <img
        className="w-8 h-8 rounded-full mr-2"
        src={profile}
        alt="User Avatar"
      />
      <div className="bg-gray-800 text-white p-2 rounded-lg">
        <p className="text-sm">Hello! How are you doing?</p>
      </div>
    </div>
    <div className="flex items-center mb-2 justify-end">
      <div className="bg-gray-600 text-white p-2 rounded-lg">
        <p className="text-sm">I'm good, thanks! How about you?</p>
      </div>
      <img
        className="w-8 h-8 rounded-full ml-2"
        src={profile}
        alt="User Avatar"
      />
    </div>
    <div className="flex items-center mb-2">
      <img
        className="w-8 h-8 rounded-full mr-2"
        src={profile}
        alt="User Avatar"
      />
      <div className="bg-gray-800 text-white p-2 rounded-lg">
        <p className="text-sm">I'm doing great, thank you!</p>
      </div>
    </div>
    <div className="flex items-center mb-2 justify-end">
      <div className="bg-gray-600 text-white p-2 rounded-lg">
        <p className="text-sm">I'm good, thanks! How about you?</p>
      </div>
      <img
        className="w-8 h-8 rounded-full ml-2"
        src={profile}
        alt="User Avatar"
      />
    </div>
    <div className="flex items-center mb-2">
      <img
        className="w-8 h-8 rounded-full mr-2"
        src={profile}
        alt="User Avatar"
      />
      <div className="bg-gray-800 text-white p-2 rounded-lg">
        <p className="text-sm">I'm doing great, thank you!</p>
      </div>
    </div>
  </div>
</div>


            {/* Message input section */}
            <div className="flex items-center bg-[#1D1E24] p-2">
        <input
          type="text"
          placeholder="Type a message..."
          className="flex-grow rounded-xl py-3 px-4 bg-[#16171B] text-white focus:outline-none"
        />
        <div className="flex items-center ml-1">
          <button className="mr-2">
            <FiPaperclip className="text-white text-xl" />
          </button>
          <button className="mr-2">
            <FiSmile className="text-white text-xl" />
          </button>
          <button onClick={() => handleSendMessage('')} className="text-black bg-[#F3FC8A] py-2 px-2 rounded-3xl flex justify-center items-center">
            <FiSend className="text-black text-xl" />
          </button>
        </div>
      </div>
    </div>
  );
};

export default ChatRoom;
