import React, { useState } from 'react';
import profile from '../../assets/about.jpg';

const AllList = ({ onChatClick }) => {
  const [searchQuery, setSearchQuery] = useState('');
  const [friendChats, setFriendChats] = useState([]);

  const handleSearch = (event) => {
    setSearchQuery(event.target.value);
  };

  const startChatWithFriend = (friend) => {
    // Check if the friend is already in the chat list
    if (friendChats.some((chat) => chat.id === friend.id)) {
      return; // Don't add duplicate entries
    }

    // Add the friend to the chat list
    setFriendChats((prevChats) => [...prevChats, friend]);

    // Pass the selected friend to the onChatClick callback
    onChatClick(friend);
  };

  // Filter the friend chats based on the search query
  const filteredChats = friendChats.filter((chat) =>
    chat.name.toLowerCase().includes(searchQuery.toLowerCase())
  );

  return (
    <div>
      <div className="p-5">
        <input
          className="w-full py-3 px-2 rounded-xl bg-[#16171B] text-white focus:outline-none"
          type="text"
          placeholder="Search"
          value={searchQuery}
          onChange={handleSearch}
        />
      </div>

      <ul>
        {filteredChats.map((chat) => (
          <li key={chat.id} className="flex items-center p-3 hover:bg-gray-700 mx-3 rounded-xl">
            <img
              src={profile}
              alt={chat.name}
              className="w-10 h-10 rounded-full mr-2"
            />
            <div>
              <h3 className="font-medium text-white">{chat.name}</h3>
            </div>
            <button
              className="ml-auto text-sm text-gray-500"
              onClick={() => startChatWithFriend(chat)}
            >
              Start Chat
            </button>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default AllList;
