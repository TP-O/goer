# Goer

During course registration, the Edusoft server is overloaded. As a result, request processings are also negatively affected. Moreover, sending a bunch of redundant requests for supporting UI and validation by the client (website) makes us wait a very long time to register for a course. In order not to complicate, this tool ignores these redundancies and only sends requests for course registration.

## Install

Which file should you [download](https://github.com/TP-O/goer/releases)?
- Linux: goer
- Windows: goer.exe
- MacOS: mgoer

## Options
![options](/img/options.png)

## Example
[How to get course ID](https://youtu.be/nPnCHI7AVZg)

```bash
$ goer \
    -i ITITIU19180 \
    -p Mypassword \
    -I "IT092IU02  01|IT092IU|Principles of Programming Languages|02|4|0|01/01/0001|0|0|0| |0|ITIT19CS31" \
    -I "IT093IU02  01|IT093IU|Web Application Development|02|4|0|01/01/0001|0|0|0| |0|ITIT19CS31"
```
