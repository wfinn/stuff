# nodiffy
*nodiffies you if the output of a command has changed*

## What is this?
A simple tool to monitor stuff. Configure crontab to run nodiffy with your command, title and message to get a push notification with the diff if the output of the command changes.  
nodiffy logs every version of the ouput in *~/.nodiffy/*.
A nice side effect of this is that you can link files from there to have an up-to-date version where you need it.
## sendpush

You need to have a program called *sendpush* which accepts a title as first and message as second parameter.
I use [simplepush for easy push notifications](https://simplepush.io/), they have a script called send-encrypted.sh.
Using that my *sendpush* looks like this:
```sh
#!/bin/sh
send-encrypted.sh -k MySimplepushKey -p MyRandomPassword -s MySecretSalt -t "$1" -m "$2"
```
Of course you can create your own *sendpush* if you don't want to use Simplepush.

## Usage
$1: command to be monitored (tip: sort the output if possible)

$2: name, used as title and file name

$3: message to be send

$4: arguments to diff

$3 & $4 are optional

### Example
```sh
nodiffy date "datecheck" "times have changed"
```
## TODO
- getopts (command is the default message title)
- support categories (simplepush event types)
