# Introduction

## Description

**Service of states** - state machine for exchanging states of process between initiators/orchestrators/executors.

There are four base entities in this service:
- Models
- Processes
- Steps
- States

![Alt text](../images/Entities.png?raw=true "Main entities")

## Models

Model is a graph of some workflow with its steps. 


For example, let's see on imaginary `5-STEPS MODEL`

![Alt text](../images/Model.png?raw=true "Main entities")

So, we can create any Model, give name to it and use it further.

We see 5 steps inside this Model (steps S1-S5). Steps are joined with themselves into certain sequence. 

We can create models using API `/api/createModel`.

## Processes

Processes are instances of concrete models. Processes save steps from models and their states. Each process is identified by its process_token (UUID).

We can create process using API `/api/createProcess`.


![Alt text](../images/Process_initial_state.png?raw=true "Main entities")

We see that Process inherited steps from Model it's based on and each Step is in certain State. Right after creating all steps are in `NONE` state.

## States

Real process executor gets `process_token` from initiator/orhestrator and writes states for steps of this process. 

Executors can set state of steps using API `/api/setState`.

Let's see example of process after some of its steps have been set.

![Alt text](../images/Process_after_some_time.png?raw=true "Main entities")

> [!NOTE]
> It is similar to Airflow or Camunda workflows excluding any implementation of executors/scripts inside the service. All scripts, services and executors should be outside of the service. You have to write own orchestrator or tasks executors based on this states machine.

## Which states are allowed?

Each step can be in one of these states

![Alt text](../images/States.png?raw=true "Main entities")


# Installation

## Step1. Database

Prepare Postgre database. Create user in it, create schema with access for this user (as owner).
```sql
CREATE ROLE <username>;
CREATE SCHEMA <schema_name> AUTHORIZATION <username>;
```

## Step2. Configuration

Distributive contains file `D:\projects\stathemv2\config_example.yml` 
Copy this file and rename to `config.yml`

File contains settings of application:
```yaml
# ------------------
# database settings
# ------------------

DB_CONN_STRING: "postgres://postgres:<pass>@<user>:<port>/postgres?sslmode=disable&search_path=<schema>"
DB_SCHEMA: "<schema>"
DB_RECREATE_TABLES: true
DB_QUERY_TIMEOUT: 3000s


# ------------------
# service settings
# ------------------

APP_PORT: <some port>
APP_LIMIT_BODY_SIZE_IN_BYTES: 1048576

```

Fill up `database settings` section according to previous chapter info.

Fill up all `APP_PORT` setting. 

> [!IMPORTANT]
> Pay attention to `DB_RECREATE_TABLES: true`. It should be true just one time to create needed tables in database. From start of using this parameter has to be `DB_RECREATE_TABLES: false`.


## Step3. Roles and users

Take file `./users_example.txt`, copy it and rename to `./users.txt`.

Put users of service into `./users.txt`.

`Users.txt` has the format:
```txt
<username1>,<password_hash1>,<right1>:<rightN>
<username2>,<password_hash2>,<right1>:<rightM>
```

List of avalaible rights (each rigth gives access to one endpoint): 
`createModel`, `getModel`, `createProcess`, `finishProcess`, `getProcesses`, `getProcess`, `setState`, `getStates`


For instance:
```txt
user1,#somehash123,getProcesses:createProcess:finishProcess
user2,#somehash567,setState:getStates
user3,#somehash901,getStates
```
To generate hash of password use CLI utility `keygen` or `keygen.exe` which is included into distributive.

## Step4. Start service

1. Start `service` (`service.exe`) for the first time (with `DB_RECREATE_TABLES: true`)
2. Check up that tables in DB have been created (`events`, `models`, `processes`)
3. Stop service
4. Change setting `DB_RECREATE_TABLES: true` -> `DB_RECREATE_TABLES: false`
5. Start `service` (`service.exe`) again in production mode

## Step5. Log of execution

Log is written into console. So you can gather it from there.

# API specification

### Authentication

You should use Basic Auth method of authentication with User and Password.

### <span background=blue>Create Model</span>

`[POST]` `/api/createModel`


Request body: 
```json
    {
        "model_name": "3 STEPS MODEL",
        "model_nodes": ["S1","S2","S3"],
        "model_edges": [{"S1":"S2"},{"S2":"S3"}],
        "model_payload": "any json content"
    }
``` 
Response: `200-OK`

> [!NOTE]
> When we create new model with the same name as existing model then existing model gets deactivated and new model becomes active.

### Get model

`[POST]` `/api/getModel`

