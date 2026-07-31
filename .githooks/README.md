# Git hooks

This repo is public — `.githooks/pre-commit` runs `gitleaks protect --staged`
before every commit and blocks it if a secret is detected.

Git does not enable a repo's hooks directory automatically on clone (for
security reasons), so **run this once after cloning**:

```
git config core.hooksPath .githooks
```

Requires `gitleaks` installed locally (e.g. `sudo pacman -S gitleaks` on
Arch, or see https://github.com/gitleaks/gitleaks#installing).
