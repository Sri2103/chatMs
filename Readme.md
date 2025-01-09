# Chat Application Architecture

This repository contains the architecture and design of a chat application built using microservices. The architecture is divided into several components, services, and use cases that define the interactions within the system.

## Table of Contents

- [Architecture Overview](#architecture-overview)
- [Components](#components)
- [Use Cases](#use-cases)
- [Getting Started](#getting-started)

## Architecture Overview

The architecture is structured into three main layers:

1. **Frontend**: The user interface built using React and Flutter.
2. **Backend**: Microservices developed in Go.
3. **Database Layer**: Data storage using MongoDB and Postgres.

### Top-Level Architecture

![Top-Level Architecture](path/to/toplevel.png)

The top-level architecture includes the following services:

- **API Gateway**: Forwards requests to the appropriate services.
- **WebSocket Server**: Manages real-time communication.
- **User Service**: Handles user authentication and management.
- **Room Service**: Manages chat rooms.
- **Chat Service**: Handles message sending and retrieval.
- **Notification Service**: Sends notifications to users.
- **Status Service**: Manages user online status.

### Components

![Components](path/to/components.png)

The components of the chat application include:

- **Frontend (React/Flutter)**:
  - Web Chat UI
  - Mobile Chat UI

- **Backend (Go Microservices)**:
  - Auth Service
  - Chat Service
  - Room Management
  - Notification Service
  - Reporting Service
  - Audit Log Service
  - WebSocket Gateway
  - REST API Gateway

- **Database Layer (MongoDB/Postgres)**:
  - User DB
  - Message DB
  - Room DB
  - Audit DB
  - Report DB

## Use Cases

![Use Cases](path/to/usecase.png)

The primary use cases for the chat application include:

- Register Account
- Login
- Send Message
- Receive Message
- Join Room
- Leave Room
- Track Online Status
- Typing Indicator
- View Message History
- Create Sub-Room
- Manage Room Hierarchy
- View Room Reports
- View User Reports
- Audit Admin Actions
- Generate Usage Reports
- Invite User to VIP Room
- Restrict Sub-Room Access

## Getting Started

To get started with the chat application, clone the repository and follow the instructions in the respective service directories.

```bash
git clone https://github.com/Sri2103/chatMs.git
cd chatMs

