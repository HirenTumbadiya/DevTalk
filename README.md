# DevTalk
ChatApp is a real-time chat application built using React, Go, and MongoDB. It provides users with a platform to engage in instant messaging with other users, creating private or group chats.

# Features
User Registration and Authentication: Users can create an account and authenticate themselves to access the chat application.
Real-time Messaging: Users can send and receive messages in real-time, enabling instant communication.
Message Notifications: Users receive notifications for new messages, ensuring they never miss important conversations.
Message History: ChatApp stores message history, allowing users to view past conversations and scroll through previous messages.
User Profile: Users have personalized profiles where they can update their information and profile picture.
Online/Offline Status: Users can see the online/offline status of other users, providing visibility into their availability for communication.
Search Functionality: Users can search for other users to start new conversations or find existing chats.
Emoji Support: ChatApp supports a wide range of emojis to make conversations more expressive and engaging.

# Technologies Used
Front-end: React, HTML, CSS
Back-end: Go (Golang)
Database: MongoDB

# Getting Started
To run the ChatApp locally, follow these steps:

# Clone the repository:
git clone https://github.com/HirenTumbadiya/DevTalk.git

Install the required dependencies for the front-end and back-end:
For the front-end, navigate to the client directory and run:

cd client
npm install

For the back-end, navigate to the server directory and run:

cd server
go get ./...

Configure the database connection:
  Create a MongoDB database and obtain the connection URI.
  
Build and run the application:
For the front-end, within the client directory, run:

npm run build
For the back-end, within the server directory, run:

go run main.go
Access the ChatApp in your browser at http://localhost:3000.

# Acknowledgments
Special thanks to the authors and contributors of the libraries and frameworks used in this project for their invaluable work.

# Contact
If you have any questions or suggestions, feel free to reach out to us at tumbadiyahiren@gmail.com.
