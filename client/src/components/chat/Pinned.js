import React from 'react';

const PinnedList = () => {
  const pinnedChats = [
    { id: 1, title: 'Chat 1', lastMessage: 'Last message 1' },
    { id: 2, title: 'Chat 2', lastMessage: 'Last message 2' },
    { id: 3, title: 'Chat 3', lastMessage: 'Last message 3' },
    // Add more pinned chats here
  ];

  return (
    <div>
      <h2 className="flex justify-center text-center items-center text-3xl font-semibold text-white p-5">Pinned List</h2>
      {pinnedChats.length > 0 ? (
        <ul className='px-5'>
          {pinnedChats.map((chat) => (
            <li key={chat.id} className="flex items-center justify-between p-2 mb-2 bg-gray-200 rounded-lg">
              <div>
                <h3 className="font-bold">{chat.title}</h3>
                <p className="text-sm text-gray-500">{chat.lastMessage}</p>
              </div>
              <button className="text-red-500 font-bold">Unpin</button>
            </li>
          ))}
        </ul>
      ) : (
        <p className="text-gray-500">No pinned chats</p>
      )}
    </div>
  );
};

export default PinnedList;
