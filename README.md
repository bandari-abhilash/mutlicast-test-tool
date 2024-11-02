# Multicast Communication Tool

## Overview

The Multicast Communication Tool is a Go application that allows users to send and receive multicast messages over a network. It provides a simple command-line interface for selecting network interfaces and configuring multicast settings.

## Features

- List available network interfaces
- Send multicast messages
- Receive multicast messages
- Handle user input for interface selection and multicast configuration
- Graceful recovery from panics

## Requirements

- Go 1.16 or later
- A network interface that supports multicast

## Installation

1. Clone the repository:
   ```bash
   git clone https://github.com/yourusername/multicast-communication-tool.git
   cd multicast-communication-tool
   ```

2. Build the application:
   ```bash
   go build -o multicast-tool main.go
   ```

3. Run the application:
   ```bash
   ./multicast-tool
   ```

## Usage

1. **Select a Network Interface**: The application will prompt you to select a network interface from the list of available interfaces on your machine.

2. **Choose an Operation**: After selecting an interface, you can choose to either send or receive multicast messages:
   - **1**: Multicast Sender (Will initiate multicast data)
   - **2**: Multicast Receiver (Will receive multicast data)
   - **3**: Exit

3. **Configure Multicast Settings**: If you choose to send or receive messages, you will be prompted to enter:
   - Multicast address (between `224.0.0.0` to `239.255.255.255`)
   - Port number (between `1` and `6445`)

4. **Sending Messages**: If you select the sender option, the application will continuously send a test message to the specified multicast address.

5. **Receiving Messages**: If you select the receiver option, the application will listen for incoming multicast messages and print them to the console.


## Error Handling

The application includes basic error handling for common issues, such as invalid multicast addresses and network errors. If an error occurs, the application will provide feedback and prompt the user to try again.

## Contributing

Contributions are welcome! If you have suggestions for improvements or new features, please open an issue or submit a pull request.

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.

## Acknowledgments

- Thanks to the Go community for their support and resources.