pwdify
======

A command line utility to password protect static web pages.


### TODO

- [ ] Output each file as it is completed on the status screen
- [ ] Everything should account appropriately for window size (WindowSizeMsg)
- [ ] Write some tests already...
- [ ] GH action for releases
  - [ ] https://goreleaser.com/
- [ ] Update the README with installation and usage

- [ ] Create a GH action for community usage
- [ ] Non-interactive and "quiet" output when a tty doesn't exist (or --quiet/-q)
- [ ] docs/ site for documentation and website

### Notes

- window.crypto.subtle only works in secure context (https, localhost, etc)

### Questions

- [ ] Should files always be overwritten or can they be output separately?
- [ ] Should there be an option to encrypt all html files in a directory (recursively)?

### References

- [Color Scheme](https://color.adobe.com/Blockboster%20Look-color-theme-925247)
- <https://clig.dev/>
