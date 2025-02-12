# New Configuration Service

The *resource-service* is a Keptn core component and used to manage resources for Keptn project-related entities, 
i.e., project, stage, and service. The entity model is shown below. To store the resources with version control, a git 
repository is used that is mounted as emptyDir volume.  Besides, this service has functionality to upload the git repository 
to any Git-based service such as GitLab, GitHub, Bitbucket, etc.

## Entity model

```
------------          ------------          ------------
|          | 1        |          | 1        |          |
| Project  |----------|  Stage   |----------| Service  |
|          |        * |          |        * |          |
------------          ------------          ------------
  1 \                   1  \                   1  \
     \ *                    \ *                    \ *
   ------------           ------------           ------------ 
   |          |           |          |           |          | 
   | Resource |           | Resource |           | Resource |  
   |          |           |          |           |          |  
   ------------           ------------           ------------ 
```

## Installation

The *resource-service* is installed as a part of [keptn](https://keptn.sh)

## Deploy in your Kubernetes cluster

To deploy the current version of the *resource-service* in your Keptn Kubernetes cluster, 
use the file `deploy/service.yaml` from this repository and apply it.

```console
kubectl apply -f deploy/service.yaml
```

## Delete in your Kubernetes cluster

To delete a deployed *configuration-service*, use the file `deploy/service.yaml` from this repository 
and delete the Kubernetes resources:

```console
kubectl delete -f deploy/service.yaml
```
