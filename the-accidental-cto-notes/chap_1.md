# 📘 **Chapter 1 Cheat Sheet – Server Crash & Monolith Basics**

---

## **🚨 The Crash**

* 3:14 AM: Everything down.
* `ssh` slow → deep system issue.
* `htop` shows:

  * **CPU 100%**
  * **RAM full**
  * **Swap full**
* Root cause: **Server had only 512MB RAM** → completely overloaded.

---

## **🧠 Server Anatomy (Restaurant Analogy)**

### **CPU = Chef**

* Executes tasks.
* More cores → more parallel work.
* 1-core chef → overwhelmed easily.

### **RAM = Countertop**

* Workspace for active tasks.
* Too small = constant trips to pantry → huge slowdown.
* Full RAM triggers **swap** → extremely slow.

### **Disk = Pantry**

* Large but slow storage.
* Not meant for active work.

### **Resource Contention**

* App, DB, OS all fighting for CPU + RAM.
* Leads to total freeze/crash.

### **Diagnosis Tools**

* `ssh` → access server.
* `htop` →

  * CPU bars → chef load.
  * Mem → countertop usage.
  * Swap → emergency usage (bad sign).
  * Process list → resource hogs.

---

## **🏗️ Monolith Architecture**

### **What it is**

* Entire app (all features) in one codebase.
* “One giant cookbook with every recipe.”

### **Why Start with a Monolith**

* Fastest to build.
* Easy to test.
* Easy to deploy.
* Ideal for small team + early product validation.

### **Monolith Problems**

* Tight coupling → one change affects everything.
* One bug can break entire system.
* Hard to onboard devs.
* Can’t scale individual components.
* High-traffic parts overload the whole app.

---

## **💡 Key Lessons**

* First server WILL fail — learn from it.
* Understand CPU, RAM, Disk deeply.
* Learn `ssh` + `htop` early — essential tools.
* Monolith is good early, bad later — architecture evolves with scale.

---
