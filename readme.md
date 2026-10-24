# 🐜 Lem-In — Ant Farm Pathfinding

A Go implementation of an ant farm simulation that finds efficient routes through a network of rooms and tunnels, then coordinates ant movements to bring all ants from the start room to the end room in as few turns as possible.

## 📌 Overview

Lem-In models a colony as a graph:

- **Rooms** are the nodes of the graph.
- **Tunnels** are the connections between rooms.
- **Ants** move through the colony toward the exit.
- **Start and end rooms** define the beginning and destination of the journey.

The challenge is not only to find a path, but also to coordinate multiple ants while avoiding occupied rooms and respecting tunnel and movement restrictions.

## ✨ Features

- Parse colony descriptions from an input file.
- Identify the start room, end room, rooms, and tunnels.
- Validate colony data and handle invalid input.
- Find paths between the start and end rooms.
- Coordinate ant movements across available paths.
- Respect room occupancy and movement constraints.
- Print the original colony description followed by each turn of ant movements.
- Use Go standard library packages only, as required by the project.

## 🛠️ Technologies

- **Language:** Go
- **Core concepts:** Graphs, pathfinding, parsing, validation, scheduling, and simulation
- **Dependencies:** Go standard library only

## 🚀 Getting Started

### Prerequisites

Install Go on your machine.

Check that Go is available:

```bash
go version
```

### Clone the repository

```bash
git clone <YOUR_REPOSITORY_URL>
cd lem-in
```

Replace `<YOUR_REPOSITORY_URL>` with your actual repository URL.

### Run the program

Pass a colony description file as an argument:

```bash
go run . test0.txt
```

You can also build an executable:

```bash
go build -o lem-in .
./lem-in test0.txt
```

Make sure the input file exists in the current directory, or provide its correct path.

## 📄 Input Format

A colony file contains the number of ants, room definitions, and tunnel connections.

### Example

```text
3
##start
start 0 0
room1 1 0
room2 2 0
##end
end 3 0
start-room1
room1-room2
room2-end
```

### Input rules

- The first line specifies the number of ants.
- Each room has a name and two integer coordinates.
- `##start` identifies the starting room.
- `##end` identifies the destination room.
- A tunnel connects two existing rooms using the format `room1-room2`.
- Room names cannot start with `L` or `#`.
- Duplicate rooms, invalid links, and malformed input must be handled safely.

## 📤 Output Format

The program prints the colony description, followed by the ants that move during each turn.

Example movement output:

```text
L1-room1
L1-room2 L2-room1
L1-end L2-room2 L3-room1
```

`L1` represents ant number 1, and the name after the hyphen identifies its destination room.

Only ants that move during a turn should be printed.

## 🧠 Algorithmic Challenges

The project focuses on several important computer science concepts:

- **Graph traversal:** Explore rooms and tunnels to find routes.
- **Pathfinding:** Identify routes from the start to the end.
- **Path optimization:** Reduce the number of turns needed to move all ants.
- **Movement scheduling:** Coordinate ants so they do not enter occupied intermediate rooms.
- **Input validation:** Detect malformed data and invalid colony structures.
- **Simulation:** Process ant movements turn by turn until all ants reach the exit.

Finding one shortest path is not always enough. Using multiple suitable paths can help ants reach the exit faster.

## 🧪 Testing

Run the project's available tests:

```bash
go test ./...
```

If test cases are provided as input files, run the program against them:

```bash
go run . test0.txt
go run . test1.txt
```

Check both valid colonies and invalid inputs, including missing start or end rooms, unknown tunnel endpoints, duplicate rooms, and colonies with no route to the exit.

## 📁 Project Structure

The exact structure depends on the implementation.

```text
lem-in/
├── README.md
├── go.mod
├── main.go
├── test0.txt
└── ...
```

Update this tree to match the actual files in your repository.

## 🎯 Learning Outcomes

This project provides practical experience with:

- Graph algorithms and pathfinding
- Parsing and validating structured input
- String manipulation
- Data structures and program design
- Simulation and scheduling
- Writing reliable Go programs
- Testing and debugging

## 🌟 Possible Bonus

A visualizer can display the colony and animate ants moving through rooms and tunnels.

The room coordinates can be used to position rooms in the visualization.

## 👤 Author

- GitHub: [Your GitHub Profile](https://github.com/ussefbouchraa)

