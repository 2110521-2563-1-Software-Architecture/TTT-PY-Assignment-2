# Assignment #1: gRPC and REST API Implementation
# Member

1. Natcha Manasuntorn 6030177021
2. Suchut Sapsathien 6030609921
3. Karnkitti Kittikamron 6031006621
4. Yanika Dontong 6031010021
5. Natthanon Manop 6031013021

# Screenshots of Swagger for APIs
![](https://i.imgur.com/3yPHIS9.png)


# Source codes
## Server
### index.js
```javascript
var express = require("express");
var Library = require("./mock.js");
var cors = require("cors");
const PORT = 50050;

var app = express();

app.use(express.json());
app.use(express.urlencoded({ extended: false }));
app.use(cors());

function error(status, msg) {
  var err = new Error(msg);
  err.status = status;
  return err;
}

app.get("/api/books/allbook", function (req, res, next) {
  res.send(Library.getAllBooks());
});

app.get("/api/books/:id", function (req, res, next) {
  const bookId = req.params.id;

  const book = Library.getBookByID(bookId);

  if (book) res.send(book);
  else next();
});

app.post("/api/addbook", function (req, res, next) {
  const { id, title, author } = req.body;

  const addedBook = Library.addBook(id, title, author);
  res.status(201);
  res.send(addedBook);
});

app.delete("/api/deletebook/:id", function (req, res, next) {
  const bookId = req.params.id;
  Library.deleteBookByID(bookId);
  res.status(200);
  res.send({ status: "success" });
});

app.use(function (err, req, res, next) {
  res.status(err.status || 500);
  res.send({ error: err.message });
});

app.use(function (req, res) {
  res.status(404);
  res.send({ error: "Oops, can't find that" });
});

app.listen(PORT, () => {
  console.log(`Express started on port ${PORT}`);
});

```
### mock.js
```javascript
class Book {
  constructor(id, title, author) {
    this._id = id;
    this._title = title;
    this._author = author;
  }
}

class Library {
  constructor() {
    var b = new Book(123, "A Tale of Two Cities", "Charles Dickens");
    this._books = [b];
  }

  getBookByID(id) {
    return this._books.find((book) => {
      if (book._id == id) return book;
    });
  }

  getAllBooks() {
    return this._books;
  }

  addBook(id, title, author) {
    var book = new Book(parseInt(id), title, author);
    this._books.push(book);
    return book;
  }

  deleteBookByID(id) {
    this._books = this._books.filter(function (obj) {
      return obj._id !== parseInt(id);
    });
    return;
  }
}

module.exports = new Library();


```
## Client
### client.js
```javascript
const request = require("request");
const commomPath = "http://localhost:50050/api";

function getAllBook() {
  request.get(commomPath + "/books/allbook", function (error, response, body) {
    if (error) {
      return console.log(error);
    }
    console.log(JSON.parse(body));
  });
}
function getBookByID(id) {
  request.get(commomPath + `/books/${parseInt(id)}`, function (
    error,
    response,
    body
  ) {
    if (error) {
      return console.log(error);
    }
    console.log(JSON.parse(body));
  });
}
function addBook(id, title, author) {
  request.post(
    {
      headers: { "content-type": "application/json" },
      url: commomPath + "/addbook",
      body: JSON.stringify({ id, title, author }),
    },
    function (error, response, body) {
      if (error) {
        return console.log(error);
      }
      console.log(JSON.parse(body));
    }
  );
}
function deleteBookByID(id) {
  request.delete(commomPath + `/deletebook/${parseInt(id)}`, function (
    error,
    response,
    body
  ) {
    if (error) {
      return console.log(error);
    }
    console.log(JSON.parse(body));
  });
}

process.argv.shift(); 
process.argv.shift(); 

var todo = process.argv.shift();

switch (todo) {
  case "list":
    getAllBook();
    break;
  case "get":
    getBookByID(process.argv[0]);
    break;
  case "insert":
    addBook(process.argv[0], process.argv[1], process.argv[2]);
    break;
  case "delete":
    deleteBookByID(process.argv[0]);
    break;
}

```
# Compare how to call the methods based on gRPC and REST API side-by-side, e.g. in a Table format as shown below.

| Functions  | gRPC  | REST API |
| :--------:  | :--------: | :--------: |
| List books | `client.List(ctx, &pb.Empty{})` | `request.get("http://localhost:50050/api/books/allbook", function (error, response, body) { ... });`|
| Insert book | `client.Insert(ctx, book)` | `request.post({headers: { "content-type": "application/json" }, url: "http://localhost:50050/api/addbook", body: JSON.stringify({ title, author })}, function (error, response, body) { ... });` |
| Delete book | `client.Delete(ctx, &pb.BookIdRequest{int32(id)})` | `request.delete("http://localhost:50050/deletebook/${parseInt(id)}", function (error, response, body) { ... });` |
| Get book | `client.Get(ctx, &pb.BookIdRequest{int32(id)})` | `request.get("http://localhost:50050/api/books/${parseInt(id)}", function (error, response, body) { ... });` |
| Watch | `client.Watch(ctx, &pb.Empty{})` | - |


# What are the main differences between REST API and gRPC?
 * gRPC สามารถทำงานได้เร็วกว่าเนื่องจากทำงานบน protocol HTTP/2 ที่ส่งข้อมูลเป็น Protobuf(binary) ซึ่งมีขนาดเล็ก ต่างจาก REST API ที่ทำงานบน protocol HTTP/1.1 ที่ส่งข้อมูลเป็น JSON(text) ซึ่งมีขนาดใหญ่กว่า
 * gRPC สามารถทำ streaming(ส่งข้อมูลสองทาง) ได้ ต่างจาก REST API ที่ส่ง request จาก Client ไปยัง Server ได้เท่านั้น
 * gRPC มี API Contract ที่เข้มงวด ต้องกำหนดรูปแบบของ service เป็นไฟล์ .proto เพื่อใช้งาน ต่างจาก REST API ที่สามารถเรียกใช้งานผ่าน URL ได้เลย
 * gRPC สร้างใช้งานกับ Web ที่มีการติดตั้ง gRPC เท่านั้น ไม่สามารถเรียกผ่าน Web browser ต่างจาก REST API ที่สามารถเข้าจาก Web browser ได้

| Feature | gRPC | REST API |
| :--------: | :--------: | :--------: |
|  Protocol  |  HTTP/2 (fast) | HTTP/1.1 (slow)|
|  Payload     |    Protobuf (binary, small)     | JSON (text, large) |
|  API contract   |   Strict, required (.proto)  |   Loose, optional (OpenAPI)|
|  Code generation   |  Built-in (protoc)  |  Third-party tools (Swagger)    |
|  Streaming     |  Bidirectional streaming    |   Client -> server request only  |
|  Browser support     |  Limited (require gRPC-web) |  Yes   |

# What is the benefits of introduce interface in front of the gRPC and REST API of the book services?
* สามารถเรียกใช้งาน service ได้สะดวกมากยิ่งขึ้น
* สามารถเรียกใช้ได้ด้วยคำสั่งที่เข้าใจง่ายและสื่อความหมาย
* สามารถทำงานได้โดย client ไม่รู้การทำงานเบื้องหลังของฝั่ง server
# Based on the introduced interface, compare how to call the methods based on gRPC and REST API side-by-side, e.g. in a Table format as shown below.



| Functions | gRPC | REST API |
| :--------: | :--------: | :--------: |
| List books     | `go run client.go list` | `node client.js list`     |
| Insert book     | `go run client.go insert <id> <title> <author>` | `node client.js insert <id> <title> <author>`    |
| Delete book     | `go run client.go delete <id>`    | `node client.js delete <id>`     |
| Get book     | `go run client.go get <id>`   | `node client.js get <id>`     |
| Watch     | `go run client.go watch`  | -     |


# Draw a component diagram representing the book services with and without interfaces.
![](https://i.imgur.com/ltCYqnO.jpg)
