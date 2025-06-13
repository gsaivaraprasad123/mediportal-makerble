
# 🏥 MediPortal

MediPortal is a scalable web app for managing patients and doctors with secure role-based access. Built using Go (Gin), PostgreSQL, and Docker, it’s ideal for hospitals or clinics needing a minimal system for patient record management.

---

## 📦 Tech Stack

- **Backend**: Golang + Gin
- **Database**: PostgreSQL
- **ORM**: GORM
- **Auth**: JWT
- **DevOps**: Docker + Docker Compose

---

## 🌐 Live Deployment (Optional)

You can pull the Docker image directly:

```bash
docker pull saivaraprasad03/makerble-app:latest
````

Then run:

```bash
docker run -p 8080:8080 \
  -e DB_URL=postgres://postgres:<password>@host.docker.internal:5432/mediportal?sslmode=disable \
  -e JWT_SECRET=your_jwt_secret \
  saivaraprasad03/makerble-app:latest
```

---

## 📁 Project Structure

```
mediportal/
├── cmd/                # App entry point (main.go)
│   └── main.go
├── config/             # env loading, DB config
├── controllers/        # HTTP handlers
├── routes/             # All route group definitions
├── middleware/         # JWT Auth, logging, recovery
├── models/             # Structs and GORM models
├── repository/         # DB logic (repository pattern)
├── services/           # Business logic (use-cases)
├── utils/              # Common utilities
├── docs/               # Swagger / Postman JSON
├── test/               # Unit and integration tests
├── Dockerfile
├── docker-compose.yml
├── go.mod / go.sum
├── .env
└── README.md

```

---

## ⚙️ Environment Setup

### 📄 1. Create a `.env` File

Create a file named `.env` at the root of the project:

```env
PORT=8080
DB_URL=postgres://postgres:postgres@db:5432/mediportal?sslmode=disable
JWT_SECRET=supersecretkey
```

You don’t need to load `.env` manually — it is auto-used by the app inside Docker via `godotenv`.

---

### 🐳 2. Run with Docker Compose

Run the app locally using Docker Compose:

```bash
docker-compose up --build
```

The API will be available at `http://localhost:8080`

---

### 🐘 3. PostgreSQL Admin Tool (Optional)

Access Postgres via:

* **TablePlus**, **pgAdmin**, or CLI
* Default Credentials:

```text
Host: localhost
Port: 5432
User: postgres
Password: postgres
Database: mediportal
```

---

## 🧪 API Documentation

### 🔐 Auth Routes

#### ✅ Register

`POST /api/signup`

```json
{
  "email": "recep3@test.com",
  "password": "recep3",
  "role": "receptionist"
}
```

#### 🔓 Login

`POST /api/login`

```json
{
  "email": "doctor2@test.com",
  "password": "doctor2"
}
```

---

### 👤 Receptionist Routes

#### ➕ Create Patient

`POST /api/patients`

```json
{
  "name": "Hritik",
  "age": 56,
  "gender": "Male",
  "condition": "AIDS"
}
```

#### 📄 Get All Patients

`GET /api/patients`

#### ✏️ Update Patient

`PUT /api/patients/:id`

```json
{
  "name": "Hritik",
  "age": 44,
  "gender": "Male",
  "condition": "AIDS"
}
```

#### ❌ Delete Patient

`DELETE /api/patients/:id`

---

### 🩺 Doctor Routes

#### 👁 View All Patients

`GET /api/doctor/patients`

#### 📝 Update Patient Condition

`PUT /api/doctor/patients/:id`

```json
{
  "condition": "Recovered from Covid-19 using CoVaxin"
}
```

---

## 🔁 Postman Collection (API Tester)

You can import the complete API JSON collection into [Postman](https://www.postman.com/) using:

```
[🔗 Postman Collection](https://universal-shuttle-500501.postman.co/workspace/Postman-API-Fundamentals-Studen~6ab8db95-d6ab-43eb-a649-ee2b00c8258d/collection/31896790-decfd559-fc5e-40e7-ab8f-eb0fce431a1f?action=share&creator=31896790)
```

Replace the token after login in the `Authorization` tab for protected routes.

---

## 🧾 Sample `.env` for Local Dev

```env
PORT=8080
DB_URL=postgres://postgres:postgres@db:5432/mediportal?sslmode=disable
JWT_SECRET=myjwtsecret
```

No manual loading is needed — it's used inside Docker.

---

## 👨‍💻 For Contributors

1. Fork the repo
2. Clone locally and create a `.env`
3. Run with `docker-compose up`
4. Create a new branch for feature/bug
5. Submit a pull request!

---

## 📽 Demo Video

[🔗 Watch Demo](https://example.com/video-demo) 

---

## 📄 License

This project is under [MIT License](LICENSE).

---

## 🙋‍♂️ Author

Maintained by **Sai Vara Prasad**
[GitHub](https://github.com/gsaivaraprasad123) • [LinkedIn](https://www.linkedin.com/in/sai-vara-prasad-gandhe-83056a278/)

```

---

```
