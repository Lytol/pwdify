pwdify
======

A command line utility to password protect static web pages.


### TODO

- [ ] A file must be selected to continue on the files screen
- [ ] Show keyboard help for each model screen
- [ ] Select all / Select none for files
- [ ] Everything should account appropriately for window size (WindowSizeMsg)
- [ ] GH action for releases

- [ ] Create a GH action for community usage
- [ ] Password should be able to be specified via file: --password-file
- [ ] Non-interactive and "quiet" output when a tty doesn't exist (or --quiet/-q)

### Notes

- window.crypto.subtle only works in secure context (https, localhost, etc)

### Questions

- [ ] Should files always be overwritten or can they be output separately?
- [ ] Should there be an option to encrypt all html files in a directory (recursively)?

### References

- [Color Scheme](https://color.adobe.com/Blockboster%20Look-color-theme-925247)
- <https://clig.dev/>
