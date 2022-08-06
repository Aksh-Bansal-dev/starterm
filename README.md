# starterm

Fancy start screen for terminal.

![starterm](https://raw.githubusercontent.com/Aksh-Bansal-dev/starterm/main/assets/logo.png)

## How to use

Add `.starterm.json` in the home directory

```json
{
  "heading": "",
  "items": [
    {
      "name": "Tmux",
      "description": "Terminal multiplexer",
      "key": "1",
      "cmd": ["gnome-terminal", "--tab", "--", "tmux"]
    },
    {
      "name": "Nvim",
      "description": "Text editor",
      "key": "2",
      "cmd": ["gnome-terminal", "--tab", "--", "nvim"]
    }
  ]
}
```
