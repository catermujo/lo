# wiz - Iterators optimizer for wizardry

![Go Version](https://img.shields.io/badge/Go-%3E%3D%201.18-%23007d9c)
![Build Status](https://github.com/catermujo/wiz/actions/workflows/test.yml/badge.svg)
[![Coverage](https://img.shields.io/codecov/c/github/catermujo/wiz)](https://codecov.io/gh/catermujo/wiz)
[![License](https://img.shields.io/github/license/catermujo/wiz)](./LICENSE)

Iterator-style library based on https://github.com/samber/lo.

Design goals:
- Adopt go iterators wherever possible
- Subset functionality from the original lib to reduce dependencies (namely, `x/text`)
