# iwf-golang-sdk
WIP Golang SDK for [iWF workflow engine](https://github.com/indeedeng/iwf)

Contribution is welcome!


# Contribution
See [contribution guide](CONTRIBUTION.md)

# Development Plan

## 1.0

- [x] Start workflow API
- [x] Executing `start`/`decide` APIs and completing workflow
- [x] Parallel execution of multiple states
- [ ] Timer command
- [ ] Signal command
- [ ] SearchAttributeRW
- [ ] DataObjectRW
- [ ] StateLocal
- [ ] Signal workflow API
- [ ] Get workflow DataObjects/SearchAttributes API
- [x] Get workflow result API
- [ ] Search workflow API
- [ ] Cancel workflow API
- [ ] Reset workflow API
- [ ] Command type(s) for inter-state communications (e.g. internal channel)
- [X] AnyCommandCompleted Decider trigger type
- [ ] More workflow start options: IdReusePolicy, cron schedule, retry
- [x] StateOption: Start/Decide API timeout and retry policy
- [ ] Reset workflow by stateId/StateExecutionId

## 1.1
- [ ] More workflow start options: initial search attributes/memo
- [ ] Decider trigger type: AnyCommandClosed
- [ ] WaitForMoreResults in StateDecision
- [ ] Skip timer API for testing/operation
- [ ] LongRunningActivityCommand
- [ ] Failing workflow details
- [ ] Auto ContinueAsNew
- [ ] StateOption: more AttributeLoadingPolicy
- [ ] StateOption: more CommandCarryOverPolicy