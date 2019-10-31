# ntci [![Build Status](https://travis-ci.org/andy-zhangtao/ntci.svg?branch=master)](https://travis-ci.org/andy-zhangtao/ntci) [![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=andy-zhangtao_ntci&metric=alert_status)](https://sonarcloud.io/dashboard?id=andy-zhangtao_ntci)

A New Tiny CI Tool

## Agents

+ ci-agent
> Listen the request from git repository, and parse it.

CI Agent listen on 8000 as a default port. User change port use `CI_WEB_PORT` env variable. 

User use `CI_WEB_LOG_LEVEL` change log level, there are five levels: debug(default), info, warn and error.