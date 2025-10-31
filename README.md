# 🔐 Secure Login System (Go)

A simple command-line **Secure Login System** built with **Go (Golang)**.  
This project allows you to **register users, log in, view all registered users (admin-only), save data to JSON, and load it back**.  
Perfect for learning file handling, structs, and basic authentication logic in Go.

---

## 📋 Features

- ✅ **Register** new users  
- 🔑 **Login** with username and password  
- 👁️ **Show users** (admin password required)  
- 💾 **Save data** to `users.json`  
- 📂 **Load data** from `users.json`  
- 🚪 **Exit** safely  

---

## 🧠 How It Works

The program runs in the terminal and waits for your input.  
It provides a simple text-based menu with these commands:

```
bash
1. register
2. login
3. show user
4. save
5. load
6. exit

---
```
bash
Enter command: register
Enter user name: John
Enter password: 12345
✅ User John registered successfully!

---

```
bash
Enter command: login
Enter user name: John
Enter password: 12345
✅ Login successful! Welcome, John!

---

```
bash
Example Usage
Welcome to Secure Login System!
---------------------------------
1. register
2. login
3. show user
4. save
5. load
6. exit
---------------------------------
Enter command: register
Enter user name: Alice
Enter password: hello123
✅ User Alice registered successfully!



