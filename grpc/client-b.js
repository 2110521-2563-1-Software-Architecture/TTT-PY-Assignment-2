var grpc = require("grpc");
const { PerformanceObserver, performance } = require("perf_hooks");
var booksProto = grpc.load("books.proto");
var client = new booksProto.books.BookService(
  "127.0.0.1:50051",
  grpc.credentials.createInsecure()
);

function printResponse(error, response) {
  if (error) console.log("Error: ", error);
  // else console.log(response);
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
async function start() {
  process.argv.shift(); // skip node.exe
  process.argv.shift(); // skip name of js file
  let n = parseInt(process.argv[0]) || 1;

  let t1 = performance.now();
  for (let i = 0; i < n; i++) {
    let book = {
      id: i,
      title: "Book" + i.toString(),
      author: "Author",
    };
    await listBooks();
    await insertBook(book.id, book.title, book.author);
    await getBook(book.id);
    await deleteBook(book.id);
  }
  let t2 = performance.now();
  console.log(t2 - t1);
}

start();
