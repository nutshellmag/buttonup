## Buttonup
Buttonup is an alternative front-end for Buttondown subscriptions. It is intended for only one subscription and allows newsletter owners to provide an alternative, more appealing landing page.

## Building
0. Install Git and [Golang](https://golang.org/dl)
1. Clone this repository (`git clone https://github.com/nutshellmag/buttonup.git`)
2. Set the `BUTTONDOWN_KEY` to have a value of your Buttondown URL.
3. Build the server app (`go build -o buttonup`)
4. Run the server (`./buttonup`)

If using this in production for something other than Nutshell, you should review the contents of `main.go` to make any relevant changes that are necessary for your deployment.

## Acknowledgements
This codebase is licensed under the Apache 2.0 license, which can be found in the `LICENSE` folder in the root of this directory.