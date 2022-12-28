# iWF Golang SDK
Golang SDK for [iWF workflow engine](https://github.com/indeedeng/iwf)

# Contribution
See [contribution guide](CONTRIBUTION.md)

# Development Plan

## 1.0

- [x] Start workflow API
- [x] Executing `start`/`decide` APIs and completing workflow
- [x] Parallel execution of multiple states
- [x] Timer command
- [x] Signal command
- [x] SearchAttributeRW
- [x] DataObjectRW
- [x] StateLocal
- [x] Signal workflow API
- [ ] Get workflow DataObjects/SearchAttributes API
- [x] Get workflow result API
- [ ] Search workflow API
- [x] Stop workflow API
- [x] Reset workflow API
- [x] Command type(s) for inter-state communications (e.g. internal channel)
- [X] AnyCommandCompleted Decider trigger type
- [x] More workflow start options: IdReusePolicy, cron schedule, retry
- [x] StateOption: Start/Decide API timeout and retry policy
- [x] Reset workflow by stateId/StateExecutionId
- [x] More workflow start options: initial search attributes

## Future
- [ ] Decider trigger type: AnyCommandClosed
- [ ] WaitForMoreResults in StateDecision
- [ ] Skip timer API for testing/operation
- [ ] LongRunningActivityCommand
- [ ] Failing workflow details
- [ ] Auto ContinueAsNew
- [ ] StateOption: more AttributeLoadingPolicy
- [ ] StateOption: more CommandCarryOverPolicy