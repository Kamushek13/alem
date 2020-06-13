#! /bin/bash
 curl -s https://01.alem.school/api/graphql-engine/v1/graphql --data '{"query":"{user(where:{githubLogin:{_eq:\"Kamushek13\"}}){id}}"}' | sed 's/[^0-9]//g'
