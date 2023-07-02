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
      <h2 className="flex justify-center text-center items-center text-3xl font-semibold text-white p-5">Trash List</h2>
      {trashItems.length > 0 ? (
        <ul className='px-5'>
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
