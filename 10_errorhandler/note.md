# Go Error Handling Guide

## ความแตกต่างระหว่าง `errors.Is` และ `errors.As`

### `errors.Is` - ตรวจสอบค่า Error
ใช้เช็คว่า error **ตรงกับค่า error ที่กำหนด**หรือไม่

```go
if errors.Is(err, io.EOF) {
    // จัดการ EOF error
}
```

**ใช้เมื่อ:** ต้องการรู้ว่า "error นี้เป็น error อันนี้หรือเปล่า?"

### `errors.As` - ตรวจสอบชนิด Error และดึงข้อมูล
ใช้เช็คว่า error เป็น**ชนิดที่กำหนด**หรือไม่ และดึง error นั้นออกมาใช้

```go
var pathErr *os.PathError
if errors.As(err, &pathErr) {
    // pathErr มีข้อมูลของ error แล้ว
    fmt.Println(pathErr.Path)
}
```

**ใช้เมื่อ:** ต้องการรู้ว่า "error นี้เป็นชนิดนี้หรือเปล่า และถ้าใช่ขอข้อมูลด้วย?"

### สรุปความแตกต่าง

| | `errors.Is` | `errors.As` |
|---|---|---|
| **เช็คอะไร** | ค่า error (value) | ชนิด error (type) |
| **Parameter** | `errors.Is(err, target)` | `errors.As(err, &target)` |
| **Return** | `bool` | `bool` และใส่ข้อมูลใน target |
| **ตัวอย่าง** | เป็น `io.EOF` หรือเปล่า? | เป็น `*os.PathError` หรือเปล่า? |

---

## ความแตกต่างระหว่าง `errors.New` และ `fmt.Errorf`

### `errors.New()` - Error แบบ Static

```go
var ErrNotFound = errors.New("not found")
var ErrUnauthorized = errors.New("unauthorized")
```

**ลักษณะ:**
- สร้าง error ที่เป็น**ค่าคงที่** (constant error)
- ข้อความไม่เปลี่ยนแปลง
- สร้างครั้งเดียวตอน package initialization
- ใช้เปรียบเทียบด้วย `errors.Is()` หรือ `==` ได้

**ใช้เมื่อ:**
- Error ที่ซ้ำๆ ในหลายที่
- Sentinel errors (errors ที่กำหนดไว้ล่วงหน้า)
- ต้องการเปรียบเทียบ error

```go
func GetUser(id int) error {
    if id < 0 {
        return ErrNotFound // ใช้ error เดิม
    }
    return nil
}

// ตรวจสอบ
if errors.Is(err, ErrNotFound) {
    // จัดการ not found
}
```

### `fmt.Errorf()` - Error แบบ Dynamic

```go
func FirstError() (string, error) {
    return "", fmt.Errorf("this is first error")
}
```

**ลักษณะ:**
- สร้าง error **ใหม่ทุกครั้ง**ที่เรียก
- สามารถใส่ตัวแปรเข้าไปใน message ได้
- รองรับ **error wrapping** ด้วย `%w`

**ใช้เมื่อ:**
- ต้องการใส่ข้อมูลเพิ่มเติม (context)
- Wrap error จาก layer อื่น
- Error message ที่เปลี่ยนตามสถานการณ์

```go
func GetUser(id int) error {
    if id < 0 {
        return fmt.Errorf("invalid user id: %d", id)
    }
    return nil
}

// Wrapping error
func ProcessUser(id int) error {
    user, err := db.GetUser(id)
    if err != nil {
        return fmt.Errorf("failed to get user %d: %w", id, err)
    }
    return nil
}
```

### สรุปความแตกต่าง

| | `errors.New` | `fmt.Errorf` |
|---|---|---|
| **สร้างเมื่อไหร่** | ครั้งเดียว (package init) | ทุกครั้งที่เรียก function |
| **ข้อความ** | คงที่ | ใส่ตัวแปรได้ |
| **Wrapping** | ไม่ได้ | ได้ (ใช้ `%w`) |
| **Performance** | เร็วกว่า | ช้าเล็กน้อย |
| **Use case** | Sentinel errors | Dynamic errors + wrapping |

---

## Bad Error Handling Patterns

### 1. ❌ ไม่เช็ค Error (Ignoring Errors)
```go
// แย่
file, _ := os.Open("config.json")
json.Unmarshal(data, &config)
```

```go
// ดี
file, err := os.Open("config.json")
if err != nil {
    return err
}
```

### 2. ❌ Error Message ไม่มี Context
```go
// แย่
if err != nil {
    return errors.New("error occurred")
}
```

```go
// ดี - ใส่ context ว่าเกิดอะไร
if err != nil {
    return fmt.Errorf("failed to open config file: %w", err)
}
```

### 3. ❌ เช็ค `err.Error() != ""`
```go
// แย่ - จะ panic ถ้า err เป็น nil
if creditErr.Error() != "" {
    // ...
}
```

```go
// ดี
if creditErr != nil {
    // ...
}
```

