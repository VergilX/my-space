# localDump

This is a project which allows me to dump files in my local network. Being local, I'm hoping that it wouldn't be susceptible to network attacks. Otherwise I would have to study security first.

## Features

- local clipboard (one-time-use)
- pastebin like service
- sftp system
- streaming
- cloud (optional)


## API Endpoints

| **Method**      | **URL**            | **Handler**   | **Action**                                |
|-----------------|--------------------|---------------|-------------------------------------------|
| GET             | /v1/status         | checkStatus   | Checks status of website                  |
| POST            | /v1/register       | registerUser  | Registers user into local network         |
| GET             | /v1/login          | loginUser     | Logs in using credentials                 |
| GET             | /v1/logout         | logoutUser    | Logs out using credentials                |
| GET             | /v1/clipboard      | getClipboard  | Gets content from clipboard               |
| GET             | /v1/view_all_paste | getAllPaste   | Gets all pastes of user                   |
| GET             | /v1/getpaste/:id   | getPaste      | Gets specified paste                      |
| POST            | /v1/newpaste       | writeNewPaste | Writes new paste                          |
| SFTP            | /v1/send           | sendFile      | Prepares to send file to listening server |
| SFTP            | /v1/receive        | receiveFile   | Receives file from sftp port              |
| HLS & MPEG-DASH | /v1/stream_client  | streamClient  | Send stream data                          |
| HLS & MPEG-DASH | /v1/stream_server  | streamServer  | Receive stream data                       |

## Tech Stack

1. FrontEnd: Vue js
2. BackEnd: Golang


## Todo

- [x] Make login and logout basic functioning
- [ ] Modify userAuth middleware
- [ ] Modify errors to make proper handlers
- [ ] Redirection after errors
- [ ] Create proper JSON request and response structures and make proper JSON returns
- [ ] (if required) Create validator package and modify funcs to use it