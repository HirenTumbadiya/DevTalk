import React, { useState } from 'react';
import Sidebar from '../components/Sidebar';
import MainWindow from '../components/MainWindow';
import AllList from '../components/chat/AllLists';
import FriendList from '../components/chat/FriendList';

const ChatRoomPage = () => {
  const [selectedOption, setSelectedOption] = useState("all");
  const [selectedChat, setSelectedChat] = useState(null);

  const handleOptionClick = (option) => {
    setSelectedOption(option);
    setSelectedChat(null); // Reset the selected chat when changing options
  };

  const handleChatClick = (chat) => {
    setSelectedChat(chat);
  };

  const renderFirstPart = () => {
    if (selectedOption === 'all') {
      return <FriendList onChatClick={handleChatClick} />;
    }
    // Render other components based on selectedOption
    return null;
  };

  return (
    <div className="flex">
      <Sidebar onOptionClick={handleOptionClick} />
      <div className="flex-1">
        <MainWindow selectedChat={selectedChat} renderFirstPart={renderFirstPart} />
      </div>
    </div>
  );
};

export default ChatRoomPage;
