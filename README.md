nix-purity-check
================

Data set of paths and hashes from Nix database for purity check experiment

```
sqlite3 /nix/var/nix/db/db.sqlite 'select path,hash,deriver from ValidPaths where path not like "%.drv";' > $(date -u +%s).txt
```
