# PTCB

`ptcb` is command line tool that allows pasting text from clipboard at runtime

### Examples
`cat README.md | cpcb && echo $(ptcb)` - Copies content of `README.md` to clipboard and echoes it back to terminal
<br />
`docker --version | cpcb && echo $(ptcb)` - Copies docker version to clipboard and echoes it back to terminal
<br />
`ls . | cpcb && echo $(ptcb)` - Copies content of current directory to clipboard and echoes it back to terminal

### Installation
Download the repository using git
<br />
`git clone https://github.com/Streamer272/cli-tools.git && cd cli-tools/ptcb`
<br />
Build project
<br />
`bash build.sh`
<br />
Add to /usr/bin directory
<br />
`bash add-to-path.sh`
