import React, { useState, useEffect } from 'react';
import axios from 'axios';
import { FaChevronUp, FaChevronDown } from 'react-icons/fa';

const FriendList = ({ onChatClick }) => {
  const currentUserId = localStorage.getItem("id");
  const [searchTerm, setSearchTerm] = useState('');
  const [searchResults, setSearchResults] = useState(null);
  const [notification, setNotification] = useState(null);
  const [friendsList, setFriendsList] = useState([]);
  const [isExpanded, setIsExpanded] = useState(true);

  const handleSearch = async () => {
    try {
      const response = await axios.get(`http://localhost:8000/users/search?username=${searchTerm}`);
      console.log(response);
      setSearchResults(response.data);
    } catch (error) {
      console.error('Error searching for users:', error);
      setSearchResults([]); // Set searchResults to an empty array when no users are found or there is an error
    }
  };
  
  const handleSendRequest = async (receiver) => {
    try {
      const response = await axios.post('http://localhost:8000/friend-requests/send', {
        UserID: currentUserId,
        FriendID: receiver,
      });
      console.log(response);
      
      // Check the response data and handle it accordingly
      if (response.data === '') {
        setNotification('Friend request sent successfully');
      } else {
        setNotification('Error sending friend request');
      }
    } catch (error) {
      console.error('Error sending friend request:', error);
    }
  };

  const getFriendsList = async () => {
    try {
      const response = await axios.get(`http://localhost:8000/friends?userId=${currentUserId}`);
      console.log(response);
      setFriendsList(response.data);
    } catch (error) {
      console.error('Error getting friends list:', error);
    }
  };

  useEffect(() => {
    // Perform the search when the search term changes
    handleSearch();
  }, [searchTerm]);

  useEffect(() => {
    // Fetch the friends list when the component mounts
    getFriendsList();
  }, []);
  const toggleExpand = () => {
    setIsExpanded((prevExpanded) => !prevExpanded);
  };

  const handleChatWithFriend = (friend) => {
    // Pass the selected friend to the onChatClick callback
    onChatClick(friend);
  };

  return (
    <div>
      <div className="p-5">
        <input
          className='w-full py-3 px-2 rounded-xl bg-[#16171B] text-white focus:outline-none'
          type="text"
          placeholder="Search for users"
          value={searchTerm}
          onChange={(e) => setSearchTerm(e.target.value)}
        />
      </div>
      
      {searchResults !== null && searchResults.length > 0 ? (
        <ul>
          {searchResults.map((user) => (
            <li key={user.id} className="flex items-center p-3 hover:bg-gray-700 mx-3 rounded-xl">
              {/* <img
                src={user.image} // Replace with the actual image source for the user
                alt={user.name}
                className="w-10 h-10 rounded-full mr-4"
              /> */}
              <div>
                <p className="font-bold text-white text-xl">{user.username}</p>
              </div>
              <button
                onClick={() => handleSendRequest(user.id)}
                className="ml-auto bg-blue-500 hover:bg-blue-700 text-white py-2 px-4 rounded"
              >
                <i className="fas fa-user-plus">Send</i>
              </button>
            </li>
          ))}
        </ul>
      ) : (
        searchResults === null ? (
          <p className='h-full w-full flex justify-center items-center'>Loading...</p>
        ) : (
          <p className='px-5'>No users found</p>
        )
      )}

      <div>
        <div className="flex items-center justify-between px-5 cursor-pointer" onClick={toggleExpand}>
        <p className='text-white'>Friends List</p>
        {isExpanded ? (
          <FaChevronUp size={20} color={"#fafafa"} />
        ) : (
          <FaChevronDown size={20} color={"#fafafa"} />
        )}
      </div>
      {isExpanded && (
        <>
        {friendsList.length > 0 ? (
          <ul>
            {friendsList.map((friend) => (
              <li key={friend.id} className="flex items-center p-3 hover:bg-gray-700 mx-3 rounded-xl cursor-pointer">
                <div>
                <p className="font-bold text-white text-xl">{friend.username}</p>
              </div>
              <button
                onClick={() => handleChatWithFriend(friend)}
                className="ml-auto bg-blue-500 hover:bg-blue-700 text-white py-2 px-4 rounded"
              >
                <i className="fas fa-user-plus">Chat</i>
              </button>
              </li>
            ))}
          </ul>
        ) : (
          <p className='p-5 text-white'>No friends found</p>
        )}
      </>
      )}
      </div>
      
      {notification && (
        <div>
          <p>{notification}</p>
        </div>
      )}
    </div>
  );
};

export default FriendList;
