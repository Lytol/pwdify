pwdify
======

A command line utility to password protect static web pages.


### Notes

- [ ] Password should be able to be specified via command line argument
  - [ ] --password-env
  - [ ] --password-file
  - [ ] Skip the password screen if specified
- [ ] Files should be able to be specified via command line argument (and stdin, newline separated)
  - [ ] This should work for quarto post-render
  - [ ] Skip the files screen if specified
- [ ] A file must be selected to continue on the files screen
- [ ] Show keyboard help for each model screen
- [ ] Select all / Select none for files
- [ ] Accept a directory rather than assuming CWD
- [ ] Everything should account appropriately for window size (WindowSizeMsg)

### Questions

- [ ] Should files always be overwritten or can they be output separately?
- [ ] Should there be an option to encrypt all html files in a directory (recursively)?

### References

- [Color Scheme](https://color.adobe.com/Blockboster%20Look-color-theme-925247)