### 4. ❌ ใช้ Switch กับ Error Value โดยตรง
```go
// แย่ - ใช้ไม่ได้กับ wrapped errors
switch creditErr {
case hello.LowerPayment:
    fmt.Println("lower payment")
}
```

```go
// ดี - ใช้ errors.Is
if errors.Is(creditErr, hello.LowerPayment) {
    fmt.Println("lower payment")
}
```

### 5. ❌ Panic สำหรับ Error ปกติ
```go
// แย่
file, err := os.Open("file.txt")
if err != nil {
    panic(err) // อย่า panic สำหรับ error ธรรมดา!
}
```

```go
// ดี
if err != nil {
    return nil, err
}
```

### 6. ❌ ไม่ใช้ `%w` ตอน Wrap Error
```go
// แย่ - error chain หาย
if err != nil {
    return fmt.Errorf("database error: %v", err)
}
```

```go
// ดี - เก็บ error chain ไว้
if err != nil {
    return fmt.Errorf("database error: %w", err)
}
```

### 7. ❌ ไม่ใช้ defer สำหรับ Cleanup
```go
// แย่
func Process() error {
    f, err := os.Open("file.txt")
    if err != nil {
        return err
    }
    
    data, err := process(f)
    if err != nil {
        return err // file ไม่ถูกปิด!
    }
    
    f.Close()
    return nil
}
```

```go
// ดี
func Process() error {
    f, err := os.Open("file.txt")
    if err != nil {
        return err
    }
    defer f.Close() // ปิดเสมอ
    
    return process(f)
}
```

### 8. ❌ Return ทั้ง Error และ Data
```go
// แย่ - สับสน
func GetUser(id int) (*User, error) {
    user := &User{ID: id}
    if err := db.Load(user); err != nil {
        return user, err // อย่า return user ตอน error!
    }
    return user, nil
}
```

```go
// ดี - return nil ตอน error
func GetUser(id int) (*User, error) {
    user := &User{ID: id}
    if err := db.Load(user); err != nil {
        return nil, err
    }
    return user, nil
}
```

---

## Go Idioms & Best Practices

### 1. Check Errors Immediately
```go
f, err := os.Open("file.txt")
if err != nil {
    return err
}
defer f.Close()
```

### 2. Early Returns (Guard Clauses)
```go
func Validate(s string) error {
    if s == "" {
        return errors.New("empty string")
    }
    if len(s) > 100 {
        return errors.New("too long")
    }
    // Happy path ไม่ซ้อนกัน
    return nil
}
```

### 3. Accept Interfaces, Return Structs
```go
// ดี - input ยืดหยุ่น, output ชัดเจน
func Process(r io.Reader) *Result {
    // ...
}
```

### 4. Use defer for Cleanup
```go
func ReadFile() error {
    f, err := os.Open("file.txt")
    if err != nil {
        return err
    }
    defer f.Close()
    
    // ทำงานกับ file
    return nil
}
```

### 5. Zero Values Are Useful
```go
var s string       // "" (empty string)
var n int          // 0
var b bool         // false
var buf bytes.Buffer // พร้อมใช้งานทันที
```

---

## ตัวอย่างที่ดี: Combining Sentinel Errors และ Dynamic Errors

```go
// Sentinel errors - ใช้ errors.New
var (
    ErrInsufficientFunds = errors.New("insufficient funds")
    ErrAccountClosed = errors.New("account closed")
    ErrInvalidAmount = errors.New("invalid amount")
)

// Dynamic errors with context - ใช้ fmt.Errorf
func Transfer(from, to int, amount float64) error {
    if amount <= 0 {
        return ErrInvalidAmount
    }
    
    balance, err := getBalance(from)
    if err != nil {
        return fmt.Errorf("failed to get balance for account %d: %w", from, err)
    }
    
    if balance < amount {
        return ErrInsufficientFunds
    }
    
    if err := debit(from, amount); err != nil {
        return fmt.Errorf("failed to debit account %d: %w", from, err)
    }
    
    if err := credit(to, amount); err != nil {
        return fmt.Errorf("failed to credit account %d: %w", to, err)
    }
    
    return nil
}

// การใช้งาน
err := Transfer(123, 456, 100.0)
if err != nil {
    if errors.Is(err, ErrInsufficientFunds) {
        log.Println("Not enough money in account")
    } else if errors.Is(err, ErrInvalidAmount) {
        log.Println("Amount must be positive")
    } else {
        log.Printf("Transfer failed: %v", err)
    }
}
```

---

## Golden Rules

1. **ทุก error ต้องถูกจัดการ** - return, handle, หรือ log พร้อมเหตุผล
2. **ใช้ `%w` สำหรับ wrapping** - เพื่อเก็บ error chain
3. **Check error ทันที** - อย่ารอเช็คหลายบรรทัด
4. **ใส่ context ใน error message** - บอกว่าเกิดอะไร ที่ไหน
5. **Clear is better than clever** - เขียนให้อ่านง่าย ไม่ต้องเก่ง

---

**สรุป:** Error handling ที่ดีทำให้ code ทนทาน debug ง่าย และ maintain ได้ในระยะยาว