# Security Policy Admissions

A simple validating admissions controller/ webhook to deny deployment of anything which:

1. Has `priviledged: true`
2. Has `allowEscalation: true`
3. Has `runAsNonRoot: false`
4. Has `readOnlyRootFilesystem: false`
5. Doesn't drop all capabilities

This will run against any deployment, including against third party deployments via helm, to gate my Kubernetes clusters in the way I want them

## Why not `AdmissionConfiguration` (or some variation thereof)

Mainly I'd like to be a lot stricter, and a lot more flexible, than the out-of-the-box stuff allows. I also want to be able to run against things other than pods in the future, such as my own CRDs
