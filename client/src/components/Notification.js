import React from 'react';

const Notification = () => {
  const notifications = [
    { id: 1, text: 'Notification 1' },
    { id: 2, text: 'Notification 2' },
    { id: 3, text: 'Notification 3' },
    // Add more notifications here
  ];

  return (
    <div>
      <h2 className="text-xl font-bold mb-4">Notifications</h2>
      {notifications.length > 0 ? (
        <ul>
          {notifications.map((notification) => (
            <li key={notification.id} className="p-2 mb-2 bg-gray-200 rounded-lg">
              {notification.text}
            </li>
          ))}
        </ul>
      ) : (
        <p className="text-gray-500">No notifications</p>
      )}
    </div>
  );
};

export default Notification;
