var express = require("express");
var Library = require("./mock.js");
var cors = require("cors");
const PORT = 50050;

var app = express();

// Middleware for parsing
app.use(express.json()); // for parsing application/json
app.use(express.urlencoded({ extended: false })); // for parsing application/x-www-form-urlencoded
app.use(cors());
// custom error handler
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
// Middleware will be trigger in order

// middleware with an arity of 4 are considered
// error handling middleware. When you next(err)
// it will be passed through the defined middleware
// in order, but ONLY those with an arity of 4, ignoring
// regular middleware.
app.use(function (err, req, res, next) {
  res.status(err.status || 500);
  res.send({ error: err.message });
});

// custom JSON 404 middleware. Since it's placed last
// it will be the last middleware called, if all others
// invoke next() and do not respond.
app.use(function (req, res) {
  res.status(404);
  res.send({ error: "Oops, can't find that" });
});

app.listen(PORT, () => {
  console.log(`Express started on port ${PORT}`);
});
