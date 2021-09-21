# cremem
cremem is a tool for recording a command and interactively choosing and executing it.

# Setup
## Load the zsh script
Add this line to your .zshrc.
```shell
source <(cremem script)
```
- Enable completion.
- Register commands automatically.
- Enable input to the buffer.

# Available Commands
Choose a command and input it to the buffer.  
`$ cremem`
## register
Register a command.  
`$ cremem register 'ls -FAlhtr'`  
Basically, don't use it.
## show
Show a recorded command.  
`$ cremem show`  
Basically, don't use it.
## remove
Remove recorded commands.  
`$ cremem remove`

# Config
## Config directory
Config is saved in $XDG_CONFIG_HOME/cremem.  
If $XDG_CONFIG_HOME is either not set or empty, a default equal to $HOME/.config should be used.  
It can also be specified by $CREMEM_CONFIG_PATH.

config.json
```json
{
  "ignoreCommands": [
    "ls",
    "cd",
    "mv",
    "cp",
    "rm"
  ]
}
```

# Data
## Data directory
Data is saved in $XDG_DATA_HOME/cremem.  
If $XDG_DATA_HOME is either not set or empty, a default equal to $HOME/.local/share should be used.  
It can also be specified by $CREMEM_DATA_PATH.

