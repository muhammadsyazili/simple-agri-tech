# Solutions & Running Instructions

This document provides instructions on how to run the solutions for each problem, along with the specific SQL queries requested for Problem 3.

**Prerequisite:**
Ensure you have `make` installed and PostgreSQL running.

## Problem 1: Polycarp (CLI)
Polycarp doesn't like integers that are divisible by 3 or end with the digit 3 in their decimal representation. Integers that meet both conditions are disliked by Polycarp, too.

Polycarp starts to write out the positive (greater than 0) integers which he likes: 1, 2, 4, 5, 7, 8, 10, 11, 14, 16, ... . Output the $k$-th element of this sequence (the elements are numbered from 1).

**Input**

The first line contains one integer $t$ ($1 \le t \le 100$) — the number of test cases. Then $t$ test cases follow.Each test case consists of one line containing one integer $k$ ($1 \le k \le 1000$).

**Output**

For each test case, output in a separate line one integer $x$ — the $k$-th element of the sequence that was written out by Polycarp.

**How to Run:**
```bash
make run-p1
```
Follow the prompts to enter the number of test cases and the value `k` for each case.

---

## Problem 2: Palindrome (CLI)
A CLI application that checks if a given string is a palindrome.

**Reference:**
[CODING INTERVIEW PALINDROME](https://www.youtube.com/watch?v=DXQuiPKl79Y&t=1812s)

**How to Run:**
```bash
make run-p2
```
Enter a string when prompted.

---

## Problem 3: User Spending (API)

### A. Munculkan data country mana aja yang spend nya terbanyak
*Raw SQL Query:*
```sql
SELECT 
    u.country,
    SUM(s.total_buy) as total_spend
FROM 
    users u
JOIN 
    spendings s ON u.id = s.user_id
GROUP BY 
    u.country
ORDER BY 
    total_spend DESC
LIMIT 1;
```

### B. Munculkan data jumlah tipe kartu kredit terbanyak
*Raw SQL Query:*
```sql
SELECT 
    credit_card_type,
    COUNT(*) as total_count
FROM 
    users
GROUP BY 
    credit_card_type
ORDER BY 
    total_count DESC
LIMIT 1;
```

### C & D. Endpoints
An API server implemented with Gin and GORM (PostgreSQL).

**Configuration (Optional):**
By default, the application connects to:
`host=localhost user=postgres password=postgres dbname=postgres port=5432`

To customize this, copy `.env.example` to `.env` and adjust the values:
```bash
cp .env.example .env
```

**Setup (Run only once):**
1. **Migration (Fresh):** Drops existing tables and re-creates them.
   ```bash
   make migrate-fresh
   ```
2. **Seeding:** Populates the database with 1000 users and spendings.
   ```bash
   make seed
   ```

**How to Run API Server:**
```bash
make run-api
```
The server will start on port `8080`.

**Endpoints:**
- `GET /top-spending-users`: Returns users from the country with the highest total spending.
  ```bash
  curl -X GET http://localhost:8080/top-spending-users
  ```
- `POST /users`: Creates a new user.
  ```bash
  curl -X POST http://localhost:8080/users \
  -H "Content-Type: application/json" \
  -d '{
      "first_name": "John",
      "last_name": "Doe",
      "country": "Indonesia",
      "credit_card_type": "visa-electron",
      "credit_card": "444444444444"
  }'
  ```

---

## Problem 4: Sorting (CLI)
A CLI application that generates random numbers and sorts them using the **Selection Sort** algorithm.

**Reference:**
[#4 LOGICAL CONCEPT & HOW SELECTION SORTING WORKS | PROGRAMMING ALGORITHM 2](https://www.youtube.com/watch?v=s-IDU_zAefA&t=319s)

**How to Run:**
```bash
make run-p4
```
Enter the number of elements to generate when prompted.
