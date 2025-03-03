# Oak 🌳

A simple CLI utility for developer notetaking. Oak wraps a number of common terminal-based tools and leverages Dropbox for file syncing.

### Dependencies
Oak is a lightweight wrapper around `fd`, `rg`, `fzf`, `bat`, and `nvim`.
It is written in Python 3.9 but has no dependencies.


### Installation
```bash
curl -s https://raw.githubusercontent.com/cdkini/oak/refs/heads/main/oak > oak 
chmod +x oak
sudo cp oak /usr/local/bin # Or wherever you keep your scripts
```

You must also set your `OAK_ROOT` env var to point to an existing Dropbox folder.

### API
```bash
oak add <NAME> <TAGS?>          # Add new note
oak daily                       # Open daily note
oak delete <NAME?>              # Delete note
oak grep <QUERY?>               # Search notes by contents
oak open <QUERY?>               # Open notes by title 
oak list                        # List files (in reverse chronological order)
```
