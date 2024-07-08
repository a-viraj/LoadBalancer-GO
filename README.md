
# Go Load Balancer

 This project demonstrates a simple load balancer implemented in Go that distributes requests to multiple backend servers. Each server responds with a unique message.

# Project Structure

```
project-root/
├── loadbalancer/
│ ├── main.go
├── server1/
│ ├── server1.go
├── server2/
│ ├── server2.go
└── server3/
├── server3.go

```
## Prerequisites

- [Go](https://golang.org/dl/) (version 1.19 or later)

## Setup and Run

> Follow the steps below to build and run the load balancer and servers.

### Load Balancer

1. **Navigate to the load-balancer directory**

    ```sh
    cd loadbalancer
    ```
2. **Build the load balancer** 
    ```sh
    go build -o loadbalancer main.go
    ```
3. **Run the load Balancer**
    ```sh
    ./loadbalancer
    ```
### Servers

> Repeat the following steps for each server (server1, server2, and server3):

1. **Navigate the server directory**
    ```sh
    cd server1
    ```
    > Change the server to server2 and server3 for the respective servers
2. **Build the servers** 
    ```sh
    go build -o server1 server1.go
    ```
3. **Run the load Balancer**
    ```sh
    ./server1
    ```  



## Access the load-balancer

> Once the load balancer and servers are up and running, you can access the load balancer at `http://localhost:8080`. The load balancer will distribute incoming requests to the backend servers in a round-robin fashion.

### Services
* Load Balancer

   * Port: 8000
* Server 1

    * Port: 8001
* Server 2

    * Port: 8002
* Server 3

    * Port: 8003

## Troubleshooting
> If you encounter issues with the setup, here are some steps to debug:

1. Check Logs

    Examine the logs of each application to look for errors.

2. Test Connectivity

    Use tools like curl or telnet to test connectivity to the servers from your local machine.

3. Network Configuration

    Ensure that no other services are running on the same ports (8000, 8001, 8002, 8003).

4. Firewall Rules

    Verify that your local firewall rules are not blocking the necessary ports.
## Author
* a-viraj
