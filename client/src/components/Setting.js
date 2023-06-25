import React from 'react';

const Setting = () => {
  const settings = [
    { id: 1, title: 'Setting 1' },
    { id: 2, title: 'Setting 2' },
    { id: 3, title: 'Setting 3' },
    // Add more settings here
  ];

  return (
    <div>
      <h2 className="text-xl font-bold mb-4">Settings</h2>
      <ul>
        {settings.map((setting) => (
          <li key={setting.id} className="p-2 mb-2 bg-gray-200 rounded-lg">
            {setting.title}
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Setting;
