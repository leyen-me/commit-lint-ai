English | [简体中文](./README_zh-CN.md)

# Git Hooks Code Standards Checker

A Git Hooks-based automation tool designed to verify code submission standards before commits, ensuring team code quality and commit message consistency.

# Usage

- Copy the commit-msg file to the .git/hooks directory
- Copy the main.exe file to your project directory. If you want to copy it to another directory, please pay attention to the path in commit-msg

## Key Features

- Commit Message Format Validation
- Automated Code Style Checking
- Basic Code Quality Verification
- Sensitive Information Leak Detection
- Branch Naming Convention Enforcement

## Technical Highlights

- Built on Git Hooks mechanism, particularly pre-commit and commit-msg hooks
- Zero intrusion implementation, easy to promote within teams
- Configurable, supporting custom validation rules
- Fail-fast mechanism for early issue detection

## Use Cases

- Standardize team code submission process
- Prevent low-quality code commits
- Ensure commit message standardization
- Reduce code review workload

## Advantages

- Automated execution without manual intervention
- Unified development standards across teams
- Enhanced code quality
- Reduced production issues

## Benefits

- Streamlines the development workflow
- Maintains consistent coding standards
- Improves team collaboration
- Reduces technical debt
- Speeds up code review process