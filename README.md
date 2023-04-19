# Final Project

Membuat simple project yang bernama MyGram. Pada Mygram ini Terdapat 4 model yaitu :

1. Users
2. Photos
3. Comment
4. Socialmedia

### Dokumentasi API

Project ini berjalan secara local host dengan port local host :

```bash
localhost:8000
```

Dengan Endpoint :

- Users
  | Method | Endpoint | Description |
  |--------|----------|-------------|
  | POST | /login | Login |
  | POST | /register | Register |

- Photos
- (Authentication)
  | Method | Endpoint | Description |
  |--------|----------|-------------|
  | GET | /photos | Get all photos |
  | GET | /photos/:id | Get photos by id |
  | POST | /photos | Create photos book |
  | PUT | /photos/:id | Update photos by id |
  | DELETE | /photos/:id | Delete photos by id |

- Comment
- (Authentication)
  | Method | Endpoint | Description |
  |--------|----------|-------------|
  | GET | /comments | Get all comment |
  | GET | /comments/:id | Get comment by photo id |
  | POST | /comments | Create comment |
  | PUT | /comments/:id | Update comment by id |
  | DELETE | /comments/:id | Delete comment by id |

- Socialmedia
  (Authentication)
  | Method | Endpoint | Description |
  |--------|----------|-------------|
  | GET | /socialmedia | Get all socialmedia|
  | GET | /socialmedia/:id | Get socialmedia by id |
  | POST | /socialmedia | Create socialmedia |
  | PUT | /socialmedia/:id | Update socialmedia by id |
  | DELETE | /socialmedia/:id | Delete socialmedia by id |

### Running Project
- Jalankan file yang bernama `main.go`
- import file yang bernama `my-gram.postman_collection` pada Postman
