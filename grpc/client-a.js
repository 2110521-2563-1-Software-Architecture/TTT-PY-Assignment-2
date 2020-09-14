var grpc = require("grpc");
const { PerformanceObserver, performance } = require("perf_hooks");
var booksProto = grpc.load("books.proto");

var client = new booksProto.books.BookService(
  "127.0.0.1:50051",
  grpc.Credentials.createInsecure()
);

function printResponse(error, response) {
  if (error) console.log("Error: ", error);
  else console.log(response);
}

function listBooks() {
  client.list({}, function (error, books) {
    printResponse(error, books);
  });
}

function insertBook(id, title, author) {
  var book = {
    id: parseInt(id),
    title: title,
    author: author,
  };
  client.insert(book, function (error, empty) {
    printResponse(error, empty);
  });
}

function getBook(id) {
  client.get(
    {
      id: parseInt(id),
    },
    function (error, book) {
      printResponse(error, book);
    }
  );
}

function deleteBook(id) {
  client.delete(
    {
      id: parseInt(id),
    },
    function (error, empty) {
      printResponse(error, empty);
    }
  );
}

function watchBooks() {
  var call = client.watch({});
  call.on("data", function (book) {
    console.log(book);
  });
}

var processName = process.argv.shift();
var scriptName = process.argv.shift();
var command = process.argv.shift();

if (command == "list") listBooks();
else if (command == "insert") {
  var n = process.argv[0];
  for (i = 0; i < n; i++) {
    var t0 = performance.now();
    addBook(i, "bookName", "authorName");
    var t1 = performance.now();
    console.log(t1 - t0 + " milliseconds.");
  }
} else if (command == "get") getBook(process.argv[0]);
else if (command == "delete") deleteBook(process.argv[0]);
else if (command == "watch") watchBooks();
