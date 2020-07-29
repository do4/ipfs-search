#!/bin/sh

# Note: this script assumes that the target indexes exist (ipfs_files_v0, ipfs_directories_v0, ipfs_invalids_v0)
# Note: for now, this is a pseudo-script

# Start crawler, create new indexes, test crawling

# Disable crawler

# Register snapshot repository

# Restore snapshot

# Set refresh_interval=-1 and number_of_replicas=0 for efficient reindexing.
PUT /ipfs_invalids_v0/_settings
{
  "index": {
    "refresh_interval": -1,
    "number_of_replicas": 0
  }
}
PUT /ipfs_directories_v0/_settings
{
  "index": {
    "refresh_interval": -1,
    "number_of_replicas": 0
  }
}
PUT /ipfs_files_v0/_settings
{
  "index": {
    "refresh_interval": -1,
    "number_of_replicas": 0
  }
}

# Reindex invalids
POST _reindex
{
  "source": {
    "index": "ipfs_v5",
    "type": "invalid"
  },
  "dest": {
    "index": "ipfs_invalids_v0",
    "type": "_doc"
  }
}

# Reindex directories
POST _reindex
{
  "source": {
    "index": "ipfs_v5",
    "type": "directory"
  },
  "dest": {
    "index": "ipfs_directories_v0",
    "type": "_doc"
  }
}

# Reindex files
POST _reindex
{
  "source": {
    "index": "ipfs_v5",
    "type": "file"
  },
  "dest": {
    "index": "ipfs_files_v0",
    "type": "_doc"
  }
}

