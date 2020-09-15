const {
  Worker,
  isMainThread,
  parentPort,
  workerData,
} = require("worker_threads");
const request = require("request");
const commomPath = "http://localhost:50050/api";
const { performance } = require("perf_hooks");

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
      // console.log(JSON.parse(body));
    }
  );
}

process.argv.shift(); // skip node.exe
process.argv.shift(); // skip name of js file
let n = parseInt(process.argv[0]) || 1;

if (isMainThread) {
  // for (let i = 1; i <= 20; i++) {
  console.log(n);
  var t0 = performance.now();
  for (let j = 0; j < n; j++) {
    new Worker(__filename, { workerData: { num: j } });
  }
  var t1 = performance.now();
  console.log(t1 - t0);
  // }
} else {
  addBook(workerData.num, "title" + workerData.num, "author" + workerData.num);
}
