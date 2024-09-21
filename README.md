pwdify
======

A command line utility to password protect static web pages.


### Notes

- [ ] User should be able to select multiple files
  - [ ] Use list rather than filepicker and recursively include all .html files with checkboxes
  - [ ] Select all / Select none
- [ ] Encrypt selected files using password
  - [ ] Use WebCrypto API
  - [ ] Store password in localstorage
  - [ ] Password form template
- [ ] Password should be able to be specified via command line argument
- [ ] Files should be able to be specified via command line argument (and stdin, newline separated)
- [ ] Show keyboard help for each model screen

### Questions

- [ ] Should files always be overwritten or can they be output separately?
- [ ] Should there be an option to encrypt all html files in a directory (recursively)?

### References

- [Color Scheme](https://color.adobe.com/Blockboster%20Look-color-theme-925247)