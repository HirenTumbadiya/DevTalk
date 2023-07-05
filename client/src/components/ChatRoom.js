import React, { useState, useEffect } from 'react';
import { FiPaperclip, FiSmile, FiSend } from 'react-icons/fi';
import profile from "../assets/about.jpg";
import axios from 'axios';

const ChatRoom = ({ selectedChat }) => {
  console.log(selectedChat)
  const [messages, setMessages] = useState([]);
  const [inputValue, setInputValue] = useState('');
  const [socket, setSocket] = useState(null);
  const [username, SetUserName] = useState();
  const userID = localStorage.getItem('id');

  useEffect(() => {
    SetUserName(selectedChat ? selectedChat.username : '');
  }, [selectedChat])

  const handleSendMessage = async () => {
    if (inputValue.trim() === '') {
      return;
    }

    if (!selectedChat) {
      console.log('No selected chat');
      return;
    }

    let recipientId = selectedChat.userId;

    if (selectedChat.userId === userID) {
      recipientId = selectedChat.friendId;
    } else {
      recipientId = selectedChat.userId
    }

    if (!recipientId) {
      console.log('No recipient ID available');
      return;
    }

    const newMessage = {
      id: Date.now().toString(),
      senderId: userID,
      recipientId: recipientId,
      message: inputValue,
      createdAt: new Date(),
    };

    // Log the message with sender ID and recipient ID
    console.log('Sending message:', newMessage);

    // Update the messages state with the new message
    setMessages(prevMessages => [...prevMessages, newMessage]);

    // Clear the input field after sending the message
    setInputValue('');

    try {
      const response = await fetch('http://localhost:8000/chat/send-message', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          senderId: newMessage.senderId,
          recipientId: newMessage.recipientId,
          message: newMessage.message
        }),
      });

      if (response.ok) {
        console.log('Message sent successfully');
      } else {
        console.log('Failed to send message');
      }
    } catch (error) {
      console.error('Error sending message:', error);
    }

    // Send the message through the WebSocket connection
    socket.send(JSON.stringify(newMessage));
  };


  // WebSocket connection initialization
  useEffect(() => {
    // Get the user ID from local storage
    const userID = localStorage.getItem('id');
    const newSocket = new WebSocket(`ws://localhost:8000/chat?senderId=${userID}`);
    console.log('WebSocket connection:', newSocket);

    newSocket.addEventListener('open', () => {
      console.log('WebSocket connection established');
    });

    newSocket.addEventListener('close', () => {
      console.log('WebSocket connection closed');
    });

    newSocket.addEventListener('error', (error) => {
      console.error('WebSocket connection error:', error);
    });

    // Handle incoming messages
    newSocket.addEventListener('message', (event) => {
      console.log('Received WebSocket message:', event.data);

      try {
        const message = JSON.parse(event.data);

        // Update the messages state with the received message
        setMessages(prevMessages => [...prevMessages, message]);
      } catch (error) {
        // Handle the case when the message is not valid JSON
        const newMessage = {
          id: Date.now().toString(),
          senderId: 'sender-id',
          recipientId: 'recipient-id',
          message: event.data,
          createdAt: new Date(),
        };

        // Log the message with sender ID and recipient ID
        console.log('Received plain text message:', newMessage);

        setMessages(prevMessages => [...prevMessages, newMessage]);
      }
    });

    // Set the socket state variable
    setSocket(newSocket);

    // Fetch chat history
    const fetchChatHistory = async () => {
      const senderId = userID // Replace with the appropriate sender ID
      const recipientID = selectedChat ? selectedChat.userId : null;
      try {
        const response = await axios.get(`http://localhost:8000/chat/history?senderId=${senderId}&recipientID=${recipientID}`);
        console.log(response);

        if (response.status === 200) {
          const chatHistory = response.data;
          setMessages(chatHistory);
        } else {
          console.log('Failed to fetch chat history');
        }
      } catch (error) {
        console.error('Error fetching chat history:', error);
      }
    };

    fetchChatHistory();

    // Cleanup function to close the WebSocket connection when the component unmounts
    return () => {
      newSocket.close();
    };
  }, []);

  if (!selectedChat) {
    return (
      <div className="flex items-center justify-center h-full">
        <p className="text-gray-500">Please select a chat to start messaging.</p>
      </div>
    );
  }

  return (
    <div className="flex flex-col h-full">
      {/* Upper bar */}
      <div className="flex items-center justify-between bg-black p-2">
        <div className="flex items-center">
          <img
            className="w-8 h-8 rounded-full mr-2"
            src={profile}
            alt="Profile Picture"
          />
          <div>
            <h2 className="font-bold text-white">{username}</h2>
            {/* <p className="text-gray-400 text-sm">Last seen: 8:30pm</p> */}
          </div>
        </div>
        <div className="flex items-center bg-[#1D1E24]">
          {/* Include your search bar component here */}
        </div>
      </div>

      {/* Main chat content */}
      <div className="flex-grow bg-[#1D1E24] flex flex-col-reverse overflow-y-scroll">
        <div className="p-4">
          {messages.map((message) => (
            <div
              key={message.id}
              className={`flex items-center mb-2 ${
                message.senderId === userID ? 'justify-end' : ''
                }`}
            >
              {message.senderId !== userID && (
                <img
                  className="w-8 h-8 rounded-full mr-2"
                  src={profile}
                  alt="User Avatar"
                />
              )}
              <div
                className={`${
                  message.senderId === userID
                    ? 'bg-gray-600 text-white'
                    : 'bg-gray-800 text-white'
                  } p-2 rounded-lg`}
              >
                <p className="text-sm">{message.message}</p>
              </div>
              {message.senderId === userID && (
                <img
                  className="w-8 h-8 rounded-full ml-2"
                  src={profile}
                  alt="User Avatar"
                />
              )}
            </div>
          ))}
        </div>
      </div>

      {/* Message input section */}
      <div className="flex items-center bg-[#1D1E24] p-2">
        <input
          type="text"
          placeholder="Type a message..."
          className="flex-grow rounded-xl py-3 px-4 bg-[#16171B] text-white focus:outline-none"
          value={inputValue}
          onChange={(e) => setInputValue(e.target.value)}
        />
        <div className="flex items-center ml-1">
          <button className="mr-2">
            <FiPaperclip className="text-white text-xl" />
          </button>
          <button className="mr-2">
            <FiSmile className="text-white text-xl" />
          </button>
          <button
            onClick={handleSendMessage}
            className="text-black bg-[#F3FC8A] py-2 px-2 rounded-3xl flex justify-center items-center"
          >
            <FiSend className="text-black text-xl" />
          </button>
        </div>
      </div>
    </div>
  );
};

export default ChatRoom;
