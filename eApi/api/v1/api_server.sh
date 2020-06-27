#!/bin/bash

swagger serve --host=0.0.0.0 --port=9000 --no-open api.swagger.json &

hostname -I