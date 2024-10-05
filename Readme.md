# ASCII Art Web Dockerized

## Description

**ASCII Art Web** is a web application that enables users to create stunning ASCII art from any text. With various banner styles to choose from, the application offers a seamless and engaging user experience, now fully containerized with Docker.

## Features

- **Text to ASCII Art**: Generate ASCII art from user-inputted text.
- **Multiple Banner Styles**: Choose from several unique styles, including Shadow, Standard, Thinkertoy, and Chap.
- **Download Feature**: Save your generated ASCII art directly to your device.
- **Custom Error Handling**: User-friendly error messages enhance the experience.
- **Responsive Design**: Works beautifully on both desktop and mobile devices.
- **Dockerized**: Easily deploy using Docker for consistent environments.

## Requirements

- Docker
- A modern web browser (Chrome, Firefox, Safari, etc.)

## Installation

1. Clone the repository:

   ```bash
   git clone https://learn.zone01oujda.ma/git/muboutoub/ascii-art-web-export-file.git
   cd ascii-art-web-export-file
   ```

2. Build the Docker image:

   ```bash
   docker build -t ascii-art-web .
   ```

3. Run the application in a Docker container:

   ```bash
   docker run -p 8080:8080 ascii-art-web
   ```

   > Or use `sh build.sh` to build the image and run the container at once.

## Usage

Once the application is running, navigate to `http://localhost:8080`, enter your text, select your preferred banner style, and view the generated ASCII art in real-time! You can also download your artwork for easy sharing or storage.