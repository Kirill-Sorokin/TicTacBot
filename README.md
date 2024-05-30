# TicTacBot

**Overview**

TicTacBot is an intelligent Tic Tac Toe game implemented in Go using the GTK library for the graphical user interface. The bot is designed to play intelligently, providing a challenging opponent for the player. The game includes a simple and intuitive interface, a score tracker, and a restart button to reset the game at any time.

**Features**

- Intelligent bot opponent that plays strategically to block the player and win.
- Simple and clean graphical user interface built with GTK.
- Score tracking to keep track of player and bot wins.
- Randomized turn order to start each game, making each game unique.
- Restart button to reset the game and play again.

**Installation**

To install and run TicTacBot, follow these steps:

1. **Clone the repository:**

    ```sh
    git clone https://github.com/Kirill-Sorokin/TicTacBot.git
    cd TicTacBot
    ```

2. **Install the GTK development libraries:**

    On macOS, use Homebrew:

    ```sh
    brew install gtk+3
    ```

    On Ubuntu or Debian-based distributions:

    ```sh
    sudo apt-get install libgtk-3-dev
    ```

3. **Build and run the application:**

    ```sh
    go build main.go
    ./main
    ```

**Usage**

- Run the application.
- The game will start with a random turn order.
- Click on the cells to make your move.
- The bot will automatically make its move after yours.
- The game will indicate the winner or if it's a draw.
- Click the "Restart" button to reset the game and play again.
- Track the scores of player and bot wins at the top of the window.

**License**

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
