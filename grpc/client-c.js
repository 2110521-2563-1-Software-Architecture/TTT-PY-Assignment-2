const async = require("async");
const { performance } = require("perf_hooks");

var grpc = require("grpc");

var booksProto = grpc.load("books.proto");

var client = new booksProto.books.BookService(
  "127.0.0.1:50051",
  grpc.credentials.createInsecure()
);

function printResponse(error, response) {
  if (error) console.log("Error: ", error);
  else console.log(response);
}

function insertBook(id, title, author) {
  var book = {
    id: parseInt(id),
    title: title,
    author: author,
  };
  client.insert(book, function (error, empty) {
    printResponse(error, empty);
    var t1 = performance.now();
    if (t1 - t0 > maxTimestamp) {
      maxTimestamp = t1 - t0;
    }
    console.log(maxTimestamp);
  });
}

///////////////////////

process.argv.shift(); // skip node.exe
process.argv.shift(); // skip name of js file
let n = parseInt(process.argv[0]) || 1;

let list = [];

for (let i = 0; i < n; i++) {
  list.push(function func(callback) {
    insertBook(i, "title" + i, "author" + i);
    callback(null, i);
  });
}

let maxTimestamp = -1;

var t0 = performance.now();
async.parallel(list);
