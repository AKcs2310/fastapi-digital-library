from fastapi import FastAPI, HTTPException, Request
from pydantic import BaseModel, Field, validator
import time
from typing import Dict

app = FastAPI(title="Digital Library API")

# In-memory database
books_db: Dict[int, dict] = {}


# ---------------- BOOK MODEL ----------------
class Book(BaseModel):
    id: int
    title: str = Field(..., min_length=1)
    author: str
    year: int
    isbn: str

    @validator("year")
    def validate_year(cls, v):
        if v < 1000 or v > 2026:
            raise ValueError("Year must be between 1000 and 2026")
        return v

    @validator("isbn")
    def validate_isbn(cls, v):
        if len(v) not in [10, 13]:
            raise ValueError("ISBN must be 10 or 13 characters long")
        return v


# ---------------- MIDDLEWARE ----------------
@app.middleware("http")
async def log_requests(request: Request, call_next):
    start_time = time.time()

    user_agent = request.headers.get("User-Agent", "Unknown")
    print(f"[LOG] Request received from: {user_agent}")

    response = await call_next(request)

    process_time = time.time() - start_time
    response.headers["X-Process-Time"] = str(process_time)

    return response


# ---------------- CRUD ROUTES ----------------
@app.post("/books", tags=["library"], summary="Add a new book", description="Adds a new book to the digital library")
def create_book(book: Book):
    if book.id in books_db:
        raise HTTPException(status_code=400, detail="Book with this ID already exists")

    books_db[book.id] = book.dict()
    return {"message": "Book added successfully", "book": book}


@app.get("/books", tags=["library"], summary="Get all books", description="Fetch all books from the digital library")
def get_books():
    return list(books_db.values())


@app.get("/books/{book_id}", tags=["library"], summary="Get a book", description="Fetch a book by its ID")
def get_book(book_id: int):
    if book_id not in books_db:
        raise HTTPException(status_code=404, detail="Book not found")

    return books_db[book_id]


@app.put("/books/{book_id}", tags=["library"], summary="Update a book", description="Update book details by ID")
def update_book(book_id: int, book: Book):
    if book_id not in books_db:
        raise HTTPException(status_code=404, detail="Book not found")

    books_db[book_id] = book.dict()
    return {"message": "Book updated successfully", "book": book}


@app.delete("/books/{book_id}", tags=["library"], summary="Delete a book", description="Delete a book from the library")
def delete_book(book_id: int):
    if book_id not in books_db:
        raise HTTPException(status_code=404, detail="Book not found")

    del books_db[book_id]
    return {"message": "Book deleted successfully"}
