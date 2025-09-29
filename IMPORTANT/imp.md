# 🔹 GORM `DB` Methods

## 1. **Create**

```go
result := config.DB.Create(&user)
```

* **Input**: Pointer to a struct, slice of structs, or map.
* **Output**: `*gorm.DB` (chainable query object).
* **Usage**: Inserts new record(s) into the database.
* **Check result**:

  * `result.Error` → error if any
  * `result.RowsAffected` → number of inserted rows

---

## 2. **Find**

```go
config.DB.Find(&users)
```

* **Input**: Pointer to slice of structs (`&[]User{}`) or pointer to a struct.
* **Output**: `*gorm.DB`
* **Usage**: Loads all records into given slice.

---

## 3. **First**

```go
config.DB.First(&user, id)
```

* **Input**:

  * `&user` → pointer to struct
  * `id` → primary key OR conditions (`map`, `struct`, `string query`, etc.)
* **Output**: `*gorm.DB`
* **Usage**: Gets the **first matching row**.
* Example:

  ```go
  config.DB.First(&user, 1)             // find by ID
  config.DB.First(&user, "email = ?", "alice@example.com")
  ```

---

## 4. **Take**

```go
config.DB.Take(&user)
```

* **Input**: Pointer to struct + conditions.
* **Output**: `*gorm.DB`
* **Usage**: Like `First` but doesn’t order by primary key.

---

## 5. **Last**

```go
config.DB.Last(&user)
```

* Same as `First` but fetches the **last record** (by primary key order).

---

## 6. **Where**

```go
config.DB.Where("name = ?", "Alice").Find(&users)
```

* **Input**: Condition string with placeholders, or map/struct.
* **Output**: `*gorm.DB` (chainable).
* **Usage**: Adds filtering to queries.

---

## 7. **Updates**

```go
config.DB.Model(&user).Updates(models.User{Name: "Bob", Email: "bob@example.com"})
```

* **Input**:

  * Struct (non-zero fields only)
  * Map (`map[string]interface{}`)
* **Output**: `*gorm.DB`
* **Usage**: Updates multiple fields at once.

---

## 8. **Update**

```go
config.DB.Model(&user).Update("name", "Charlie")
```

* **Input**: Field name (string), new value (any type).
* **Output**: `*gorm.DB`
* **Usage**: Updates a **single field**.

---

## 9. **Delete**

```go
config.DB.Delete(&user)
```

* **Input**: Pointer to struct OR model type + condition.
* **Output**: `*gorm.DB`
* **Usage**: Deletes record(s).
* Example:

  ```go
  config.DB.Delete(&User{}, 1)               // delete by ID
  config.DB.Delete(&User{}, "email = ?", "x@example.com")
  ```

---

## 10. **Save**

```go
config.DB.Save(&user)
```

* **Input**: Pointer to struct.
* **Output**: `*gorm.DB`
* **Usage**: Updates the entire record (if ID exists) or inserts if not.

---

## 11. **Model**

```go
config.DB.Model(&user).Update("name", "Zed")
```

* **Input**: Model struct or pointer.
* **Output**: `*gorm.DB`
* **Usage**: Specifies the table/model to work with before chaining.

---

## 12. **Count**

```go
var count int64
config.DB.Model(&User{}).Where("email LIKE ?", "%@gmail.com").Count(&count)
```

* **Input**: Pointer to `int64`
* **Output**: `*gorm.DB`
* **Usage**: Count rows matching query.

---

## 13. **Raw**

```go
config.DB.Raw("SELECT * FROM users WHERE name = ?", "Alice").Scan(&user)
```

* **Input**: Raw SQL query string, variables.
* **Output**: `*gorm.DB`
* **Usage**: Execute raw SQL, load into struct.

---

## 14. **Exec**

```go
config.DB.Exec("UPDATE users SET name = ? WHERE id = ?", "Alice", 1)
```

* **Input**: SQL string + arguments.
* **Output**: `*gorm.DB`
* **Usage**: Run raw SQL **without returning rows** (INSERT, UPDATE, DELETE).

---

# 📌 Quick Table

| Method    | Input(s)                    | Return    | Purpose                   |
| --------- | --------------------------- | --------- | ------------------------- |
| `Create`  | struct/slice                | \*gorm.DB | Insert new records        |
| `Find`    | slice/struct                | \*gorm.DB | Fetch all matches         |
| `First`   | struct + id/condition       | \*gorm.DB | Fetch first row           |
| `Take`    | struct + condition          | \*gorm.DB | Fetch one row (no order)  |
| `Last`    | struct                      | \*gorm.DB | Fetch last row            |
| `Where`   | condition string/map/struct | \*gorm.DB | Add filter                |
| `Updates` | struct/map                  | \*gorm.DB | Update multiple fields    |
| `Update`  | field + value               | \*gorm.DB | Update single field       |
| `Delete`  | struct + id/condition       | \*gorm.DB | Delete rows               |
| `Save`    | struct                      | \*gorm.DB | Insert or update full row |
| `Model`   | model reference             | \*gorm.DB | Choose target model       |
| `Count`   | pointer to int64            | \*gorm.DB | Count rows                |
| `Raw`     | SQL + args                  | \*gorm.DB | Raw SQL → scan            |
| `Exec`    | SQL + args                  | \*gorm.DB | Raw SQL exec (no scan)    |

---

⚡ TL;DR → `config.DB` methods are all about:

1. **Create** new rows
2. **Read** rows (`Find`, `First`, `Where`)
3. **Update** rows (`Update`, `Updates`, `Save`)
4. **Delete** rows (`Delete`)
5. **Custom queries** (`Raw`, `Exec`)
