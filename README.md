## 📝 Project
Minha Saúde app

## Description
💳 This application aims to offer a new disruptive way to check our personal health.

## Tech infos
* Language:
  - [Golang](https://golang.org/)  

### Endpoints

##### Login into API
`POST http://localhost:3001/user/login`

##### Prescriptions Upload
`POST http://localhost:3001/prescription/:user_uuid/upload`

##### Get users exams
`GET http://localhost:3001/user/:user-uuid/health-checks`

##### Download PDF
`GET http://localhost:3001/user/:user-uuid/health-checks/:uuid`

## ❗ Requirements
To run this application you have to install (if you don't have already installed) the follow programs:
* <b>In your computer</b>:
   * Docker 🐳 [click here](https://docs.docker.com/get-docker/)
<br>

## ▶️ Start application

#### Permissions first:  

* For <b>Unix</b> enviroment, run the comand:  
<b>```chmod +x .docker/entrypoint.sh```</b>  

* For <b>Windows</b> enviroment, run the comand:   
<b>```dos2unix +x .docker/entrypoint.sh```</b>  

### 💻 Start:
* Now, in your terminal, you can run:  <br>
<b>```docker-compose up```</b>

<br><br>
