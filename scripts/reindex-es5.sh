#!/bin/sh

# Note: this script assumes that the target indexes exist (ipfs_files_v0, ipfs_directories_v0, ipfs_invalids_v0)
# Note: for now, this is a pseudo-script

# Register snapshot repository
PUT /_snapshot/ipfs_s3_v0
{
  "type": "s3",
  "settings": {
    "bucket": "ipfs-search-snapshots",
    "storage_class": "standard_ia",
    "compress": true,
    "role_arn": "arn:aws:iam::123456789012:role/TheSnapshotRole"
  }
}

# Register output snapshot repository
PUT /_snapshot/ipfs_s3_v1
{
  "type": "s3",
  "settings": {
    "bucket": "ipfs-search-snapshots-v1",
    "storage_class": "standard_ia",
    "compress": true,
    "role_arn": "arn:aws:iam::123456789012:role/TheSnapshotRole"
  }
}

# Restore snapshot
POST /_snapshot/ipfs_s3_v0/$SRC_SNAP/_restore

# Invalids index
PUT /ipfs_invalids_v0
{
  "settings": {
    "index": {
      "refresh_interval": "-1",
      "mapping": {
          "total_fields": {
              "limit": "8192"
          }
      },
    "number_of_shards" : "5",
    "number_of_replicas": "0"
    }
  },
  "mappings": {
    "properties": {
      "_doc": {
          "dynamic_templates": [
              {
                  "default_noindex": {
                      "match": "*",
                      "mapping": {
                          "index": "no",
                          "doc_values": false,
                          "include_in_all": false,
                          "norms": false
                      }
                  }
              }
          ],
          "properties": {
             "error": {
                "type": "text",
                "index": false
             }
          }
      }
    }
  }
}

# Directories index
PUT /ipfs_directories_v0
{
  "settings": {
    "index": {
      "refresh_interval": "-1",
      "mapping": {
          "total_fields": {
              "limit": "8192"
          }
      },
    "number_of_shards" : "5",
    "number_of_replicas": "0"
    }
  },
  "mappings": {
    "properties": {
      "_doc": {
          "dynamic": "strict",
          "properties": {
              "first-seen": {
                  "type": "date",
                  "format": "date_time_no_millis"
              },
              "last-seen": {
                  "type": "date",
                  "format": "date_time_no_millis"
              },
              "links":  {
                  "dynamic":  true,
                  "properties": {
                      "Hash": {
                          "type": "keyword",
                          "index": true
                      },
                      "Name": {
                          "type": "text"
                      },
                      "Size": {
                         "type": "long",
                         "ignore_malformed": true
                      },
                      "Type": {
                         "type": "keyword"
                      }
                   }
              },
              "size": {
                  "type": "long",
                  "ignore_malformed": true
              },
              "references":  {
                  "properties": {
                      "name": {
                          "type": "text",
                          "index": true
                      },
                      "hash": {
                          "type": "keyword",
                          "index": true
                      },
                      "parent_hash": {
                          "type": "keyword",
                          "index": true
                      }
                  }
              }
          }
      }
    }
  }
}

# Files index
PUT /ipfs_files_v0
{
  "settings": {
    "index": {
      "refresh_interval": "-1",
      "mapping": {
          "total_fields": {
              "limit": "8192"
          }
      },
        "analysis": {
            "filter": {
                "shingle_filter": {
                    "type": "shingle",
                    "min_shingle_size": 5,
                    "max_shingle_size": 5,
                    "output_unigrams": false
                },
                "minhash_filter": {
                    "type": "min_hash",
                    "hash_count": 1,
                    "bucket_count": 512,
                    "hash_set_size": 1,
                    "with_rotation": true
                }
            }
        },
        "analyzer": {
            "fingerprint_analyzer": {
                "tokenizer": "standard",
                "filter": [
                  "shingle_filter",
                  "minhash_filter"
                ]
            }
        }
    "number_of_shards" : "5",
    "number_of_replicas": "0"
    }
  },
  "mappings": {
    "properties": {
     "_doc": {
        "dynamic": "strict",
        "dynamic_templates": [
            {
                "default_noindex": {
                    "match": "*",
                    "mapping": {
                        "index": false,
                        "doc_values": false,
                  "norms": false
                    }
                }
            }
        ],
        "properties": {
            "first-seen": {
                "type": "date",
                "format": "date_time_no_millis"
            },
            "last-seen": {
                "type": "date",
                "format": "date_time_no_millis"
            },
            "content":  {
                "type": "text",
                "term_vector": "with_positions_offsets"
            },
            "fingerprint": {
                "type": "text",
                "analyzer": "fingerprint_analyzer",
            },
            "ipfs_tika_version": {
              "type": "keyword"
            },
            "language": {
              "properties": {
                "confidence": {
                  "type": "keyword"
                },
                "language": {
                  "type": "keyword"
                },
                "rawScore": {
                  "type": "double"
                }
              }
            },
            "metadata":  {
                "dynamic":  "true",
                "properties": {
                    "title" : {
                        "type": "text"
                    },
                    "name": {
                        "type": "text"
                    },
                    "author": {
                        "type": "text"
                    },
                    "description": {
                        "type": "text"
                    },
                    "producer": {
                        "type": "text"
                    },
                    "publisher": {
                        "type": "text"
                    },
                    "isbn": {
                        "type": "keyword"
                    },
                    "language": {
                        "type": "keyword"
                    },
                    "resourceName": {
                      "type": "keyword"
                    },
                    "keywords": {
                       "type": "text"
                    },
                    "xmpDM:album": {
                        "type": "text"
                    },
                    "xmpDM:albumArtist": {
                        "type": "text"
                    },
                    "xmpDM:artist": {
                        "type": "text"
                    },
                    "xmpDM:composer": {
                        "type": "text"
                    },
                    "Content-Type": {
                        "type": "keyword"
                    },
                    "X-Parsed-By": {
                        "type": "keyword"
                    },
                    "created": {
                        "type": "date",
                        "format": "date_optional_time"
                    },
                    "date": {
                        "type": "date",
                        "format": "date_optional_time"
                    },
                    "modified": {
                        "type": "date",
                        "format": "date_optional_time"
                    }
                }
            },
            "urls": {
                "type": "keyword"
            },
            "size": {
                "type": "long",
                "ignore_malformed": true
            },
            "references":  {
                "properties": {
                    "name": {
                        "type": "text"
                    },
                    "hash": {
                        "type": "keyword"
                    },
                    "parent_hash": {
                        "type": "keyword"
                    }
                }
            }
        }
      }
    }
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

# Make snapshot of new indexes
PUT /_snapshot/ipfs_s3_v1/200729-1910
{
  "indices": "ipfs_invalids_v0,ipfs_directories_v0,ipfs_files_v0"
}
