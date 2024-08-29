from flask import Flask, render_template, request, redirect, url_for, session, flash
import os
import requests
from auth import auth
from dotenv import load_dotenv

app = Flask(__name__)
app.secret_key = 'my_secret_key'  # my secret key

#* load strong secret key :)
load_dotenv(dotenv_path='../.env')
app.secret_key = os.getenv('FLASK_SECRET_KEY', 'my_secret_key')  # development key


# configuration
API_BASE_URL = os.getenv("API_BASE_URL", "http://localhost:8080") 

app.register_blueprint(auth)


@app.route('/')
def index():
    # get filter and sort parameters from request
    genre = request.args.get('genre')
    sort_by = request.args.get('sort_by')

    # build query parameters for api request
    params = {}
    if genre:
        params['genre'] = genre
    if sort_by:
        params['sort_by'] = sort_by

    response = requests.get(f'{API_BASE_URL}/books', params=params)  # fetch books from go backend
    if response.status_code == 200:
        books = response.json()
    else:
        books = []

    return render_template('index.html', books=books)

# book details route
@app.route('/books/<int:id>')
def book_details(id):
    response = requests.get(f'{API_BASE_URL}/books/{id}')  # fetch book details
    if response.status_code == 200:
        book = response.json()
    else:
        book = None
    return render_template('book_details.html', book=book)

# add book route
@app.route('/add-book', methods=['GET', 'POST'])
def add_book():
    if request.method == 'POST':
        # check if the user is logged in
        if 'user_id' not in session:
            return redirect(url_for('login'))  # redirect to login 
        
        # collect form data
        new_book = {
            'title': request.form['title'],
            'author': request.form['author'],
            'description': request.form['description'],
            'cover_image': request.form.get('cover_image')  # optional field
        }
        # send post request to backend
        response = requests.post(f'{API_BASE_URL}/books', json=new_book)
        if response.status_code == 201:
            flash('Book added successfully!', 'success')
            return redirect(url_for('index'))
        else:
            flash('Failed to add book', 'danger')
            return redirect(url_for('add_book'))
    return render_template('add_book.html')

# add review route
@app.route('/books/<int:id>/add-review', methods=['POST'])
def add_review(id):
    if 'user_id' not in session:
        return redirect(url_for('login'))  # redirect to login 

    review = {
        'user_id': session['user_id'],
        'rating': request.form['rating'],
        'review_text': request.form['review_text']
    }

    response = requests.post(f'{API_BASE_URL}/books/{id}/reviews', json=review)
    if response.status_code == 201:
        flash('Review added successfully!', 'success')
    else:
        flash('Failed to add review', 'danger')

    return redirect(url_for('book_details', id=id))

# filter and sort books
@app.route('/filter', methods=['GET'])
def filter_books():
    genre = request.args.get('genre')
    return redirect(url_for('index', genre=genre))

@app.route('/sort', methods=['GET'])
def sort_books():
    sort_by = request.args.get('sort_by')
    return redirect(url_for('index', sort_by=sort_by))

if __name__ == '__main__':
    app.run(debug=True, port=5000)
    