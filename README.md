# iWF Golang SDK
[![Build status](https://github.com/indeedeng/iwf-golang-sdk/actions/workflows/ci-integ-test.yml/badge.svg?branch=main)](https://github.com/indeedeng/iwf-golang-sdk/actions/workflows/ci-integ-test.yml)

Golang SDK for [iWF workflow engine](https://github.com/indeedeng/iwf)

See [samples](https://github.com/indeedeng/iwf-golang-samples) for how to use this SDK.
# Contribution
See [contribution guide](CONTRIBUTION.md)

# Development Plan

## 1.0

- [x] Start workflow API
- [x] Executing `start`/`decide` APIs and completing workflow
- [x] Parallel execution of multiple states
- [x] Timer command
- [x] Signal command
- [x] SearchAttribute
- [x] DataAttributes
- [x] StateExecutionLocal
- [x] Signal workflow API
- [x] Get workflow result API
- [x] Search workflow API
- [x] Describe workflow API
- [x] Stop workflow API
- [x] Reset workflow API
- [x] Command type(s) for inter-state communications (e.g. internal channel)
- [x] More workflow start options: IdReusePolicy, cron schedule, retry
- [x] StateOption: Start/Decide API timeout and retry policy
- [x] Reset workflow by stateId/StateExecutionId
- [x] More workflow start options: initial search attributes

## 1.1

- [x] Skip timer API for testing/operation
- [x] Decider trigger type: any command combination

## 1.2
- [x] API improvements to reduce boilerplate code

## 1.3
- [x] Support failing workflow with results
- [x] Improve workflow uncompleted error return(canceled, failed, timeout, terminated)

### 2.0

- [x] Renaming some concepts/APIs with breaking changes(see releaste notes)
- [x] Support workflow RPC
- [x] PARTIAL_WITH_EXCLUSIVE_LOCK persistence loading type
