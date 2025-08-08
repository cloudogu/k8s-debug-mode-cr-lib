# k8s-debug-mode-cr-lib Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [v0.2.3] - 2025-08-08
### Fixed
- update conditions via status

## [v0.2.1] - 2025-08-07
### Changed
- pass message and reason to Condition helper

## [v0.2.0] - 2025-08-07
### Added
- New phase 'Failed' to indicate a failed debug mode request cr
- Make CR a singleton by using name rule in x-kubernetes-validation
### Fixed
- correct group name without doubled domain

## [v0.1.1] - 2025-08-05
### Fixed
- Fixing Debugmode API, CRD and REST-client for proper debug-mode-operator usability.

## [v0.1.0] - 2025-08-04
### Added
- Initialize debug-mode-cr-lib
  - scaffolding, api and cr client added