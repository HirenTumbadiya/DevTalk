import React from 'react';
import ChatRoom from './ChatRoom';

const MainWindow = ({ selectedChat, renderFirstPart }) => {
  return (
    <div className=" w-full bg-black flex flex-grow p-2 h-screen">
      <div className=" w-1/4 bg-[#20232B]">
        {/* Render the content of the first part */}
        {renderFirstPart()}
      </div>
      <div className="flex-grow bg-gray-300">
        <ChatRoom selectedChat={selectedChat} />
      </div>
    </div>
  );
};

export default MainWindow;
