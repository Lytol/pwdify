pwdify
======

A command line utility to password protect static web pages.

### TODO

- [ ] Write some tests already...
- [ ] GH action for releases
  - [ ] https://goreleaser.com/
- [ ] Update the README with installation and usage

- [ ] Create a GH action for community usage
- [ ] Non-interactive and "quiet" output when a tty doesn't exist (or --quiet/-q)
- [ ] Output each file as it is completed on the status screen
- [ ] Everything should be full-screen and appropriately for window size (WindowSizeMsg)
- [ ] docs/ site for documentation and website
- [ ] Status TUI screen should be cancellable with `ctrl-c`

### Notes

- window.crypto.subtle only works in secure context (https, localhost, etc)

### Questions

- [ ] Should files always be overwritten or can they be output separately?
- [ ] Should there be an option to encrypt all html files in a directory (recursively)?

### References

- [Color Scheme](https://color.adobe.com/Blockboster%20Look-color-theme-925247)
- <https://clig.dev/>
