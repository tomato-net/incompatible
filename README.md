# incompatible

A linter that checks your go.mod to ensure there are no direct dependencies with non-modular packages greater than v1.

## Rationale

In the pre-modular world, packages would maintain the same import path no matter which major version of the package is
required. This could, and did, lead to dependencies bumping major versions fairly silently with potentially breaking 
changes.

Now dependencies are typically modules, that must define their major version in the import path for versions greater than
v1, such as for example, `example.org/me/mymod/v2`.  However, some dependencies are still packages and this may be 
undesirable for the reasons above. 

This linter gives you the opportunity to block the import of packages.

## Usage

```bash
$ incompatible
```
