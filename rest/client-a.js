const request = require("request");
const commomPath = "http://localhost:50050/api";
const { PerformanceObserver, performance } = require("perf_hooks");
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
      return;
      // console.log(JSON.parse(body));
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

process.argv.shift(); // skip node.exe
process.argv.shift(); // skip name of js file

var todo = process.argv.shift();
switch (todo) {
  case "list":
    getAllBook();
    break;
  case "get":
    getBookByID(process.argv[0]);
    break;
  case "insert":
    var n = parseInt(process.argv[0]) || 1;
    var t0 = performance.now();
    for (j = 0; j < n; j++) {
      addBook(j, "bookName", "authorName");
    }
    var t1 = performance.now();
    console.log(t1 - t0);

    break;
  case "delete":
    deleteBookByID(process.argv[0]);
    break;
}
