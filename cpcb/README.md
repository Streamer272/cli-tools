# CPCB

`cpcb` is command line tool that allows copying texts to clipboard using pipe (`|`)

### Examples
`cat README.md | cpcb` - Copies content of `README.md` to clipboard
<br />
`docker --version | cpcb` - Copies docker version to clipboard
<br />
`ls . | cpcb` - Copies content of current directory to clipboard

### Installation
Download the repository using git
<br />
`git clone https://github.com/Streamer272/cli-tools.git && cd cli-tools/cpcb`
<br />
Build project
<br />
`bash build.sh`
<br />
Add to /usr/bin directory
<br />
`bash add-to-path.sh`
