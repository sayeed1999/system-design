# 📘 **Chapter 2 Cheat Sheet – The WhatsApp → Dukaan MVP**

---

## **🚨 The Real-World Problem**

Small shops during lockdown were using **WhatsApp** to run their entire business:

* Blurry **PDF catalogs**
* Customers manually typing orders
* Sellers juggling dozens of chats
* Payment **screenshots**
* No structure, no tracking, constant mistakes

💡 **Insight:**
Small sellers don’t need Shopify.
They need something **simpler than WhatsApp**, built for commerce.

---

## **🛹 MVP Philosophy (Skateboard Mentality)**

### **MVP is not a small version of the final product.**

It’s a **test** of a hypothesis.

### **Hypothesis:**

“If we give sellers a 1-minute online store creator, will they use it?”

### **48-hour MVP contained only 3 features:**

1. **Create store** → phone number + OTP
2. **Add products** → name, price, single image
3. **Share link** → `mydukaan.io/mystore`

No:

* Payments
* Inventory
* Categories
* Analytics
* Dashboard

Just a skateboard to validate demand.

---

## **⚙️ Tech Stack (Speed-Optimized)**

### **Language → Python**

* Very fast to build
* Clean, readable
* Massive libraries

### **Framework → Django**

* “Batteries included”
* Provides:

  * Built-in Admin Panel
  * ORM
  * Auth
  * URL routing
  * Forms

Huge time-saver for an MVP.

### **Database → PostgreSQL**

* Strong, relational, reliable
* Great for structured commerce data
* Powerful future features (LISTEN/NOTIFY)

### **Final MVP Stack**

> **Python + Django + PostgreSQL**

---

## **🖥️ Server Setup (DigitalOcean $5 Droplet)**

### Why DO?

* Cheap
* Simple
* Reliable enough for MVP
* Fast provisioning

### What they installed:

* **Ubuntu** → base OS
* **Nginx** → web server (waiter)
* **Gunicorn** → app server (kitchen)
* **Django** → main application (chef)

**Flow:**
Client → Nginx → Gunicorn → Django → Postgres → Back to user

---

## **📌 Key Lessons**

* Build **fast**, not perfect.
* Don’t over-engineer an MVP.
* Choose a **framework that reduces decisions** (Django).
* You need **just enough** server to test your idea.
* Understand basic web infra:

  * **Nginx = handles requests**
  * **Gunicorn = runs Python app**
  * **Django = application logic**

---
