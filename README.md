### VICTR Language Challenge 2018

See each of the challenge directories for specific instructions.

Challenge 1: [Roman Numeral Calculator](challenge_01/README.md)
Challenge 2: [Basic Web Server](challenge_02/README.md)


#### Testing

The go test runner is pretty bare-bones, e.g. it doesn't have a `--watch` option. Thankfully, Jest has `jest-runner-go` so it can be used for more full featured testing.

```
$ yarn install # or npm install
$ jest --watch
```

