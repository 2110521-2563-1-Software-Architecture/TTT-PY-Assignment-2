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
