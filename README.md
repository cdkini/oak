# Oak

A simple CLI utility for developer notetaking. 


### Dependencies
Oak is a lightweight wrapper around `fd`, `rg`, `fzf`, `bat`, and `nvim`. 


### Installation
```bash
git clone https://github.com/cdkini/oak.git
cd oak
sudo cp oak /usr/local/bin # Or wherever you keep your scripts
```

You must also set your `OAK_ROOT` env var to point to a notes folder.
I like using `~/Dropbox` to keep things in sync between my machines.

### API
```
oak add <NAME?>     # Add new note
oak daily           # Open daily note
oak delete <NAME?>  # Delete note
oak grep <QUERY?>   # Search notes (by contents)
oak find <QUERY?>   # Search notes (by title) 
```

