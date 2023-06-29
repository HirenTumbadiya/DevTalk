import React, { useState } from 'react';
import { BrowserRouter as Router, Routes, Route, Link } from 'react-router-dom';
import Sidebar from '../components/Sidebar';
import MainWindow from '../components/MainWindow';
import AllList from '../components/chat/AllLists';
import FriendList from '../components/chat/FriendList';
import Notification from '../components/Notification';
import ArchivedList from '../components/chat/ArchivedList';
import Pinned from '../components/chat/Pinned';
import TrashList from '../components/chat/TrashList';
import Setting from '../components/Setting';

const ChatRoomPage = () => {
  const [selectedOption, setSelectedOption] = useState('all');
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
      return <AllList onChatClick={handleChatClick} />;
    } else if (selectedOption === 'notification') {
      return <Notification />;
    }else if (selectedOption === 'user-list') {
      return <FriendList />;
    }else if (selectedOption === 'archived') {
      return <ArchivedList />;
    }else if (selectedOption === 'pinned') {
      return <Pinned />;
    }else if (selectedOption === 'trash') {
      return <TrashList />;
    }else if (selectedOption === 'settings') {
      return <Setting />;
    }
    return null;
  };

  return (
    <Router>
      <div className="flex">
        <Sidebar onOptionClick={handleOptionClick} />
        <div className="flex-1">
          <Routes>
            <Route path="/" element={<MainWindow selectedChat={selectedChat} renderFirstPart={renderFirstPart} />} />
            <Route path="/notification" element={<Notification />} />
          </Routes>
        </div>
      </div>
    </Router>
  );
};

export default ChatRoomPage;
