# CowsayGenerator 🐮

CowsayGenerator is a fun and lightweight Go program that generates motivational or humorous quotes using OpenAI's GPT API and displays them with the classic `cowsay` terminal program. It's perfect for brightening your day whenever you open a terminal!

## Features

- 🤖 **AI-Powered Quotes**: Generates unique, funny, or motivational quotes using OpenAI's GPT.
- 🐄 **Cowsay Integration**: Pairs the generated quotes with the delightful ASCII cows from `cowsay`.
- 🚀 **Terminal Startup**: Easily set it up to run automatically when you open a terminal.

---

## Prerequisites

### 1. **Install Required Tools**

- **Go**: Install [Go](https://go.dev/) (version 1.18 or later).
- **Cowsay**: Install the `cowsay` program.
  - On Arch Linux:
    ```
    sudo pacman -S cowsay
    ```
  - On Debian/Ubuntu:
    ```
    sudo apt install cowsay
    ```
  - On macOS:
    ```
    brew install cowsay
    ```

### 2. **OpenAI API Key**

- Sign up at [OpenAI](https://openai.com/) and get an API key.
- Add the API key as an environment variable (add to shell config file to persist):
  ```
  export OPENAI_API_KEY=your_api_key_here
  ```

---

## Installation

1. **Clone the Repository**

   ```
   git clone https://github.com/ksande91/cowsay-generator.git
   cd cowsay-generator
   ```

2. **Build the Program**

   ```
   go build -o cowsaygenerator
   ```

3. **Run the Program**
   ```
   ./cowsaygenerator
   ```

---

## Usage

### 1. **Run It Manually**

From the folder where `cowsaygenerator` resides:

```
./cowsaygenerator
```

### 2. **Run It Automatically on Terminal Startup**

To make the program run every time you open a terminal:

#### For **Bash**:

Edit your `~/.bashrc` file:

```
nano ~/.bashrc
```

Add this line:

```
/path/to/cowsaygenerator
```

Save and reload:

```
source ~/.bashrc
```

#### For **Zsh**:

Edit your `~/.zshrc` file:

```
nano ~/.zshrc
```

Add this line:

```
/path/to/cowsaygenerator
```

Save and reload:

```
source ~/.zshrc
```

#### For **Fish**:

Edit your `~/.config/fish/config.fish` file:

```
nano ~/.config/fish/config.fish
```

Add this line:

```
/path/to/cowsaygenerator
```

Save and reload:

```
source ~/.config/fish/config.fish
```

---

## Example Output

Here’s what you’ll see when you run the program:

```
 __________________________
< "Don’t moooove too fast, enjoy the grass!" >
 --------------------------
        \   ^__^
         \  (oo)\_______
            (__)\       )\/\
                ||----w |
                ||     ||
```

---

## License

This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.
