# NSYSS

`nsyss` is command line tool that makes creating new systemctl service much easier 

### Examples
`nsyss --name your_service_name --script your_bash_script_path --description your_service_description` - 
Creates new systemctl service with name `your_service_name`, start bash script `your_bash_script_path` and description `your_service_description`

### Installation
Download the repository using git
<br />
`git clone https://github.com/Streamer272/cli-tools.git && cd cli-tools/nsyss`
<br />
Build project
<br />
`bash build.sh`
<br />
Add to /usr/bin directory
<br />
`bash add-to-path.sh`
