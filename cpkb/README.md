# CPKB

`cpkb` is command line tool that allows copying texts to clipboard using pipe (`|`)

### Examples
`cat README.md | cpbk` - Copies content of `README.md` to clipboard
<br />
`docker --version | cpbk` - Copies docker version to clipboard
<br />
`ls . | cpbk` - Copies content of current directory to clipboard

### Installation
Download the repository using git
<br />
`git clone https://github.com/Streamer272/cpkb.git && cd cpkb`
<br />
Build main.go
<br />
`go build .`
<br />

##### Windows
1. This PC
2. Advanced system settings
3. Environmental variables
4. New
5. Paste path to built binary
6. Apply

##### Linux
1. `bash add-to-path.sh`
2. Enter root password
