# GoMinder
Gominder is a simple command-line reminder tool written in Go. It allows users to set reminders by specifying a time and a message, and it will display a notification when the specified time is reached.

## Usage
    ```bash
    gominder <time> <message>

`<time>`: The time at which the reminder should be triggered.
`<message>`: The message or description for the reminder.

## Installation
    ```bash
    go install

Ensure that the Go bin directory is in your system's PATH environment variable to run gominder globally.

## Dependencies
Gominder relies on the following third-party packages:

gen2brain/beeep: A cross-platform package for displaying desktop notifications.
olebedev/when: A natural language date/time parser with pluggable rules.

## How it works
Gominder uses the when package for parsing the provided time and scheduling the reminder. It also utilizes the beeep package to display desktop notifications.

When a reminder is set, Gominder checks for a specific environment variable to avoid running multiple instances of the same reminder.

## Contributing
Feel free to contribute to Gominder by opening issues or submitting pull requests. Bug reports, feature requests, and feedback are all welcome.

## License
This project is licensed under the MIT License - see the LICENSE file for details.
