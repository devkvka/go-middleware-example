#!/bin/bash

curl localhost:9499

curl localhost:9499/secure

curl localhost:9499/something


curl -H "X-Auth-Token: secure" localhost:9499

curl -H "X-Auth-Token: secure" localhost:9499/secure

curl -H "X-Auth-Token: secure" localhost:9499/something


