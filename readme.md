# Openshift RBAC lookup

## Building and Running
You can either download the binary from the release or clone and build
## Download the binary
Download the latest Release for your architecture [here](https://github.com/mirroredge/oc-roles/releases)
## Build from source
This project requires go modules to be turned on to build it.

Clone the repo, then run inside the repo run 
```go get -v -t -d ./... && go build .```
 this will create an executable oc-roles

## Running
There are two commands to run one to get all the users for a given role, and another to get all roles for a given user.

### Roles for a User
Log into your Kubernetes or openshift cluster so a valid kube config exists in your ~/.kube directory
Run ```./oc-roles user-roles <username>```

### Users for a Role
Log into your Kubernetes or openshift cluster so a valid kube config exists in your ~/.kube directory
Run ```./oc-roles roles-user <rolename>```

## Output
You can specify table or json output by using the -o or --output flag. Currently the supported values are ```table``` or ```json```

```./oc-roles -o json roles-user <rolename>```

## Note
This was just a one day side project, it is not intended for production use 