#!/usr/bin/env bash
go-wrapper "download" \
    && go-wrapper "install" \
    && go-wrapper "run"