#!/usr/bin/env bash

set -e

echo checking dependencies:

if [[ 0 != $(go version >>/dev/null 2>&1; echo $?) ]]; then
    (>&2 echo "go...   NOT ok.   You need to install go. try: brew install go")
    return 16
else
    echo "go...          ok"
fi

if [[ 0 != $(jq --version >>/dev/null 2>&1; echo $?) ]]; then
    (>&2 echo "jq...          NOT ok.   You need to install jq. try: brew install jq")
    return 16
else
    echo "jq...          ok"
fi

if [[ 0 != $(http --version >>/dev/null 2>&1; echo $?) ]]; then
    (>&2 echo "httpie...      NOT ok.   You need to install httpie. try: brew install httpie")
    return 16
else
    echo "httpie...      ok"
fi

if [[ 0 != $(terraform -v >>/dev/null 2>&1; echo $?) ]]; then
    (>&2 echo "terraform...   NOT ok.   You need to install terraform. try: brew install terraform")
    return 16
else
    echo "terraform...   ok"
fi

if [[ 0 != $(minikube version >>/dev/null 2>&1; echo $?) ]]; then
    (>&2 echo "minikube...   NOT ok.   You need to install minikube. try: brew install terraform")
    return 16
else
    echo "minikube...    ok"
fi

if [[ 0 != $(kubectl version >>/dev/null 2>&1; echo $?) ]]; then
    (>&2 echo "kubectl...   NOT ok.   You need to install kubectl. try: brew install kubectl")
    return 16
else
    echo "kubectl...     ok"
fi


if [[ 0 != $(helm version >>/dev/null 2>&1; echo $?) ]]; then
    (>&2 echo "helm...   NOT ok.   You need to install helm. try: see readme")
    return 16
else
    echo "helm...        ok"
fi

if [[ 0 != $(docker -v >>/dev/null 2>&1; echo $?) ]]; then
    (>&2 echo "docker...   NOT ok.   You need to install docker. try: brew install docker")
    return 16
else
    echo "docker...      ok"
fi

if [[ 0 != $(docker-compose -v >>/dev/null 2>&1; echo $?) ]]; then
    (>&2 echo "dockerComp...   NOT ok.   You need to install dockerComp.")
    return 16
else
    echo "dockerComp...  ok"
fi
