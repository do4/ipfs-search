package config

// Write JSON here to keep correspondence with browser-based settings

const fileSettingsJSON = `{
  "index": {
    "refresh_interval": "15m",
    "mapping": {
      "total_fields": {
        "limit": "8192"
      }
    },
    "query": {
      "default_field": [
        "content",
        "fingerprint",
        "first-seen",
        "ipfs_tika_version",
        "language.confidence",
        "language.language",
        "language.rawScore",
        "last-seen",
        "metadata.Content-Type",
        "metadata.X-Parsed-By",
        "metadata.author",
        "metadata.created",
        "metadata.date",
        "metadata.description",
        "metadata.isbn",
        "metadata.keywords",
        "metadata.language",
        "metadata.modified",
        "metadata.name",
        "metadata.producer",
        "metadata.publisher",
        "metadata.resourceName",
        "metadata.title",
        "metadata.xmpDM:album",
        "metadata.xmpDM:albumArtist",
        "metadata.xmpDM:artist",
        "metadata.xmpDM:composer",
        "references.hash",
        "references.name",
        "references.parent_hash",
        "size",
        "urls"
      ]
    },
    "analysis": {
      "filter": {
        "shingle_filter": {
          "type": "shingle",
          "min_shingle_size": "5",
          "max_shingle_size": "5",
          "output_unigrams": "false"
        },
        "minhash_filter": {
          "type": "min_hash",
          "hash_count": "1",
          "bucket_count": "512",
          "hash_set_size": "1",
          "with_rotation": "true"
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
    },
    "number_of_shards": "20",
    "number_of_replicas": "0"
  }
}`

const fileMappingJSON = `{
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
    "content": {
      "type": "text",
      "term_vector": "with_positions_offsets"
    },
    "fingerprint": {
      "type": "text",
      "analyzer": "fingerprint_analyzer"
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
    "metadata": {
      "dynamic": "true",
      "properties": {
        "created": {
          "type": "date",
          "format": "date_optional_time",
          "ignore_malformed": true
        },
        "creation-date": {
          "type": "keyword",
          "index": false,
          "doc_values": false
        },
        "creationdate": {
          "type": "keyword",
          "index": false,
          "doc_values": false
        },
        "title": {
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
    "references": {
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
}`

const dirSettingsJSON = `{
  "index": {
    "refresh_interval": "15m",
    "number_of_shards": "20",
    "number_of_replicas": "0"
  }
}`

const dirMappingJSON = `{
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
    "links": {
      "dynamic": true,
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
    "references": {
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
}`

const invalidSettingsJSON = `{
  "index": {
    "refresh_interval": "15m",
    "number_of_shards": "20",
    "number_of_replicas": "0"
  }
}`

const invalidMappingJSON = `{
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
}`
