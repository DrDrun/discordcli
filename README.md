# discord-cli
Minimalistic Command-Line Interface for Discord

I haven't tried to mess with the AUR yet, if I ever will, so it is unlikely that anyone has found this.

Reagrdless, most of this isn't my work. Most of it was done by a github user that goes by Rivalo.
All I have done is implemented private messaging, and while I plan to do more, that does not change
how little work done here that is mine.

Questions can be answered at my discord server, which at the time of editing is empty:

https://discord.gg/qp2Q8jB

### How to Install the Master branch?
Currently the easiest working way to install is to use the Go tools. I'm looking at using GCCGO and makefiles to reduce installation steps, and make setting PATHS unnecessary.
* Install the Go Tools and setup the `$GOPATH` (There are loads of tutorial for this part)
* `$ go get -u github.com/Rivalo/discord-cli`
* Go to the `bin` folder inside your `$GOPATH`
* `./discord-cli`

### (Master) Configuration Settings
Configuration files are being stored in JSON format and are automatically created when you first run discord-cli. Do not change the 'key' value inside `{"key":"value"}`, this is the part that discord-cli uses for parsing, missing keys will definitely return errors.

| Setting       | Function         |
| ------------- |-------------|
| username      | Discord Username (emailaddress) |
| password      | Discord Password |
| messagedefault| (true or false) Display messages automatically|
| messages   | Amount of Messages kept in memory |

### (Master) Chat Commands
When inside a text channel, the following commands are available:

| Command       | Function         |
| ------------- |-------------|
| :q      | Quits discord-cli |
| :g      | Change listening Guild|
| :c      | Change listening Channel inside Guild |
| :m [n]      | Display last [n] messages: ex. `:m 2` displays last two messages |
| :p      | Pulls up the private channel menu |
