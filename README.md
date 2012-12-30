nix-purity-check
================

Data set of paths and hashes from Nix database for purity check experiment

sqlite3 /nix/var/nix/db/db.sqlite  'select path,hash from ValidPaths limit 10000000;' > $(date +%s).txt
