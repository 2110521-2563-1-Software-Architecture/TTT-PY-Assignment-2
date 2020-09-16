const request = require("request");
const commomPath = "http://localhost:50050/api";
let async = require("async");
const { performance } = require("perf_hooks");

process.argv.shift(); // skip node.exe
process.argv.shift(); // skip name of js file
let n = parseInt(process.argv[0]) || 1;

let list = [];

for (let i = 0; i < n; i++) {
  list.push(function func(callback) {
    addBook(i, "title" + i, "author" + i);
    callback(null, i);
  });
}

let maxTimestamp = -1;

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
      console.log({});
      var t1 = performance.now();
      if (t1 - t0 > maxTimestamp) {
        maxTimestamp = t1 - t0;
      }
      console.log(maxTimestamp);
    }
  );
}

var t0 = performance.now();
async.parallel(list);
