#!/bin/bash

BQ_DATASET_NAME=prometheus
BQ_TABLE_NAME=embersword
GCP_PROJECT_ID=ember-sword-399413
bq --location=EU mk --dataset $GCP_PROJECT_ID:$BQ_DATASET_NAME
bq mk --table \
  --schema ./bq-schema.json \
  --time_partitioning_field timestamp \
  --time_partitioning_type DAY $GCP_PROJECT_ID:$BQ_DATASET_NAME.$BQ_TABLE_NAME

