# awscm

AWS Credentials Manager, a cli tool for switching AWS profiles.

## Example

Accidentally run `aws-cli` command in the incorrect account:

```shell
$ aws ec2 describe-instances --instance-id i-xxxxxx
An error occurred (UnauthorizedOperation) when calling the DescribeInstances operation: You are not authorized to perform this operation.
```

List available profiles:

```shell
$ awscm ls
home
work
```

Switch accounts and rerun command:

```shell
$ awscm use work
$ aws ec2 describe-instances --instance-id i-0a6d8ff0f31xxxxxx
{
    "Reservations": [
        // Output truncated
    ]
}
```

## API

```
$ awscm
awscm is a tool for setting an AWS profile to use.

Usage:
  awscm [command]

Available Commands:
  help        Help about any command
  init        print installation instructions
  ls          List available AWS profiles
  output      use switches to an AWS output format
  region      use switches to an AWS region
  status      Show current settings
  use         use switches to an AWS profile

Flags:
  -h, --help          help for awscm

Use "awscm [command] --help" for more information about a command.
```

## Install

Install with [Homebrew](https://brew.sh/) (macOS only):

```shell
$ brew install jamesroutley/tap/awscm
$ awscm init >> ~/.bashrc  # or ~/.zshrc
```

Install with `go get` (requires `go >= 1.8`):

```shell
$ go get github.com/jamesroutley/awscm
$ awscm init >> ~/.bashrc  # or ~/.zshrc
```
