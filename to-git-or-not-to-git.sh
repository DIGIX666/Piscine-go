#!/bin/bash

curl -s https://zone01normandie.org/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{login:{_eq:\"Thox66\"}}){id}}"}' | cut -c24-26