# nodiffy
nodiffies you if the output of a command has changed

$1: command to be monitored (tip: sort the output if possible)

$2: name, used as title and file name

$3: message to be send

$4: arguments to diff

$3 & $4 are optional

You need to have a program called sendpush which accepts a title as first and message as second parameter.
I use [simplepush.io](https://simplepush.io/) for notifications.
