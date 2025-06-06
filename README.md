# pwd-tool - Password Management Tool

A simple command-line password management tool written in Go that securely stores passwords in a local SQLite database.

## ✨ Features

*   **Generate Password**: Generate strong random passwords of specified length (including uppercase/lowercase letters, numbers, and special characters).
*   **Save Password**: Associate passwords with names and optional URLs and store them in the database.
*   **List Passwords**: Display all saved password records.
*   **Delete Password**: Delete specific password records by ID.
*   **Local Storage**: Store passwords in a local file (`passwords.db`) using SQLite database.

## 🚀 Usage

### Compilation

Make sure you have Go environment installed (>= 1.18). Run in the project root directory:
