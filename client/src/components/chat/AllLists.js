import React, { useState } from 'react';
import profile from '../../assets/about.jpg'

const AllList = () => {
  const [searchQuery, setSearchQuery] = useState('');

  const handleSearch = (event) => {
    setSearchQuery(event.target.value);
    // Perform search/filter logic here
  };

  const friendChats = [
    {
      id: 1,
      name: 'John Doe',
      lastMessage: 'Hello there!',
      lastMessageTime: '9:30 AM',
      profileImage: profile,
    },
    {
      id: 2,
      name: 'Jane Smith',
      lastMessage: 'How are you?',
      lastMessageTime: '8:45 AM',
      profileImage: profile,
    },
    // Add more friend chat objects here
  ];

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
              src={chat.profileImage}
              alt={chat.name}
              className="w-10 h-10 rounded-full mr-2"
            />
            <div>
              <h3 className="font-medium text-white">{chat.name}</h3>
              <p className="text-sm text-gray-500">{chat.lastMessage}</p>
            </div>
            <span className="text-sm ml-auto text-gray-500">{chat.lastMessageTime}</span>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default AllList;
