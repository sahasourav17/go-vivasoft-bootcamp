### 1. What is `struct` in Go? Explain it. (Go তে Struct কি ? ব্যাখ্যা করুন ।)

A struct in Go is a composite data type that groups together variables under a single type. These variables (also called fields) can be of different types, and a struct is used to define and represent real-world objects or concepts.

```go
type Person struct {
    Name string
    Age  int
    City string
}
```

Here, Person is a struct with three fields: Name, Age, and City

### 2. What are the methods in G0? (Go তে মেথডস কি ? )

In Go, a method is a function that is associated with a specific type (usually a struct). Methods allow you to define behaviors for your types.

A method is defined using a receiver—a variable that binds the method to the type.

```go
type Rectangle struct {
    Width, Height float64
}

// Method for Rectangle
func (r Rectangle) Area() float64 {
    return r.Width * r.Height
}
```

Here, the Area method is associated with the Rectangle type.

### 3. Does Go support Type Inheritance? (Go তে কি টাইপ ইনহেরিটেন্স সাপোর্ট করে ?)

No, Go does not support classical type inheritance as seen in object-oriented languages like Java or C++. Instead, Go uses composition (embedding one type into another) to achieve code reuse.

```go
type Animal struct {
    Name string
}

type Dog struct {
    Animal  // Embedding Animal struct
    Breed string
}
```

### 4. Does Go support Method Overloading? (Go তে কি মেথড ওভারলোডিং সাপোর্ট করে ? )

No, Go does not support method overloading (i.e., defining multiple methods with the same name but different parameters). Each method must have a unique name within its type.

### 5. How is a Struct created in Go? (Go তে একটি Struct কিভাবে তৈরি করা হয় ?)

```go
type Employee struct {
    Name   string
    Salary float64
    Age    int
}

// creating an instance of Employee struct
e := Employee{Name: "John", Salary: 50000, Age: 30}
```

### 6. What is a Pointer? (পয়েন্টার কি ব্যাখ্যা করুন ।)

A pointer in Go is a variable that stores the memory address of another variable. Pointers are used to access and modify data without copying it, enabling efficient memory usage.

```go
x := 10
p := &x // 'p' is a pointer to the address of 'x'
fmt.Println(*p) // Dereferences the pointer to get the value of 'x'
```

### 7. What is the default value of a Pointer?(একটা পয়েন্টারের ডিফল্ট ভ্যালু কি?)

he default value of a pointer in Go is nil, which means it does not point to any memory address.

```go
var p *int
fmt.Println(p) // Output: <nil>
```

### 8. How is a Pointer represented? (পয়েন্টার কিভাবে রিপ্রেজেন্ট করা হয়?)

A pointer in Go is represented using the \* and & symbols:

- `*` is used to dereference a pointer (access the value it points to).

- `&` is used to get the memory address of a variable

```go
x := 42
p := &x  // 'p' stores the address of 'x'
fmt.Println(*p) // Dereference 'p' to get the value of 'x'
```
