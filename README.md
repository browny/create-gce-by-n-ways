
# Create Compute Engine Instances in N ways

**Note**: You need to replace the variables enclosed by `<>` (e.g. `<YOUR_PROJECT_ID>, <SERVICE_ACCOUNT_KEY>`), NOT just copy and paste

**Prerequisites**: Create service accounts with `compute.instanceAdmin.v1` role, download the JSON key file created by Cloud Console

## N ways

1. **Cloud Console**

   https://cloud.google.com/compute/docs/quickstart-linux

2. **Cloud Shell**

   ```bash
   # no need to authenticate gcloud, it has been authed by default
   gcloud config list
    
   # get command from Equivalent command line
   gcloud compute instances create vm-by-cloud-shell --zone=asia-east1-b --machine-type=f1-micro
   ```

3. **API Explorer**

   ```bash
   # use Equivalent REST
   https://cloud.google.com/compute/docs/reference/rest/v1/instances/insert
   ```

4. **Local Shell**

   ```bash
   # run an environment with gcloud SDK
   docker pull google/cloud-sdk:latest
   docker run -ti google/cloud-sdk:latest bash

   # authentication and set project
   export PROJECT_ID=<YOUR_PROJECT_ID>
   gcloud auth login
   gcloud config set project $PROJECT_ID
   gcloud config list

   # create compute engine instance
   gcloud compute instances create vm-by-localhost --zone=asia-east1-b --machine-type=f1-micro
   ```

5. **Curl**

   ```bash
   
   # run an environment with gcloud SDK and mount demo files under /lab directory
   docker pull google/cloud-sdk:latest
   docker run -v `pwd`:/lab -ti google/cloud-sdk:latest bash

   # authentication (user account or service account, choose one)
   # user account
   gcloud auth login 
   # service account (recommended)
   gcloud auth activate-service-account --key-file <SERVICE_ACCOUNT_KEY>

   # prepare API request body
   export PROJECT_ID=<YOUR_PROJECT_ID>
   sed "s/PROJECT_ID/$PROJECT_ID/" /lab/vm.json > /lab/tmp.json

   # create compute engine instance
   curl -X POST \
    -H "Authorization: Bearer "$(gcloud auth print-access-token) \
    -H "Content-Type: application/json; charset=utf-8" \
    --data @/lab/tmp.json \
    https://www.googleapis.com/compute/v1/projects/$PROJECT_ID/zones/asia-east1-b/instances

   # get operation (use the operation ID returned from above response)
   curl -H "Authorization: Bearer "$(gcloud auth print-access-token) \
    https://www.googleapis.com/compute/v1/projects/$PROJECT_ID/zones/asia-east1-b/operations/<operation_id>
   ```

6. **Client library (Golang)**

   ```bash
   # create service accounts with `compute.instanceAdmin.v1` role, download the JSON key file
   created by Cloud Console

   # run an environment with golang runtime and mount demo files under /lab directory (read only)
   docker pull golang:latest
   docker run -v `pwd`:/lab -ti golang:latest bash

   # prepare API request body (replace PROJECT_ID with your project ID)
   export PROJECT_ID=<YOUR_PROJECT_ID>
   sed "s/PROJECT_ID/$PROJECT_ID/" /lab/vm.json > /lab/tmp.json

   # build command
   cd /lab/go
   # (optional) go mod vendor
   go build -v -mod=vendor

   # run command
   export GOOGLE_APPLICATION_CREDENTIALS=/lab/<SERVICE_ACCOUNT_KEY>
   ./govm $PROJECT_ID /lab/tmp.json
   ```

7. **Deployment Manager**

   ```bash
   # run an environment with gcloud SDK and mount demo files under /lab directory
   docker pull google/cloud-sdk:latest
   docker run -v `pwd`:/lab -ti google/cloud-sdk:latest bash

   # authentication and set project
   export PROJECT_ID=<YOUR_PROJECT_ID>
   gcloud auth login
   gcloud config set project $PROJECT_ID
   gcloud config list

   # prepare deployment manager config
   sed "s/PROJECT_ID/$PROJECT_ID/" /lab/vm.yaml > /lab/tmp.yaml
   cat /lab/tmp.yaml

   # optional
   gcloud services list --available

   # apply deployment manager config
   gcloud services enable deploymentmanager.googleapis.com
   (optional) gcloud deployment-manager deployments delete dpm-single-vm
   gcloud deployment-manager deployments create dpm-single-vm --config /lab/tmp.yaml
   ```
