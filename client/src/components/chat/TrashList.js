import React from 'react';

const TrashList = () => {
  const trashItems = [
    { id: 1, title: 'Item 1' },
    { id: 2, title: 'Item 2' },
    { id: 3, title: 'Item 3' },
    // Add more trash items here
  ];

  return (
    <div>
      <h2 className="text-xl font-bold mb-4">Trash List</h2>
      {trashItems.length > 0 ? (
        <ul>
          {trashItems.map((item) => (
            <li key={item.id} className="flex items-center justify-between p-2 mb-2 bg-gray-200 rounded-lg">
              <span>{item.title}</span>
              <button className="text-red-500 font-bold">Delete</button>
            </li>
          ))}
        </ul>
      ) : (
        <p className="text-gray-500">Trash is empty</p>
      )}
    </div>
  );
};

export default TrashList;
