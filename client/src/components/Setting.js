import React from 'react';
import { FiUser, FiLock, FiShield, FiLogOut } from 'react-icons/fi';

const Setting = () => {
  const settings = [
    { id: 1, title: 'Profile', icon: <FiUser className="mr-2" /> },
    { id: 2, title: 'Privacy', icon: <FiLock className="mr-2" /> },
    { id: 3, title: 'Security', icon: <FiShield className="mr-2" /> },
    { id: 4, title: 'Logout', icon: <FiLogOut className="mr-2" /> },
  ];

  return (
    <div>
      <h2 className="flex justify-center text-center items-center text-3xl font-semibold text-white p-5">Settings</h2>
      <ul className='px-5'>
        {settings.map((setting) => (
          <li key={setting.id} className="p-2 mb-2 bg-gray-200 rounded-lg flex items-center">
            {setting.icon}
            <span>{setting.title}</span>
          </li>
        ))}
      </ul>
    </div>
  );
};

export default Setting;
