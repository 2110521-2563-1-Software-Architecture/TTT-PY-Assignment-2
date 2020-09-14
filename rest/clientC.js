// var worker = new Worker("clientCWorker.js");

// worker.postMessage("Happy Birthday");

// worker.addEventListener("message", function (e) {
//   console.log(e.data);
// });

// worker.postMessage("Happy Birthday");

const {
  Worker,
  isMainThread,
  parentPort,
  workerData,
} = require("worker_threads");
const request = require("request");
const commomPath = "http://localhost:50050/api";

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

if (isMainThread) {
  for (let i = 0; i < 10; i++) {
    new Worker(__filename, { workerData: { num: i } });
  }
} else {
  addBook(workerData.num, "title" + workerData.num, "author" + workerData.num);
}
