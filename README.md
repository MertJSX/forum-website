# Forum website
## Project Overview
A full-stack forum website designed for educational discussions, built with modern web technologies:
- Frontend: Vite + React + TypeScript (Fast, type-safe client application)
- Backend: Golang with GoFiber (High-performance web framework)
- Database: SQLite (Lightweight, file-based database)

## How to setup
- Clone the project
- Run ```make setup``` in /server
- Run ```npm install``` in /client
- Start the client server with the command ```npm run dev```
- Start the backend server with the command ```make start```
- (optional) You can make tests using the command ```make test```

## Key Features
### User System
- User registration and authentication
- JWT (JSON Web Token) based login system
- User profiles with activity tracking
- Follow system to connect with other users
### Forum Functionality
- Create, edit, and delete posts
- Categorize discussions by topics
- Comment system for threaded discussions
- Upvote system for quality content curation

## Technical highlights
- Type-safe codebase with TypeScript and Go
- RESTful API design
- Responsive UI components
- Efficient state management
- Secure authentication flow

## Database tables
### users
stores user accounts
### posts
stores user posts
### comments
stores user comments
### post_upvotes
manages post voting system
### followers
manages follower system
