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
| GET             | /v1/clipboard      | getClipContent| Gets content from clipboard               |
| POST            | /v1/setclip        | setClipText   | Sets content of clipboard                 |
| GET             | /v1/view_all_paste | getAllPaste   | Gets all pastes of user                   |
| GET             | /v1/getpaste/:id   | getPaste      | Gets specified paste                      |
| POST            | /v1/newpaste       | addNewPaste   | Writes new paste                          |
| SFTP            | /v1/send           | sendFile      | Prepares to send file to listening server |
| SFTP            | /v1/receive        | receiveFile   | Receives file from sftp port              |
| HLS & MPEG-DASH | /v1/stream_client  | streamClient  | Send stream data                          |
| HLS & MPEG-DASH | /v1/stream_server  | streamServer  | Receive stream data                       |

## Tech Stack

1. FrontEnd: Vue js
2. BackEnd: Golang


## Todo

#### User authentication
- [x] Make login and logout basic functioning
- [x] Create database structure and design
- [x] Implement basic database design
- [ ] Create proper JSON request and response structures and make proper JSON returns
    - [x] Define request and response internals
    - [x] Create proper error functions
- [x] Create validator and error function failedValidation
- [x] Change all instance of app.badRequestError with appropriate stuff
- [x] trace should only be shown if a flag is enabled
- [x] integrate context if required
- [x] change db structure to keep UserId as primary key
- [x] in protection middleware, keep UserId in context
- [x] Integrate database with backend
- [x] Modify userAuth middleware
- [ ] Add indexing
- [ ] Add bool to exists variable in authentication
- [ ] Clean invalid sessions at the end of the day

doubts:
    - How will you check if the error from DB is duplciated username or something else?

#### Clipboard
- [x] Implement getClipboard() handler to return clip context as JSON
- [x] Implement setClipText() handler to set clipboard context
- [ ] Timer to clean clips at the end of the day

#### Paste
- [x] Issue: check if paste can be edited even after deletion
- [ ] Change timing to local time. Currently in UTC

#### FTP

#### Streaming