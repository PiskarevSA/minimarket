#!/usr/bin/bash
time act --secret GITHUB_TOKEN="$(gh auth token)" -j statictest
time act --secret GITHUB_TOKEN="$(gh auth token)" -j build