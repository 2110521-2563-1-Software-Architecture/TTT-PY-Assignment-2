const { performance } = require("perf_hooks");
const request = require("request");
const commomPath = "http://localhost:50050/api";

class Client {
  async getAllBook() {
    request.get(commomPath + "/books/allbook", function (
      error,
      response,
      body
    ) {
      if (error) {
        return console.log(error);
      }
      // console.log(JSON.parse(body));
    });
  }
  async getBookByID(id) {
    request.get(commomPath + `/books/${parseInt(id)}`, function (
      error,
      response,
      body
    ) {
      if (error) {
        return console.log(error);
      }
      // console.log(JSON.parse(body));
    });
  }
  async addBook(id, title, author) {
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
  async deleteBookByID(id) {
    request.delete(commomPath + `/deletebook/${parseInt(id)}`, function (
      error,
      response,
      body
    ) {
      if (error) {
        return console.log(error);
      }
      // console.log(JSON.parse(body));
    });
  }
}

async function start() {
  process.argv.shift(); // skip node.exe
  process.argv.shift(); // skip name of js file
  let n = parseInt(process.argv[0]) || 1;

  let t1 = performance.now();
  for (let i = 0; i < n; i++) {
    let client = new Client();
    let book = {
      id: i,
      title: "Book" + i.toString(),
      author: "Author",
    };
    await client.getAllBook();
    await client.addBook(book.id, book.title, book.author);
    await client.getBookByID(book.id);
    await client.deleteBookByID(book.id);
  }
  let t2 = performance.now();
  console.log(t2 - t1);
}

start();
