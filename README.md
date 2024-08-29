
---

![received_1097514964147756.(jpg?1652351843](https://bookshop-uk-prod-images.storage.googleapis.com/spree/affiliate_profiles/banner_images/10531/original/received_1097514964147756.jpg?1652351843)
# BookNook - Book Review Platform


Welcome to **BookNook**, a simple web application that allows users to browse, add, rate, and review books. This platform provides an easy way for users to discover new books, share their thoughts, and see what others have to say.

## Table of Contents

- [Features](#features)
- [Technologies Used](#technologies-used)
- [Folder Structure](#folder-structure)
- [Setup Instructions](#setup-instructions)
- [API Endpoints](#api-endpoints)
- [Database Schema](#database-schema)
- [Optional Features](#optional-features)
- [Contributing](#contributing)
- [License](#license)

## Features

- **Book Browsing**: View a list of books with titles, authors, and average ratings.
- **Book Details**: Click on a book to see more details, including a description, user reviews, and ratings.
- **Add Books**: Users can contribute by adding new books to the platform.
- **Submit Reviews**: Users can rate books and leave detailed reviews.
- **Filter and Sort**: Filter books by genre or author, and sort by highest rating or most reviewed.

## Technologies Used

- **Backend**: Go (Golang) for RESTful API
- **Frontend**: Python Flask for rendering templates and handling user interactions
- **Database**: MySQL for storing book and review data
- **HTML/CSS/JavaScript**: For the frontend UI

## Folder Structure

```plaintext
booknook/
├── api/ - Go backend API folder
│   ├── main.go                   # entry point for the api
│   ├── handlers/                 # request handlers folder
│   │   ├── bookHandler.go        # handles book-related routes
│   │   └── reviewHandler.go      # handles review-related routes
│   ├── models/                   # database models folder
│   │   ├── book.go               # book model definition
│   │   └── review.go             # review model definition
│   ├── routes/                   # api routes folder
│   │   └── routes.go             # defines api routes
│   ├── database/                 # database connection folder
│   │   └── db.go                 # manages database connection
│   ├── utils/                    # utility functions folder
│   │   └── utils.go              # common utility functions
│   ├── templates/                # html templates folder (Go templates)
│   │   ├── index.html            # homepage template
│   │   ├── book_details.html     # book details page template
│   │   ├── add_book.html         # add new book page template
│   │   └── login.html            # login page template
│   └── static/                   # static assets folder
│       ├── css/                  # css styles folder
│       │   └── styles.css        # main stylesheet
│       └── js/                   # javascript folder
│           ├── book-details.js   # javascript for book details page
│           ├── book-list.js      # javascript for book list page
│           └── review.js         # javascript for review page
├── sql/ - MySQL                  # sql scripts folder
│   ├── schema.sql                # database schema script
│   └── seed.sql                  # initial seed data script
├── README.md                     # project setup and instructions
└── .gitignore                    # git ignore file
```

## Setup Instructions

### Prerequisites

- **Go**: Ensure you have Go installed (version 1.16 or later).
- **Python**: Ensure you have Python installed (version 3.8 or later).
- **MySQL**: Ensure MySQL is installed and running.

### Backend Setup (Go API)

**Navigate to the `api/` directory**:
   ```bash
   cd booknook/api
   ```

**Install Go dependencies**:
   ```bash
   go mod tidy
   ```

**Set up the database**:
   - Create a MySQL database.
   - Update `db.go` with your MySQL connection details.

**Run the API**:
   ```bash
   go run main.go
   ```
   The API will run on `http://localhost:8080`.

### Frontend Setup (Flask)

**Navigate to the `bff/` directory**:
   ```bash
   cd booknook/bff
   ```

**Create a virtual environment**:
   ```bash
   python -m venv venv
   ```

**Activate the virtual environment**:
   - On macOS/Linux:
     ```bash
     source venv/bin/activate
     ```

**Install Flask dependencies**:
   ```bash
   pip install -r requirements.txt
   ```

**Run the Flask app**:
   ```bash
   flask run
   ```
   The app will run on `http://localhost:5000`.

### Database Setup

**Run the schema script**:
   ```bash
   mysql -u your_username -p your_database_name < sql/schema.sql
   ```

**(Optional) Seed the database**:
   ```bash
   mysql -u your_username -p your_database_name < sql/seed.sql
   ```

## API Endpoints

- **POST /books**: Add a new book.
- **GET /books**: Retrieve a list of all books.
- **GET /books/{id}**: Retrieve details for a specific book, including reviews.
- **POST /books/{id}/reviews**: Add a review to a specific book.

## Database Schema

### Tables

- **Books**:
  - `id`: Primary key
  - `title`: Book title
  - `author`: Book author
  - `description`: Book description
  - `cover_image`: URL to the book cover image (optional)
  - `average_rating`: Average rating of the book

- **Reviews**:
  - `id`: Primary key
  - `book_id`: Foreign key to the Books table
  - `user_id`: ID of the user who left the review
  - `rating`: Rating given by the user (1-5)
  - `review_text`: Review text

## Optional Features

- **User Authentication**: Implement user login and registration.
- **User Profiles**: Allow users to view the books they've reviewed.
- **Search**: Add search functionality to find books by title or author.
- **Recommendations**: Suggest similar books based on genre or ratings.
- **Responsive Design**: Ensure the application is mobile-friendly.

## Contributing

We welcome contributions! Please fork the repository, create a new branch, and submit a pull request.

---
