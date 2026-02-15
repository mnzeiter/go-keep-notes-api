
# Go Keep Notes – Go + Gin Notes API

A lightweight and efficient REST API built with **Go 1.25**, **Gin**, and **MongoDB 7**.  
Designed to help users create, browse, update, and manage notes with a clean and simple structure.

---

## 🔧 Features

- Create, read, update, and delete notes (CRUD)
- MongoDB-backed persistent storage
- Clean project architecture using internal packages
- Fast and minimal API powered by Gin
- Dockerized for easy setup and deployment

---

## 🛠 Tech Stack

- **Go 1.25**
- **Gin Web Framework**
- **MongoDB 7**
- **Docker & Docker Compose**

---

## 🚀 Usage

1. Clone the repository:

   ```bash
   git clone <repo-url>
   ```

2. Start services using Docker:

   ```bash
   docker-compose up --build
   ```

3. Your API is now live at:

   ```
   http://localhost:8080
   ```

---

## 🧪 Endpoints

### ➕ Create Note  
**POST** `/notes`

```json
{
  "title": "Hello",
  "content": "World"
}
```

---

### 📄 Get All Notes  
**GET** `/notes`

---

### ✏️ Update Note  
**PUT** `/notes/<id>`


{
  "title": "Updated Title",
  "content": "Updated Content"
}

---

### 🗑️ Delete Note  
**DELETE** `/notes/<id>`

---

## 📷 API Test Screenshot

Add your API testing screenshot here:

```
assets/api-test.png
```

To display it:

```markdown
![API Test](assets/api-test.png)
```

---

## 🌐 Connect with Me
- 💼 [LinkedIn](https://linkedin.com/in/mozeiter)  
- 🌍 [Portfolio Website](https://mohammadalzeiter.com)  
- 📧 Email: mohammadalzeiter@outlook.com



## ✨ A clean and modern notes API built with Go, Gin, and MongoDB.
```

