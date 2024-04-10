# IM-cli-chat

This is a simple chat client side implemented in Go, using WebSockets for real-time communication. This is just the client side, and will not function properly without the server ([IM-cli-chat-server](https://github.com/ivermoka/IM-cli-chat-server/))

<img width="1438" alt="Screenshot 2024-04-10 at 12 52 13" src="https://github.com/ivermoka/IM-cli-chat-server/assets/119415554/7a8b5e7b-705d-47f9-9071-db1b1ec13821">


This project is primarily made for IM at Elvebakken Upper Secondary School, and hosted at a VM on our server. The downside of this, is that it can only be accessed to people on the same network as the VM, or by using a VPN. You could open a port to the public, but for security reasons, I chose to have it run locally on our network. 

This is the client side part of the IM-cli-chat project. The server side repo is [here](https://github.com/ivermoka/IM-cli-chat-server).

### System sketch (in norwegian)
![system drawio](https://github.com/ivermoka/IM-cli-chat-server/assets/119415554/fefe61a0-e38e-4693-812f-11bbcd76fa99)


### Running it on your own

#### Prerequisites
* Have [go](https://go.dev/) version **1.22.1** (may work on different versions) installed on your system. **(If you choose to run program using go)**
* Have [git](https://git-scm.com/) installed on your system **(if you choose to run program using go)**.
* OS cabable of running executables (MacOS, Linux(?)) **(if you choose to run program using the built executable in the repo)**

  
#### Running it

##### Running it yourself with go
* Clone this repo (```git clone https://github.com/ivermoka/IM-cli-chat-server```)

##### Running the executable **(may not work, depending on your OS)**
* Download the executable from this repo (enter file then click download)
* Run it locally (MacOS & Linux: ```./IM-cli-chat```)

### Usage

* This client side program is used to connect to the server ([IM-cli-chat-server](https://github.com/ivermoka/IM-cli-chat-server)
* Once connected, clients can send messages to the server, which will be broadcasted to all connected clients.


### Features

* WebSocket-based communication
* Real-time messaging
* CLI program (used in terminal)



### Dependencies
* [gocui](https://github.com/jroimartin/gocui): A GUI implementation for go.

### Contributing

Contributions are welcome! If you'd like to contribute to this project, feel free to open an issue or submit a pull request.
