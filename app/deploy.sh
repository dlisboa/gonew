#!/bin/sh

echo 'deploying...'

# Add deploy steps here
image="gcr.io/${GCP_PROJECT_ID}/${GCP_SERVICE_NAME}"
gcloud builds submit --tag $image
gcloud run deploy ${GCP_SERVICE_NAME} \
  --image $image \
  --region southamerica-east1 \
  --allow-unauthenticated
