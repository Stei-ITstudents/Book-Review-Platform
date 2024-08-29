from flask import Blueprint, render_template, redirect, url_for, request, flash, session
import requests

auth = Blueprint('auth', __name__)

# registration route
@auth.route('/register', methods=['GET', 'POST'])
def register():
    if request.method == 'POST':
        # collect form data
        user_data = {
            'username': request.form['username'],
            'password': request.form['password'],
        }
        # send post request to backend for registration
        response = requests.post('http://localhost:8080/register', json=user_data)
        if response.status_code == 201:
            flash('Registration successful! Please log in.', 'success')
            return redirect(url_for('auth.login'))
        else:
            flash('Registration failed. Please try again.', 'danger')
    return render_template('register.html')

# login route
@auth.route('/login', methods=['GET', 'POST'])
def login():
    if request.method == 'POST':
        # collect form data
        credentials = {
            'username': request.form['username'],
            'password': request.form['password'],
        }
        # send post request to backend for authentication
        response = requests.post('http://localhost:8080/login', json=credentials)
        if response.status_code == 200:
            session['user'] = response.json()  # assuming the backend returns user info
            flash('Login successful!', 'success')
            return redirect(url_for('index'))
        else:
            flash('Login failed. Please check your credentials.', 'danger')
    return render_template('login.html')

# logout route
@auth.route('/logout')
def logout():
    session.pop('user', None)
    flash('You have been logged out.', 'info')
    return redirect(url_for('auth.login'))
