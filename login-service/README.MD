# Login service

### Features Created:

1. **Login Page:**
   - **Function:** `login(w http.ResponseWriter, r *http.Request)`
   - **Description:** This function serves the `login.html` file to the user when they navigate to `/login`. It parses the HTML file and executes the template, sending the resulting HTML to the user's browser.
   - **How to Call:** Access `https://localhost:8080/login` in your web browser.

2. **Login Submission:**
   - **Function:** `loginSubmit(w http.ResponseWriter, r *http.Request)`
   - **Description:** This function handles form submissions from the login page. It reads the `email` and `password` form values, checks them against a hardcoded user map (`users`), and returns a success or failure message based on whether the credentials match.
   - **How to Call:** Submit a form with `email` and `password` fields to `https://localhost:80/login-submit`.

### Detailed Explanation of How to Call These Features:

1. **Serving the Login Page:**

    - When a user navigates to `https://localhost:80/login`, the `handler` function routes the request to the `login` function.
    - The `login` function then looks for a file named `login.html` in the same directory and parses it.
    - If the file is found and successfully parsed, it sends the HTML content to the user's browser.

    **Example URL to access the login page:**
    ```plaintext
    https://localhost:80/login
    ```

2. **Handling Login Submission:**

    - When a user submits the login form (presumably located in `login.html`), the form should be configured to POST the data to `/login-submit`.
    - The `handler` function routes the request to `loginSubmit`.
    - The `loginSubmit` function retrieves the `email` and `password` form values, checks them against the `users` map, and returns an appropriate response.

    **Example URL to submit login data:**
    ```plaintext
    https://localhost:80/login-submit
    ```


### Running the Server:

1. **Ensure you have `cert.pem` and `key.pem` files generated and placed in the same directory as your Go code.**
2. **Run your server:**

    ```sh
    go run main.go
    ```

3. **Access the features in your web browser:**

    - **Login Page:** `https://localhost:80/login`
    - **Submit Login:** Form in the login page submits to `https://localhost:80/login-submit`