Request body: 
```json
    {
        "model_name":"3 STEPS MODEL"
    }
```
Response: `200-OK`
```json
    {
        "model_id": 2,
        "model_name": "3 STEPS MODEL",
        "model_nodes": ["S1","S2","S3"],
        "model_edges": [
            {"S1": "S2"},
            {"S2": "S3"}
        ],
        "model_payload": "",
        "model_creator": "postman_user",
        "model_created_at": "2026-10-06T16:42:58.284868+03:00",
        "model_deactivator": "",
        "model_deactivated_at": null
    }
```
### Get models

`[POST]` `/api/getModels`

Request body: `empty`

Response: `200-OK`

```json
    [
        <array of responses getModel api>
    ]
```

### Create process

`[POST]` `/api/createProcess`

Request body: 
```json
    {
        "model_name": "3 STEPS MODEL",
        "process_payload":"any json content"
    }
```

Response: `200-OK`

```json
    {
        "process_token": "bb0bc094-26e4-4a94-a05b-c03c30c20a53"
    }
```

> [!NOTE]
> During process creation all steps of this process set into initial state `NONE` 

### Get process

`[POST]` `/api/getProcess`

Request body: 
```json
    {
        "process_token": "bb0bc094-26e4-4a94-a05b-c03c30c20a53"
    }
```

Response: `200-OK`

```json
    {
        "model_name": "3 STEPS MODEL",
        "process_token": "bb0bc094-26e4-4a94-a05b-c03c30c20a53",
        "process_params": "",
        "process_payload": "calculation_id:123",
        "process_created_by": "postman_user",
        "process_created_at": "2026-10-08T13:34:49.203455+03:00",
        "process_finished_by": "",
        "process_finished_at": null
    }
```

### Get processes

`[POST]` `/api/getProcesses`

Request body: `empty`

Response: `200-OK`

```json
    [
        <array of responses getProcess api>
    ]
```

### Set state

`[POST]` `/api/setState`

Request body: 
```json
    {
        "process_token":"bb0bc094-26e4-4a94-a05b-c03c30c20a53",
        "step_name": "S1",
        "state_name":"RUN",
        "payload": "any json content"
    }
```
> [!NOTE]
> You can use one of these states `NONE`, `RUN`, `SKIP`, `TIMEOUT`, `DONE`, `ERROR` in accordance to [States](#which-states-are-allowed).

Response: `200-OK`

```json
    {
        "process_token": "bb0bc094-26e4-4a94-a05b-c03c30c20a53",
        "step_name": "S1",
        "state_name": "RUN",
        "checks_results": {
            "is_step_name_right": true,
            "is_state_name_right": true,
            "is_step_sequence_right": true,
            "is_state_sequence_right": true,
            "is_set_while_process_run": true
        },
        "created_by": "postman_user",
        "created_at": "2026-10-08T13:40:10.792872+03:00"
    }
```

### Get states

`[POST]` `/api/getStates`

Request body: 
```json
    {
        "process_token":"bb0bc094-26e4-4a94-a05b-c03c30c20a53",
        "filter": {
            "is_set_while_process_run": true,
            "is_actual_only": false,
            "is_step_name_right": true,
            "is_state_name_right": true,
            "is_step_sequence_right": true,
            "is_state_sequence_right": true,
            "is_child_step_name": "S2"
        }
    }
```

Response: `200-OK`

```json
    [
        {
            "process_token": "bb0bc094-26e4-4a94-a05b-c03c30c20a53",
            "step_name": "S1",
            "state_name": "RUN",
            "checks_results": {
                "is_step_name_right": true,
                "is_state_name_right": true,
                "is_step_sequence_right": true,
                "is_state_sequence_right": true,
                "is_set_while_process_run": true
            },
            "created_by": "postman_user",
            "created_at": "2026-10-08T13:40:10.792872+03:00"
        }
    ]
```

## Filters

How do filters work when we do `api/getStates` request?

| Filter | True | False | Don't use this filter |
| --- | --- | --- | --- |
| `is_set_while_process_run` | Selects all states which were set while `process was active` (not finished) | Selects all states which were set `after process had been finished` | Selects all |
| `is_actual_only` | Selects all `actual states` for every step of process | Selects all `not actual states` for every step of process | Selects all |
| `is_step_name_right` | Selects all states with `step_name exists in model` | Selects all states with `step_name doesn't exist in model` | Selects all |
| `is_step_sequence_right` | Selects all states with `steps were set in right steps sequence`  | Selects all `state with steps were set in wrong steps sequence` | Selects all |
| `is_state_name_right` | Selects all states with `state_name exists in states model` | Selects all states with `state_name doesn't exist in state model` | Selects all |
| `is_state_sequence_right` | Selects all `states were set in right states sequence` | Selects all `states were set in wrong states sequence` | Selects all |
