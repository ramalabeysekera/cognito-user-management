# Cognito User Management CLI

## Overview
The Cognito User Management CLI is a command-line tool designed to simplify the management of AWS Cognito user pools. This tool provides an intuitive interface for developers and administrators to perform various operations on Cognito users and user pools, such as creating, updating, deleting, and listing users, as well as managing user attributes and groups.

## Features
- **User Pool Selection**: Interactively select a Cognito User Pool from your AWS account.
- **User Creation**: Create new users in a Cognito User Pool with options for temporary or permanent passwords.
- **Bulk User Creation**: Import users from a CSV file and create them in bulk.
- **AWS Profile Selection**: Choose an AWS profile from your local configuration for authentication.
- **Interactive CLI**: User-friendly prompts for seamless interaction.

## Prerequisites
- Go 1.23 or later
- AWS credentials configured in your local environment
- IAM permissions to manage Cognito user pools

## Installation
1. Clone the repository:
    ```bash
    git clone https://github.com/ramalabeysekera/cognitousermanagement.git
    ```
2. Navigate to the project directory:
    ```bash
    cd cognitousermanagement
    ```
3. Build the executable:
    ```bash
    go build -o cognitousermanagement.exe
    ```

## Usage
Run the CLI tool using the following command:
```bash
./cognitousermanagement.exe
```

### Commands
#### `createuser`
Create a new user in a Cognito User Pool.

**Options:**
- `--permanentpassword`: Set the password as permanent during user creation.
- `--bulk`: Create multiple users from a CSV file.

**Example:**
```bash
./cognitousermanagement.exe createuser --permanentpassword=true
```

#### `root`
The root command provides an overview of the tool and its functionalities.

## File Structure
- `cmd/`: Contains the CLI command definitions.
- `config/`: Handles configuration loading.
- `pkg/common/`: Implements core functionalities like user creation and password management.
- `pkg/helpers/`: Provides utility functions for interactive prompts and CSV handling.
- `pkg/selections/`: Manages user pool selection logic.

## CSV File Format
For bulk user creation, the CSV file should have the following format:
```
username,password
```
Example:
```
john_doe,P@ssw0rd123
alice_smith,Secur3P@ss!
```

## Contributing
Found a bug or have a feature request? Please open an issue on GitHub:
https://github.com/ramalabeysekera/cognitousermanagement/issues

## License
This project is unlicensed. See the LICENSE file for details.

## Author
Ramal Abeysekera