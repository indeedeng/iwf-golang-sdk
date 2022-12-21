# iwf-golang-sdk
WIP Golang SDK for [iWF workflow engine](https://github.com/indeedeng/iwf)

Contribution is welcome!


# Contribution
### How to update IDL and the generated code
1. Install openapi-generator using Homebrew if you haven't. See more [documentation](https://openapi-generator.tech/docs/installation)
2. Check out the idl submodule by running the command: `git submodule update --init --recursive`
3. Run the command `git submodule update --remote --merge` to update IDL to the latest commit
4. Run `make idl-code-gen` to refresh the generated code